package main

import (
	file "advent_of_code/shared"
	"testing"
)

func TestMeasure(t *testing.T) {
	expect := 4

	got := measure(7, 9)

	if got != expect {
		t.Errorf("Wrong number of ways to win returned. Got %v, expected %v \n",
			got, expect)
	}
}

func TestSumWaystoWin(t *testing.T) {
	expect := 288

	got := sumWaysToWin("../test_input.txt")

	if got != expect {
		t.Errorf("Wrong number of ways to win returned. Got %v, expected %v \n",
			got, expect)
	}
}

func TestMapResults(t *testing.T) {
	lines := file.Lines("../test_input.txt")
	expect := []RaceResult{
		{
			time:     7,
			distance: 9,
		},
		{
			time:     15,
			distance: 40,
		},
		{
			time:     30,
			distance: 200,
		},
	}

	got := mapResults(lines)

	for i, element := range got {
		if element.time != expect[i].time || element.distance != expect[i].distance {
			t.Errorf("Wrong time or distance found. Got time: %v and distance: %v, expected time: %v and distance: %v \n",
				element.time, element.distance, expect[i].time, expect[i].distance)
		}
	}
}
