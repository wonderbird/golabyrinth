package main

import "testing"

type Coordinate struct {
	row    int
	column int
}

type Board struct {
	robot Coordinate
}

func TestPlaceARobot(t *testing.T) {
	got := Board{
		robot: Coordinate{row: 100, column: 100},
	}

	want := Board{
		robot: Coordinate{row: 100, column: 100},
	}

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}
