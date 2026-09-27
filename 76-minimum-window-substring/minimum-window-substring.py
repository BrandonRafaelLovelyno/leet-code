class Solution:
    def minWindow(self, s: str, t: str) -> str:
        freq_t, freq = {}, {}
        for c in t:
            if c not in freq_t:
                freq_t[c] = 1
                freq[c] = 0
            else:
                freq_t[c] += 1

        left, right, ans_left, ans_right = 0, 0, 0, -1
        required, confirmed = len(freq_t), 0
        
        while right < len(s):
            c = s[right]
            if c not in freq:
                right += 1
                continue

            freq[c] += 1
            if freq[c] == freq_t[c]:
                confirmed += 1
            
            if confirmed == required:                
                while left <= right and confirmed == required:
                    if ans_right == -1 or right - left < ans_right - ans_left:
                        ans_left, ans_right = left, right
                    
                    if s[left] not in freq:
                        left += 1
                        continue

                    freq[s[left]] -= 1
                    if freq[s[left]] < freq_t[s[left]]:
                        confirmed -= 1
                    left += 1

            right += 1

        return s[ans_left:ans_right+1]
