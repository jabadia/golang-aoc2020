package main

import (
	"aoc/utils"
	"fmt"
)

var sampleInput = []int{
	1721,
	979,
	366,
	299,
	675,
	1456,
}

func main() {
	utils.CheckSolution(fmt.Sprintf("%v", sampleInput), solve(sampleInput), 241861950)
	input := utils.ReadInts("input.txt")
	fmt.Printf("%d\n", solve(input))
}

func solve(input []int) int {
	for i, expense := range input {
		for j, expense2 := range input[i+1:] {
			for _, expense3 := range input[j+1:] {
				if expense+expense2+expense3 == 2020 {
					return expense * expense2 * expense3
				}
			}
		}
	}
	return -1
}
