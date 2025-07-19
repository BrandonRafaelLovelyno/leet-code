import (
    "sort"
)

func numRescueBoats(people []int, limit int) int {
    sorted_people := people
    sort.Ints(sorted_people)
    
    p1, p2 := 0, len(people) - 1
    ans := 0

    for p1 <= p2 {
        if sorted_people[p2] + sorted_people[p1] > limit{
            p2--
        }else{
            p2--
            p1++
        }
        ans++;
    }

    return ans
}