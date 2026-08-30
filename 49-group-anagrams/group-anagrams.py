class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        hi_map = {}
        ans = []

        for s in strs:
            freq = {}
            for c in s:
                if c not in freq:
                    freq[c] = 1
                else:
                    freq[c] += 1
            
            sorted_keys = sorted(freq)
            hashs = ''.join(key+str(freq[key]) for key in sorted_keys)
            
            if hashs in hi_map:
                ans[hi_map[hashs]].append(s)
            else:
                hi_map[hashs] = len(ans)
                ans.append([s])
        
        return ans