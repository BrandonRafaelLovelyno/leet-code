class Solution:
    def findKDistantIndices(self, nums: List[int], key: int, k: int) -> List[int]:
        ans = []

        j = 0
        for i, n in enumerate(nums):
            if n == key:
                j = max(j, i - k)
                while j <= i + k and j < len(nums):
                    ans.append(j)
                    j += 1

        return ans
