package piscine

import "github.com/01-edu/z01"

var solutions [][]rune

func isSafe(board [][]int, row int, col int) bool {
	for i := 0; i < row; i++ {
		if board[i][col] == 1 {
			return false
		}
	}

	for i, j := row-1, col-1; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if board[i][j] == 1 {
			return false
		}
	}

	for i, j := row-1, col+1; i >= 0 && j < 8; i, j = i-1, j+1 {
		if board[i][j] == 1 {
			return false
		}
	}

	return true
}

func Backtracking(board [][]int, row int) {
	if row == 8 {
		tmp := make([]rune, 8)
		for i := range board {
			for j := range board {
				if board[i][j] == 1 {
					tmp[i] = intToChar(j + 1)
				}
			}
		}
		solutions = append(solutions, tmp)
		return
	}
	for col := 0; col < 8; col++ {
		if isSafe(board, row, col) {
			board[row][col] = 1
			Backtracking(board, row+1)
			board[row][col] = 0
		}
	}
}

func EightQueens() {
	solutions = make([][]rune, 0)
	board := make([][]int, 8)
	for i := 0; i < 8; i++ {
		board[i] = make([]int, 8)
		for j := 0; j < 8; j++ {
			board[i][j] = 0
		}
	}
	Backtracking(board, 0)
	for i := 0; i < len(solutions); i++ {
		for j := 0; j < 8; j++ {
			z01.PrintRune(solutions[i][j])
		}
		z01.PrintRune('\n')
	}
}

// ION MI DIAVAZEIS TA COMMENTS MOU
