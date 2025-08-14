import "strconv"

func largestGoodInteger(num string) string {
    digit := -1

    for i :=0 ; i < len(num) -2 ;i++{
        n, d := num[i] , int(num[i] - '0')

        if n == num[i + 1] && n == num[i + 2] && d > digit{
            digit = d
        } 
    }

    if digit == -1{
        return ""
    }

    sDigit := strconv.Itoa(digit)
    return sDigit + sDigit + sDigit
}