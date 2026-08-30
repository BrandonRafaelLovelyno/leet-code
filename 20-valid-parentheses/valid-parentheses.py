class Solution:
    def isValid(self, s: str) -> bool:
        stack = []
        brackets = {
            ")": "(",
            "}": "{",
            "]": "["
        }

        for c in s:
            if c in brackets:
                if len(stack) == 0 or brackets[c] != stack[-1]:
                    return False
                stack.pop()
            else:
                stack.append(c)

        return len(stack) == 0