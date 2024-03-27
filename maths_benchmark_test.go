package tws

import (
	"math/rand"
	"testing"
)

func makeRandSlice(size int, min int64, max int64) []int64 {
	res := make([]int64, size)
	for i := 0; i < size; i++ {
		res = append(res, min+rand.Int63n(max)+1)
	}
	return res
}

func BenchmarkCalcDiscrepancy(b *testing.B) {
	b.StopTimer()
	arr := makeRandSlice(500000, 500, 5000000)
	b.StartTimer()
	for i := 0; i < 100; i++ {
		_, _ = calcDiscrepancy(rand.Float64(), &arr)
	}
}

func BenchmarkCalcDiscrepancy2(b *testing.B) {
	b.StopTimer()
	arr := makeRandSlice(500000, 500, 5000000)
	b.StartTimer()
	for i := 0; i < 10000; i++ {
		_, _ = calcDiscrepancy2(rand.Float64(), &arr, 1.22797832233232445879793222)
	}
}

func BenchmarkCalcDiscrepancy3(b *testing.B) {
	b.StopTimer()
	arr := makeRandSlice(5000000, 500, 5000000)
	b.StartTimer()
	sum := int64(0)
	for _, v := range arr {
		sum += v
	}
	for i := 0; i < 100000; i++ {
		_, _ = calcDiscrepancy3(rand.Float64(), 1.22797832233232445879793222, len(arr), sum)
	}
}

func BenchmarkCalcArrDiscrepancy(b *testing.B) {
	b.StopTimer()
	arr := makeRandSlice(500000, 500, 5000000)
	b.StartTimer()
	for i := 0; i < 100; i++ {
		_ = calcArrDiscrepancy(&arr)
	}
}
