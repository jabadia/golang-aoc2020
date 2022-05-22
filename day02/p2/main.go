package main

import (
	"aoc/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	sample := utils.ReadLines("sample.txt")
	utils.CheckSolution(fmt.Sprintf("%v", sample), solve(sample), 1)
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
		p0, _ := strconv.Atoi(parts[0])
		p1, _ := strconv.Atoi(parts[1])
		if (password[p0-1] == letter[0]) != (password[p1-1] == letter[0]) {
			validPasswords++
		}
	}
	return validPasswords
}
