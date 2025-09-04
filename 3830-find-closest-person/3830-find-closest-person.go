import "math"

func findClosest(x int, y int, z int) int {
    xz, yz := math.Abs(float64(x-z)), math.Abs(float64(y-z))

    if xz > yz {
        return 2
    }else if xz < yz {
        return 1
    }else{
        return 0
    }
}