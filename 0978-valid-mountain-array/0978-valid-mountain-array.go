func validMountainArray(arr []int) bool {
	n :=  len(arr)

    if n < 3{
        return false
    }

    p1 := 0
    for p1 < n-1 && arr[p1] < arr[p1 + 1]{
        p1++
    }

    if p1 == 0 || p1 == n-1{
        return false
    }

    for p1 < n-1 && arr[p1] > arr[p1 + 1]{
        p1++
    }

    return p1 == n-1
}