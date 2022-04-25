package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix [][]int

func New(s string) (Matrix, error) {
	rows := strings.Split(s, "\n")
	var matrixLength = 0
	var matrixHeight = len(rows)
	var newMatrix = make(Matrix, matrixHeight)

	for rowIndex, row := range rows {
		trimmedRow := strings.TrimSpace(row)
		columnCells := strings.Split(trimmedRow, " ")
		if matrixLength == 0 {
			matrixLength = len(columnCells)
		}
		if matrixLength != len(columnCells) {
			return Matrix{}, errors.New("uneven column lengths")
		}

		newMatrix[rowIndex] = make([]int, matrixLength)
		for columnIndex, strCellValue := range columnCells {
			cellValue, err := strconv.Atoi(strCellValue)
			if err != nil {
				return Matrix{}, err
			}
			newMatrix[rowIndex][columnIndex] = cellValue
		}
	}
	return newMatrix, nil
}

func (m *Matrix) Set(row, column, value int) bool {
	if row < 0 || column < 0 {
		return false
	}
	nRows := len(*m)
	nColumns := len((*m)[0])
	if row > nRows-1 || column > nColumns-1 {
		return false
	}
	(*m)[row][column] = value
	return true
}

func (m Matrix) Cols() [][]int {
	nColumns := len(m[0])
	nRows := len(m)
	mCopy := make(Matrix, nColumns)
	for i := range mCopy {
		mCopy[i] = make([]int, nRows)
	}
	for ri, row := range mCopy {
		for ci := range row {
			mCopy[ri][ci] = m[ci][ri]
		}
	}
	return mCopy
}

func (m Matrix) Rows() [][]int {
	mCopy := make(Matrix, len(m))
	for ri, columns := range m {
		mCopy[ri] = make([]int, len(columns))
		copy(mCopy[ri], m[ri])
	}
	return mCopy
}
