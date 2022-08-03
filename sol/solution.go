package sol

func rotate(matrix [][]int) {
	n := len(matrix)
	small := n / 2
	large := small
	if n%2 != 0 {
		large = small + 1
	}
	for i := 0; i < large; i++ {
		for j := 0; j < small; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-1-j][i]
			matrix[n-1-j][i] = matrix[n-1-i][n-1-j]
			matrix[n-1-i][n-1-j] = matrix[j][n-1-i]
			matrix[j][n-1-i] = temp
		}
	}
}
