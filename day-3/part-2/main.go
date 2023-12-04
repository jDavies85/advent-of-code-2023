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

func GetEngineMap(lines []string) [][]string {
	var result = [][]string{}
	for _, line := range lines {
		var tempSlice = []string{}
		for _, char := range line {
			tempSlice = append(tempSlice, string(char))
		}
		result = append(result, tempSlice)
	}

	return result
}

func IsNumber(str string) bool {
	var result = nonNumericRegex.ReplaceAllString(str, "")
	return result != ""
}

func CaptureNumber(str []string, index int) int {
	startingIndex := 0
	for i := index; i >= 0; i-- {
		if IsNumber(str[i]) {
			//fmt.Println(str[i])
			if i == 0 {
				startingIndex = i
			}
		} else {
			startingIndex = i + 1
			break
		}
	}
	numberString := ""
	for i := startingIndex; i < len(str); i++ {
		if IsNumber(str[i]) {
			//fmt.Println(str[i])
			numberString = numberString + str[i]
		} else {
			break
		}
	}
	number, err := strconv.Atoi(numberString)

	if err != nil {
		fmt.Println("Error during conversion")
	}
	return number
}

func CalculateGearPower(y int, x int, input [][]string) int {
	maxY := len(input) - 1
	maxX := len(input[0]) - 1
	adjacentParts := []int{}

	if x+1 <= maxX {
		if IsNumber(input[y][x+1]) {
			number := CaptureNumber(input[y], x+1)
			if len(adjacentParts) == 0 || adjacentParts[len(adjacentParts)-1] != number {
				adjacentParts = append(adjacentParts, number)
			}
		}
	}
	if x+1 <= maxX && y+1 <= maxY {
		if IsNumber(input[y+1][x+1]) {
			number := CaptureNumber(input[y+1], x+1)
			if len(adjacentParts) == 0 || adjacentParts[len(adjacentParts)-1] != number {
				adjacentParts = append(adjacentParts, number)
			}
		}
	}
	if y+1 <= maxY {
		if IsNumber(input[y+1][x]) {
			number := CaptureNumber(input[y+1], x)
			if len(adjacentParts) == 0 || adjacentParts[len(adjacentParts)-1] != number {
				adjacentParts = append(adjacentParts, number)
			}
		}
	}
	if y+1 <= maxY && x-1 >= 0 {
		if IsNumber(input[y+1][x-1]) {
			number := CaptureNumber(input[y+1], x-1)
			if len(adjacentParts) == 0 || adjacentParts[len(adjacentParts)-1] != number {
				adjacentParts = append(adjacentParts, number)
			}
		}
	}
	if x-1 >= 0 {
		if IsNumber(input[y][x-1]) {
			number := CaptureNumber(input[y], x-1)
			if len(adjacentParts) == 0 || adjacentParts[len(adjacentParts)-1] != number {
				adjacentParts = append(adjacentParts, number)
			}
		}
	}
	if x-1 >= 0 && y-1 >= 0 {
		if IsNumber(input[y-1][x-1]) {
			number := CaptureNumber(input[y-1], x-1)
			if len(adjacentParts) == 0 || adjacentParts[len(adjacentParts)-1] != number {
				adjacentParts = append(adjacentParts, number)
			}
		}
	}
	if y-1 >= 0 {
		if IsNumber(input[y-1][x]) {
			number := CaptureNumber(input[y-1], x)
			if len(adjacentParts) == 0 || adjacentParts[len(adjacentParts)-1] != number {
				adjacentParts = append(adjacentParts, number)
			}
		}
	}
	if y-1 >= 0 && x+1 <= maxX {

		if IsNumber(input[y-1][x+1]) {
			number := CaptureNumber(input[y-1], x+1)
			if len(adjacentParts) == 0 || adjacentParts[len(adjacentParts)-1] != number {
				adjacentParts = append(adjacentParts, number)
			}
		}
	}
	if len(adjacentParts) > 1 {
		return adjacentParts[0] * adjacentParts[1]
	}
	return 0
}

func main() {
	var lines = ReadFile("../day_3_input.txt")
	var engineMap = GetEngineMap(lines)
	var numbers = FindGears(engineMap)
	var sum int = 0
	for _, n := range numbers {
		sum = sum + n
	}
	fmt.Println(sum)
}

func FindGears(engineMap [][]string) []int {
	var result = []int{}
	for y := 0; y < len(engineMap); y++ {
		for x := 0; x < len(engineMap[0]); x++ {
			if engineMap[y][x] == "*" {
				var gearPower = CalculateGearPower(y, x, engineMap)
				if gearPower > 0 {
					result = append(result, gearPower)
				}
				//fmt.Println(gearPower)
			}
		}
	}
	return result
}
