package main

import (
	"fmt"
	"github.com/nozgurozturk/aoc/util/gopher/io"
	"github.com/nozgurozturk/aoc/util/gopher/math"
	"os"
	"strconv"
)

const FileName = "../INPUT"

var commands = map[string]struct {
	x int
	y int
}{
	"forward": {1, 0},
	"up":      {0, -1},
	"down":    {0, 1},
}

func main() {
	input := io.ReadPairs(FileName, " ", math.MinUInt, math.MaxUInt)
	aimOn, _ := strconv.ParseBool(os.Args[1])

	h, v := Dive(input, aimOn)

	fmt.Println(h * v)
}

func Dive(input []math.Pair, aimOn bool) (h int, v int) {
	aim := 0

	for _, pair := range input {
		command := pair.First
		value := pair.Second

		dir, ok := commands[command]
		if !ok {
			continue
		}

		hValue := value * dir.x
		vValue := value * dir.y

		h += hValue

		if aimOn {
			vValue = aim * hValue
		}

		v += vValue
		aim += value * dir.y

	}

	return
}
