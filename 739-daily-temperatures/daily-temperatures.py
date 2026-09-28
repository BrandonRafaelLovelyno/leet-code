class Solution:
    def dailyTemperatures(self, temperatures: list[int]) -> list[int]:
        ans = [0 for _ in range(len(temperatures))]

        stack = []
        for i, temp in enumerate(temperatures):
            if stack and temp > stack[-1][1]:
                while stack and stack[-1][1] < temp:
                    j, t = stack.pop()
                    ans[j] = i - j
            stack.append((i, temp))
        
        return ans
