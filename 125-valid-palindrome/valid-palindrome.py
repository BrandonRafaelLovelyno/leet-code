class Solution:
    def isPalindrome(self, s: str) -> bool:
        clean_s = ''.join(c.lower() for c in s if c.isalnum())
        
        p_start, p_end = 0, len(clean_s) - 1

        while p_start < p_end:
            if clean_s[p_start] != clean_s[p_end]:
                return False
            p_start, p_end = p_start + 1, p_end - 1
        
        return True