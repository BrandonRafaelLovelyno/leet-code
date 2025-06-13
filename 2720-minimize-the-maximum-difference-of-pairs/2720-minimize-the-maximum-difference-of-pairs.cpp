#include <vector>
#include <algorithm>

class Solution {
public:
    int minimizeMax(std::vector<int>& nums, int p) {
        std::sort(nums.begin(), nums.end());

        int n = nums.size();
        int res = 0;

        int left = 0;
        int right = nums[n - 1] - nums[0];

        while (left <= right) {
            int mid = (right + left) / 2;

            if (canFormPairs(nums, mid, p)) {
                res = mid;
                right = mid - 1;
            } else {
                left = mid + 1;
            }
        }

        return res;
    }

private:
    bool canFormPairs(const std::vector<int>& nums, int max_diff, int target_p) {
        int count = 0;
        int i = 0;

        while (i < nums.size() - 1) {
            if (nums[i+1] - nums[i] <= max_diff) {
                count++;
                i += 2;
            } else {
                i++;
            }
        }
        return count >= target_p;
    }
};