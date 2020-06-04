package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

//Cell object
type Cell struct {
	value     int
	column    int
	row       int
	isVisible bool // true: se ve el número y false: no se muestra el núm
}

func test(pepitosupercoolchico string) string {
	return "hola " + pepitosupercoolchico
}

func board() [][]int {
	board := [][]int{
		{2, 1, 3, 4, 5, 6, 7, 8, 9},
		{4, 5, 6, 7, 8, 9, 1, 2, 3},
		{7, 8, 9, 1, 2, 3, 4, 5, 6},
		{2, 1, 4, 3, 6, 5, 8, 9, 7},
		{3, 6, 5, 8, 9, 7, 2, 1, 4},
		{8, 9, 7, 2, 1, 4, 3, 6, 5},
		{5, 3, 1, 6, 4, 2, 9, 7, 8},
		{6, 4, 2, 9, 7, 8, 5, 3, 1},
		{9, 7, 8, 5, 3, 1, 6, 4, 2}}
	return board
}

func main() {

	js := mewn.String("./frontend/dist/app.js")
	css := mewn.String("./frontend/dist/app.css")

	app := wails.CreateApp(&wails.AppConfig{
		Width:  1024,
		Height: 768,
		Title:  "sudoku",
		JS:     js,
		CSS:    css,
		Colour: "#131313",
	})
	app.Bind(board)
	app.Bind(test)
	app.Run()
}

// [][]int{
// 	{2, 1, 3, 4, 5, 6, 7, 8, 9},
// 	{4, 5, 6, 7, 8, 9, 1, 2, 3},
// 	{7, 8, 9, 1, 2, 3, 4, 5, 6},
// 	{2, 1, 4, 3, 6, 5, 8, 9, 7},
// 	{3, 6, 5, 8, 9, 7, 2, 1, 4},
// 	{8, 9, 7, 2, 1, 4, 3, 6, 5},
// 	{5, 3, 1, 6, 4, 2, 9, 7, 8},
// 	{6, 4, 2, 9, 7, 8, 5, 3, 1},
// 	{9, 7, 8, 5, 3, 1, 6, 4, 2}}

// func createboard() {
// 	boardWithobjects []
// 	for row = 1, row <= 9, row++ {
// 		createrow(row, value[row])
// 		push row to board
// 	}
// 	return board[rows[cells]]
// }

// func CreateRow(row int, value) {
// 	row []
// 	for col = 1, col <= 9, col++ {
// 		cell{value,row,col,true}
// 		push cell to row (row = append(row, Cell))
// 	}
// 	return row[cells]
// }

// CreateRowOfCells a row
func CreateRowOfCells(row []int, indexOfRow int) []Cell {
	var newRow []Cell

	for col := 0; col < 9; col++ {
		var newCell = Cell{row[col], col, indexOfRow, true}
		newRow = append(newRow, newCell)
	}
	return newRow
}
