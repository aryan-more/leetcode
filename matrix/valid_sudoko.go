package main

import (
	"math"
)
// https://leetcode.com/problems/valid-sudoku/
func validSudoko(board [][]byte) bool {

	ignore := '.'
	horizontalCache := map[int][]byte{}
	verticalCache := map[int][]byte{}
	zonalCache := map[int][]byte{}

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			box := board[i][j]
			if box != byte(ignore) {
				horizontal := horizontalCache[i]
				for _, exist := range horizontal {
					if exist == box {
						return false
					}
				}
				horizontalCache[i] = append(horizontalCache[i], box)

				vertical := verticalCache[j]
				for _, exist := range vertical {
					if exist == box {
						return false
					}
				}
				verticalCache[j] = append(verticalCache[j], box)

				zone := int(math.Floor(float64(i/3))*3 + math.Floor(float64(j/3)))
				zonal := zonalCache[zone]
				for _, exist := range zonal {
					if exist == box {
						return false
					}
				}
				zonalCache[zone] = append(zonalCache[zone], box)

			}
		}
	}

	return true
}

// func main() {
// 	invalidBoard := [][]byte{{'8', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

// 	validBoard := [][]byte{{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
// 		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
// 		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
// 		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
// 		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
// 		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
// 		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
// 		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
// 		{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}

// 	zonalInvalid := [][]byte{{'.', '.', '.', '.', '5', '.', '.', '1', '.'},
// 		{'.', '4', '.', '3', '.', '.', '.', '.', '.'},
// 		{'.', '.', '.', '.', '.', '3', '.', '.', '1'},
// 		{'8', '.', '.', '.', '.', '.', '.', '2', '.'},
// 		{'.', '.', '2', '.', '7', '.', '.', '.', '.'},
// 		{'.', '1', '5', '.', '.', '.', '.', '.', '.'},
// 		{'.', '.', '.', '.', '.', '2', '.', '.', '.'},
// 		{'.', '2', '.', '9', '.', '.', '.', '.', '.'},
// 		{'.', '.', '4', '.', '.', '.', '.', '.', '.'}}

// 	fmt.Println(validSudoko(invalidBoard))
// 	fmt.Println(validSudoko(validBoard))
// 	fmt.Println(validSudoko(zonalInvalid))
// }
