func maxArea(height []int) int {
	var maxArea int = 0

	var p1, p2 int = 0, len(height) - 1
	for p1 < p2 {
		maxArea = max(maxArea, (min(height[p1], height[p2]) * (p2 - p1)))
		if height[p1] < height[p2] {
			p1++
		} else {
			p2--
		}
	}

	return maxArea
}