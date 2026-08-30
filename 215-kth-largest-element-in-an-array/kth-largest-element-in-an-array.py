class Solution:
    def findKthLargest(self, nums: List[int], k: int) -> int:
        heap = self.buildHeap(nums[:k])

        for i in range(k, len(nums)):
            if nums[i] > heap[0]:
                heap[0] = nums[i]
                self.heapify(heap, 0, len(heap))
        
        return heap[0]

    def buildHeap(self, heap):
        for i in range(len(heap)//2, -1 , -1):
            self.heapify(heap, i, len(heap))
        return heap

    def heapify(self, heap, i, n):
        smallest = i
        left = 2 * i + 1
        right = 2 * i + 2

        if left < n and heap[left] < heap[smallest]:
            smallest = left
        if right < n and heap[right] < heap[smallest]:
            smallest = right
        
        if smallest != i:
            heap[i], heap[smallest] = heap[smallest], heap[i]
            self.heapify(heap, smallest, n)