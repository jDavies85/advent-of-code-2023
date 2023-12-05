package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

func SumWinningScores(input string) int {
	result := 0
	var gameSplit = strings.Split(input, "|")
	var cardSplit = strings.Split(gameSplit[0], ":")
	winningNumbers := ExtractScores(cardSplit[1])
	playerNumbers := ExtractScores(gameSplit[1])

	for _, item := range playerNumbers {
		if Search(winningNumbers, item) {
			if result == 0 {
				result = 1
			} else {
				result = result * 2
			}
		}
	}
	return result
}

func ExtractScores(str string) []int {
	playerNumbers := []int{}
	str = strings.TrimSpace(str)
	split := strings.Split(str, " ")
	for _, item := range split {
		item = strings.TrimSpace(item)
		if item != "" {
			number, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println("Error during conversion")
			}
			playerNumbers = append(playerNumbers, number)
		}
	}

	return playerNumbers
}

func Search(arr []int, x int) bool {
	front := 0
	back := len(arr) - 1

	for front <= back {
		if arr[front] == x || arr[back] == x {
			return true
		}
		front++
		back--
	}
	return false
}

func main() {
	lines := ReadFile("../day_4_input.txt")
	score := 0
	for _, item := range lines {
		score += SumWinningScores(item)
	}

	fmt.Println(score)
}
