from collections import deque

class Solution:
    def orangesRotting(self, grid: list[list[int]]) -> int:
        row, col = len(grid), len(grid[0])
        visited = [[False for _ in range(col)] for _ in range(row)]

        queue, orange = deque([]), 0
        for y in range(row):
            for x in range(col):
                if grid[y][x] == 2:
                    queue.append((x, y, 0))
                if grid[y][x] == 1:
                    orange += 1
        
        ans, done = 0, 0
        while queue:
            x, y, time = queue.popleft()

            if visited[y][x]:
                continue
            
            ans = max(ans, time)
            visited[y][x] = True

            if grid[y][x] == 1:
                done += 1

            if y - 1 >= 0 and grid[y - 1][x] == 1:
                queue.append((x, y - 1, time + 1))
            if y + 1 < row and grid[y + 1][x] == 1:
                queue.append((x, y + 1, time + 1))
            if x - 1 >= 0 and grid[y][x - 1] == 1:
                queue.append((x - 1, y, time + 1))
            if x + 1 < col and grid[y][x + 1] == 1:
                queue.append((x + 1, y, time + 1))

        return ans if done == orange else -1

