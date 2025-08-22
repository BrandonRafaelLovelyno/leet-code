func minimumArea(grid [][]int) int {
    left, top := int(1e4), int(1e4)
    right, bottom := -1,-1

    for i := 0; i < len(grid) ; i++{
        for j:=0 ; j< len(grid[i]); j++{
            if grid[i][j] == 1 {
                if j < left {
                    left = j
                }
                
                if j > right {
                    right = j
                }

                if i < top {
                    top = i
                }
                
                if i > bottom{
                    bottom = i
                }
            }
        }
    }

    return (bottom - top + 1) * (right - left + 1)
}
