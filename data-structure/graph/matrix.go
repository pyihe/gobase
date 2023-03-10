package graph

import (
	"errors"
	"fmt"

	"github.com/pyihe/gobase/data-structure/queue"
)

/*	图的存储结构: 邻接矩阵
	用两个数组来表示图，一个一维数组存储图中的顶点信息，一个二维数组存储图中的边或弧的信息
	对于无向图来说，邻接矩阵是以左上角到右下角对角线为轴的轴对称图形

	缺点: 对于稀疏图形来说，存在空间浪费

	代码均是针对无向图, 有向图对于边进行方向的判断与识别
*/

// MatrixGraph 图的邻接矩阵存储方式
type MatrixGraph struct {
	vexs        [MaxVex]VertexType       // 顶点数组
	arc         [MaxVex][MaxVex]EdgeType // 边的二维数组
	numVertexes int                      // 顶点数量
	numEdges    int                      // 边数量
}

// NewMatrixGraph 创建邻接矩阵图,
func NewMatrixGraph() (*MatrixGraph, error) {
	m := &MatrixGraph{}
	fmt.Printf("输入顶点数和边数: \n")
	_, err := fmt.Scanf("%d %d", &m.numVertexes, &m.numEdges)
	if err != nil {
		return nil, err
	}
	if m.numVertexes <= 0 || m.numVertexes > MaxVex {
		return nil, errors.New("invalid vertex num")
	}
	if m.numEdges < 0 || m.numEdges > m.numVertexes*(m.numVertexes-1)/2 {
		return nil, errors.New("invalid edge num")
	}
	// 初始化邻接矩阵
	for i := 0; i < m.numVertexes; i++ {
		for j := 0; j < m.numVertexes; j++ {
			m.arc[i][j] = Infinity
		}
	}

	for k := 0; k < m.numEdges; k++ {
		i, j, w := 0, 0, 0
		fmt.Printf("输入边(Vi, Vj)上的下标i, 下标j和权重w:\n")
		_, err = fmt.Scanf("%d %d %d", &i, &j, &w)
		if err != nil {
			return nil, err
		}
		m.arc[i][j] = EdgeType(w)
		m.arc[j][i] = EdgeType(w)
	}
	return m, nil
}

// HasEdge 判断两个顶点i和j之间是否有边
func (m *MatrixGraph) HasEdge(i, j int) bool {
	if i < 0 || i >= m.numEdges {
		return false
	}
	if j < 0 || j >= m.numEdges {
		return false
	}
	return m.arc[i][j] != Infinity
}

// GetDegree 获取顶点i的度
func (m *MatrixGraph) GetDegree(i int) int {
	if i < 0 || i >= m.numEdges {
		return 0
	}
	degree, l := 0, m.arc[i]
	for j, d := range l {
		if j == m.numVertexes {
			break
		}
		if d != Infinity {
			degree += 1
		}
	}
	return degree
}

// GetAdjacency 获取顶点i的邻接点
func (m *MatrixGraph) GetAdjacency(i int) (result []int) {
	if i < 0 || i >= m.numEdges {
		return
	}
	l := m.arc[i]
	for j, d := range l {
		if d == Infinity {
			continue
		}
		result = append(result, j)
	}
	return result
}

// DFSTraverse 邻接矩阵的深度遍历
func (m *MatrixGraph) DFSTraverse() {
	visited := [MaxVex]bool{}
	for i := 0; i < m.numVertexes; i++ {
		if !visited[i] {
			m.dfs(i, visited)
		}
	}
}

func (m *MatrixGraph) dfs(i int, visited [MaxVex]bool) {
	visited[i] = true
	fmt.Printf("%v ", m.vexs[i])
	for j := 0; j < m.numVertexes; j++ {
		if m.arc[i][j] != EdgeType(Infinity) && !visited[j] {
			m.dfs(j, visited)
		}
	}
}

// BFSTraverse 邻接矩阵的广度优先遍历
func (m *MatrixGraph) BFSTraverse() {
	var (
		visited = [MaxVex]bool{}
		mQueue  = queue.NewLinkQueue()
	)
	for i := 0; i < m.numVertexes; i++ {
		if !visited[i] {
			visited[i] = true            // 设置顶点访问状态
			fmt.Printf("%v ", m.vexs[i]) // 打印顶点
			_ = mQueue.EnQueue(i)        // 将顶点入队列
			for mQueue.Len() > 0 {
				v, ok := mQueue.DeQueue()
				if !ok {
					return
				}
				i = v.(int)
				for j := 0; j < m.numVertexes; j++ {
					if m.arc[i][j] != Infinity && !visited[j] {
						visited[j] = true
						fmt.Printf("%v ", m.vexs[j])
						_ = mQueue.EnQueue(j)
					}
				}
			}
		}
	}
}
