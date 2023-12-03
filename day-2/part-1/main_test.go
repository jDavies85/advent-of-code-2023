package main

import (
	"testing"
)

func TestReadFile(t *testing.T) {
	//arrange
	path := "test_input.txt"
	expect := []string{"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"}

	//act
	got := ReadFile(path)

	//assert
	for i, element := range got {
		if element != expect[i] {
			t.Errorf("Failed to parse file %v. Got %v on line %v, expected %v \n",
				path, element, i, got[i])
		}
	}
}

func TestGameResults(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	expect := []string{"Game 1", "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"}

	got := GameResults(input)

	for i, element := range got {
		if element != expect[i] {
			t.Errorf("Failed to parse %v. Got %v , expected %v \n",
				input, element, got[i])
		}
	}
}

func TestGetRoundResult(t *testing.T) {
	input := "3 blue, 4 red"
	expect := RoundResult{blue: 3, red: 4}

	got := GetRoundResult(input)

	if got.blue != 3 {
		t.Errorf("Failed to parse %v. Got %v , expected %v \n",
			input, got.blue, expect.blue)
	}
	if got.red != 4 {
		t.Errorf("Failed to parse %v. Got %v , expected %v \n",
			input, got.red, expect.red)
	}
}

func TestRoundResults(t *testing.T) {
	input := "3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	expect := []RoundResult{RoundResult{blue: 3, red: 4}, RoundResult{red: 1, green: 2, blue: 6}, RoundResult{green: 2}}

	got := RoundResults(input)

	for i, element := range got {
		if element.red != expect[i].red {
			t.Errorf("Failed to parse %v. Got %v , expected %v \n",
				input, element, expect[i])
		}
		if element.blue != expect[i].blue {
			t.Errorf("Failed to parse %v. Got %v , expected %v \n",
				input, element, expect[i])
		}
		if element.green != expect[i].green {
			t.Errorf("Failed to parse %v. Got %v , expected %v \n",
				input, element, expect[i])
		}
	}
}

func TestSumValidGames(t *testing.T) {
	path := "test_input.txt"
	expect := 8

	got := SumValidGames(path)

	if expect != got {
		t.Errorf("Failed to extract number from %v. Got %v, expected %v \n",
			path, got, expect)
	}
}
