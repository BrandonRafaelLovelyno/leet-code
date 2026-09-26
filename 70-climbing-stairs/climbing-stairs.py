class Solution:
    def climbStairs(self, n: int) -> int:
        memo = [-1 for _ in range(n)]
        def dp(i):
            if i < n and i in memo and memo[i] != -1:
                return memo[i]
            elif i == n:
                return 1
            elif i > n:
                return 0

            memo[i] = dp(i+1) + dp(i+2)

            return memo[i]
        
        dp(0)

        return memo[0]