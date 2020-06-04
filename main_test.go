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
		row []int
	}
	tests := []struct {
		name     string
		args     args
		expected []Cell
	}{
		{"Devuelve un array con 9 int", args{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}}, []Cell{
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
			assert.Equal(t, test.expected, CreateRowOfCells(test.args.row))
		})
	}
}
