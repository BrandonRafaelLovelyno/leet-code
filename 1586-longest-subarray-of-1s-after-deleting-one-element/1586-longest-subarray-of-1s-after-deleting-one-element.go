import "math"

func longestSubarray(nums []int) int {
	ans, summary := 0, []int{}

	i := 0
	for i < len(nums) {
		j := i + 1
		for j < len(nums) && nums[j] == nums[i] {
			j++
		}
		summary = append(summary, j-i)
		i = j
	}

	if zeroFirst := nums[0] == 0; zeroFirst {
		i = 1
	} else {
		i = 0
	}

	for i < len(summary) {
		ans = int(math.Max(float64(ans), float64(summary[i])))

		if i+2 < len(summary) {
			if summary[i+1] < 2 {
				ans = int(math.Max(float64(ans), float64(summary[i]+summary[i+2])))
			}
			i += 2
		} else {
			break
		}
	}

	if len(summary) < 2 && ans != 0 {
		ans--
	}

	return ans
}