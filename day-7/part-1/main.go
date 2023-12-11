package main

import (
	file "advent_of_code/shared"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type CardHand struct {
	bid  int
	hand []KeyValuePair
}

type KeyValuePair struct {
	key   int
	value string
}

func sortHands(hands []CardHand) []CardHand {
	for _, h := range hands {
		sort.Slice(h.hand, func(i, j int) bool {
			return h.hand[i].key < h.hand[j].key
		})
	}
	return hands
}

func mapCardHands(lines []string) []CardHand {
	ordering := []KeyValuePair{
		{
			key:   0,
			value: "A",
		},
		{
			key:   1,
			value: "K",
		},
		{
			key:   2,
			value: "Q",
		},
		{
			key:   3,
			value: "J",
		},
		{
			key:   4,
			value: "T",
		},
		{
			key:   5,
			value: "9",
		},
		{
			key:   6,
			value: "8",
		},
		{
			key:   7,
			value: "7",
		},
		{
			key:   8,
			value: "6",
		},
		{
			key:   9,
			value: "5",
		},
		{
			key:   10,
			value: "4",
		},
		{
			key:   11,
			value: "3",
		},
		{
			key:   12,
			value: "2",
		},
	}
	result := []CardHand{}
	for _, line := range lines {
		cardHand := CardHand{}
		split := strings.Split(line, " ")
		cardHand.bid = getNumber(split[1])
		cardHand.hand = make([]KeyValuePair, len(split[0]))

		for i, r := range split[0] {
			currentCard := string(r)
			cardHand.hand[i] = KeyValuePair{
				key: slices.IndexFunc(ordering, func(kvp KeyValuePair) bool {
					return kvp.value == currentCard
				}),
				value: currentCard,
			}
		}
		result = append(result, cardHand)
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

func main() {
	lines := file.Lines("../test_input.txt")
	cardHands := mapCardHands(lines)
	cardHands = sortHands(cardHands)
}
