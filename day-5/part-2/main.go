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

type SeedRange struct {
	start int
	end   int
}

type ResourceMap struct {
	name          string
	resourceRange []ResourceRange
}

type ResourceRange struct {
	destinationRangeStart int
	sourceRangeStart      int
	rangeLength           int
}

func GetSeedRanges(str string) []SeedRange {
	result := []SeedRange{}
	split1 := strings.Split(str, ":")
	split2 := strings.Split(strings.TrimSpace(split1[1]), " ")

	for i := 0; i < len(split2); i += 2 {
		startIndex := GetNumber(split2[i])
		endIndex := startIndex + GetNumber(split2[i+1])

		result = append(result, SeedRange{start: startIndex, end: endIndex})
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

func GetDestination(src int, res []ResourceRange) int {
	for _, r := range res {
		if src >= r.sourceRangeStart && src <= (r.sourceRangeStart-1)+r.rangeLength {
			return src + (r.destinationRangeStart - r.sourceRangeStart)
		}
	}
	return src
}

func GetSeedLocation(seed int, maps []ResourceMap) int {
	location := seed
	for _, _map := range maps {
		location = GetDestination(location, _map.resourceRange)
	}
	return location
}

func GetLowestLocationNumber(seedRanges []SeedRange, maps []ResourceMap) int {
	result := 0
	for _, seedRange := range seedRanges {
		for i := seedRange.start; i < seedRange.end; i++ {
			location := GetSeedLocation(i, maps)
			if result == 0 || location < result {
				result = location
			}
		}
	}

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
	lines := ReadFile("../day_5_input.txt")
	seedRanges := GetSeedRanges(lines[0])
	maps := MapInput(lines)

	lowest := GetLowestLocationNumber(seedRanges, maps)

	fmt.Println(lowest)
}
