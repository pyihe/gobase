package go_skipList

import (
	"fmt"
	"sync"
)

type SkipList interface {
	Len() int                                                                                                                 // ZCARD: 获取有序集合的成员数
	InsertByEle(member string, score float64, addData interface{}) Node                                                       // ZADD: 向有序集合添加一个成员，或者更新已存在成员的分数
	InsertByEleArray(scoreMemAddValues ...interface{}) ([]Node, error)                                                        // ZADD: 向有序集合添加多个成员，或者更新已存在成员的分数
	IncrBy(member string, increment float64) (bool, error)                                                                    // ZINCRBY: 有序集合中对指定成员的分数加上增量 increment
	DeleteByMem(member ...string) (ok bool, err error)                                                                        // ZREM: 移除有序集合中的一个或多个成员
	DeleteByScore(startScore, endScore float64) (removed uint)                                                                // ZREMRANGEBYSCORE: 移除有序集合中分数区间[startScore, endScore]的所有成员
	DeleteByRank(startRank, endRank uint) (removed uint)                                                                      // ZREMRANGEBYRANK: 移除有序集合中给定的排名区间的所有成员
	UpdateScore(member string, score float64) (ok bool, err error)                                                            // 更新指定成员的分数
	UpdateAddData(member string, data interface{}) (ok bool, err error)                                                       // 更新指定成员的附加数据
	GetRank(members ...string) (rankInfo map[string]uint, err error)                                                          // ZRANK: 返回有序集合中指定成员的索引
	GetElementByMem(members ...string) (nodes []Node, err error)                                                              // 根据成员名获取成员信息
	GetElementByRank(ranks ...uint) (nodes []Node, err error)                                                                 // 根据成员排名获取成员信息
	SafeRange(startRank, endRank uint, isReverse bool, do func(mem string, score float64, addData interface{}) error) error   // safe range
	UnsafeRange(startRank, endRank uint, isReverse bool, do func(mem string, score float64, addData interface{}) error) error // unsafe range
}

type Node interface {
	Name() string
	Score() float64
	Data() interface{}
}

type zset struct {
	lock     sync.Mutex
	dict     map[string]*ele
	skipList *skipList
}

func NewSkipList() SkipList {
	return &zset{
		lock:     sync.Mutex{},
		dict:     make(map[string]*ele),
		skipList: newSkipList(),
	}
}

func (z *zset) Len() int {
	return int(z.skipList.length)
}

func (z *zset) InsertByEle(member string, score float64, addData interface{}) Node {
	z.lock.Lock()
	defer z.lock.Unlock()

	data, ok := z.dict[member]
	switch ok {
	case true:
		// 已经存在，分数有变更时才更新
		if data.score != score {
			_ = z.skipList.zslUpdateScore(data.score, member, score)
			data.score = score
		}

	default:
		_ = z.skipList.zslInsert(score, member)
		data = &ele{
			mem:   member,
			score: score,
			data:  addData,
		}
		z.dict[member] = data
	}

	return data
}

func (z *zset) InsertByEleArray(scoreMemAddValues ...interface{}) ([]Node, error) {
	valueCount := len(scoreMemAddValues)
	if valueCount%3 != 0 {
		return nil, errInvalidScoreMemArray
	}

	z.lock.Lock()
	defer z.lock.Unlock()

	var result []Node
	var mem string
	var score float64
	var addData interface{}
	for i := 0; i < valueCount; i += 3 {
		mem = scoreMemAddValues[i].(string)
		addData = scoreMemAddValues[i+2]
		switch t := scoreMemAddValues[i+1].(type) {
		case int:
			score = float64(t)
		case float64:
			score = t
		default:
			return nil, errInvalidScoreType
		}
		// 判断是否已经添加了，如果是，则直接更新
		data, ok := z.dict[mem]
		if ok {
			if score != data.score { // 如果分数不同，需要更新分数
				z.skipList.zslUpdateScore(data.score, mem, score)
				data.score = score
			}
			// 更新addData
			data.data = addData
		} else { // 如果不存在，则直接insert
			node := z.skipList.zslInsert(score, mem)
			if node == nil {
				continue
			}
			data = &ele{
				mem:   mem,
				score: score,
				data:  addData,
			}
			z.dict[mem] = data
		}
		result = append(result, data)
	}
	return result, nil
}

// IncrBy 对指定成员的分数加上增量increment
func (z *zset) IncrBy(member string, increment float64) (bool, error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	data, ok := z.dict[member]
	if !ok {
		return false, errNoKey
	}
	z.skipList.zslUpdateScore(data.score, member, data.score+increment)
	data.score = data.score + increment
	return true, nil
}

