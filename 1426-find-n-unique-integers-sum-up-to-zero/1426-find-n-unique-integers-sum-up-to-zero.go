func sumZero(n int) []int {
    ans := []int{}
    for i := n/2; i >0; i--{
        ans = append(ans, i)
        ans = append(ans, i * -1)
    }
    
    if n % 2 == 1{
        ans=append(ans,0)
    }
    return ans
}