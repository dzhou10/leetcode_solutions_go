package main

import (
	"fmt"
	"testing"
)

func TestN054SpiralMatrix(t *testing.T) {
	var n054 N054SpiralMatrix
	matrix1 := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	matrix2 := [][]int{[]int{2, 3}}
	matrix3 := [][]int{[]int{7}, []int{9}, []int{6}}
	fmt.Print("N054SpiralMatrix:\t", n054.spiralOrder(matrix1), n054.spiralOrder(matrix2), n054.spiralOrder(matrix3), "\n")
}
