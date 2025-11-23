from collections import Counter

class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq_dict = Counter(nums)

        ans = []
        i = 0
        for num, freq in sorted(freq_dict.items(), key = lambda x: x[1], reverse = True):
            ans.append(num)
            i+=1
            if i == k:
                break
        return ans