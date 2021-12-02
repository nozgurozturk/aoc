package main

import (
	"aoc/pkg/math"
	"testing"
)

func TestPuzzle(t *testing.T) {
	tests := []struct {
		input              []math.Pair
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
				{"down", 2},
			},
			15,
			10,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			h, v := Dive(test.input)

			if h != test.expectedHorizontal {
				t.Errorf("expected horizontal %6d, result %6d\n", test.expectedHorizontal, h)
			}

			if v != test.expectedVertical {
				t.Errorf("expected vertical %6d, result %6d\n", test.expectedVertical, v)
			}
		})
	}
}
