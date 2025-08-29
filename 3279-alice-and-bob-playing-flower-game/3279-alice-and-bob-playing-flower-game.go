func flowerGame(n int, m int) int64 {
    nOdd, nEven, mOdd, mEven := 0,0,0,0

    if n % 2 == 1{
        nOdd = n/2 + 1
    }else{
        nOdd = n/2
    }
    nEven = n - nOdd

    if m % 2 == 1 {
        mOdd = m/2 + 1
    }else{
        mOdd = m/2
    }
    mEven = m - mOdd

    return int64(nOdd * mEven + mOdd * nEven)
}