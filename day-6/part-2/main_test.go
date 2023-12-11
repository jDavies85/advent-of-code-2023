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

func TestMapResult(t *testing.T) {
	lines := file.Lines("../test_input.txt")
	expect := RaceResult{
		time:     71530,
		distance: 940200,
	}

	got := mapResult(lines)

	if got.time != expect.time || got.distance != expect.distance {
		t.Errorf("Wrong time or distance found. Got time: %v and distance: %v, expected time: %v and distance: %v \n",
			got.time, got.distance, expect.time, expect.distance)
	}
}
