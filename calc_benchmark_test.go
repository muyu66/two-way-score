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

// 测试大数据集
func BenchmarkHigh(b *testing.B) {
	var nodes []Node
	for i := 1; i < MaxNodeCount; i++ {
		d := Node{
			RaterId:  rand.Int63n(100) + 1,
			TargetId: rand.Int63n(100) + 1,
			Deep:     rand.Int63n(16) + 1,
			Score:    rand.Int63n(10) + 1,
		}
		nodes = append(nodes, d)
	}
	_, _ = Calc(nodes)
}

// 测试大数据集+大用户数
func BenchmarkHigh2(b *testing.B) {
	var nodes []Node
	for i := 1; i < MaxNodeCount; i++ {
		d := Node{
			RaterId:  rand.Int63n(MaxUserCount) + 1,
			TargetId: rand.Int63n(MaxUserCount) + 1,
			Deep:     rand.Int63n(16) + 1,
			Score:    rand.Int63n(10) + 1,
		}
		nodes = append(nodes, d)
	}
	_, _ = Calc(nodes)
}

// 测试大数据集+大复杂深度
func BenchmarkHigh3(b *testing.B) {
	var nodes []Node
	for i := 1; i < MaxNodeCount; i++ {
		d := Node{
			RaterId:  rand.Int63n(100) + 1,
			TargetId: rand.Int63n(100) + 1,
			Deep:     rand.Int63n(MaxDeep) + 1,
			Score:    rand.Int63n(10) + 1,
		}
		nodes = append(nodes, d)
	}

	_, _ = Calc(nodes)
}

// 测试大数据集+大用户数+大复杂深度
func BenchmarkHigh4(b *testing.B) {
	var nodes []Node
	for i := 1; i < MaxNodeCount; i++ {
		d := Node{
			RaterId:  rand.Int63n(MaxUserCount) + 1,
			TargetId: rand.Int63n(MaxUserCount) + 1,
			Deep:     rand.Int63n(MaxDeep) + 1,
			Score:    rand.Int63n(10) + 1,
		}
		nodes = append(nodes, d)
	}

	_, _ = Calc(nodes)
}
