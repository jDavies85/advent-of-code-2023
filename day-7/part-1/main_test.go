package main

import (
	file "advent_of_code/shared"
	"testing"
)

func TestMapCards(t *testing.T) {
	lines := file.Lines("../test_input.txt")
	expect := []CardHand{
		{
			bid: 765,
			hand: []KeyValuePair{
				{11, "3"},
				{12, "2"},
				{4, "T"},
				{11, "3"},
				{1, "K"},
			},
		},
		{
			bid: 684,
			hand: []KeyValuePair{
				{4, "T"},
				{9, "5"},
				{9, "5"},
				{3, "J"},
				{9, "5"},
			},
		},
		{
			bid: 28,
			hand: []KeyValuePair{
				{1, "K"},
				{1, "K"},
				{8, "6"},
				{7, "7"},
				{7, "7"},
			},
		},
		{
			bid: 220,
			hand: []KeyValuePair{
				{1, "K"},
				{4, "T"},
				{3, "J"},
				{3, "J"},
				{4, "T"},
			},
		},
		{
			bid: 483,
			hand: []KeyValuePair{
				{2, "Q"},
				{2, "Q"},
				{2, "Q"},
				{3, "J"},
				{0, "A"},
			},
		},
	}

	got := mapCardHands(lines)

	for i, element := range got {
		if element.bid != expect[i].bid {
			t.Errorf("Wrong bid amount found. Got: %v, expected: %v \n",
				element.bid, expect[i].bid)
		}
		for j, innerElement := range element.hand {
			if innerElement != expect[i].hand[j] {
				t.Errorf("Wrong card value found. Got: %v, expected: %v \n",
					innerElement, expect[i].hand[j])
			}
		}
	}
}
