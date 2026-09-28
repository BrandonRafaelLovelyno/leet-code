from collections import deque

class Solution:
    def findCheapestPrice(self, n: int, flights: list[list[int]], src: int, dst: int, k: int) -> int:
        graph = {}
        for u, v, price in flights:
            graph.setdefault(u, []).append((v, price))

        best = [float('inf')] * n
        ans = float('inf')
        queue = deque([(src, 0, 0)]) 

        while queue:
            u, cost, stops = queue.popleft()

            if stops > k:
                continue

            for v, price in graph.get(u, []):
                new_cost = cost + price
                if new_cost >= best[v]:
                    continue
                best[v] = new_cost
                if v == dst:
                    ans = min(ans, new_cost)
                else:
                    queue.append((v, new_cost, stops + 1))

        return ans if ans != float('inf') != -1 else -1