package middle

// setZeroes 矩阵置零
// 给定一个 m x n 的矩阵，如果一个元素为 0 ，则将其所在行和列的所有元素都设为 0 。请使用 原地 算法。
func setZeroes(matrix [][]int) {
	// 记录需要置零的列
	cols := make(map[int]struct{})
	for i := 0; i < len(matrix); i++ {
		rowZero := false
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] == 0 {
				rowZero = true
				cols[j] = struct{}{}
			}
		}
		// 当前行置零
		if rowZero {
			matrix[i] = make([]int, len(matrix[i]))
		}

	}

	for j := range cols {
		for k := 0; k < len(matrix); k++ {
			matrix[k][j] = 0
		}
	}
}
