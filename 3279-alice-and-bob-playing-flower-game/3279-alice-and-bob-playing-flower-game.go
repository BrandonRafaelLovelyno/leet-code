func flowerGame(n int, m int) int64 {
	nOdd := (n + 1) / 2
	nEven := n / 2
	mOdd := (m + 1) / 2
	mEven := m / 2

	return int64(nOdd*mEven + mOdd*nEven)
}