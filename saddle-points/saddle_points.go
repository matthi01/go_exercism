package matrix

import (
	"fmt"
	"strconv"
	"strings"
)

// Matrix is the entire matrix
// type Matrix []Rows

// Pair specifies a location in the matrix
type Pair struct {
	col, row int
}

type Column []int
type Row []int
type Matrix []Row

// New also still requires a comment
func New(s string) (*Matrix, error) {
	strRows := strings.Split(s, "\n")
	// fmt.Println("original:", s)
	// fmt.Println("rows", strRows)
	// var m Matrix
	var r Matrix
	for _, tr := range strRows {
		var c Column
		for _, tc := range tr {
			n, err := strconv.Atoi(string(tc))
			if err != nil {
				continue
			}
			c = append(c, n)
		}
		// fmt.Println("c:", c)
		r = append(r, c)
	}
	// m = rows
	// fmt.Println(r)
	return &r, nil
}

// Saddle still needs a comment...
func (m *Matrix) Saddle() []Pair {
	fmt.Println("matrix:", *m)
	var saddlePairs []Pair

	for ri, row := range *m {
		fmt.Println("row:", row)
		fmt.Println("ri:", ri)
		for ci, col := range row {
			fmt.Println("col:", col)
			fmt.Println("ci:", ci)
		}
	}

	return saddlePairs
}

// Sum returns the sum of any given number of integers
func Sum(ns ...int) int {
	var sum int
	for _, n := range ns {
		sum += n
	}
	return sum
}

// RowSum needs a comment
func RowSum(r Row) int {
	return 0
}

// ColumnSum needs a comment
func ColumnSum() int {
	return 0
}
