package graph

import (
	"fmt"

	"github.com/pyihe/gobase/data-structure/queue"
)

/*	图的存储结构: 邻接表
	用数组和链表作为图的存储方式
	1. 顶点用一维数组存储(也可以用单链表来存储，只是数组更利于顶点的读取);
	一维数组每个元素存储顶点信息的同时, 还需要存储指向顶点第一个邻接点的指针, 以便于查找该顶点的边信息。
	2. 每个顶点的所有邻接点构成一个线性表, 由于邻接点个数不定, 所以用单链表存储,
	无向图时单链表称为顶点的边表, 有向图则称为顶点的出边表(顶点作为弧尾)

	代码均是针对无向图, 有向图对于边进行方向的判断与识别
*/

// 一维数组存储的顶点元素
type vertexNode struct {
	data      VertexType // 顶点数据
	firstEdge *edgeNode
}

// 顶点的边信息
type edgeNode struct {
	adjvex int       // 邻接点在顶点数组中的下标
	weight int       // 边的权重
	next   *edgeNode // 下一个邻接点
}

// ListGraph 图的邻接表存储方式
type ListGraph struct {
	adjList     [MaxVex]*vertexNode // 顶点数组
	posInfo     map[VertexType]int  // 每个顶点在数组中的位置信息
	numVertexes int                 // 顶点数
	numEdges    int                 // 边数
}

func NewListGraph() (*ListGraph, error) {
	m := &ListGraph{
		posInfo: make(map[VertexType]int),
	}
	fmt.Printf("输入顶点数和边数: \n")
	_, err := fmt.Scanf("%d %d", &m.numVertexes, &m.numEdges)
	if err != nil {
		return nil, err
	}

	// 输入每个顶点的序号
	for i := 0; i < m.numVertexes; i++ {
		node := &vertexNode{}
		fmt.Printf("输入第[%d]个顶点的数据(整型): \n", i)
		_, err = fmt.Scanf("%d", &node.data)
		if err != nil {
			return nil, err
		}
		m.adjList[i] = node
		m.posInfo[node.data] = i
	}

	for k := 0; k < m.numEdges; k++ {
		var (
			i, j VertexType
			w    int
		)
		fmt.Printf("输入边(Vi, Vj)上的顶点序号以及权重: \n")
		_, err = fmt.Scanf("%d %d %d", &i, &j, &w)
		if err != nil {
			return nil, err
		}
		edge := &edgeNode{
			adjvex: m.posInfo[j],
			weight: w,
			next:   m.adjList[m.posInfo[i]].firstEdge,
		}
		m.adjList[m.posInfo[i]].firstEdge = edge

		// 无向图边是相互的
		edge = &edgeNode{
			adjvex: m.posInfo[i],
			weight: w,
			next:   m.adjList[m.posInfo[j]].firstEdge,
		}
		m.adjList[m.posInfo[j]].firstEdge = edge
	}
	return m, nil
}

func (m *ListGraph) findVertex(i VertexType) (vertex *vertexNode) {
	for _, v := range m.adjList {
		if v == nil {
			break
		}
		if v.data == i {
			vertex = v
			break
		}
	}
	return
}

func (m *ListGraph) HasEdge(i, j VertexType) bool {
	iVertex := m.findVertex(i)
	if iVertex == nil {
		return false
	}
	firstEdge := iVertex.firstEdge
	for firstEdge != nil {
		if m.adjList[firstEdge.adjvex].data == j {
			return true
		}
		firstEdge = firstEdge.next
	}
	return false
}

func (m *ListGraph) GetDegree(i VertexType) (n int) {
	iVertex := m.findVertex(i)
	if iVertex == nil {
		return
	}
	firstEdge := iVertex.firstEdge
	for firstEdge != nil {
		n += 1
		firstEdge = firstEdge.next
	}
	return
}

func (m *ListGraph) GetAdjacency(i VertexType) (result []VertexType) {
	iVertex := m.findVertex(i)
	if iVertex == nil {
		return
	}
	firstEdge := iVertex.firstEdge
	for firstEdge != nil {
		result = append(result, m.adjList[firstEdge.adjvex].data)
		firstEdge = firstEdge.next
	}
	return
}

func (m *ListGraph) dfs(i int, visited *[MaxVex]bool) {
	visited[i] = true
	fmt.Printf("%v ", m.adjList[i].data)
	p := m.adjList[i].firstEdge
	for p != nil {
		if !visited[p.adjvex] {
			m.dfs(p.adjvex, visited)
		}
		p = p.next
	}
}

func (m *ListGraph) DFSTraverse() {
	var visited = [MaxVex]bool{}
	for i := 0; i < m.numVertexes; i++ {
		if visited[i] {
			continue
		}
		m.dfs(i, &visited)
	}
}

func (m *ListGraph) BFSTraverse() {
	var (
		p       *edgeNode
		visited [MaxVex]bool
		mQueue  = queue.NewLinkQueue()
	)
	for i := 0; i < m.numVertexes; i++ {
		if visited[i] {
			continue
		}
		visited[i] = true
		fmt.Printf("%v ", m.adjList[i].data)
		_ = mQueue.EnQueue(i)
		for mQueue.Len() > 0 {
			v, _ := mQueue.DeQueue()
			i = v.(int)
			p = m.adjList[i].firstEdge
			for p != nil {
				if !visited[p.adjvex] {
					visited[p.adjvex] = true
					fmt.Printf("%v ", m.adjList[p.adjvex].data)
					_ = mQueue.EnQueue(p.adjvex)
				}
				p = p.next
			}
		}
	}
}
