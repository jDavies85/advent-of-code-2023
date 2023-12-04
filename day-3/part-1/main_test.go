package main

import (
	"testing"
)

func TestFindEngineParts(t *testing.T) {
	var lines = ReadFile("../test_input.txt")
	var engineMap = GetEngineMap(lines)
	expect := []int{467, 35, 633, 617, 592, 755, 664, 598}
	got := FindEngineParts(engineMap)

	for i, element := range got {
		if element != expect[i] {
			t.Errorf("Failed to parse file. Got %v at element %v, expected %v \n",
				element, i, expect[i])
		}
	}
}
