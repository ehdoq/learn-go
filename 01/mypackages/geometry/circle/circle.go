package circle

import (
	"math"
)

func Area(r float64) float64 {
	return math.Pi * r * r
}

func Perimeter(r float64) float64 {
	return math.Pi * 2 * r
}
