import (
	"fmt"
	"math/bits"
)

func makeTheIntegerZero(num1 int, num2 int) int {
	for i:=0; num1 - i*num2 > 0; i++{
		if bits.OnesCount(uint(num1-i*num2)) <= i && num1-i*num2 >= i {
			return i
		}
	}
	return -1
}