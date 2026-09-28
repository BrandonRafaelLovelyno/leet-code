class Solution:
    def topKFrequent(self, nums: list[int], k: int) -> list[int]:
        freq = {}
        for n in nums:
            if n not in freq:
                freq[n] = 1
            else:
                freq[n] += 1
        
        bucket = [[] for _ in range(len(nums) + 1)]
        for n in freq:
            bucket[freq[n]].append(n)

        ans = []
        for i in range(len(nums), -1, -1):
            if not bucket[i]:
                continue
            ans.extend(bucket[i])

        return ans[:k]
            


        

            