package intutils

import "math"

func Abs(a int) int {
	if a >= 0 {
		return a
	} else {
		return -a
	}
}

func Ceil(x float64) int {
	return int(math.Ceil(x))
}

func Pow(x int, f int) int {
	return int(math.Pow(float64(x), float64(f)))
}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}