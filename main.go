package main

import (
	"fmt"
	"math/rand"

	"github.com/Jonas-Wendland/Learn-Golang/calculator"
)

func main() {
	fmt.Println("Hello, World!", rand.Intn(10))
	fmt.Println("23 + 42 =", calculator.Add(23, 42))
}
