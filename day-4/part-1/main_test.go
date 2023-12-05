package main

import "testing"

func TestSearchContainsItem(t *testing.T) {
	arr := []int{83, 86, 6, 31, 17, 9, 48, 53}
	x := 6
	expect := true

	got := Search(arr, x)

	if got != expect {
		t.Errorf("Array does not contain %v. Got %v, expected %v \n",
			x, got, expect)
	}

}

func TestSearchDoesNotContainItem(t *testing.T) {
	arr := []int{83, 86, 6, 31, 17, 9, 48, 53}
	x := 99
	expect := false

	got := Search(arr, x)

	if got != expect {
		t.Errorf("Array does not contain %v. Got %v, expected %v \n",
			x, got, expect)
	}

}
