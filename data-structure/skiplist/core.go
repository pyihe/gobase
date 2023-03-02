package go_skipList

import (
	"math/rand"
	"time"
)

/*
	跳表特性:
	1. 底层有序链表
	2. 索引链表
	3. 节点
*/

var (
	rad = rand.New(rand.NewSource(time.Now().UnixNano()))
)

type zLexRangeSpec struct {
	min   string // 最小key
	max   string // 最大key
	minex int    // 是否包含最小
	maxex int    // 是否包含最大
}

type zRangeSpec struct {
	min   float64 // 分数最小值
	max   float64 // 分数最大值
	minex int     // 是否包含最小值
	maxex int     // 是否包含最大值
}

// 链表中的每个节点
type skipNode struct {
	eleKey   string       // 节点name
	score    float64      // 分值
	backward *skipNode    // 后退指针
	level    []*skipLevel // 索引层
}

// 每一层的索引
type skipLevel struct {
	forward *skipNode // 前进指针
	span    uint      // 层跨越的节点数量
}

// 有序链表本身
type skipList struct {
	header *skipNode
	tail   *skipNode
	length uint
	level  int
}

// 元素节点
type ele struct {
	mem   string      // 成员名
	score float64     // 分值
	data  interface{} // 附加数据
}

// 创建一个跳表节点
func newSkipNode(level int, score float64, key string) *skipNode {
	var node = &skipNode{
		eleKey: key,
		score:  score,
		level:  make([]*skipLevel, level),
	}
	for i := 0; i < level; i++ {
		node.level[i] = &skipLevel{
			forward: nil,
			span:    0,
		}
	}
	return node
}

// 创建跳表
func newSkipList() *skipList {
	var zsl = &skipList{
		level:  1,
		length: 0,
		header: newSkipNode(maxLevel, 0, ""),
	}
	for j := 0; j < maxLevel; j++ {
		zsl.header.level[j] = &skipLevel{
			forward: nil,
			span:    0,
		}
	}
	// zsl.header.backward = nil
	// zsl.tail = nil
	return zsl
}

// 随机生成索引层数
func randomLevel() int {
	var level = 1
	for float64(rad.Int()&0xFFFF) < skipListP*0xFFFF {
		level += 1
	}
	if level > maxLevel {
		level = maxLevel
	}
	return level
}

func (e *ele) Name() string {
	return e.mem
}

func (e *ele) Score() float64 {
	return e.score
}

func (e *ele) Data() interface{} {
	return e.data
}

