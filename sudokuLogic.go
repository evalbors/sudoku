package main

import (
	"math/rand"
	"time"
)

// INPUT CHECKER
func makeInput(board [9][9]int, fila int, columna int, number int) [9][9]int {
	if number > 0 && number <= 9 {
		board[fila][columna] = number
		return board
	}
	return board
}

func numberInRow(row [9]int, number int) bool {
	for i := 0; i < 9; i++ {
		if number == row[i] {
			return false
		}
	}
	return true
}

func transformColumnToArray(board [9][9]int, numberCol int) [9]int {
	column := [9]int{}
	for x := 0; x < 9; x++ {
		column[x] = board[x][numberCol]
	}
	return column
}

func numberInColumn(col [9]int, number int) bool {
	for i := 0; i < 9; i++ {
		if number == col[i] {
			return false
		}
	}
	return true
}

func quadrantDefiner(board [9][9]int, quadrantNumber int) [3][3]int {
	quadrant := [3][3]int{}
	if quadrantNumber == 1 {
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				quadrant[x][y] = board[x][y]
			}
		}
	}
	if quadrantNumber == 2 {
		for x := 0; x < 3; x++ {
			contador := 0
			for y := 3; y < 6; y++ {
				quadrant[x][contador] = board[x][y]
				contador++
			}
		}
	}
	if quadrantNumber == 3 {
		for x := 0; x < 3; x++ {
			contador := 0
			for y := 6; y < 9; y++ {
				quadrant[x][contador] = board[x][y]
				contador++
			}
		}
	}
	if quadrantNumber == 4 {
		contadorX := 0
		for x := 3; x < 6; x++ {
			contadorY := 0
			for y := 0; y < 3; y++ {
				quadrant[contadorX][contadorY] = board[x][y]
				contadorY++
			}
			contadorX++
		}
	}
	if quadrantNumber == 5 {
		contadorX := 0
		for x := 3; x < 6; x++ {
			contadorY := 0
			for y := 3; y < 6; y++ {
				quadrant[contadorX][contadorY] = board[x][y]
				contadorY++
			}
			contadorX++
		}
	}
	if quadrantNumber == 6 {
		contadorX := 0
		for x := 3; x < 6; x++ {
			contadorY := 0
			for y := 6; y < 9; y++ {
				quadrant[contadorX][contadorY] = board[x][y]
				contadorY++
			}
			contadorX++
		}
	}
	if quadrantNumber == 7 {
		contadorX := 0
		for x := 6; x < 9; x++ {
			contadorY := 0
			for y := 0; y < 3; y++ {
				quadrant[contadorX][contadorY] = board[x][y]
				contadorY++
			}
			contadorX++
		}
	}
	if quadrantNumber == 8 {
		contadorX := 0
		for x := 6; x < 9; x++ {
			contadorY := 0
			for y := 3; y < 6; y++ {
				quadrant[contadorX][contadorY] = board[x][y]
				contadorY++
			}
			contadorX++
		}
	}
	if quadrantNumber == 9 {
		contadorX := 0
		for x := 6; x < 9; x++ {
			contadorY := 0
			for y := 6; y < 9; y++ {
				quadrant[contadorX][contadorY] = board[x][y]
				contadorY++
			}
			contadorX++
		}
	}
	return quadrant
}

func numberInQuadrant(quadrant [3][3]int, number int) bool {
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			if number == quadrant[x][y] {
				return false
			}

		}
	}
	return true
}

func verifyAllInputBooleans(row bool, col bool, quadrant bool) bool {
	if row == false {
		return false
	}
	if col == false {
		return false
	}
	if quadrant == false {
		return false
	}
	return true
}

//SUDOKU CHECKER
func verifyAllRows(board [9][9]int) bool {
	for x := 0; x < 9; x++ {
		row := board[x]
		contador := 0
		for y := 0; y < 9; y++ {
			contador = contador + row[y]
		}
		if contador > 45 {
			return false
		}
		if contador < 45 {
			return false
		}
	}
	return true
}

func verifyAllColumns(board [9][9]int) bool {

	for x := 0; x < 9; x++ {
		columna := transformColumnToArray(board, x)
		contador := 0
		for y := 0; y < 9; y++ {
			contador = contador + columna[y]
		}
		if contador > 45 {
			return false
		}
		if contador < 45 {
			return false
		}
	}
	return true
}

func verifyAllQuadrants(board [9][9]int) bool {
	for q := 1; q < 10; q++ {
		quadrant := quadrantDefiner(board, q)
		contador := 0
		for x := 0; x < 3; x++ {
			for y := 0; y < 3; y++ {
				contador = contador + quadrant[x][y]
			}
		}
		if contador != 45 {
			return false
		}
	}
	return true
}

func verifyMatch(board [9][9]int) bool {
	if verifyAllColumns(board) == verifyAllRows(board) == verifyAllQuadrants(board) == true {
		return true
	}
	return false
}

//SUDOKU MAKER
func randomInteger(min, max int) int {
	return min + rand.Intn(max-min)
}

func randomInput() int {
	rand.Seed(time.Now().UnixNano())
	randomNumber := randomInteger(1, 10)
	return randomNumber
}

func quadrantPicker(x int, y int) int {
	var quadrantNumber int
	if x < 3 && y < 3 {
		quadrantNumber = 1
	}
	if x < 3 && y >= 3 && y < 6 {
		quadrantNumber = 2
	}
	if x < 3 && y >= 6 && y < 9 {
		quadrantNumber = 3
	}
	if x >= 3 && x < 6 && y < 3 {
		quadrantNumber = 4
	}
	if x >= 3 && x < 6 && y >= 3 && y < 6 {
		quadrantNumber = 5
	}
	if x >= 3 && x < 6 && y >= 6 && y < 9 {
		quadrantNumber = 6
	}
	if x >= 6 && x < 9 && y < 3 {
		quadrantNumber = 7
	}
	if x >= 6 && x < 9 && y >= 3 && y < 6 {
		quadrantNumber = 8
	}
	if x >= 6 && x < 9 && y >= 6 && y < 9 {
		quadrantNumber = 9
	}
	return quadrantNumber
}

func sudokuMaker(board [9][9]int) [9][9]int {
	sudoku := board
	for x := 0; x < 9; x++ {
		for y := 0; y < 9; y++ {
			contador := 0
			contadorUltra := 0
			for {
				if contador == 2000 {
					sudoku[x] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
					contador = 0
					y = 0
				}
				if contadorUltra == 4000 {
					sudoku[x] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
					sudoku[x-1] = [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}
					contadorUltra = 0
					contador = 0
					y = 0
					x = x - 1
				}
				number := randomInput()
				row := numberInRow(sudoku[x], number)
				col := numberInColumn(transformColumnToArray(sudoku, y), number)
				quadrantNumber := quadrantPicker(x, y)
				quadrant := numberInQuadrant(quadrantDefiner(sudoku, quadrantNumber), number)
				superTrue := verifyAllInputBooleans(row, col, quadrant)
				if superTrue == true {
					sudoku = makeInput(sudoku, x, y, number)
					break
				}
				if superTrue == false {
					contadorUltra++
					contador++
				}
			}
		}
	}
	//hidden := hideNumbers(sudoku)
	return sudoku
}

func hideNumbers(sudoku [9][9]int) [9][9]int {
	for i := 0; i < 70; i++ {
		x := randomInteger(0, 9)
		y := randomInteger(0, 9)
		sudoku[x][y] = 0
	}
	return sudoku
}
