#include <vector>
#include <unordered_map>
using namespace std;

class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
        unordered_map<int, int> seen; // Stores {value: index}
        
        for (int i = 0; i < nums.size(); ++i) {
            int complement = target - nums[i];
            
            // Check if the complement exists in the map
            if (seen.count(complement)) {
                return {seen[complement], i}; // Found the pair
            }
            
            // Store the current number and its index
            seen[nums[i]] = i;
        }
        
        return {}; // Problem guarantees a solution, so this is a fallback
    }
};