package composite

import "math"

type Point struct {
	X int
	Y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}

func (point *Point) DistanceFromZero() float64 {
	return math.Sqrt(
		math.Pow(float64(point.X), 2) +
			math.Pow(float64(point.Y), 2))
}
