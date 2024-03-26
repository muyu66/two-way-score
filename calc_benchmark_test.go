package tws

import (
	"math/rand"
	"testing"
)

const (
	MaxNodeCount = 1000000
	MaxUserCount = 1000000
	MaxDeep      = 60
)

// nodes数量公式 = maxLeaf ^ (maxDeep - 1)
func generateNodes(targetId int, deep int64, nodes *[]Node, maxDeep int64, maxLeaf int) {
	deep++
	if maxDeep <= deep {
		return
	}
	for i := 1; i <= maxLeaf; i++ {
		rId := 2 * i * int(deep)
		if rId == targetId {
			continue
		}
		dd := Node{
			RaterId:  rId,
			TargetId: targetId,
			Deep:     deep,
			Score:    rand.Int63n(10) + 1,
		}
		*nodes = append(*nodes, dd)
		generateNodes(rId, deep, nodes, maxDeep, maxLeaf)
	}
}

// 测试小数据集 625
func BenchmarkSmall(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 5, 5)
	b.StartTimer()
	_, _ = Calc(&nodes)
}

// 测试大数据集 390412
func BenchmarkBig(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 5, 25)
	b.StartTimer()
	_, _ = Calc(&nodes)
}

// 测试大复杂深度 393215
func BenchmarkBig2(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 19, 2)
	b.StartTimer()
	_, _ = Calc(&nodes)
}

// 测试大数据集+大复杂深度 256124
func BenchmarkBig3(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 7, 8)
	b.StartTimer()
	_, _ = Calc(&nodes)
}
