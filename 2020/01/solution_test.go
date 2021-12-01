package main

import (
	"fmt"
	"testing"
)

func TestReportAndRepair(t *testing.T) {

	tests := []struct {
		input           []int
		numberOfEntries int
		expected        int
	}{
		{
			[]int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			2,
			514579,
		}, {
			[]int{
				1721,
				979,
				366,
				299,
				675,
				1456,
			},
			3,
			241861950,
		},
	}

	for _, test := range tests {
		testName := fmt.Sprintf("For given %d entiries", test.numberOfEntries)
		t.Run(testName, func(t *testing.T) {
			result := ReportAndRepair(test.input, test.numberOfEntries)
			if result != test.expected {
				t.Errorf("expected %6d, result %6d\n", test.expected, result)
			}
		})
	}

}
