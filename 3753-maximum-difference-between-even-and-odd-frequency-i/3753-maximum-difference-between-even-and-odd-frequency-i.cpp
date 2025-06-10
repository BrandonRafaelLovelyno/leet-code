#include <string>
#include <limits>
#include <algorithm>
#include <cmath>

class Solution {
public:
    int maxDifference(std::string s) {
        int chars[26] = {0};
        
        for (char ch : s) {
            int index = ch - 'a';
            if (index >= 0 && index < 26) {
                chars[index]++;
            }
        }

        int odd = 0;
        int even = std::numeric_limits<int>::max();

        for (int i = 0; i < 26; i++) {
            int current_freq = chars[i];
            
            if (current_freq == 0) {
                continue;
            }

            if (current_freq % 2 == 1) {
                odd = std::max(odd, current_freq);
            } else {
               even = std::min(even, current_freq);
            }
        }
        
        return odd - even;
    }
};