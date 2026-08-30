class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        curr_min, curr_max, best = float('inf'), 0, 0
        for price in prices:
            if price > curr_max:
                curr_max = price
            if price < curr_min:
                curr_min = price
                curr_max = 0 
            if curr_max - curr_min > best:
                best = curr_max - curr_min
        return best
