package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sample := utils.ReadLines("sample.txt")
	utils.CheckSolution(fmt.Sprintf("%v", sample), solve(sample), 2)
	input := utils.ReadLines("input.txt")
	fmt.Printf("%d\n", solve(input))
}

func solve(input []string) int {
	var validPasswords int
	for _, line := range input {
		parts := strings.Split(line, ": ")
		policy, password := parts[0], parts[1]
		parts = strings.Split(policy, " ")
		minMax, letter := parts[0], parts[1]
		parts = strings.Split(minMax, "-")
		min, _ := strconv.Atoi(parts[0])
		max, _ := strconv.Atoi(parts[1])
		// fmt.Printf("[%d-%d] [%s] [%s]\n", min, max, letter, password)
		count := strings.Count(password, letter)
		if min <= count && count <= max {
			validPasswords++
		}
	}
	return validPasswords
}
