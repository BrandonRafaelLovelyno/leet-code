class Solution:
    def lengthOfLongestSubstring(self, s: str) -> int:
        seen = set()
        longest = 0

        start, i = 0, 0
        while i < len(s):
            if s[i] not in seen:
                seen.add(s[i])
                longest = max(longest, i - start + 1)
            else:
                while s[i] in seen:
                    seen.remove(s[start])
                    start += 1
                seen.add(s[i])
            i += 1
        
        return longest

