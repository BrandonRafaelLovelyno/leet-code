import "math"

type DictKey struct {
    a int
    b int
}

var dict = make(map[DictKey]float64)

func soupServings(n int) float64 {
    u := int(math.Ceil(float64(n)/25))

    if u > 500 {
        return 1
    }

    return dp(u,u)
}

func dp(a, b int) float64{
    key := DictKey{a: a, b: b}
    d, ok := dict[key]
    if ok {
        return d
    }

    if b <= 0 && a > 0 {
        return 0
    }

    var ans float64 = 0
    if a <= 0 {
        if b > 0 {
            ans = 1
        } else if b <= 0  {
            ans = 0.5
        }
        dict[key] = ans
        return ans
    }

    ans = 0.25 * (dp(a - 4, b) + dp(a-3,b-1) + dp(a-2,b-2) +  dp(a-1,b-3))
    dict[key] = ans
    
    return ans
}