package middle

// lengthOfLongestSubstring 返回无重复字符的最长子串
// 给定一个字符串 s ，请你找出其中不含有重复字符的 最长子串的长度。
func lengthOfLongestSubstring(s string) int {
	max := 0
	temp := map[rune]int{}
	// 使用双置针，start记录字符串起始位置,end记录字符串结束位置
	var start, end int
	for end < len(s) {
		r := rune(s[end])
		// 检查字符串是否已存在
		if _, ok := temp[r]; ok {
			// 更新起始位置
			if temp[r]+1 > start {
				start = temp[r] + 1
			}
		}

		temp[r] = end
		// 计算字符串长度
		if length := end - start + 1; length > max {
			max = length
		}
		end++

	}

	return max
}
