package tws

import (
	"errors"
	"math"
)

/**
 * [性能存在问题 应该使用 @see calcDiscrepancy3]
 * 计算数字与数组离散程度，返回值越接近0，说明num在arr里越接近中心
 */
func calcDiscrepancy(num float64, arr *[]int64) (float64, error) {
	if len(*arr) == 0 {
		return 0, errors.New("数组要求必须有值")
	}

	// 计算平均值
	sum := int64(0)
	for _, v := range *arr {
		sum += v
	}
	mean := float64(sum) / float64(len(*arr))

	discrepancy := calcArrDiscrepancy(arr)

	// 计算num与平均值之差的标准差
	diffFromMean := math.Abs(num - mean)
	if discrepancy == 0 {
		return 0, nil
	}
	return diffFromMean / discrepancy, nil
}

/**
 * [性能存在问题 应该使用 @see calcDiscrepancy3]
 * 等同于 calcDiscrepancy，节省discrepancy计算成本
 */
func calcDiscrepancy2(num float64, arr *[]int64, discrepancy float64) (float64, error) {
	if len(*arr) == 0 {
		return 0, errors.New("数组要求必须有值")
	}

	// 计算平均值
	sum := int64(0)
	for _, v := range *arr {
		sum += v
	}
	mean := float64(sum) / float64(len(*arr))

	// 计算num与平均值之差的标准差
	diffFromMean := math.Abs(num - mean)
	if discrepancy == 0 {
		return 0, nil
	}
	return diffFromMean / discrepancy, nil
}

/**
 * 等同于 calcDiscrepancy，节省discrepancy、arr计算成本
 */
func calcDiscrepancy3(num float64, discrepancy float64, arrLen int, arrSum int64) (float64, error) {
	if arrLen == 0 {
		return 0, errors.New("数组要求必须有值")
	}

	// 计算平均值
	mean := float64(arrSum) / float64(arrLen)

	// 计算num与平均值之差的标准差
	diffFromMean := math.Abs(num - mean)
	if discrepancy == 0 {
		return 0, nil
	}
	return diffFromMean / discrepancy, nil
}

/**
 * [性能存在问题 不要嵌入于大数据遍历]
 * 计算数组离散程度
 */
func calcArrDiscrepancy(arr *[]int64) float64 {
	// 计算平均值
	var sum int64 = 0
	for _, v := range *arr {
		sum += v
	}
	mean := float64(sum) / float64(len(*arr))

	// 计算所有元素与平均值之差的平方和
	var varianceSum float64
	for _, v := range *arr {
		varianceSum += math.Pow(float64(v)-mean, 2)
	}

	// 计算方差并进一步得到标准差（离散程度）
	variance := varianceSum / float64(len(*arr)-1)
	return math.Sqrt(variance)
}
