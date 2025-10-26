class Solution:
    def totalMoney(self, n: int) -> int:
        weeks, days = divmod(n, 7)
        sum_weeks, sum_days = 28 * weeks + 7 * (weeks - 1) * weeks // 2, 0
        for i in range(0, days):
            sum_days += 1 + weeks + i
        return sum_weeks + sum_days
