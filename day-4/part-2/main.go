package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ScratchCard struct {
	cardNumber     int
	winningNumbers []int
	playerNumbers  []int
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

func GetScratchCards(lines []string) []ScratchCard {
	result := []ScratchCard{}
	for _, item := range lines {
		gameSplit := strings.Split(item, "|")
		cardSplit := strings.Split(gameSplit[0], ":")
		cardNumber := strings.Split(cardSplit[0], " ")
		number, err := strconv.Atoi(cardNumber[len(cardNumber)-1])
		if err != nil {
			fmt.Println("Error during conversion")
		}
		winningNumbers := ExtractScores(cardSplit[1])
		playerNumbers := ExtractScores(gameSplit[1])
		result = append(result, ScratchCard{cardNumber: number, winningNumbers: winningNumbers, playerNumbers: playerNumbers})
	}

	return result
}

func CountWinningNumbers(card ScratchCard) int {
	result := 0
	//fmt.Println("Processing Card ", card.cardNumber)
	for _, item := range card.playerNumbers {
		if Search(card.winningNumbers, item) {
			result++
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

func ProcessCards(cards []ScratchCard, startingCards []ScratchCard) int {
	scratchCards := []ScratchCard{}
	cardCount := 0
	for _, card := range cards {
		cardCount++
		wins := CountWinningNumbers(card)
		for i := 0; i < wins; i++ {
			cardCopy := startingCards[card.cardNumber+i]
			scratchCards = append(scratchCards, cardCopy)
		}
	}
	if len(scratchCards) > 0 {
		cardCount += ProcessCards(scratchCards, startingCards)
	}

	return cardCount
}

func main() {
	lines := ReadFile("../day_4_input.txt")
	cardCount := 0
	scratchCards := GetScratchCards(lines)

	processed := ProcessCards(scratchCards, scratchCards)

	fmt.Println(processed, cardCount)
}
