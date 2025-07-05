class Solution:
    def findLucky(self, arr: List[int]) -> int:
        my_dict = {}
        for n in arr:
            if n in my_dict:
                my_dict[n] += 1
            else:
                my_dict[n] = 1
        
        ans = -1
        for n, freq in my_dict.items():
            if n == freq and n > ans:
                ans = n
        
        return ans