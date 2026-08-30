class Solution:
    def isPalindrome(self, s: str) -> bool:
        clean_s = ''.join(c.lower() for c in s if c.isalnum())
        
        stack = [clean_s[i] for i in range(len(clean_s)//2)]

        start_i = len(clean_s) // 2
        if len(clean_s) % 2 != 0:
            start_i += 1

        for i in range(start_i, len(clean_s)):
            if clean_s[i] != stack.pop():
                return False
        
        return True