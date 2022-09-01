package rotate

func rotate(matrix [][]int) {
	length := len(matrix)
	for i := 0; i < length/2; i++ {
		for j := 0; j < length-i*2-1; j++ {
			t := matrix[i][i+j]
			matrix[i][i+j] = matrix[length-1-i-j][i]
			matrix[length-1-i-j][i] = matrix[length-1-i][length-1-i-j]
			matrix[length-1-i][length-1-i-j] = matrix[i+j][length-1-i]
			matrix[i+j][length-1-i] = t
		}
	}
}
