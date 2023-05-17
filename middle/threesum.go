package middle

import "sort"

// threeSum 三数之和
func threeSum(nums []int) [][]int {
	items := make([][]int, 0, 64)
	n := len(nums)
	sort.Ints(nums)
	// a+b+c = 0
	for first := 0; first < n; first++ {
		// 需要和上一次枚举的数不相同，避免计算重复数据，提高效率
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := n - 1
		target := -1 * nums[first]
		// 枚举 b
		for second := first + 1; second < n; second++ {
			// 需要和上一次枚举的数不相同，避免计算重复数据，提高效率
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			// 需要保证 b 的指针在 c 的指针的左侧
			for second < third && nums[second]+nums[third] > target {
				third--
			}
			// 如果指针重合，随着 b 后续的增加
			// 就不会有满足 a+b+c=0 并且 b<c 的 c 了，可以退出循环
			if second == third {
				break
			}
			if nums[second]+nums[third] == target {
				items = append(items, []int{nums[first], nums[second], nums[third]})
			}

		}
	}
	return items
}
