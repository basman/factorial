package main

import (
	"fmt"
	"testing"
)

func TestFactorial_1to5(t1 *testing.T) {
	// arrange
	tbl := []struct {
		n uint64
		f uint64
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 24},
		{5, 120},
		{6, 720},
		{9, 362880},
		{12, 479001600},
	}

	for _, tt := range tbl {
		t1.Run(fmt.Sprintf("Test N=%v", tt.n), func(t *testing.T) {
			// act
			res := Factorial(tt.n)

			// assert
			if res != tt.f {
				t.Errorf("Wrong result for %v!: %v, expected %v", tt.n, res, tt.f)
			}
		})
	}
}
