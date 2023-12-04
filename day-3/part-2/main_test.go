package main

import (
	"testing"
)

func TestCaptureNumber(t *testing.T) {
	input := []string{"4", "6", "7", ".", ".", ".", ".", "1", "1", "4"}

	expect := 114
	got := CaptureNumber(input, 9)

	if got != expect {
		t.Errorf("Failed to capture number. Got  %v, expected %v \n",
			got, expect)
	}

}

func TestFindGears(t *testing.T) {
	var lines = ReadFile("../test_input.txt")
	var engineMap = GetEngineMap(lines)
	expect := []int{16345, 451490}
	got := FindGears(engineMap)

	for i, element := range got {
		if element != expect[i] {
			t.Errorf("Failed to parse file. Got %v at element %v, expected %v \n",
				element, i, expect[i])
		}
	}
}
