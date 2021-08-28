package main

func merge(v, v1, v2 []int) {
	i, j, k := 0, 0, 0
	for i < len(v1) && j < len(v2) {
		if v1[i] <= v2[j] {
			v[k] = v1[i]
			k++
			i++
		} else {
			v[k] = v2[j]
			k++
			j++
		}
	}
	for i < len(v1) {
		v[k] = v1[i]
		k++
		i++
	}
	for j < len(v2) {
		v[k] = v2[j]
		k++
		j++
	}
}

func mergeSort(vector []int) {
	vectorSize := len(vector)
	if vectorSize < 2 {
		return
	}
	vectorMid := int(vectorSize / 2)
	var (
		v1 = make([]int, vectorMid)
		v2 = make([]int, vectorSize-vectorMid)
	)
	for i := 0; i < vectorSize; i++ {
		if i < vectorMid {
			v1[i] = vector[i]
		} else {
			v2[i-vectorMid] = vector[i]
		}
	}
	mergeSort(v1)
	mergeSort(v2)
	merge(vector, v1, v2)
}
