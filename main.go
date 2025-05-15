package main

import (
	"fmt"
	"math/rand"

	"github.com/Jonas-Wendland/Learn-Golang/calculator"
	"github.com/Jonas-Wendland/Learn-Golang/composite"
)

func main() {
	fmt.Println("Hello, World!", rand.Intn(10))
	fmt.Println("23 + 42 =", calculator.Add(23, 42))
	fmt.Println(calculator.Divide(23, 5))

	fmt.Println(calculator.Sum(1, 10))

	fmt.Println(calculator.SumUntil(1000))

	fmt.Println(calculator.IsSqaureNumber(25))

	fmt.Println(calculator.MultiplyFromAToB(1, 10))

	fmt.Println(composite.Add(23, 42))

	point := composite.Point{X: 3, Y: 7}
	fmt.Println(point.X, point.Y)

	composite.DemoCollections()
}
