package main

import (
	"fmt"
	"github.com/nozgurozturk/aoc/util/gopher/io"
	"github.com/nozgurozturk/aoc/util/gopher/math"
)

const FileName = "../INPUT"

func main() {
	drawnNumbers := io.ReadIntegers(FileName, 0, 2)
	matrices := io.ReadMatrices(FileName, 2, math.MaxUInt)

	score := GiantSquid(drawnNumbers, matrices)

	fmt.Println(score)
}

func GiantSquid(drawnNumbers []int, matrices [][][]int) int {
	panic("not implemented")
}
