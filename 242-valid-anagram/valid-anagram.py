class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        seen = {}
        for c in s:
            if c in seen:
                seen[c]+=1
            else:
                seen[c]=1
        
        for c in t:
            if c not in seen:
                return False
            seen[c] -= 1
    
        for c in seen:
            if seen[c] != 0:
                return False
        
        return True

