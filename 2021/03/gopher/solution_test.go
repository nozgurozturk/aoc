package main

import (
	"testing"
)

func TestBinaryDiagnostic(t *testing.T) {
	tests := []struct {
		input                     []string
		expectedPowerConsumption  int
		expectedLifeSupportRating int
	}{
		{
			[]string{
				"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010",
			},
			198,
			230,
		},
	}

	for _, test := range tests {
		t.Run("Testing BinaryDiagnostic", func(t *testing.T) {
			p, l := BinaryDiagnostic(test.input)

			if p != test.expectedPowerConsumption {
				t.Errorf("expected expectedPowerConsumption %d, result %d\n", test.expectedPowerConsumption, p)
			}

			if l != test.expectedLifeSupportRating {
				t.Errorf("expected expectedLifeSupportRating %d, result %d\n", test.expectedLifeSupportRating, l)
			}
		})
	}
}

