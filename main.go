package main

import "os"

import "fmt"

func main() {
	var board [9][9]string
	var row [82]int
	var col [82]int
	cell := 0
	for i, word := range os.Args {
		if i > 0 {
			for j, letter := range word {
				board[i-1][j] = string(letter)
				if string(letter) == "." {
					row[cell] = i - 1
					col[cell] = j
					cell++
				}
			}
		}
	}
	row[cell] = 100
	col[cell] = 100
	if Sudoku(row, col, board, 0) == board {
		fmt.Println("Error")
	} else {
		fmt.Println(Sudoku(row, col, board, 0))
	}

}

func Valid(array [9][9]string) bool {
	for row := 0; row < 9; row++ {
		for col := 0; col < 9; col++ {
			for i := 0; i < 9; i++ {
				if i != col && array[row][col] != "." {
					if array[row][col] == array[row][i] {
						return false
					}
				}
			}
		}
	}
	for col := 0; col < 9; col++ {
		for row := 0; row < 9; row++ {
			for i := 0; i < 9; i++ {
				if i != row && array[row][col] != "." {
					if array[row][col] == array[i][col] {
						return false
					}
				}
			}
		}
	}
	for i := 0; i < 9; i += 3 {
		for j := 0; j < 9; j += 3 {
			for row := i; row < i+3; row++ {
				for col := j; col < j+3; col++ {
					for index := i; index < i+3; index++ {
						for jindex := j; jindex < j+3; jindex++ {
							if (row == index && col == jindex) || array[row][col] == "." {

							} else {
								if array[row][col] == array[index][jindex] {
									return false
								}
							}

						}
					}
				}
			}
		}
	}

	return true
}

func Sudoku(row [82]int, col [82]int, array [9][9]string, i int) [9][9]string {
	if i < 0 {
		return array
	}
	if row[i] > 10 {
		return array
	}

	array[row[i]][col[i]] = string(Atoi(array[row[i]][col[i]]) + 49)
	if array[row[i]][col[i]] > "9" {
		array[row[i]][col[i]] = "."
		return Sudoku(row, col, array, i-1)
	}

	if Valid(array) {
		return Sudoku(row, col, array, i+1)
	} else {
		return Sudoku(row, col, array, i)
	}
}

func Atoi(s string) int {
	var i int
	var isnegative bool = false
	for _, letter := range s {
		if letter == '-' {
			if isnegative || i != 0 {
				return 0
			} else {
				isnegative = true
			}
		} else if letter == '1' {
			i = (i * 10) + 1
		} else if letter == '2' {
			i = (i * 10) + 2
		} else if letter == '3' {
			i = (i * 10) + 3
		} else if letter == '4' {
			i = (i * 10) + 4
		} else if letter == '5' {
			i = (i * 10) + 5
		} else if letter == '6' {
			i = (i * 10) + 6
		} else if letter == '7' {
			i = (i * 10) + 7
		} else if letter == '8' {
			i = (i * 10) + 8
		} else if letter == '9' {
			i = (i * 10) + 9
		} else if letter == '0' {
			i = i * 10
		} else if letter == ' ' {
			return 0
		} else {
			return 0
		}
	}
	if isnegative {
		i *= -1
	}
	return i
}
