class Solution:
    def averageOfSubtree(self, root: TreeNode) -> int:
        ans = 0

        def traverse(node):
            nonlocal ans

            if not node.left and not node.right:
                ans += 1
                return (node.val, 1)

            val, count = node.val, 1

            if node.left:
                child_val, child_count = traverse(node.left)
                val += child_val
                count += child_count

            if node.right:
                child_val, child_count = traverse(node.right)
                val += child_val
                count += child_count

            if val // count == node.val:
                ans += 1

            return (val, count)

        traverse(root)

        return ans