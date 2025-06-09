class Solution {
public:
    int findKthNumber(int n, int k) {
        int curr = 1;
        k--;

        while(k > 0){
            long long steps = countSteps(curr, n);

            if(steps <= k){
                curr++;
                k = k - steps;
            }else{
                curr = curr * 10;
                k--;
            }
        }

        return curr;
    }
private:
    long long countSteps(long long curr, long long n){
        long long steps = 0;

        long long first = curr;
        long long last = curr + 1;
        while(first <= n){
            steps += std::min(n + 1, last) - first;
            first = first * 10;
            last = last * 10;
        }

        return steps;
    }
};