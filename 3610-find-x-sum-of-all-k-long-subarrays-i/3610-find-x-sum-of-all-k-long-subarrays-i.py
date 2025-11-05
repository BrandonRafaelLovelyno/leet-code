from collections import Counter

class Solution:
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        n = len(nums)
        freq = Counter(nums[:k])

        ans = []
        for i in range(n-k+1):
            top = sorted(freq.items(), key = lambda i: (i[1],i[0]), reverse=True)[:x]
            ans.append(sum(f * n for f, n in top))
            
            if i + k < n:
                left, new = nums[i], nums[i+k]
                freq[left] -= 1
                freq[new] = 1 if new not in freq else freq[new] + 1

        return ans