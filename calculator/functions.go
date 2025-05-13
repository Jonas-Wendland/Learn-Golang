package calculator

import "math"

func Abs(value int) int {
	if value >= 0 {
		return value
	}

	return -value
}

func IsSqaureNumber(value int) bool {

	if sqrt := math.Sqrt(float64(value)); sqrt != float64(int(sqrt)) {
		return false
	}
	return true
}
