package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var nonNumericRegex = regexp.MustCompile(`[^0-9]`)

func ReadFile(path string) []string {
	readFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileLines []string

	for fileScanner.Scan() {
		fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	return fileLines
}

func ClearString(str string) string {
	return nonNumericRegex.ReplaceAllString(str, "")
}

func ExtractNumber(str string) int {
	var clearedString = ClearString(str)
	first := string(clearedString[0:1])
	last := string(clearedString[len(clearedString)-1:])
	number, err := strconv.Atoi(first + last)

	if err != nil {
		fmt.Println("Error during conversion")
		return 0
	}

	return number
}

func main() {
	fmt.Println("Hello Day 1")
	lines := ReadFile("../day_1_input.txt")
	var numbers = []int{}
	for _, line := range lines {
		var number = ExtractNumber(line)
		numbers = append(numbers, number)
	}

	var sum int = 0
	for _, n := range numbers {
		sum = sum + n
	}
	fmt.Println(sum)
}
