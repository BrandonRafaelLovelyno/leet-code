class Solution:
    def productExceptSelf(self, nums: List[int]) -> List[int]:
        n = len(nums)
        ans = [1] * n

        for i in range(1, n):
            ans[i] = ans[i-1] * nums[i-1]
        
        running_product = nums[n-1]
        for i in range(n-2, -1, -1):
            ans[i] *= running_product
            running_product *= nums[i]

        return ans