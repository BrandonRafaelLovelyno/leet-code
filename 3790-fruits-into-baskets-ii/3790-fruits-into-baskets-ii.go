func numOfUnplacedFruits(fruits []int, baskets []int) int {
    placed := 0
    for _, f := range fruits{
        for j, b := range baskets{
            if f <= b {
                placed = placed + 1
                baskets[j] = -1
                break
            }
        }
    }

    return len(fruits) - placed
}