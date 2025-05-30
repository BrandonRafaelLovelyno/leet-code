class Solution:
    def check_duplicate(self, ss):
        for c in ss:
            cc = ss.count(c)
            if cc > 1:
                return True

        return False

    def lengthOfLongestSubstring(self, s):
        ans = 1 if len(s) > 0 else 0
        
        i, j = 0, 1
        while j < len(s):
            if self.check_duplicate(s[i : j + 1]):
                if j == i + 1:
                    j += 1
                else:
                    i += 1
            else:
                temp = j - i + 1
                if temp > ans:
                    ans = temp
                
                j += 1

        return ans
