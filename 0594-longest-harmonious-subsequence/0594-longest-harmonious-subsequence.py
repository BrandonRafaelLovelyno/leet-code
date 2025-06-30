class Solution:
    def findLHS(self, nums: List[int]) -> int:
        my_dict = {}
        for n in nums:
            if n in my_dict:
                my_dict[n] = my_dict[n] + 1
            else:
                my_dict[n] = 1
        
        ans = 0
        for n in my_dict:
            if n - 1 in my_dict:
                ans = max(ans, my_dict[n] + my_dict[n-1])
        
        return ans

