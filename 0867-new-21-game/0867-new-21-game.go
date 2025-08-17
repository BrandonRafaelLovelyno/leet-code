func new21Game(n int, k int, maxPts int) float64 {
	if k == 0 {
		return 1.0
	}

	dp,sw := make([]float64,n+1),1.0
	dp[0] = 1.0

	for i := 1; i <= n; i++{
		dp[i] = sw / float64(maxPts)
		
		if i < k{
			sw += dp[i]
		}

		if i - maxPts >= 0 {
			 if i - maxPts < k {
                sw -= dp[i - maxPts]
            }
		}
	} 

	ans := 0.0 
	for i := k ; i <= n ;i++{
		ans += dp[i]
	} 
	
	return ans
}