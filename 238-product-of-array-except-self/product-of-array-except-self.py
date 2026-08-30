class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        ans = [1]

        for i in range(1, len(nums)):
            ans.append(ans[-1] * nums[i-1])
        
        suffix = 1
        for i in range(len(nums) - 2, -1, -1):
            suffix *= nums[i+1]
            ans[i] *= suffix
        
        return ans