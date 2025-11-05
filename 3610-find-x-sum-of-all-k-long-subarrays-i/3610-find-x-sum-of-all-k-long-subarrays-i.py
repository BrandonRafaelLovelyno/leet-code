class Solution:
    def findXSum(self, nums: List[int], k: int, x: int) -> List[int]:
        freq = {}
        answer = []
        left, right = 0, k-1

        while right < len(nums):
            if left == 0:
                for i in range(k):
                    freq[nums[i]] = 1 if nums[i] not in freq else freq[nums[i]] + 1
            else:
                freq[nums[left - 1]] -= 1
                freq[nums[right]] = 1 if nums[right] not in freq else freq[nums[right]] + 1
            count, ans = 0, 0
            for k, v in sorted(freq.items(), key= lambda i: (i[1],i[0]), reverse=True):
                if count == x: break
                ans += v * k
                count += 1
            answer.append(ans)
            right+=1
            left+=1

        return answer