class Solution:
    def kthCharacter(self, k: int) -> str:
        seq = [0]
        while len(seq) < k:
            shift = []
            for s in seq:
                next_s = (s + 1) % 26
                shift += [next_s]
            seq += shift
            print(seq)
        
        return chr(seq[k - 1] + 97)
