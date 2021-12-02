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
	for _, pair := range input {
		command := pair.First

		if command == "forward" {
			h += pair.Second
		}
		if command == "up" {
			v -= pair.Second
		}
		if command == "down" {
			v += pair.Second
		}
	}

	return
}
