package graph

import "math"

type (
	// VertexType 顶点类型
	VertexType int

	// EdgeType 边上的权值类型
	EdgeType int
)

const (
	MaxVex   = 100         // 最大顶点数
	Infinity = math.MaxInt // 边的权值的最大值
)
