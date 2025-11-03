class Solution:
    def minCost(self, colors: str, neededTime: List[int]) -> int:
        ans = 0
        curr_color, max_group_time, sum_group_time = colors[0], 0, 0

        for c, t in zip(colors, neededTime):
            if c != curr_color:
                ans += sum_group_time - max_group_time
                sum_group_time, max_group_time = t, t
                curr_color = c
            else:
                max_group_time = max(max_group_time, t)
                sum_group_time += t
        
        ans += sum_group_time - max_group_time

        return ans