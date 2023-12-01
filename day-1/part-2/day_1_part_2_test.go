package main

import "testing"

func TestTryParseInt(t *testing.T) {
	//arrange
	str := "a"
	expect1, expect2 := 0, false

	//act
	got1, got2 := TryParseInt(str)

	//assert
	if expect1 != got1 || expect2 != got2 {
		t.Errorf("Failed to parse %v. Got1 %v, expected1 %v, Got2 %v, expected2 %v \n",
			str, got1, expect1, got2, expect2)
	}
}

func TestExtractNumber(t *testing.T) {
	//arrange
	str := "eighthree"
	expect := 83

	//act
	got := ExtractNumber(str)

	//assert
	if expect != got {
		t.Errorf("Failed to extract number from %v. Got %v, expected %v \n",
			str, got, expect)
	}
}

func TestProcessInput(t *testing.T) {
	//arrange
	filePath := "part_2_test_input.txt"
	expect := 281

	//act
	got := ProcessInput(filePath)

	//assert
	if expect != got {
		t.Errorf("Error processing input from file: %v. Got %v, expected %v \n",
			filePath, got, expect)
	}
}