func (zsl *skipList) zslInsert(score float64, mem string) *skipNode {
	var x *skipNode
	var update [maxLevel]*skipNode
	var rank [maxLevel]uint
	var i, level int

	if len(mem) == 0 {
		return nil
	}

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		switch {
		case i == zsl.level-1:
			rank[i] = 0
		default:
			rank[i] = rank[i+1]
		}
		for x.level[i].forward != nil && (x.level[i].forward.score < score || (x.level[i].forward.score == score && x.level[i].forward.eleKey < mem)) {
			rank[i] += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	level = randomLevel()
	if level > zsl.level {
		for i = zsl.level; i < level; i++ {
			rank[i] = 0
			update[i] = zsl.header
			update[i].level[i].span = zsl.length
		}
		zsl.level = level
	}
	x = newSkipNode(level, score, mem)
	for i = 0; i < level; i++ {
		x.level[i].forward = update[i].level[i].forward
		update[i].level[i].forward = x

		x.level[i].span = update[i].level[i].span - (rank[0] - rank[i])
		update[i].level[i].span = (rank[0] - rank[i]) + 1
	}

	for i = level; i < zsl.level; i++ {
		update[i].level[i].span++
	}
	if update[0] == zsl.header {
		x.backward = nil
	} else {
		x.backward = update[0]
	}

	if x.level[0].forward != nil {
		x.level[0].forward.backward = x
	} else {
		zsl.tail = x
	}
	zsl.length++
	return x
}

func (zsl *skipList) zslDeleteNode(x *skipNode, update []*skipNode) {
	var i int
	for i = 0; i < zsl.level; i++ {
		if update[i].level[i].forward == x {
			update[i].level[i].span += x.level[i].span - 1
			update[i].level[i].forward = x.level[i].forward
		} else {
			update[i].level[i].span -= 1
		}
	}
	if x.level[0].forward != nil {
		x.level[0].forward.backward = x.backward
	} else {
		zsl.tail = x.backward
	}
	for zsl.level > 1 && zsl.header.level[zsl.level-1].forward == nil {
		zsl.level--
	}
	zsl.length--
}

func (zsl *skipList) zslDelete(score float64, key string) bool {
	var x *skipNode
	var update = make([]*skipNode, maxLevel)
	var i int

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (x.level[i].forward.score < score ||
			(x.level[i].forward.score == score && x.level[i].forward.eleKey < key)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	if x != nil && score == x.score && x.eleKey == key {
		zsl.zslDeleteNode(x, update)
		return true
	}
	return false
}

func (zsl *skipList) zslUpdateScore(curScore float64, key string, newScore float64) *skipNode {
	var update = make([]*skipNode, maxLevel)
	var x *skipNode
	var i int
	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (x.level[i].forward.score < curScore || (x.level[i].forward.score == curScore && x.level[i].forward.eleKey < key)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	if (x.backward == nil || x.backward.score < newScore) && (x.level[0].forward == nil || x.level[0].forward.score > newScore) {
		x.score = newScore
		return x
	}
	zsl.zslDeleteNode(x, update)
	newNode := zsl.zslInsert(newScore, x.eleKey)
	x = nil
	return newNode
}

func zslValueGteMin(value float64, spec *zRangeSpec) bool {
	if spec.minex != 0 {
		return value > spec.min
	}
	return value >= spec.min
}

func zslValueLteMax(value float64, spec *zRangeSpec) bool {
	if spec.maxex != 0 {
		return value < spec.max
	}
	return value <= spec.max
}

func (zsl *skipList) zslIsInRange(spec *zRangeSpec) bool {
	var x *skipNode

	if spec.min > spec.max || (spec.min == spec.max && (spec.minex != 0 || spec.maxex != 0)) {
		return false
	}
	x = zsl.tail
	if x == nil || !zslValueGteMin(x.score, spec) {
		return false
	}
	x = zsl.header.level[0].forward
	if x == nil || !zslValueLteMax(x.score, spec) {
		return false
	}
	return true
}

func (zsl *skipList) zslFirstInRange(spec *zRangeSpec) *skipNode {
	var x *skipNode
	var i int

	if !zsl.zslIsInRange(spec) {
		return nil
	}

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && !zslValueGteMin(x.level[i].forward.score, spec) {
			x = x.level[i].forward
		}
	}

	x = x.level[0].forward
	if !zslValueLteMax(x.score, spec) {
		return nil
	}
	return x
}

func (zsl *skipList) zslLastInRange(spec *zRangeSpec) *skipNode {
	var x *skipNode
	var i int

	if !zsl.zslIsInRange(spec) {
		return nil
	}

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && zslValueLteMax(x.level[i].forward.score, spec) {
			x = x.level[i].forward
		}
	}
	if !zslValueGteMin(x.score, spec) {
		return nil
	}
	return x
}

func (zsl *skipList) zslDeleteRangeByScore(spec *zRangeSpec, dict map[string]*ele) uint {
	var x *skipNode
	var update = make([]*skipNode, maxLevel)
	var i int
	var removed uint = 0

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && ((spec.minex != 0 && x.level[i].forward.score <= spec.min) || (spec.minex == 0 && x.level[i].forward.score < spec.min)) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward
	for x != nil && ((spec.maxex != 0 && x.score < spec.max) || (spec.maxex == 0 && x.score <= spec.max)) {
		var next = x.level[0].forward
		zsl.zslDeleteNode(x, update)
		delete(dict, x.eleKey)
		removed++
		x = next
	}
	return removed
}

func zslLexValueGteMin(key string, spec *zLexRangeSpec) bool {
	if spec.minex != 0 {
		return key > spec.min
	}
	return key >= spec.min
}

func zslLexValueLteMax(value string, spec *zLexRangeSpec) bool {
	if spec.maxex != 0 {
		return value < spec.max
	}
	return value <= spec.max
}

func (zsl *skipList) zslIsInLexRange(spec *zLexRangeSpec) bool {
	var x *skipNode

	if spec.min > spec.max || (spec.min == spec.max && (spec.minex != 0 || spec.maxex != 0)) {
		return false
	}
	x = zsl.tail
	if x == nil || !zslLexValueGteMin(x.eleKey, spec) {
		return false
	}
	x = zsl.header.level[0].forward
	if x == nil || !zslLexValueLteMax(x.eleKey, spec) {
		return false
	}
	return true
}

func (zsl *skipList) zslFirstInLexRange(spec *zLexRangeSpec) *skipNode {
	var x *skipNode
	var i int

	if !zsl.zslIsInLexRange(spec) {
		return nil
	}
	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && !zslLexValueGteMin(x.level[i].forward.eleKey, spec) {
			x = x.level[i].forward
		}
	}

	x = x.level[0].forward
	if !zslLexValueLteMax(x.eleKey, spec) {
		return nil
	}
	return x
}

