package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// func TestName(t *testing.T) {
// 	type args struct {
// 	}
// 	tests := []struct {
// 		name     string
// 		args     args
// 		expected bool
// 	}{
// 		{"Test name", args{}},
// 	}
// 	for _, test := range tests {
// 		t.Run(test.name, func(t *testing.T) {
// 			assert.Equal(t, test.expected)
// 		})
// 	}
// }

func TestCreateRowOfCells(t *testing.T) {
	type args struct {
		row        [9]int
		indexOfRow int
	}
	tests := []struct {
		name     string
		args     args
		expected []Cell
	}{
		{"Devuelve un fila con 9 cells", args{[9]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, 0}, []Cell{
			Cell{0, 0, 0, true},
			Cell{1, 1, 0, true},
			Cell{2, 2, 0, true},
			Cell{3, 3, 0, true},
			Cell{4, 4, 0, true},
			Cell{5, 5, 0, true},
			Cell{6, 6, 0, true},
			Cell{7, 7, 0, true},
			Cell{8, 8, 0, true}}},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, CreateRowOfCells(test.args.row, test.args.indexOfRow))
		})
	}
}

func TestCreateBoard(t *testing.T) {
	type args struct {
		basicBoard [9][9]int
	}
	tests := []struct {
		name     string
		args     args
		expected [][]Cell
	}{
		{
			"Devuelve un board de 9 fillas de cells",
			args{[9][9]int{
				{1, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 2, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 3, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0},
				{0, 0, 0, 0, 0, 0, 0, 0, 0}}},
			[][]Cell{
				{Cell{0, 0, 0, true},
					Cell{0, 1, 0, true},
					Cell{0, 2, 0, true},
					Cell{0, 3, 0, true},
					Cell{0, 4, 0, true},
					Cell{0, 5, 0, true},
					Cell{0, 6, 0, true},
					Cell{0, 7, 0, true},
					Cell{0, 8, 0, true}},
				{Cell{0, 0, 1, true},
					Cell{0, 1, 1, true},
					Cell{0, 2, 1, true},
					Cell{0, 3, 1, true},
					Cell{0, 4, 1, true},
					Cell{0, 5, 1, true},
					Cell{0, 6, 1, true},
					Cell{0, 7, 1, true},
					Cell{0, 8, 1, true}},
				{Cell{0, 0, 2, true},
					Cell{0, 1, 2, true},
					Cell{0, 2, 2, true},
					Cell{0, 3, 2, true},
					Cell{0, 4, 2, true},
					Cell{0, 5, 2, true},
					Cell{0, 6, 2, true},
					Cell{0, 7, 2, true},
					Cell{0, 8, 2, true}},
				{Cell{0, 0, 3, true},
					Cell{0, 1, 3, true},
					Cell{0, 2, 3, true},
					Cell{0, 3, 3, true},
					Cell{0, 4, 3, true},
					Cell{0, 5, 3, true},
					Cell{0, 6, 3, true},
					Cell{0, 7, 3, true},
					Cell{0, 8, 3, true}},
				{Cell{0, 0, 4, true},
					Cell{0, 1, 4, true},
					Cell{0, 2, 4, true},
					Cell{0, 3, 4, true},
					Cell{0, 4, 4, true},
					Cell{0, 5, 4, true},
					Cell{0, 6, 4, true},
					Cell{0, 7, 4, true},
					Cell{0, 8, 4, true}},
				{Cell{0, 0, 5, true},
					Cell{0, 1, 5, true},
					Cell{0, 2, 5, true},
					Cell{0, 3, 5, true},
					Cell{0, 4, 5, true},
					Cell{0, 5, 5, true},
					Cell{0, 6, 5, true},
					Cell{0, 7, 5, true},
					Cell{0, 8, 5, true}},
				{Cell{0, 0, 6, true},
					Cell{0, 1, 6, true},
					Cell{0, 2, 6, true},
					Cell{0, 3, 6, true},
					Cell{0, 4, 6, true},
					Cell{0, 5, 6, true},
					Cell{0, 6, 6, true},
					Cell{0, 7, 6, true},
					Cell{0, 8, 6, true}},
				{Cell{0, 0, 7, true},
					Cell{0, 1, 7, true},
					Cell{0, 2, 7, true},
					Cell{0, 3, 7, true},
					Cell{0, 4, 7, true},
					Cell{0, 5, 7, true},
					Cell{0, 6, 7, true},
					Cell{0, 7, 7, true},
					Cell{0, 8, 7, true}},
				{Cell{0, 0, 8, true},
					Cell{0, 1, 8, true},
					Cell{0, 2, 8, true},
					Cell{0, 3, 8, true},
					Cell{0, 4, 8, true},
					Cell{0, 5, 8, true},
					Cell{0, 6, 8, true},
					Cell{0, 7, 8, true},
					Cell{0, 8, 8, true}}},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			assert.Equal(t, test.expected, CreateBoard(test.args.basicBoard))
		})
	}
}
