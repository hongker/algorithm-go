package middle

// longestPalindrome 最长回文子串
// TODO 使用动态规划
func longestPalindrome(s string) string {
	var sub string
	if isPalindrome(s) {
		return s
	}

	for i := 0; i < len(s)-1; i++ {
		for j := i + 1 + len(sub); j <= len(s); j++ {
			if isPalindrome(s[i:j]) && len(s[i:j]) > len(sub) {
				sub = s[i:j]
			}
		}
	}

	return sub
}

func isPalindrome(s string) bool {
	start := 0
	end := len(s) - 1
	for start < end {
		if s[start] != s[end] {
			return false
		}
		start++
		end--
	}

	return true
}
