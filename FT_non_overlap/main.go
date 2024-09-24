package main

import "fmt"

func Ft_non_overlap(intervals [][]int) int {
	if len(intervals) == 0 {
		return 0
	}

	nonOverlapCount := 1
	prevEnd := intervals[0][1]

	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= prevEnd {
			nonOverlapCount++
			prevEnd = intervals[i][1]
		} else {
			if intervals[i][1] < prevEnd {
				prevEnd = intervals[i][1]
			}
		}
	}
	return len(intervals) - nonOverlapCount
}

func main() {
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}, {3, 4}, {1, 3}})) // resultat : 1
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {2, 3}}))                 // resultat : 0
	fmt.Println(Ft_non_overlap([][]int{{1, 2}, {1, 2}, {1, 2}}))         // resultat : 2
}
