import "math"

func isPowerOfThree(n int) bool {
    largest := int(math.Pow(3,19))
    return n > 0 && largest % n == 0
}