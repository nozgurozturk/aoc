package io

import (
	"aoc/pkg/math"
	"bufio"
	"os"
	"strings"
)

func ReadStrings(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var text []string
	for scanner.Scan() {
		text = append(text, strings.TrimRight(scanner.Text(), "\n"))
	}
	return text
}

func ReadIntegers(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var numbers []int
	for scanner.Scan() {
		numbers = append(numbers, math.ParseInt(scanner.Text()))
	}
	return numbers
}
