func findDiagonalOrder(mat [][]int) []int {
	ans := []int{}

	isUp, x,y := true,0,0
	for true{
		for true {
			fmt.Println(y, x)
			ans = append(ans, mat[y][x])
			if isUp {
				if ok := moveUp(&x, &y, mat); !ok {
					break
				}
			}else{
				if ok := moveDown(&x, &y, mat); !ok {
					break
				}
			}
		}
		if x == len(mat[0]) - 1 && y == len(mat) - 1 {
			break
		}else{
			moveNewStart(&x, &y, isUp, mat)
			isUp = !isUp
		}
	}

	return ans
}

func moveUp(x, y *int, mat [][]int) bool{
	if *y - 1 >= 0 && *x + 1 <= len(mat[0]) - 1 {
		*x++
		*y--
		return true
	}else{
		return false
	}
}

func moveDown(x, y *int, mat [][]int) bool{
	if *y + 1 <= len(mat) - 1 && *x - 1 >= 0 {
		*x--
		*y++
		return true
	}else {
		return false
	}
}

func moveNewStart(x, y *int, isUp bool, mat [][]int){
	if isUp {
		if *x + 1 >= len(mat[0]){
			*y++
		}else{
			*x++
		}
	}else{
		if *y + 1 >= len(mat){
			*x++
		}else{
			*y++
		}
	}
}