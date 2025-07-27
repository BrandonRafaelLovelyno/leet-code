func countHillValley(nums []int) int {
	ans := 0
	i := 1
	for i < len(nums) - 1 {
		j := i + 1
		if nums[i-1] < nums[i] {
			for j < len(nums) - 1 && nums[j] == nums[i] {
				j++
			}
			if nums[j] < nums[i] {
				ans += 1
			}
            i = j
		} else if nums[i-1] > nums[i] {
			for j < len(nums) - 1 && nums[j] == nums[i] {
				j++
			}
			if nums[j] > nums[i] {
				ans += 1
			}
            i = j
		} else {
            i++
        }
	}
    return ans
}