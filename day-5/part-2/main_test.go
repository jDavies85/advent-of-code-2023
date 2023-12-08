package main

import "testing"

func TestGetLowestLocationNumber(t *testing.T) {
	lines := ReadFile("../test_input.txt")
	seedRanges := GetSeedRanges(lines[0])
	maps := MapInput(lines)

	expect := 46

	got := GetLowestLocationNumber(seedRanges, maps)

	if got != expect {
		t.Errorf("Wrong location number found. Got %v, expected %v \n",
			got, expect)
	}
}

func TestGetSeedRanges(t *testing.T) {
	lines := ReadFile("../test_input.txt")
	expect := []SeedRange{{start: 79, end: 92}, {start: 55, end: 67}}
	got := GetSeedRanges(lines[0])

	for i, sr := range got {
		if sr.start != expect[i].start && sr.end != expect[i].end {
			t.Errorf("Wrong seed range found. Got %v, expected %v \n",
				got, expect[i])
		}
	}
}
