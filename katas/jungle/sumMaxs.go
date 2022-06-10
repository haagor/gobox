package main

import "sort"

func sumMaxsQS(numbers []int) int {
	sort.Ints(numbers)
	return numbers[len(numbers)-1] + numbers[len(numbers)-2]
}

func sumMaxs(numbers []int) int {
	var maxs [2]int
	for _, n := range numbers {
		if n > maxs[0] {
			tmp := maxs[0]
			maxs[0] = n
			if tmp > maxs[1] {
				maxs[1] = tmp
			}
		} else if n > maxs[1] {
			maxs[1] = n
		}
	}
	return maxs[0] + maxs[1]
}
