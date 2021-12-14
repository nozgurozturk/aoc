package main

import (
	"fmt"
	"github.com/nozgurozturk/aoc/util/gopher/io"
	"github.com/nozgurozturk/aoc/util/gopher/math"
	"strings"
)

const FileName = "../INPUT"

func main() {
	findLast := false

	drawnNumbersInput := io.ReadStrings(FileName, 0, 1)
	drawNumbersSplit := strings.Split(drawnNumbersInput[0], ",")
	drawnNumbers := make([]int, len(drawNumbersSplit), len(drawNumbersSplit))
	for i, s := range drawNumbersSplit {
		drawnNumbers[i] = math.ParseInt(s)
	}

	matrices := io.ReadMatrices(FileName, 2, math.MaxUInt)
	score := GiantSquid(drawnNumbers, matrices, findLast)

	fmt.Println(score)
}

func transpose(matrix [][]int) [][]int {
	transposed := make([][]int, len(matrix))

	for i, row := range matrix {
		transposed[i] = make([]int, len(row))
		for j, _ := range matrix {
			transposed[i][j] = matrix[j][i]
		}
	}
	return transposed
}

func checkColumns(matrix [][]int) bool {

	for _, column := range transpose(matrix) {
		marked := 0
		for _, num := range column {
			if num == -1 {
				marked++
			}
		}
		if marked == len(column) {
			return true
		}
	}
	return false
}

func checkRows(matrix [][]int) bool {
	for _, row := range matrix {
		marked := 0
		for _, num := range row {
			if num == -1 {
				marked++
			}
		}
		if marked == len(row) {
			return true
		}
	}
	return false
}

func sumOfMatrix(matrix [][]int) int {
	sum := 0
	for _, row := range matrix {
		for _, num := range row {
			if num != -1 {
				sum += num
			}
		}
	}

	return sum
}

func findWinnerScore(drawnNumbers []int, matrices [][][]int) int {
	for _, number := range drawnNumbers {
		for _, matrix := range matrices {
			for _, row := range matrix {
				for columnIndex, num := range row {
					if num != number {
						continue
					}
					row[columnIndex] = -1
				}
			}
			if checkRows(matrix) || checkColumns(matrix) {
				return sumOfMatrix(matrix) * number
			}
		}
	}
	return ^0
}

func findLoserScore(drawnNumbers []int, matrices [][][]int) int {

	m := make([][][]int, len(matrices))
	copy(m, matrices)
	var n int
	var lastWinner [][]int
	for _, number := range drawnNumbers {

		for _, matrix := range m {
			fmt.Println(matrix, number)
			for _, row := range matrix {
				for columnIndex, num := range row {
					if num != number {
						continue
					}
					row[columnIndex] = -1
				}
			}
			if checkRows(matrix) || checkColumns(matrix) {
				n = number
				lastWinner = matrix
				//m = append(matrices[:matrixIndex], matrices[matrixIndex+1:]...)
			}
		}
	}

	if lastWinner != nil {
		return sumOfMatrix(lastWinner) * n
	}

	return ^0
}

func GiantSquid(drawnNumbers []int, matrices [][][]int, findLast bool) int {

	if findLast {
		return findLoserScore(drawnNumbers, matrices)
	} else {
		return findWinnerScore(drawnNumbers, matrices)
	}

}
