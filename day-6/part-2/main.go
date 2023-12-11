package main

import (
	file "advent_of_code/shared"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type RaceResult struct {
	time     int
	distance int
}

var whitespaceRegex = regexp.MustCompile(`/\s/g`)

func measure(time int, distance int) int {
	result := 0
	speed := 0
	for i := 0; i < time; i++ {
		totalDistanceTravelled := (time - i) * speed
		if totalDistanceTravelled > distance {
			result++
		}
		speed++
	}

	return result
}

func mapResult(lines []string) RaceResult {
	result := RaceResult{}
	split := strings.Split(lines[0], ":")
	split2 := strings.Replace(split[1], " ", "", -1)

	result.time = getNumber(split2)

	split3 := strings.Split(lines[1], ":")
	split4 := strings.Replace(split3[1], " ", "", -1)

	result.distance = getNumber(split4)

	return result
}

func getNumber(str string) int {
	number, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error during conversion")
	}
	return number
}

func sumWaysToWin(path string) int {
	var lines = file.Lines(path)
	var race = mapResult(lines)
	result := measure(race.time, race.distance)

	return result
}

func main() {
	var result = sumWaysToWin("../day_6_input.txt")
	fmt.Println(result)
}
