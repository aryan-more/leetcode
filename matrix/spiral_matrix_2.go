package main
// https://leetcode.com/problems/spiral-matrix-ii/
func GenerateMatrix(n int) {
	matrix := make([][]int, n)

	for row := 0; row < n; row++ {
		matrix[row] = make([]int, n)
	}
	up, left := 0, 0
	down, right := n-1, n-1

	counter := 0

	for counter < n*n {
		for index := left; index <= right; index++ {
			counter++
			matrix[up][index] = counter
		}
		up++
		for index := up; index <= down; index++ {
			counter++
			matrix[index][right] = counter
		}
		right--
		for index := right; index >= left; index-- {
			counter++
			matrix[down][index] = counter
		}
		down--
		for index := down; index >= up; index-- {
			counter++
			matrix[index][left] = counter
		}
		left++
	}

}

// func main() {
// 	generateMatrix(5)
// }
