class Solution:
    def maxDistinctElements(self, nums: List[int], k: int) -> int:
        freq = {}
        for n in nums:
            if n in freq:
                freq[n] = freq[n] + 1
            else:
                freq[n] = 1
        
        ans, high = len(nums), float('-inf')
        for n in sorted(freq.keys()):
            diff = min(n+k-high-1, 2*k)
            if freq[n] > diff + 1:
                ans -= freq[n] - (diff + 1)
            
            if high >= n-k:
                high = min(high+freq[n],n+k)
            else:
                high = min(n-k-1+freq[n],n+k)

        return ans
        
        