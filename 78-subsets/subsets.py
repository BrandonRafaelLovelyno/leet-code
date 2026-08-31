class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        tree = []
        def dfs(i, t):
            if i < len(nums):
                dfs(i+1, t+[nums[i]])
                dfs(i+1, t)
            else:
                tree.append(t)
                print(t)
                print(tree)

        dfs(0, [])

        return tree

        