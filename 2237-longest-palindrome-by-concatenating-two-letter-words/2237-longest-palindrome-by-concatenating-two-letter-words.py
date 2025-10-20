class Solution:
    def longestPalindrome(self, words: List[str]) -> int:
        non_double_freq, double_freq = {}, {}
        for word in words:
            if word[0] == word[1]:
                double_freq[word] = 1 if word not in double_freq else double_freq[word] + 1
            else:
                non_double_freq[word] = 1 if word not in non_double_freq else non_double_freq[word] + 1
        
        palindrome = 0
        for word in non_double_freq:
            if non_double_freq[word] == 0: continue
            p_word = word[1] + word[0]
            if p_word in non_double_freq:
                count = min(non_double_freq[word], non_double_freq[p_word])
                non_double_freq[word] -= count
                non_double_freq[p_word] -= count
                palindrome += count * 4
        
        single, double = 0, 0
        for word in double_freq:
            if double_freq[word] % 2 == 1 and single == 0:
                single = 2
            double += double_freq[word] //2 * 2 * 2
        
        return palindrome + single + double