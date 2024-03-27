package tws

import (
	"math/rand"
	"testing"
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

// 测试小数据集 1500152
func BenchmarkSmallPlus(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 5, 35)
	b.StartTimer()
	_, _ = Calc(&nodes)
}

// 测试大数据集 12944285
func BenchmarkBig(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 5, 60)
	b.StartTimer()
	_, _ = Calc(&nodes)
}

// 测试大数据集 11786432
func BenchmarkBigPlus(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 6, 26)
	b.StartTimer()
	_, _ = Calc(&nodes)
}

// 测试大复杂深度 12582911
func BenchmarkBig2(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 24, 2)
	b.StartTimer()
	_, _ = Calc(&nodes)
}

// 测试大复杂深度 16951984
func BenchmarkBig2Plus(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 13, 4)
	b.StartTimer()
	_, _ = Calc(&nodes)
}

// 测试大数据集+大复杂深度 15874997
func BenchmarkBig3(b *testing.B) {
	b.StopTimer()
	var nodes []Node
	generateNodes(1, 0, &nodes, 9, 8)
	b.StartTimer()
	_, _ = Calc(&nodes)
}
