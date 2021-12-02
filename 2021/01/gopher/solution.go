package main

import (
	"fmt"
	"github.com/nozgurozturk/aoc/util/gopher/io"
	"github.com/nozgurozturk/aoc/util/gopher/math"
	"os"
)

const FileName = "../INPUT"

func main() {
	depths := io.ReadIntegers(FileName)
	windowSize := math.ParseInt(os.Args[1])

	result := SonarSweep(depths, windowSize)

	fmt.Println(result)
}

func SonarSweep(depths []int, windowSize int) int {

	inc := 0
	prevSum := 0

	for i := 0; i <= len(depths)-windowSize; i++ {
		window := make([]int, windowSize, windowSize)

		for j := 0; j < windowSize; j++ {
			window[j] = depths[i+j]
		}

		currentSum := math.SumOf(window...)
		if currentSum > prevSum && prevSum > 0 {
			inc++
		}

		prevSum = currentSum
	}
	return inc
}
