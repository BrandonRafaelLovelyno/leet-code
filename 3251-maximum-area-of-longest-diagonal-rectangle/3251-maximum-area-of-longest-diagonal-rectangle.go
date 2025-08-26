func max(a, b int) int {
    if a > b {
        return a
    }else {
        return b
    }
}

func areaOfMaxDiagonal(dimensions [][]int) int {
    diagonal, area := 0, 0

    for _, rect := range dimensions{
        d := rect[0] * rect[0] + rect[1] * rect[1] 
        if d > diagonal{
            diagonal = d
            area = rect[0] * rect[1]
        }else if d == diagonal {
            area = max(area,rect[0] * rect[1])
        }
    }

    return area
}