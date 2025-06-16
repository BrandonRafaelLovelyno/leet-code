#include <climits>

class Solution {
public:
    int maximumDifference(vector<int>& nums) {
        int min = nums[0];
        int max_diff = -1;

        for(int i = 0; i < nums.size();i++){
            if(nums[i] > min){
                int diff = nums[i] - min;
                if(diff > max_diff){
                    max_diff = diff;
                }
            }else{
                min = nums[i];
            }
        }

        return max_diff;
    }
};