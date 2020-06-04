package main

import (
	"github.com/leaanthony/mewn"
	"github.com/wailsapp/wails"
)

//Cell object
type Cell struct {
	value     int  //row[col]
	column    int  //col
	row       int  //indexOfRow
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

// CreateRowOfCells a row
func CreateRowOfCells(row [9]int, indexOfRow int) []Cell {
	var newRow []Cell

	for col := 0; col < 9; col++ {
		var newCell = Cell{row[col], col, indexOfRow, true}
		newRow = append(newRow, newCell)
	}
	return newRow
}

//CreateBoard a board
func CreateBoard(basicBoard [9][9]int) [][]Cell {
	var newBoard [][]Cell
	var indexOfRow int

	for _, row := range basicBoard {
		var newRow = CreateRowOfCells(row, indexOfRow)
		indexOfRow++
		newBoard = append(newBoard, newRow)
	}

	return newBoard
}
