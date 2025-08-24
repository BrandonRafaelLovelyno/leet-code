func longestSubarray(nums []int) int {
	ans,left,isAllOnes := 0, 0, true

	for left < len(nums){
		if nums[left] == 0 {
			isAllOnes = false
			left++
			continue
		}

		right, divider,zeroCount := left,left,0
		for right < len(nums) && zeroCount < 2 {
			if nums[right] == 1 {
				ans = max(ans, (right - left + 1) - zeroCount)
			} else {
				isAllOnes = false
				zeroCount++
				if zeroCount == 1 {
					divider = right
				}
			}
			right++
		}

		left = divider + 1
	}

	if isAllOnes && ans != 0 {
		ans--
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
