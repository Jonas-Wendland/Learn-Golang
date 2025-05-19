package sorting_algorithm

func BubbleSort(array []int) {
	for i := 0; i < len(array); i++ {
		for j := 0; j < len(array)-i-1; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
			}
		}
	}

}

// func main() {
// 	numbers := []int{5, 3, 8, 4, 2}
// 	BubbleSort(numbers)
// 	fmt.Println("Sorted array:", numbers)
// }

// numbers in array implement different sorting algorithm -> (https://en.wikipedia.org/wiki/Sorting_algorithm)
// Dijkstra shortest path javascript and golang
