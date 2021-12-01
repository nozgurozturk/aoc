package main

import (
	"aoc/pkg/io"
	"fmt"
)

const FileName = "INPUT"

func main() {

	depths := io.ReadIntegers(FileName)

	result := SonarSweep(depths)

	fmt.Println(result)
}

func SonarSweep(depths []int) int {
	inc := 0
	for i := 1; i < len(depths); i++ {
		if depths[i] > depths[i-1] {
			inc++
		}
	}
	return inc
}
