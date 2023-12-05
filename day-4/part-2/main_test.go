package main

import "testing"

func TestCountWinningNumbers(t *testing.T) {
	card := ScratchCard{cardNumber: 1, winningNumbers: []int{41, 48, 83, 86, 17}, playerNumbers: []int{83, 86, 6, 31, 17, 9, 48, 53}}
	expect := 4

	got := CountWinningNumbers(card)

	if got != expect {
		t.Errorf("Wrong number of winning numbers. Got %v, expected %v \n",
			got, expect)
	}
}
