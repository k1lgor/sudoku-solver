package main

import "fmt"

func solveSudoku(board [][]int) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == 0 {
				for k := 1; k <= 9; k++ {
					if isValid(board, i, j, k) {
						board[i][j] = k
						if solveSudoku(board) {
							return true
						}
						board[i][j] = 0
					}
				}
				return false
			}
		}
	}
	return true
}

func isValid(board [][]int, row, col, num int) bool {
	for i := 0; i < 9; i++ {
		if board[i][col] == num ||
			board[row][i] == num ||
			board[3*(row/3)+i/3][3*(col/3)+i%3] == num {
			return false
		}
	}
	return true
}

func main() {
	board := [][]int{
		{3, 1, 0, 0, 0, 0, 0, 2, 0},
		{0, 0, 6, 1, 0, 9, 0, 0, 5},
		{0, 0, 0, 0, 8, 0, 0, 0, 0},
		{0, 2, 0, 8, 0, 4, 0, 5, 0},
		{0, 0, 4, 0, 7, 0, 0, 0, 0},
		{0, 0, 0, 0, 6, 0, 0, 0, 8},
		{0, 6, 0, 0, 0, 0, 9, 0, 0},
		{0, 0, 9, 4, 0, 5, 0, 0, 1},
		{0, 0, 0, 0, 0, 7, 0, 0, 0},
	}

	fmt.Println("Sudoku puzzle:")
	for _, row := range board {
		fmt.Println(row)
	}
	fmt.Println()

	solveSudoku(board)

	fmt.Println("Sudoku solution:")
	for _, row := range board {
		fmt.Println(row)
	}
}
