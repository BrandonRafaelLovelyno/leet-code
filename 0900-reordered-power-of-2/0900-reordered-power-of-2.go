var powers = make(map[int]map[int]int)

func reorderedPowerOf2(n int) bool {
    initPowers(29)
    nFreq := makeFrequencyMap(n)

    for _, freqs := range powers {
        isPower := true
        for i := 0; i <= 9; i++{
            nF, _ := nFreq[i]
            f, _ := freqs[i]

            if nF != f {
                isPower = false
            }
        }

        if isPower {
            return true
        }
    }

    return false
}

func initPowers(n int) {
    for i := 0; i <= n; i++{
        power := 1 << i
        freqs := makeFrequencyMap(power)
        powers[i] = freqs
    }
}

func makeFrequencyMap(n int) map[int]int {
    freqs := make(map[int]int)

    for n > 0 {
        digit := n % 10

        f, ok := freqs[digit]
        if ok {
            freqs[digit] = f + 1
        }else{
            freqs[digit] = 1
        }

        n /= 10
    }

    return freqs
}