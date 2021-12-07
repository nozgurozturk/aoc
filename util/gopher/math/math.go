package math

import (
	"math/bits"
	"strconv"
	"strings"
)

// ParseInt parse string to int
func ParseInt(str string) int {
	integer, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}

	return integer
}

// ParseInts parse string with a given delimiter to array of int
func ParseInts(str string, delimiter string) []int {
	strs := strings.Split(str, delimiter)
	integers := make([]int, len(strs), len(strs))

	for i, s := range strs {
		integers[i] = ParseInt(s)
	}
	return integers
}

type Pair struct {
	First  string
	Second int
}

// ParsePair parse string and integer with a given row and column delimiter to array of tuple []{string, int}
func ParsePair(str, delimiter string) Pair {
	p := strings.Split(str, delimiter)

	return Pair{p[0], ParseInt(p[1])}
}

// Max maximum of two integer
func Max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

// Min minimum of two integer
func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// SumOf sum of finite integers
func SumOf(nums ...int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

// ProductOf multiplication of finite integer
func ProductOf(nums ...int) int {
	prd := 1
	for _, num := range nums {
		prd *= num
	}
	return prd
}

// CheckSumOf checks total of n int for a given number
func CheckSumOf(sum int) func(nums ...int) bool {
	return func(nums ...int) bool {
		return sum == SumOf(nums...)
	}
}

// Combinations returns combinations of n elements for a given int array.
func Combinations(set []int, n int) (subsets [][]int) {
	length := uint(len(set))

	if n > len(set) {
		n = len(set)
	}

	// Go through all possible combinations of objects
	// from 1 (only first object in subset) to 2^length (all objects in subset)
	for subsetBits := 1; subsetBits < (1 << length); subsetBits++ {
		if n > 0 && bits.OnesCount(uint(subsetBits)) != n {
			continue
		}

		var subset []int

		for object := uint(0); object < length; object++ {
			// checks if object is contained in subset
			// by checking if bit 'object' is set in subsetBits
			if (subsetBits>>object)&1 == 1 {
				// add object to subset
				subset = append(subset, set[object])
			}
		}
		// add subset to subsets
		subsets = append(subsets, subset)
	}
	return subsets
}

// Permutations returns permutations of a given int array.
func Permutations(sets []int) [][]int {
	result := make([][]int, 0)
	for p := make([]int, len(sets)); p[0] < len(p); nextPerm(p) {
		result = append(result, getPerm(sets, p))
	}
	return result
}

func nextPerm(p []int) {
	for i := len(p) - 1; i >= 0; i-- {
		if i == 0 || p[i] < len(p)-i-1 {
			p[i]++
			return
		}
		p[i] = 0
	}
}

func getPerm(orig, p []int) []int {
	result := append([]int{}, orig...)
	for i, v := range p {
		result[i], result[i+v] = result[i+v], result[i]
	}
	return result
}
