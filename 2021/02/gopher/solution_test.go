package main

import (
	"fmt"
	"github.com/nozgurozturk/aoc/util/gopher/math"
	"testing"
)

func TestDive(t *testing.T) {
	tests := []struct {
		input              []math.Pair
		aimOn              bool
		expectedHorizontal int
		expectedVertical   int
	}{
		{
			[]math.Pair{
				{"forward", 5},
				{"down", 5},
				{"forward", 8},
				{"up", 3},
				{"down", 8},
				{"forward", 2},
			},
			false,
			15,
			10,
		},
		{
			[]math.Pair{
				{"forward", 5},
				{"down", 5},
				{"forward", 8},
				{"up", 3},
				{"down", 8},
				{"forward", 2},
			},
			true,
			15,
			60,
		},
	}

	for i, test := range tests {

		t.Run(fmt.Sprintf("Part %d", i+1), func(t *testing.T) {
			h, v := Dive(test.input, test.aimOn)

			if h != test.expectedHorizontal {
				t.Errorf("expected horizontal %d, result %d\n", test.expectedHorizontal, h)
			}

			if v != test.expectedVertical {
				t.Errorf("expected vertical %d, result %d\n", test.expectedVertical, v)
			}
		})
	}
}
