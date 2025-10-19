from collections import deque

class Solution:
    def apply_add(self,s,a):
        s_list = list(s)
        for i in range(1,len(s),2):
            s_list[i] = str((int(s_list[i]) + a) % 10)
        return "".join(s_list)
    
    def apply_rot(self,s,b):
        b_actual = b % len(s)
        return s[-b_actual:] + s[:-b_actual]

    def findLexSmallestString(self, s: str, a: int, b: int) -> str:
        queue, visited, min_s = deque([s]), {s}, s
        while queue :
            cur = queue.popleft()
            s_add, s_rot= self.apply_add(cur,a), self.apply_rot(cur,b)
            min_s = min(min_s, s_add, s_rot)

            if s_add not in visited:
                queue.append(s_add)
                visited.add(s_add)
            if s_rot not in visited:
                queue.append(s_rot)
                visited.add(s_rot)

        return min_s