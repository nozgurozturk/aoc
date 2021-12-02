package main

import (
	"aoc/pkg/io"
	"aoc/pkg/math"
	"fmt"
	"os"
	"strconv"
)

const FileName = "INPUT"

func main() {
	input := io.ReadPairs(FileName, " ")
	aimOn, _ := strconv.ParseBool(os.Args[1])

	h, v := Dive(input, aimOn)

	fmt.Println(h * v)
}

func Dive(input []math.Pair, aimOn bool) (h int, v int) {
	aim := 0
	for _, pair := range input {
		command := pair.First
		value := pair.Second

		if command == "forward" {
			h += value
			if aimOn {
				v += aim * value
			}
		}
		if command == "up" {
			if aimOn {
				aim -= value
				continue
			}
			v -= value
		}
		if command == "down" {
			if aimOn {
				aim += value
				continue
			}
			v += value
		}
	}

	return
}
