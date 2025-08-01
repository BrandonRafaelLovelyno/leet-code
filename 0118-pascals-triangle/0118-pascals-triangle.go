func generate(numRows int) [][]int {
    ans := [][]int{{1}}

    for i:= 1; i < numRows; i++{
        up:= ans[i-1]
        curr := []int{1}
        for j := 1; j < len(up);j++{
            curr = append(curr, up[j] + up[j-1])
        } 
        curr = append(curr,1)
        ans = append(ans, curr)
    }

    return ans
}