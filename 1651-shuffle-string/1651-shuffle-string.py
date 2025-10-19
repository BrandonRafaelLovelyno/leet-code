class Solution:
    def restoreString(self, s: str, indices: List[int]) -> str:
        chrs = [""] * len(indices)
        for c, ind in zip(s, indices):
            chrs[ind] = c
        return "".join(chrs)