package io

import (
	"bufio"
	"github.com/nozgurozturk/aoc/util/gopher/math"

	"os"
	"strings"
)

func ReadStrings(filename string, from, to uint) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var text []string
	for i := uint(0); scanner.Scan(); i++ {
		if i < from {
			continue
		}
		if i > to && to > 0 {
			break
		}
		text = append(text, scanner.Text())
	}
	return text
}

func ReadIntegers(filename string, from, to uint) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var numbers []int
	for i := uint(0); scanner.Scan(); i++ {
		if i < from {
			continue
		}
		if i > to {
			break
		}
		numbers = append(numbers, math.ParseInt(scanner.Text()))
	}
	return numbers
}

func ReadPairs(filename string, delimiter string, from, to uint) []math.Pair {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var pairs []math.Pair
	for i := uint(0); scanner.Scan(); i++ {
		if i < from {
			continue
		}
		if i > to {
			break
		}
		pairs = append(pairs, math.ParsePair(scanner.Text(), delimiter))
	}
	return pairs
}

func ReadMatrices(filename string, from, to uint) [][][]int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	var matrices [][][]int
	var matrix [][]int

	for i := uint(0); scanner.Scan(); i++ {
		if i < from {
			continue
		}

		if i > to {
			break
		}

		if scanner.Text() == "" {
			matrices = append(matrices, matrix)
			matrix = nil
			continue
		}

		splits := strings.Fields(strings.TrimSpace(scanner.Text()))
		var row []int
		for j := 0; j < len(splits); j++ {
			row = append(row, math.ParseInt(splits[j]))
		}
		matrix = append(matrix, row)
	}

	if matrix != nil {
		matrices = append(matrices, matrix)
	}

	return matrices
}
