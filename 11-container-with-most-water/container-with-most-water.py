class Solution:
    def maxArea(self, height: list[int]) -> int:
        ans = 0
        p1, p2 = 0, len(height) - 1

        while p1 < p2:
            ans = max(ans, (p2 - p1) * min(height[p1], height[p2]))
            if height[p1] < height[p2]:
                p1+=1
            else:
                p2-=1

        return ans