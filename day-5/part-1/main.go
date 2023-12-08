package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var nonNumericRegex = regexp.MustCompile(`[^0-9]`)

type ResourceMap struct {
	name          string
	resourceRange []ResourceRange
}

type ResourceRange struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func GetSeeds(str string) []int {
	result := []int{}
	split1 := strings.Split(str, ":")
	split2 := strings.Split(strings.TrimSpace(split1[1]), " ")
	for _, item := range split2 {
		result = append(result, GetNumber(item))
	}

	return result
}

func MapInput(lines []string) []ResourceMap {
	result := []ResourceMap{}
	for i := 2; i < len(lines)-1; i++ {
		if lines[i] != "" {
			var numberString = nonNumericRegex.ReplaceAllString(lines[i], "")
			if numberString == "" {
				result = append(result, ResourceMap{name: strings.Split(lines[i], " ")[0]})
			} else {
				rr := GetResourceRange(lines[i])
				result[len(result)-1].resourceRange = append(result[len(result)-1].resourceRange, rr)
			}
		}
	}
	return result
}

func GetResourceRange(str string) ResourceRange {
	result := ResourceRange{}
	split := strings.Split(str, " ")
	result.destinationRangeStart = GetNumber(split[0])
	result.sourceRangeStart = GetNumber(split[1])
	result.rangeLength = GetNumber(split[2])
	return result
}

func GetNumber(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error during conversion")
	}
	return number
}

func GetSeedLocation(seed int, maps []ResourceMap) int {
	result := 0

	return result
}

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

func main() {
	//lines := ReadFile("../test_input.txt")
	//seeds := GetSeeds(lines[0])
	// maps := MapInput(lines)
	// for _, seed := range seeds {

	// }
}
