package tws

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalcBasicScore(t *testing.T) {
	nodes := []Node{
		{RaterId: 2, TargetId: 1, Deep: 1, Score: 8},
		{RaterId: 5, TargetId: 1, Deep: 1, Score: 6},
		{RaterId: 5, TargetId: 2, Deep: 2, Score: 10},
		{RaterId: 3, TargetId: 2, Deep: 2, Score: 4},
		{RaterId: 4, TargetId: 3, Deep: 3, Score: 8},
		{RaterId: 5, TargetId: 3, Deep: 3, Score: 4},
	}

	expectedResult := map[int64]int{
		1: 71884,
		2: 67593,
		3: 60000,
	}

	result, _ := Calc(nodes)

	assert.Equal(t, expectedResult, result)
}

func TestCalcBasicScore2(t *testing.T) {
	nodes := []Node{
		{RaterId: 2, TargetId: 1, Deep: 1, Score: 8},
		{RaterId: 5, TargetId: 1, Deep: 1, Score: 6},
		{RaterId: 5, TargetId: 2, Deep: 2, Score: 10},
		{RaterId: 3, TargetId: 2, Deep: 2, Score: 4},
		{RaterId: 4, TargetId: 3, Deep: 3, Score: 8},
		{RaterId: 5, TargetId: 3, Deep: 3, Score: 4},
		// 追加
		{RaterId: 6, TargetId: 1, Deep: 1, Score: 2},
	}

	expectedResult := map[int64]int{
		1: 56400,
		2: 67593,
		3: 60000,
	}

	result, _ := Calc(nodes)

	assert.Equal(t, expectedResult, result)
}

func TestCalcBasicScore3(t *testing.T) {
	nodes := []Node{
		{RaterId: 2, TargetId: 1, Deep: 1, Score: 8},
		{RaterId: 5, TargetId: 1, Deep: 1, Score: 6},
		{RaterId: 5, TargetId: 2, Deep: 2, Score: 10},
		{RaterId: 3, TargetId: 2, Deep: 2, Score: 4},
		{RaterId: 4, TargetId: 3, Deep: 3, Score: 8},
		{RaterId: 5, TargetId: 3, Deep: 3, Score: 4},
		{RaterId: 6, TargetId: 1, Deep: 1, Score: 2},
		// 追加
		{RaterId: 3, TargetId: 5, Deep: 2, Score: 8},
	}
	expectedResult := map[int64]int{
		1: 58283,
		2: 76158,
		3: 55805,
		5: 80558,
	}

	result, _ := Calc(nodes)

	assert.Equal(t, expectedResult, result)
}

// 追加1评7，看看其它节点有没有变化
func TestCalcBasicScore4(t *testing.T) {
	nodes := []Node{
		{RaterId: 2, TargetId: 1, Deep: 2, Score: 8},
		{RaterId: 5, TargetId: 1, Deep: 2, Score: 6},
		{RaterId: 5, TargetId: 2, Deep: 3, Score: 10},
		{RaterId: 3, TargetId: 2, Deep: 3, Score: 4},
		{RaterId: 4, TargetId: 3, Deep: 4, Score: 8},
		{RaterId: 5, TargetId: 3, Deep: 4, Score: 4},
		{RaterId: 6, TargetId: 1, Deep: 2, Score: 2},
		{RaterId: 3, TargetId: 5, Deep: 3, Score: 8},
		// 追加
		{RaterId: 1, TargetId: 7, Deep: 1, Score: 9},
	}
	expectedResult := map[int64]int{
		1: 58283,
		2: 76158,
		3: 55805,
		5: 80558,
		7: 90582,
	}

	result, _ := Calc(nodes)

	assert.Equal(t, expectedResult, result)
}

// 追加8评7(8是空节点)，看看其它节点有没有变化
func TestCalcBasicScore5(t *testing.T) {
	nodes := []Node{
		{RaterId: 2, TargetId: 1, Deep: 2, Score: 8},
		{RaterId: 5, TargetId: 1, Deep: 2, Score: 6},
		{RaterId: 5, TargetId: 2, Deep: 3, Score: 10},
		{RaterId: 3, TargetId: 2, Deep: 3, Score: 4},
		{RaterId: 4, TargetId: 3, Deep: 4, Score: 8},
		{RaterId: 5, TargetId: 3, Deep: 4, Score: 4},
		{RaterId: 6, TargetId: 1, Deep: 2, Score: 2},
		{RaterId: 3, TargetId: 5, Deep: 3, Score: 8},
		{RaterId: 1, TargetId: 7, Deep: 1, Score: 9},
		// 追加
		{RaterId: 8, TargetId: 7, Deep: 1, Score: 2},
	}
	expectedResult := map[int64]int{
		1: 58283,
		2: 76158,
		3: 55805,
		5: 80558,
		7: 57985,
	}

	result, _ := Calc(nodes)

	assert.Equal(t, expectedResult, result)
}

