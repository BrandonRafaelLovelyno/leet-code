import heapq

class Solution:
    def networkDelayTime(self, times: list[list[int]], n: int, k: int) -> int:
        def dijkstra(start, graph):
            dist = {node: 0 if node == start else float('inf') for node in graph}

            pq = [(0, start)]

            while pq:
                d, u = heapq.heappop(pq)
                
                for v, w in graph[u]:
                    new_dist = d + w
                    if new_dist < dist[v]:
                        dist[v] = new_dist
                        heapq.heappush(pq, (dist[v], v))
            
            return dist
                        

        graph = {node: [] for node in range(1, n + 1)}
        for time in times:
            u, v, w = time
            graph[u].append((v, w))
        
        dist = dijkstra(k, graph)
        
        ans = -1
        for node in range(1, n + 1):
            if dist[node] == float('inf'):
                ans = -1
                break
            ans = max(ans, dist[node])

        return ans