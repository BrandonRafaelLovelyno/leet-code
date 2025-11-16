from collections import Counter

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        freq = {}
        for num in nums:
            freq[num] = 1 if num not in freq else freq[num] + 1
            if freq[num] > 1:
                return True
        return False