import "sort"

func sortMatrix(grid [][]int) [][]int {
    for i := len(grid); i >= 0;i--{
		j,diagonal := 0,[]int{}
		for i + j < len(grid) {
			diagonal = append(diagonal, grid[i+j][j])
			j++
		}
		sort.Ints(diagonal)
		j--
		for _, val := range diagonal {
			grid[i+j][j] = val
			j--
		}
	}

	for j := 1; j < len(grid[0]); j++{
		i,diagonal := 0,[]int{}
		for i + j < len(grid[0]) {
			diagonal = append(diagonal, grid[i][i+j])
			i++
		}
		sort.Sort(sort.Reverse(sort.IntSlice(diagonal)))
		i--
		for _, val := range diagonal {
			grid[i][i+j] = val
			i--
		}	
	}
	
	return grid
}