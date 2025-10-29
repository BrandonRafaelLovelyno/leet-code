class Solution:
    def smallestNumber(self, n: int) -> int:
        binary = format(n, 'b')
        flipped_binary = ''.join(["1" if bit == "0" else "0" for bit in binary])
        return n + int(flipped_binary, 2)