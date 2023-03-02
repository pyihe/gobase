package go_skipList

import (
	"fmt"
	"math/rand"
	"strconv"
	"testing"
)

var (
	s = NewSkipList()
)

func init() {
	//插入
	for i := -1; i <= 5; i++ {
		rand.Float64()
		s.InsertByEle("k"+strconv.Itoa(i), float64(i)*0.1, nil)
	}
	s.InsertByEleArray("k6", 1.1, nil, "k7", 2, nil, "k8", 0, nil, "k9", -1, nil)
}

func TestSkipList_Insert(t *testing.T) {
	output(s)

	s.IncrBy("k1", -1)
	s.IncrBy("k4", 1)
	s.IncrBy("k3", 0.7)
	output(s)

	//根据成员名删除
	fmt.Println(s.DeleteByMem("k1", "k5"))
	output(s)
	////根据分数范围删除
	s.DeleteByScore(1, 2)
	output(s)

	//根据排名删除
	fmt.Println(s.DeleteByRank(1, 1))
	output(s)

	//更新分数
	s.UpdateScore("k8", 100)
	s.UpdateScore("k0", 9.9)
	s.UpdateAddData("k8", "k8")
	output(s)

	//根据成员名获取排名
	info, _ := s.GetRank("k9", "k8", "k0", "k-1")
	fmt.Println(info)

	//根据成员名获取成员数据
	nodes, _ := s.GetElementByMem("k9", "k8", "k0", "k-1")
	for _, n := range nodes {
		fmt.Printf("mem: %s, score: %v, data: %v\n", n.Name(), n.Score(), n.Data())
	}

	fmt.Println()
	nodes, _ = s.GetElementByRank(1, 2, 3, 4)
	for _, n := range nodes {
		fmt.Printf("mem: %s, score: %v, data: %v\n", n.Name(), n.Score(), n.Data())
	}

}

func output(l SkipList) {
	fmt.Printf("output data is: \n")
	l.SafeRange(0, 100, false, func(mem string, score float64, addData interface{}) error {
		fmt.Printf("mem: %s, score: %.2f, data: %v\n", mem, score, addData)
		return nil
	})
}
