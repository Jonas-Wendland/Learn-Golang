package sorting_algorithm

func QuickSort(array []int, low, high int) {

	if low < high {
		p := partition(array, low, high)

		QuickSort(array, low, p-1)
		QuickSort(array, p+1, high)
	}
}

func partition(array []int, low, high int) int {
	pivot := array[high]
	i := low - 1

	for j := low; j <= high; j++ {
		if array[j] < pivot {
			i++
			array[i], array[j] = array[j], array[i]
		}
	}
	array[i+1], array[high] = array[high], array[i+1]
	return i + 1
}
