import "strconv"

func maximum69Number (num int) int {
    str := strconv.Itoa(num)
    runes := []rune(str)

    hasChanged := false
    for index, value := range(runes){
        if !hasChanged && value == '6' {
            hasChanged = true
            runes[index] = '9'
        }
    }

    ans, _ := strconv.Atoi(string(runes))
    return ans
}