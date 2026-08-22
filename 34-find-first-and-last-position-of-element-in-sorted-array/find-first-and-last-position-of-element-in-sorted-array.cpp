class Solution {
public:
    vector<int> searchRange(vector<int>& nums, int target) {
        int left = 0, right = nums.size() - 1;
        int mid = -1, first_idx = -1;

        while (left <= right) {
            mid = (right + left) / 2;
            if (nums[mid] >= target) {
                if (nums[mid] == target) {
                    first_idx = mid;
                }
                right = mid - 1;
            } else if (nums[mid] < target) {
                left = mid + 1;
            }
        }

        left = 0, right = nums.size() - 1;
        int last_idx = -1;

        while (left <= right) {
            mid = (right + left) / 2;
            if (nums[mid] <= target) {
                if (nums[mid] == target) {
                    last_idx = mid;
                }
                left = mid + 1;
            } else if (nums[mid] > target) {
                right = mid - 1;
            }
        }

        return {first_idx, last_idx};
    }
};