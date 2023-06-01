package middle

import "math"

// increasingTriplet 递增的三元子序列
func increasingTriplet(nums []int) bool {
	small := math.MaxInt
	mid := math.MaxInt
	for _, num := range nums {
		if num <= small {
			// 记录遍历过的最小值
			small = num
		} else if num <= mid {
			// 记录比small大的值
			mid = num
		} else {
			// 满足条件small < mid < num
			return true
		}
	}
	return false
}
