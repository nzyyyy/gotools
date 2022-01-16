package emath

import "math"

func AbsInt(i int) int {
	return int(math.Abs(float64(i)))
}