func (z *zset) DeleteByMem(members ...string) (bool, error) {
	if len(members) == 0 {
		return false, nil
	}

	z.lock.Lock()
	defer z.lock.Unlock()

	for _, member := range members {
		data, ok := z.dict[member]
		if !ok {
			return false, fmt.Errorf("no mem: %s", member)
		}
		ok = z.skipList.zslDelete(data.score, data.mem)
		if !ok {
			return false, fmt.Errorf("delete %s fail", member)
		}
		delete(z.dict, member)
	}
	return true, nil
}

// DeleteByScore 删除某个分数区间的（闭区间）的节点
func (z *zset) DeleteByScore(startScore, endScore float64) (removed uint) {
	spec := &zRangeSpec{
		min:   startScore,
		max:   endScore,
		minex: 0,
		maxex: 0,
	}
	z.lock.Lock()
	defer z.lock.Unlock()
	removed = z.skipList.zslDeleteRangeByScore(spec, z.dict)
	return removed
}

// DeleteByRank 删除某个排序区间（闭区间）的节点
func (z *zset) DeleteByRank(startRank, endRank uint) (removed uint) {
	if startRank == 0 {
		startRank = 1
	}
	if endRank > z.skipList.length {
		endRank = z.skipList.length
	}
	startRank = z.skipList.length - startRank + 1
	endRank = z.skipList.length - endRank + 1
	if startRank > endRank {
		startRank, endRank = endRank, startRank
	}
	z.lock.Lock()
	defer z.lock.Unlock()
	removed = z.skipList.zslDeleteRangeByRank(startRank, endRank, z.dict)
	return
}

func (z *zset) UpdateScore(member string, score float64) (ok bool, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	data, ok := z.dict[member]
	if !ok {
		return false, errNoKey
	}
	if data.score == score {
		return false, nil
	}
	_ = z.skipList.zslUpdateScore(data.score, member, score)
	data.score = score
	return true, nil
}

func (z *zset) UpdateAddData(member string, data interface{}) (ok bool, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	ele, ok := z.dict[member]
	if !ok {
		return false, errNoKey
	}
	ele.data = data
	return true, nil
}

func (z *zset) GetRank(members ...string) (rankInfo map[string]uint, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	rankInfo = make(map[string]uint)
	var rank uint
	for _, m := range members {
		ele, ok := z.dict[m]
		if !ok {
			err = errNoKey
			return
		}
		rank = z.skipList.zslGetRank(ele.score, ele.mem)
		rank = z.skipList.length - rank + 1
		rankInfo[m] = rank
	}
	return
}

func (z *zset) GetElementByMem(members ...string) (nodes []Node, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()

	for _, m := range members {
		ele, ok := z.dict[m]
		if !ok {
			err = fmt.Errorf("no mem: %s", m)
			nodes = nil
			return
		}
		nodes = append(nodes, ele)
	}
	return
}

func (z *zset) GetElementByRank(ranks ...uint) (nodes []Node, err error) {
	z.lock.Lock()
	defer z.lock.Unlock()
	for _, rank := range ranks {
		r := z.skipList.length - rank + 1
		sNode := z.skipList.zslGetElementByRank(r)
		if sNode == nil {
			nodes = nil
			err = fmt.Errorf("invalid rank: %d", rank)
			return
		}
		data, ok := z.dict[sNode.eleKey]
		if !ok {
			nodes = nil
			z.skipList.zslDelete(sNode.score, sNode.eleKey)
			err = fmt.Errorf("invalid rank: %d", rank)
			return
		}
		nodes = append(nodes, data)
	}
	return
}

func (z *zset) SafeRange(startRank, endRank uint, isReverse bool, do func(mem string, score float64, addData interface{}) error) error {
	z.lock.Lock()
	defer z.lock.Unlock()

	err := z.UnsafeRange(startRank, endRank, isReverse, do)

	return err
}

func (z *zset) UnsafeRange(startRank, endRank uint, isReverse bool, do func(mem string, score float64, addData interface{}) error) error {
	if z.skipList.length == 0 {
		return nil
	}
	if startRank == 0 {
		startRank = 1
	}
	if endRank > z.skipList.length {
		endRank = z.skipList.length
	}
	if startRank > endRank {
		startRank, endRank = endRank, startRank
	}
	loopCnt := endRank - startRank + 1
	var node *skipNode
	switch isReverse {
	case true:
		node = z.skipList.zslGetElementByRank(startRank)
	default:
		node = z.skipList.zslGetElementByRank(endRank)
	}
	if node == nil {
		return errInvalidRank
	}
	for loopCnt > 0 {
		ele := z.dict[node.eleKey]
		if err := do(node.eleKey, node.score, ele.data); err != nil {
			return err
		}
		switch isReverse {
		case true:
			node = node.level[0].forward
		default:
			node = node.backward
		}
		loopCnt--
	}
	return nil
}
