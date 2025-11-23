from collections import Counter


def create_max_heap(heap):
    n = len(heap)
    for i in range(n // 2 - 1, -1, -1):
        heapify(heap, i)


def heapify(heap, i):
    biggest = i
    n = len(heap)

    if 2 * i + 1 < n and heap[biggest][1] < heap[2 * i + 1][1]:
        biggest = 2 * i + 1
    if 2 * i + 2 < n and heap[biggest][1] < heap[2 * i + 2][1]:
        biggest = 2 * i + 2

    if biggest != i:
        heap[i], heap[biggest] = heap[biggest], heap[i]
        heapify(heap, biggest)


def pop_heap(heap):
    root = heap[0]
    last = heap.pop()

    if heap:
        heap[0] = last
        heapify(heap, 0)
    
    return root


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        freq_dict = Counter(nums)
        heap = list(freq_dict.items())
        create_max_heap(heap)

        ans = []
        for i in range(k):
            ans.append(pop_heap(heap)[0])

        return ans
