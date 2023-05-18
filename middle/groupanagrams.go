package middle

type Elem [26]int // 采用一个26为的int数组作为判断字符串是否包含相同字母的依据

// groupAnagrams 字母异位词分组
func groupAnagrams(items []string) [][]string {
	res := make([][]string, 0, len(items))

	// 使用map对数组归类
	temp := make(map[Elem][]string)
	for _, item := range items {
		elem := newElem(item)
		// 根据Elem对字符串归为异位词
		temp[elem] = append(temp[elem], item)
	}

	for _, elems := range temp {
		res = append(res, elems)
	}
	return res
}

// newElem 将字符串转换为目标数组
func newElem(str string) (elem Elem) {
	for _, r := range str {
		elem[r-'a']++
	}
	return
}
