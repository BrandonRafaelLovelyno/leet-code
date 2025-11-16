from collections import Counter

class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        freq_s = Counter(s)
        freq_t = Counter(t)

        if len(freq_s) != len(freq_t):
            return False 

        for c_s, f_s in freq_s.items():
            if c_s not in freq_t or f_s != freq_t[c_s]:
                return False
        return True