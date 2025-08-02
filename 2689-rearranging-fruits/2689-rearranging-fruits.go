package main

import "sort"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func minCost(basket1 []int, basket2 []int) int64 {
	minCostVal := int(1e9 + 7)
	dict1 := make(map[int]int)
	for _, cost := range basket1 {
		dict1[cost]++
		if cost < minCostVal {
			minCostVal = cost
		}
	}

	dict2 := make(map[int]int)
	for _, cost := range basket2 {
		dict2[cost]++
		if cost < minCostVal {
			minCostVal = cost
		}
	}

	exit1, exit2 := []int{}, []int{}
	for cost1, freq1 := range dict1 {
		freq2, _ := dict2[cost1]
		if freq1 == freq2 {
			continue
		}

		if (freq1+freq2)%2 != 0 {
			return -1
		}

		if freq1 > freq2 {
			exit := (freq1 - freq2) / 2
			for i := 0; i < exit; i++ {
				exit1 = append(exit1, cost1)
			}
		}
	}

	for cost2, freq2 := range dict2 {
		freq1, _ := dict1[cost2]
		if freq2 == freq1 {
			continue
		}

		if (freq1+freq2)%2 != 0 {
			return -1
		}

		if freq2 > freq1 {
			exit := (freq2 - freq1) / 2
			for i := 0; i < exit; i++ {
				exit2 = append(exit2, cost2)
			}
		}
	}

	if len(exit1) != len(exit2) {
		return -1
	}

	sort.Ints(exit1)
	sort.Sort(sort.Reverse(sort.IntSlice(exit2)))

	var ans int64 = 0
	i := 0
	for i < len(exit1) {
		ans += int64(min(min(exit1[i], exit2[i]), minCostVal*2))
		i++
	}

	return ans
}