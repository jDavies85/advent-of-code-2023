package main

import (
	file "advent_of_code/shared"
	"fmt"
	"strconv"
	"strings"
)

type RaceResult struct {
	time     int
	distance int
}

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

func mapResults(lines []string) []RaceResult {
	result := []RaceResult{}
	split := strings.Split(lines[0], ":")
	split2 := strings.Split(strings.TrimSpace(split[1]), " ")
	for i := 0; i < len(split2); i++ {
		if split2[i] != "" {
			result = append(result, RaceResult{time: getNumber(strings.TrimSpace(split2[i]))})
		}
	}
	//this is obviously all gross but whatevs, time is of the essence
	split3 := strings.Split(lines[1], ":")
	split4 := strings.Split(strings.TrimSpace(split3[1]), " ")
	j := 0
	for i := 0; i < len(split4); i++ {
		if split4[i] != "" {
			result[j].distance = getNumber(strings.TrimSpace(split4[i]))
			j++
		}
	}

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
	result := 1
	var lines = file.Lines(path)
	var races = mapResults(lines)
	for _, race := range races {
		result = result * measure(race.time, race.distance)
	}

	return result
}

func main() {
	var result = sumWaysToWin("../day_6_input.txt")
	fmt.Println(result)
}