// 追加2评7(2是上层节点)，看看其它节点有没有变化
func TestCalcBasicScore6(t *testing.T) {
	nodes := []Node{
		{RaterId: 2, TargetId: 1, Deep: 2, Score: 8},
		{RaterId: 5, TargetId: 1, Deep: 2, Score: 6},
		{RaterId: 5, TargetId: 2, Deep: 3, Score: 10},
		{RaterId: 3, TargetId: 2, Deep: 3, Score: 4},
		{RaterId: 4, TargetId: 3, Deep: 4, Score: 8},
		{RaterId: 5, TargetId: 3, Deep: 4, Score: 4},
		{RaterId: 6, TargetId: 1, Deep: 2, Score: 2},
		{RaterId: 3, TargetId: 5, Deep: 3, Score: 8},
		{RaterId: 1, TargetId: 7, Deep: 1, Score: 9},
		{RaterId: 8, TargetId: 7, Deep: 1, Score: 2},
		// 追加
		{RaterId: 2, TargetId: 7, Deep: 1, Score: 10},
	}
	expectedResult := map[int64]int{
		1: 58283,
		2: 76158,
		3: 55805,
		5: 80558,
		7: 75650,
	}

	result, _ := Calc(nodes)

	assert.Equal(t, expectedResult, result)
}

// 即使插入,也以最低节点计算Deep
func TestCalcBasicScore111(t *testing.T) {
	nodes := []Node{
		{RaterId: 2, TargetId: 1, Deep: 1, Score: 5},
		{RaterId: 3, TargetId: 2, Deep: 2, Score: 5},
		{RaterId: 4, TargetId: 3, Deep: 3, Score: 5},
		// 插入
		{RaterId: 5, TargetId: 2, Deep: 2, Score: 5},
	}
	expectedResult := map[int64]int{
		1: 50502,
		2: 50250,
		3: 50000,
	}

	result, _ := Calc(nodes)

	assert.Equal(t, expectedResult, result)
}

// 以插入节点计算Deep
func TestCalcBasicScore112(t *testing.T) {
	nodes := []Node{
		{RaterId: 3, TargetId: 2, Deep: 1, Score: 5},
		{RaterId: 4, TargetId: 3, Deep: 2, Score: 5},
		// 插入
		{RaterId: 5, TargetId: 2, Deep: 1, Score: 5},
	}
	expectedResult := map[int64]int{
		2: 50250,
		3: 50000,
	}

	result, _ := Calc(nodes)

	assert.Equal(t, expectedResult, result)
}

func Test0(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 1, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 2, Score: 10},
	}
	expectedResult := map[int64]float64{
		4: 11.05,
		5: 10.5,
	}
	result, _ := Calc(nodes)
	assert.Equal(t, expectedResult, result)
}

func Test01(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 1, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 2, Score: 10},
		// 追加
		{RaterId: 6, TargetId: 4, Deep: 1, Score: 10},
	}
	expectedResult := map[int64]float64{
		4: 10.87258064516129,
		5: 10.5,
	}
	result, _ := Calc(nodes)
	assert.Equal(t, expectedResult, result)
}

// 基础计算
func Test1(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 1, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 2, Score: 10},
		// 追加
		{RaterId: 6, TargetId: 4, Deep: 1, Score: 6},
	}
	expectedResult := map[int64]float64{
		4: 9.938003761240065,
		5: 11.631370849898476,
	}
	result, _ := Calc(nodes)
	assert.Equal(t, expectedResult, result)
}

func Test11(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 1, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 2, Score: 10},
		{RaterId: 6, TargetId: 4, Deep: 1, Score: 6},
		// 追加
		{RaterId: 8, TargetId: 4, Deep: 1, Score: 6},
	}
	expectedResult := map[int64]float64{
		4: 9.156143596998511,
		5: 11.19282032302755,
	}
	result, _ := Calc(nodes)
	assert.Equal(t, expectedResult, result)
}

