import "sort"

type ByXThenY [][]int

func (arr ByXThenY) Len() int {
	return len(arr)
}

func (arr ByXThenY) Less(i, j int) bool {
	if arr[i][0] != arr[j][0] {
		return arr[i][0] < arr[j][0]
	} else {
		return arr[i][1] > arr[j][1]
	}
}

func (arr ByXThenY) Swap(i, j int) {
	arr[i], arr[j] = arr[j], arr[i]
}


func numberOfPairs(points [][]int) int {
	ans := 0
	sort.Sort(ByXThenY(points))
	
	for i := 0; i < len(points); i++{
		for j := i+1; j < len(points); j++{
			left, right := points[i], points[j]

			if left[1] < right[1]{
				continue
			}

			invalid := false
			for k := i+1; k < j; k++{
				if points[k][1] <= points[i][1] && points[k][1] >= points[j][1]{
					invalid = true
					break
				} 
			}

			if !invalid{
				ans++
			}
		}
	}

    return ans
}