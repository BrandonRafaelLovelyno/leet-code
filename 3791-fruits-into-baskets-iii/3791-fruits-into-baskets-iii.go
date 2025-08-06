func numOfUnplacedFruits(fruits []int, baskets []int) int {
	placed := 0
    
    segmentTree := getSegmentTree(baskets)
    for _, f := range fruits{
        index := querySegmentTree(segmentTree, f, 1, 0, len(baskets) -1)

        if index != -1{
            placed = placed + 1
            updateSegmentTree(segmentTree, index, 1, 0, len(baskets) - 1)
        }
    }

    return len(fruits) - placed
}

func getSegmentTree(arr []int) []int{
	segmentTree := make([]int, len(arr)*4)
	buildSegmentTree(arr, segmentTree, 1, 0, len(arr)-1)
	return segmentTree
}

func buildSegmentTree(arr, segmentTree []int, node, start, end int) {
	if start == end {
		segmentTree[node] = arr[start]
        return
	}

	mid := (start + end) / 2
	buildSegmentTree(arr, segmentTree, 2*node, start, mid)
    buildSegmentTree(arr, segmentTree, 2*node+1, mid+1, end)

	val := max(segmentTree[2*node], segmentTree[2*node + 1])
	segmentTree[node] = val
}

func querySegmentTree(segmentTree []int, val, node, start, end int) int {
	if segmentTree[node] < val {
		return -1
	}

    if start == end{
        return start
    }

	mid := (start + end) / 2
	
    if segmentTree[node * 2] >= val {
        ans := querySegmentTree(segmentTree, val, node * 2, start, mid)
        return ans
    }

	return querySegmentTree(segmentTree, val, node*2+1, mid+1, end)
}

func updateSegmentTree(segmentTree []int, index, node, start, end int) {
    if start == end{
        segmentTree[node] = -1
        return
    }

    mid := (start + end) / 2
    if index <= mid{
        updateSegmentTree(segmentTree, index, 2*node, start, mid)
    }else{
        updateSegmentTree(segmentTree, index, 2 * node + 1, mid + 1, end)
    }

    segmentTree[node] = max(segmentTree[2*node], segmentTree[2*node + 1])
    return
}