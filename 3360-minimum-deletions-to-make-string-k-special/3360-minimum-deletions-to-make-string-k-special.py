import sys


class Solution:
    def minimumDeletions(self, word: str, k: int) -> int:
        freq = {}
        for char in word:
            freq[char] = freq.get(char, 0) + 1

        res = sys.maxsize
        for base_freq in freq.values():
            deleted = 0
            for f in freq.values():
                if f > base_freq + k:
                    deleted += f - (base_freq + k)
                elif f < base_freq:
                    deleted += f
            res = min(res, deleted)

        return res
