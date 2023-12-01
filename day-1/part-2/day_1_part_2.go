package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var numberStrings = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

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

func TryParseInt(str string) (int, bool) {
	number, err := strconv.Atoi(str)

	if err != nil {
		return 0, false
	}
	return number, true
}

func TryParseIntText(str string) (int, bool) {
	var parsed = 0
	for i, element := range numberStrings {
		if strings.HasPrefix(str, element) {
			parsed = i + 1
			return parsed, true
		}
	}
	return parsed, false
}

func ExtractNumber(str string) int {
	mySlice := []string{}
	var currentString = str
	for len(currentString) > 0 {
		parsed, success := TryParseInt(currentString[:1])

		if success {
			mySlice = append(mySlice, strconv.Itoa(parsed))
			currentString = currentString[1:]
		} else {
			parsed2, success2 := TryParseIntText(currentString)
			if success2 {
				mySlice = append(mySlice, strconv.Itoa(parsed2))
				currentString = currentString[1:]
			} else {
				currentString = currentString[1:]
			}
		}
	}

	number, err := strconv.Atoi(mySlice[0] + mySlice[len(mySlice)-1])
	if err != nil {
		return 0
	}

	return number
}

func ProcessInput(filePath string) int {
	lines := ReadFile(filePath)
	var numbers = []int{}
	for _, line := range lines {
		var number = ExtractNumber(line)
		numbers = append(numbers, number)
	}

	var sum int = 0
	for _, n := range numbers {
		sum = sum + n
	}
	return sum
}

func main() {
	fmt.Println("Hello Day 1 Part 2")

	var sum = ProcessInput("../day_1_input.txt")

	fmt.Println(sum)
}
