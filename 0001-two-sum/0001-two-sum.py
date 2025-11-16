class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        nums_w_index = []
        for i, num, in enumerate(nums):
            nums_w_index.append((num, i))
        nums_w_index.sort(key=lambda n: n[0])

        n = len(nums_w_index)
        left, right = 0, n - 1
        ans = []
        while left < right:
            total = nums_w_index[right][0] + nums_w_index[left][0] 
            if total < target:
                left+=1
            elif total > target:
                right -= 1
            else:
                ans.append(nums_w_index[right][1])
                ans.append(nums_w_index[left][1])
                break
        return ans

