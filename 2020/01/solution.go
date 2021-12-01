package main

import (
	"aoc/pkg/io"
	"aoc/pkg/math"
	"fmt"
	"os"
)

const FileName = "INPUT"

func main() {

	expenses := io.ReadIntegers(FileName)
	numberOfEntries := math.ParseInt(os.Args[1])

	result := ReportAndRepair(expenses, numberOfEntries)

	fmt.Println(result)
}

func ReportAndRepair(expenses []int, numberOfEntries int) int {

	checkSum2020 := math.CheckSumOf(2020)

	intArr := make([]int, numberOfEntries, numberOfEntries)
	for _, e1 := range expenses {
		intArr[0] = e1
		for _, e2 := range expenses {
			intArr[1] = e2
			if  numberOfEntries == 2 {
				if checkSum2020(intArr...) {
					return math.ProductOf(intArr...)
				}
			}
			if numberOfEntries == 3 {
				for _, e3 := range expenses {
					intArr[2] = e3
					if checkSum2020(intArr...) {
						return math.ProductOf(intArr...)
					}
				}
			}

		}
	}

	return 0
}
