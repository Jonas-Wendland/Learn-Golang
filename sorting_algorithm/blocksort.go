package sorting_algorithm

func insertSort(array []int, x, y int) {
	for i := x + 1; i <= y; i++ {
		temp := array[i]
		j := i - 1
		for j >= x && array[j] > temp {
			array[j+1] = array[j]
			j--
		}
		array[j+1] = temp
	}
}

func rotate(array []int, x, y, z int) {
	reverse := func(a []int, i, j int) {
		for i < j {
			a[i], a[j] = a[j], a[i]
			i++
			j--
		}
	}
	reverse(array, x, y)
	reverse(array, y+1, z)
	reverse(array, x, z)
}

func merge(array []int, x, y, z int) {
	i := x
	j := y + 1
	for i <= y && j <= z {
		if array[i] <= array[j] {
			i++
		} else {
			rotate(array, i, j-1, j)
			i++
			y++
			j++
		}
	}
}

func BlockSort(array []int) {
	blockSize := 16

	for i := 0; i < len(array); i += blockSize {
		x := i + blockSize - 1
		if x >= len(array) {
			x = len(array) - 1
		}
		insertSort(array, i, x)
	}

	for blockSize < len(array) {
		for i := 0; i < len(array); i += 2 * blockSize {
			x := i + blockSize - 1
			y := i + 2*blockSize - 1
			if x >= len(array) || x >= len(array) {
				if x < len(array) {
					x = len(array) - 1
				} else {
					continue
				}
			}
			merge(array, i, x, y)
		}
		blockSize *= 2
	}
}
