package main

import (
	"aoc/pkg/io"
	"aoc/pkg/math"
	"fmt"
)

const FileName = "INPUT"

func main() {
	input := io.ReadPairs(FileName, " ")

	h, v := Dive(input)

	fmt.Println(h * v)
}

func Dive(input []math.Pair) (h int, v int) {
	panic("not implemented")
}
