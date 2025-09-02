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
	for iLeft, left := range points {
		for j, right := range points[iLeft+1:] {
			iRight := iLeft + 1 + j
			if left[1] >= right[1] {
				if iRight-iLeft == 1 {
					ans++
				} else {
					for i, p := range points[iLeft+1 : iRight] {
						if p[1] > left[1] || p[1] < right[1] {
							if i+iLeft+1 == iRight-1 {
								fmt.Println(points[iLeft], points[iRight])
								ans++
							}
						} else {
							break
						}
					}
				}
			}
		}
	}

	return ans
}