func (zsl *skipList) zslLastInLexRange(spec *zLexRangeSpec) *skipNode {
	var x *skipNode
	var i int

	if !zsl.zslIsInLexRange(spec) {
		return nil
	}

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && zslLexValueLteMax(x.level[i].forward.eleKey, spec) {
			x = x.level[i].forward
		}
	}

	if !zslLexValueGteMin(x.eleKey, spec) {
		return nil
	}
	return x
}

// 所有节点分数相同的情况下，以字典序删除区间范围内的节点
func (zsl *skipList) zslDeleteRangeByLex(spec *zLexRangeSpec, dict map[string]*ele) uint {
	var x *skipNode
	var update = make([]*skipNode, maxLevel)
	var removed uint
	var i int

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && !zslLexValueGteMin(x.level[i].forward.eleKey, spec) {
			x = x.level[i].forward
		}
		update[i] = x
	}

	x = x.level[0].forward

	for x != nil && zslLexValueLteMax(x.eleKey, spec) {
		next := x.level[0].forward
		zsl.zslDeleteNode(x, update)
		delete(dict, x.eleKey)
		removed++
		x = next
	}
	return removed
}

func (zsl *skipList) zslDeleteRangeByRank(start, end uint, dict map[string]*ele) uint {
	var x *skipNode
	var update = make([]*skipNode, maxLevel)
	var i int
	var traversed, removed uint

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (traversed+x.level[i].span) < start {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		update[i] = x
	}

	traversed++
	x = x.level[0].forward
	for x != nil && traversed <= end {
		next := x.level[0].forward
		zsl.zslDeleteNode(x, update)
		delete(dict, x.eleKey)
		removed++
		traversed++
		x = next
	}
	return removed
}

func (zsl *skipList) zslGetRank(score float64, ele string) uint {
	var x *skipNode
	var i int
	var rank uint

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && (x.level[i].forward.score < score ||
			(x.level[i].forward.score == score && x.level[i].forward.eleKey <= ele)) {
			rank += x.level[i].span
			x = x.level[i].forward
		}
		if len(x.eleKey) > 0 && x.eleKey == ele {
			return rank
		}
	}
	return 0
}

func (zsl *skipList) zslGetElementByRank(rank uint) *skipNode {
	var x *skipNode
	var traversed uint
	var i int

	x = zsl.header
	for i = zsl.level - 1; i >= 0; i-- {
		for x.level[i].forward != nil && ((traversed + x.level[i].span) <= rank) {
			traversed += x.level[i].span
			x = x.level[i].forward
		}
		if traversed == rank {
			return x
		}
	}
	return nil
}
