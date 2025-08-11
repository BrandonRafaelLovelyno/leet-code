const MOD = 1e9 + 7

func productQueries(n int, queries [][]int) []int {
	powers := makePowers(n)
	prefix := makePrefixProduct(powers)

	ans := []int{}
	for _, query := range queries {
		left, right := 1, prefix[query[1]]

        if query[0] != 0 {
            left = prefix[query[0]-1]
        }

		leftInv := getModInverse(left)
		ans = append(ans, (right * leftInv)%MOD)
	}

	return ans
}

func makePowers(n int) []int {
	powers := []int{}

	for i := 0; n > 0; i++ {
		if n&1 == 1 {
			powers = append(powers, 1<<i)
		}
		n >>= 1
	}

	return powers
}

func makePrefixProduct(powers []int) []int {
	if len(powers) == 0 {
		prefix := []int{}
		return prefix
	}

	prefix := make([]int, len(powers))
	prefix[0] = powers[0]

	for i := 1; i < len(powers); i++ {
		prefix[i] = (prefix[i-1] * powers[i]) % MOD
	}

	return prefix
}

func getModInverse(n int) int {
	return modExponent(n, MOD - 2)
}

func modExponent(n, pow int) int {
    ans := 1
    n %= MOD

    for pow > 0 {
        if pow % 2 == 1{
            ans = (ans * n) % MOD 
        }

        n = (n*n) % MOD
        pow /= 2
    }

    return ans
}