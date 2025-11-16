from collections import Counter

class Solution:
    def containsDuplicate(self, nums: List[int]) -> bool:
        freq_dict = Counter(nums)

        for num, freq in freq_dict.items():
            if freq > 1:
                return True
        
        return False