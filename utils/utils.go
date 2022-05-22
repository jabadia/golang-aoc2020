package utils

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CheckSolution(input string, actualSolution int, expectedSolution int) {
	inputRunes := []rune(input)
	var inputStr string
	if len(inputRunes) > 50 {
		inputStr = string(inputRunes[:44]) + "..." + string(inputRunes[len(inputRunes)-3:])
	} else {
		inputStr = string(inputRunes)
	}
	fmt.Printf("%s = ", inputStr)
	if actualSolution == expectedSolution {
		fmt.Printf("OK %d", actualSolution)
	} else {
		fmt.Printf("FAIL %d != %d", actualSolution, expectedSolution)
	}
	fmt.Printf("\n")
}

func ReadInts(filename string) (numbers []int) {
	f, err := os.Open("input.txt")
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filename, err))
	}
	var line int
	for {
		_, err := fmt.Fscanf(f, "%d\n", &line)
		if err == io.EOF {
			return numbers
		}
		if err != nil {
			panic(fmt.Sprintf("Fscanf %v", err))
		}
		numbers = append(numbers, line)
	}
}

func ReadLines(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(fmt.Sprintf("open %s: %v", filename, err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if scanner.Err() != nil {
		panic(fmt.Sprintf("scanner %v", scanner.Err()))
	}
	return lines
}
