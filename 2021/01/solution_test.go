package main

import (
	"testing"
)

func TestSonarSweep(t *testing.T) {
	tests := []struct {
		input    []int
		expected int
	}{
		{
			[]int{
				199,
				200,
				208,
				210,
				200,
				207,
				240,
				269,
				260,
				263,
			},
			7,
		},
	}

	for _, test := range tests {
		t.Run("", func(t *testing.T) {
			result := SonarSweep(test.input)

			if result != test.expected {
				t.Errorf("expected %6d, result %6d\n", test.expected, result)
			}
		})
	}
}
