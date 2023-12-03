package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game struct {
	gameNumber   int
	roundResults []RoundResult
}

type RoundResult struct {
	red   int
	green int
	blue  int
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

func GameResults(input string) []string {
	var result = strings.Split(input, ":")
	result[1] = strings.TrimSpace(result[1])

	return result
}

func RoundResults(input string) []RoundResult {
	var result = []RoundResult{}

	var split = strings.Split(input, ":")
	split[1] = strings.TrimSpace(split[1])

	var rounds = strings.Split(split[1], ";")
	for _, round := range rounds {
		var roundResult = GetRoundResult(round)
		result = append(result, roundResult)
	}
	return result
}

func GetRoundResult(input string) RoundResult {
	var result = RoundResult{}
	var scores = strings.Split(input, ",")
	for _, score := range scores {
		score = strings.TrimSpace(score)
		var scoreSplit = strings.Split(score, " ")
		number, err := strconv.Atoi(scoreSplit[0])
		if err != nil {
			fmt.Println("Error during conversion")
		}
		if scoreSplit[1] == "red" {
			result.red = number
		}
		if scoreSplit[1] == "green" {
			result.green = number
		}
		if scoreSplit[1] == "blue" {
			result.blue = number
		}
	}
	return result
}

func IsValidGame(roundResults []RoundResult, maxCubes RoundResult) bool {
	for _, round := range roundResults {
		if round.red > maxCubes.red || round.green > maxCubes.green || round.blue > maxCubes.blue {
			return false
		}
	}
	return true
}

func SumValidGames(filePath string) int {
	//set scores
	var maxCubes = RoundResult{red: 12, green: 13, blue: 14}

	//parse file
	var lines = ReadFile(filePath)
	//get list of games with round results
	var games = []Game{}
	for i, line := range lines {
		var roundresults = RoundResults(line)
		games = append(games, Game{gameNumber: i + 1, roundResults: roundresults})
	}
	//loop through games and add valid games to valid game list
	var validGames = []int{}
	for _, game := range games {
		if IsValidGame(game.roundResults, maxCubes) {
			validGames = append(validGames, game.gameNumber)
		}
	}
	//sum game numbers
	var sum int = 0
	for _, n := range validGames {
		sum = sum + n
	}

	return sum
}

func main() {
	fmt.Println("Hello Day 2")
	var sum = SumValidGames("../day_2_input.txt")
	//var sum = SumValidGames("test_input.txt")
	fmt.Println(sum)
}
