func isPowerOfFour(n int) bool {
    if n == 0 {
        return false
    }

    isTwoPower := n & (n-1) == 0
    return isTwoPower && n & 0x55555555 == n
}