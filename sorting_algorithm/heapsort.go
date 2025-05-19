package sorting_algorithm

func heapify(array []int, n, i int) {
	parent := i
	iLeftChild := 2*i + 1
	iRightChild := 2*i + 2

	if iLeftChild < n && array[iLeftChild] > array[parent] {
		parent = iLeftChild
	}

	if iRightChild < n && array[iRightChild] > array[parent] {
		parent = iRightChild
	}

	if parent != i {
		array[i], array[parent] = array[parent], array[i]

		heapify(array, n, parent)
	}
}

func HeapSort(array []int) {

	n := len(array)

	for i := n/2 - 1; i >= 0; i-- {
		heapify(array, n, i)
	}

	for i := n - 1; i > 0; i-- {
		array[0], array[i] = array[i], array[0]

		heapify(array, i, 0)
	}
}
