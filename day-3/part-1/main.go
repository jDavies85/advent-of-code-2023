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

func IsSpecialCharacter(str string) bool {
	if str == "." {
		return false
	}
	var result = nonNumericRegex.ReplaceAllString(str, "")
	return result == ""
}

func IsEnginePart(y int, x int, input [][]string) bool {
	maxY := len(input) - 1
	maxX := len(input[0]) - 1

	if x+1 <= maxX {
		//fmt.Println(input[y][x+1]) //right
		if IsSpecialCharacter(input[y][x+1]) {
			return true
		}
	}
	if x+1 <= maxX && y+1 <= maxY {
		//fmt.Println(input[y+1][x+1]) //bottom-right
		if IsSpecialCharacter(input[y+1][x+1]) {
			return true
		}
	}
	if y+1 <= maxY {
		//fmt.Println(input[y+1][x]) //bottom
		if IsSpecialCharacter(input[y+1][x]) {
			return true
		}
	}
	if y+1 <= maxY && x-1 > 0 {
		//fmt.Println(input[y+1][x-1]) //bottom-left
		if IsSpecialCharacter(input[y+1][x-1]) {
			return true
		}
	}
	if x-1 > 0 {
		//fmt.Println(input[y][x-1]) //left
		if IsSpecialCharacter(input[y][x-1]) {
			return true
		}
	}
	if x-1 > 0 && y-1 > 0 {
		//fmt.Println(input[y-1][x-1]) //top-left
		if IsSpecialCharacter(input[y-1][x-1]) {
			return true
		}
	}
	if y-1 > 0 {
		//fmt.Println(input[y-1][x]) //top
		if IsSpecialCharacter(input[y-1][x]) {
			return true
		}
	}
	if y-1 > 0 && x+1 <= maxX {
		//fmt.Println(input[y-1][x+1]) //top-right
		if IsSpecialCharacter(input[y-1][x+1]) {
			return true
		}
	}

	return false
}

func main() {
	var lines = ReadFile("../day_3_input.txt")
	var engineMap = GetEngineMap(lines)
	var numbers = FindEngineParts(engineMap)
	var sum int = 0
	for _, n := range numbers {
		sum = sum + n
	}
	fmt.Println(sum)
}

func FindEngineParts(engineMap [][]string) []int {
	var result = []int{}
	for y := 0; y < len(engineMap); y++ {
		var currentString = ""
		var isEnginePart = false
		for x := 0; x < len(engineMap[0]); x++ {
			//is this character not a special character
			if nonNumericRegex.ReplaceAllString(engineMap[y][x], "") != "" {
				//is this the first character of a new string,
				if x == 0 || (x > 0 && nonNumericRegex.ReplaceAllString(engineMap[y][x-1], "") == "") {
					currentString = engineMap[y][x]
				} else {
					currentString = currentString + engineMap[y][x]
				}

				if !isEnginePart {
					isEnginePart = IsEnginePart(y, x, engineMap)

				} else {
					if x == len(engineMap[0])-1 {
						number, err := strconv.Atoi(currentString)

						if err != nil {
							fmt.Println("Error during conversion")
						}
						result = append(result, number)
						println("CurrentString = ", currentString, " is an engine part")
					}
				}
			} else {
				if len(currentString) > 0 {
					if isEnginePart {
						number, err := strconv.Atoi(currentString)

						if err != nil {
							fmt.Println("Error during conversion")
						}
						result = append(result, number)
						println("CurrentString = ", currentString, " is an engine part")
					} else {
						println("CurrentString = ", currentString, " is not an engine part")
					}
					currentString = ""
					isEnginePart = false
				}
			}
		}
	}
	return result
}
