package main

import (
	"fmt"

	"github.com/Jonas-Wendland/Learn-Golang/sorting_algorithm"
)

func main() {
	// fmt.Println("Hello, World!", rand.Intn(10))
	// fmt.Println("23 + 42 =", calculator.Add(23, 42))
	// fmt.Println(calculator.Divide(23, 5))

	// fmt.Println(calculator.Sum(1, 10))

	// fmt.Println(calculator.SumUntil(1000))

	// fmt.Println(calculator.IsSqaureNumber(25))

	// fmt.Println(calculator.MultiplyFromAToB(1, 10))

	// fmt.Println(composite.Add(23, 42))

	// point := composite.Point{X: 1, Y: 1}
	// fmt.Println(point.X, point.Y)

	// // composite.DemoCollections()

	// fmt.Println(point.DistanceFromZero())

	// numbers := []int{5, 3, 8, 4, 2}
	// sorting_algorithm.BubbleSort(numbers)
	// fmt.Println("Sorted array:", numbers)

	array := []int{9, 4, 3, 2, 7, 10, 23, 6, 100, 26, 55, 1}
	// sorting_algorithm.BlockSort(array)
	// fmt.Println("BlockSort:", array)

	sorting_algorithm.HeapSort(array)
	fmt.Println("Heapsort:", array)

}
