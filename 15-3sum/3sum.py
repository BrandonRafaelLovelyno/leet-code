class Solution:
    def threeSum(self, nums: list[int]) -> list[list[int]]:
        ans = []
        nums.sort()

        for i in range(0, len(nums) - 2):
            if i > 0 and nums[i - 1] == nums[i]:
                continue

            p1, p2 = i + 1, len(nums) - 1

            target = -nums[i]
            while p1 < p2:
                if nums[p1] + nums[p2] < target:
                    p1+=1
                elif nums[p1] + nums[p2] > target:
                    p2-=1
                else:
                    if len(ans) == 0 or [nums[i], nums[p1], nums[p2]] != ans[-1]:
                        ans.append([nums[i], nums[p1], nums[p2]])
                    p2 -= 1
        
        return ans


