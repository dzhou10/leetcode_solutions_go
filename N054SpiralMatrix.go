package main

import (
	"math"
)

type N054SpiralMatrix struct {
}

func (this *N054SpiralMatrix) spiralOrder(matrix [][]int) []int {
	rows := len(matrix)
	cols := 0
	if rows > 0 {
		cols = len(matrix[0])
	}
	result := make([]int, 0, rows*cols)
	maxRound := (int)(math.Min(float64(rows), float64(cols))/2.0 + 0.5)
	if rows == 0 && cols == 0 {
		maxRound = 0
	}
	row := 0
	col := 0
	for round := 0; round < maxRound; round++ {
		// Move right
		for row, col = round, round; col <= cols-round-1; col++ {
			result = append(result, matrix[row][col])
		}
		// Move down
		for col, row = cols-round-1, round+1; row <= rows-round-1; row++ {
			result = append(result, matrix[row][col])
		}
		// Move left
		for row, col = rows-round-1, cols-round-2; col > round && row > round; col-- {
			result = append(result, matrix[row][col])
		}
		// Move up
		for col, row = round, rows-round-1; row > round && col < cols-round-1; row-- {
			result = append(result, matrix[row][col])
		}
	}
	return result
}
