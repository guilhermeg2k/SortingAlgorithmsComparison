package main

func heapify(v []int, i, size int) {
	var left, right, largest int
	for i < size {
		left = (i * 2) + 1
		right = (i * 2) + 2
		largest = i
		if left < size && v[left] > v[largest] {
			largest = left
		}
		if right < size && v[right] > v[largest] {
			largest = right
		}
		if i == largest {
			break
		}
		swp := v[i]
		v[i] = v[largest]
		v[largest] = swp
		i = largest
	}
}

func make_heap(v []int, size int) {
	for i := size / 2; i >= 0; i-- {
		heapify(v, i, size)
	}
}

func heapSort(v []int, size int) {
	make_heap(v, size)
	for i := size - 1; i > 0; i-- {
		swp := v[i]
		v[i] = v[0]
		v[0] = swp
		heapify(v, 0, i)
	}
}
