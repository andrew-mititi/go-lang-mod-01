package main

import (
	"testing"
)


func TestCalculateinfo(t *testing.T){
	s := Square{
		length: 2,
		width: 3,
	}

	square := int(calculateInfo(s))

	if square != 6 {
		t.Errorf("Sum function returned incorrect value, got: %d, expected: %d", square, 6)
	}
}


func TestAdd(t *testing.T){
	got := add(5,4)
	want := 9
	if got != want {
		t.Errorf("Addition failed: Got: %d  Expected: %d", got, want)
	}
}