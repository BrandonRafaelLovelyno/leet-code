class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        chrs = [""] * len(indices)
        for i in range(len(indices)):
            chrs[indices[i]] = s[i]
        return "".join(chrs)