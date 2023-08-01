package main

import (
	"testing"
)

func TestSumInt(t *testing.T) {
	xTC := []struct {
		Case     int
		Input    []int
		Expected int
	}{
		{
			Case:     1,
			Input:    []int{1, 2, 3, 4, 5},
			Expected: 15,
		},
		{
			Case:     2,
			Input:    []int{0, 0, 0},
			Expected: 0,
		},
		{
			Case:     3,
			Input:    []int{1, -1, 4, 6, 10, -15},
			Expected: 5,
		},
	}

	for _, tc := range xTC {
		got := SumInt(tc.Input...)
		if got != tc.Expected {
			t.Errorf("sumint test case failed - caseo : %v | expected : %v | got : %v", tc.Case, tc.Expected, got)
		}
	}
}
