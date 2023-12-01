package main

import "testing"

func TestClearString(t *testing.T) {
	//arrange
	str := "1abc2"
	expect := "12"

	//act
	got := ClearString(str)

	//assert
	if expect != got {
		t.Errorf("Failed to remove non numeric characters from %v. Got %v, expected %v \n",
			str, got, expect)
	}
}

func TestExtractNumber(t *testing.T) {
	//arrange
	str := "a1b2c3d4e5f"
	expect := 15

	//act
	got := ExtractNumber(str)

	//assert
	if expect != got {
		t.Errorf("Failed to extract number from %v. Got %v, expected %v \n",
			str, got, expect)
	}
}
