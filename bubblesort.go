package main

func bubbleSort(vector []int) {
	_switch := true
	for _switch {
		_switch = false
		for i := 0; i < len(vector)-1; i++ {
			if vector[i] > vector[i+1] {
				_switch = true
				switchValues(&vector[i], &vector[i+1])
			}
		}
	}
}

func switchValues(x, y *int) {
	aux := *x
	*x = *y
	*y = aux
}
