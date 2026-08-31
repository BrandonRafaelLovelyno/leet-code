class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        tree, curr = [], []
        length = len(nums)

        def dfs(i):
            if not i < length:
                tree.append(curr.copy())
                return
            
            curr.append(nums[i])
            dfs(i+1)
            curr.pop()
            dfs(i+1)

        dfs(0)

        return tree

        