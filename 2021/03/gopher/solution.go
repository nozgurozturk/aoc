package main

import (
	"fmt"
	"github.com/nozgurozturk/aoc/util/gopher/collection"
	"github.com/nozgurozturk/aoc/util/gopher/io"
	"strconv"
	"sync"
	"sync/atomic"
)

const FileName = "../INPUT"

func main() {
	input := io.ReadStrings(FileName)

	p, l := BinaryDiagnostic(input)

	fmt.Println(p, l)
}

func BinaryDiagnostic(input []string) (p, l int) {

	p = calculatePowerConsumption(input)
	l = calculateLifeSupportRating(input)

	return p, l
}

func findMostAndLeastCommonValues(input []string, column int) (uint8, uint8) {
	zero, one := 0, 0

	for _, line := range input {
		if line[column] == '0' {
			zero++
		} else {
			one++
		}
	}

	if zero > one {
		return '0', '1'
	}
	return '1', '0'
}

func calculateRating(seq []string, position int, isMost bool) int64 {
	if len(seq) == 1 {
		g, _ := strconv.ParseInt(seq[0], 2, 16)
		return g
	}

	most, least := findMostAndLeastCommonValues(seq, position)

	// comparator
	comp := least
	if isMost {
		comp = most
	}

	return calculateRating(collection.Filter(seq, func(s string) bool {
		return comp == s[position]
	}), position+1, isMost)
}

func calculateLifeSupportRating(input []string) int {

	var o2 int64
	var co2 int64

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		o2Input := make([]string, len(input))
		copy(o2Input, input)
		atomic.AddInt64(&o2, calculateRating(o2Input, 0, true))
		wg.Done()
	}()

	go func() {
		co2Input := make([]string, len(input))
		copy(co2Input, input)
		atomic.AddInt64(&co2, calculateRating(co2Input, 0, false))
		wg.Done()
	}()

	wg.Wait()

	return int(co2 * o2)
}

func calculatePowerConsumption(input []string) int {

	pInput := make([]string, len(input))
	copy(pInput, input)

	gamaBinaryString := ""
	epsilonBinaryString := ""

	for i := 0; i < len(pInput[0]); i++ {
		most, least := findMostAndLeastCommonValues(pInput, i)
		gamaBinaryString += string(most)
		epsilonBinaryString += string(least)
	}

	g, _ := strconv.ParseInt(gamaBinaryString, 2, 16)
	e, _ := strconv.ParseInt(epsilonBinaryString, 2, 16)

	return int(g * e)
}