// 基础计算，节点8打6分，结果和节点6一致
func TestUpCalc1(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 1, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 2, Score: 10},
		{RaterId: 6, TargetId: 4, Deep: 1, Score: 6},
		// 追加
		{RaterId: 8, TargetId: 4, Deep: 1, Score: 6},
	}
	expectedResult := map[int64]float64{
		5: 1.1547005383792517,
		6: 0.5773502691896256,
		8: 0.5773502691896256,
	}
	result := make(map[id]float64)
	result2 := make(map[id]float64)
	upCalc(nodes, 2, result, result2)
	assert.Equal(t, expectedResult, result)
}

// 节点8打10分，结果和节点5一致
func TestUpCalc2(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 1, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 2, Score: 10},
		{RaterId: 6, TargetId: 4, Deep: 1, Score: 6},
		// 追加
		{RaterId: 8, TargetId: 4, Deep: 1, Score: 10},
	}
	expectedResult := map[int64]float64{
		5: 0.577350269189626,
		6: 1.154700538379251,
		8: 0.577350269189626,
	}
	result := make(map[id]float64)
	result2 := make(map[id]float64)
	upCalc(nodes, 2, result, result2)
	assert.Equal(t, expectedResult, result)
}

// 测试深度打分
func TestUpCalc3(t *testing.T) {
	nodes := []Node{
		{RaterId: 2, TargetId: 1, Deep: 1, Score: 10},
		{RaterId: 8, TargetId: 1, Deep: 1, Score: 9},
		{RaterId: 5, TargetId: 4, Deep: 2, Score: 8},
		{RaterId: 6, TargetId: 4, Deep: 2, Score: 7},
		{RaterId: 7, TargetId: 6, Deep: 3, Score: 10},
		// 插入
		{RaterId: 4, TargetId: 1, Deep: 1, Score: 4},
	}
	expectedResult := map[int64]float64{
		2: 0.7258661863112976,
		4: 1.140646864203468,
		5: 1.462304392478939,
		6: 1.462304392478939,
		8: 0.41478067789217005,
	}
	result := make(map[id]float64)
	result2 := make(map[id]float64)

	upCalc(nodes, 2, result, result2)
	assert.Equal(t, expectedResult, result)
}

// 往末端加入Node2，上级节点不变
func Test2(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 2, Score: 10},
		{RaterId: 6, TargetId: 4, Deep: 2, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 3, Score: 10},
		// 插入
		{RaterId: 4, TargetId: 2, Deep: 1, Score: 10},
	}
	expectedResult := map[int64]float64{
		2: 11.08725806451613,
		4: 10.87258064516129,
		5: 10.5,
	}
	result, _ := Calc(nodes)
	assert.Equal(t, expectedResult, result)
}

// Node3评论Node2，旁支节点不变
func Test3(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 2, Score: 10},
		{RaterId: 6, TargetId: 4, Deep: 2, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 3, Score: 10},
		{RaterId: 4, TargetId: 2, Deep: 1, Score: 10},
		// 插入
		{RaterId: 3, TargetId: 2, Deep: 1, Score: 10},
	}
	expectedResult := map[int64]float64{
		2: 10.902266701849733,
		4: 10.87258064516129,
		5: 10.5,
	}
	result, _ := Calc(nodes)
	assert.Equal(t, expectedResult, result)
}

func Test4(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 2, Score: 10},
		{RaterId: 6, TargetId: 4, Deep: 2, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 3, Score: 10},
		{RaterId: 4, TargetId: 2, Deep: 1, Score: 10},
		// 插入
		{RaterId: 3, TargetId: 2, Deep: 1, Score: 6},
	}
	expectedResult := map[int64]float64{
		2: 10.179003296712969,
		4: 13.244101457684986,
		5: 11.845434264405943,
	}
	result, _ := Calc(nodes)
	assert.Equal(t, expectedResult, result)
}

func Test5(t *testing.T) {
	nodes := []Node{
		{RaterId: 5, TargetId: 4, Deep: 2, Score: 10},
		{RaterId: 6, TargetId: 4, Deep: 2, Score: 10},
		{RaterId: 7, TargetId: 5, Deep: 3, Score: 10},
		{RaterId: 4, TargetId: 2, Deep: 1, Score: 10},
		// 插入
		{RaterId: 3, TargetId: 4, Deep: 2, Score: 6},
	}
	expectedResult := map[int64]float64{
		2: 11.017661011681684,
		4: 10.17661011681683,
		5: 11.885640646055101,
	}
	result, _ := Calc(nodes)
	assert.Equal(t, expectedResult, result)
}
