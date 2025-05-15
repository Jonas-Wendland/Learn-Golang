package composite

import (
	"fmt"
)

func DemoCollections() {
	var primesArray [5]int = [5]int{2, 3, 5, 7, 11}

	fmt.Println(primesArray)

	// primesSlice := []int{2, 3, 5, 7, 11}
	primesSlice := make([]int, 5, 7)
	primesSlice = append(primesSlice, 2)
	primesSlice = append(primesSlice, 3)
	primesSlice = append(primesSlice, 5)
	fmt.Println(primesSlice)
	fmt.Println(len(primesSlice))
	fmt.Println(cap(primesSlice))

}
