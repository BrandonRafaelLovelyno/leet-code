class Solution:
    def search(self, nums: list[int], target: int) -> int:
        n = len(nums)
        left, right, k = 0, n - 1, -1
        while left <= right:
            mid = (left + right) // 2

            if mid + 1 < n and nums[mid] > nums[mid + 1]:
                k = mid
                break

            if nums[mid] <= nums[n - 1]:
                right = mid - 1
            else:
                left = mid + 1
        
        left, right = 0, n - 1
        if nums[right] >= target:
            left = k + 1
        else:
            right = k

        ans = -1
        while left <= right:
            mid = (left + right) // 2
            if nums[mid] > target:
                right = mid - 1
            elif nums[mid] < target:
                left = mid + 1
            else:
                ans = mid
                break

        return ans
