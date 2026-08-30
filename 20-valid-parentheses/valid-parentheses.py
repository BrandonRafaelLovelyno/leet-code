class Solution:
    def isValid(self, s: str) -> bool:
        if len(s) == 1 :
            return False

        stack = []
        opening = ["(", "{", "["]

        for c in s:
            if c in opening:
                stack.append(c)
            else:
                if len(stack) == 0 or not self.isValidClosing(stack[-1], c):
                    return False
                stack.pop()

        return len(stack) == 0

    def isValidClosing(self, s, t):
        if s == "(":
            return t == ")"
        elif s == "{":
            return t == "}"
        elif s == "[":
            return t == "]"
        return False