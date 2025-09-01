func maxAverageRatio(classes [][]int, extraStudents int) float64 {
	arr := [][]float64{}
	for _, class := range classes {
		pass, total := float64(class[0]), float64(class[1])
		gain := ((pass + 1) / (total + 1)) - (pass / total)
		arr = append(arr, []float64{pass, total, gain})
	}
	
	makeHeap(arr)

	for i := 0; i < extraStudents; i++ {
		top := arr[0]
		
		top[0]++
		top[1]++
		top[2] = ((top[0] + 1) / (top[1] + 1)) - (top[0] / top[1])

		arr[0] = top
		heapify(arr, 0, len(arr))
	}

	var sumRatio float64
	for _, a := range arr {
		sumRatio += a[0] / a[1]
	}

	return sumRatio / float64(len(classes))
}

func makeHeap(arr [][]float64) {
	n := len(arr)
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, i, n)
	}
}

func heapify(arr [][]float64, i int, n int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2

	if left < n && arr[left][2] > arr[largest][2] {
		largest = left
	}

	if right < n && arr[right][2] > arr[largest][2] {
		largest = right
	}

	if largest != i {
		swap(arr, i, largest)
		heapify(arr, largest, n)
	}
}

func swap(arr [][]float64, i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}
