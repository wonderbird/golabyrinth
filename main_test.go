package main

import "testing"

func Hello() string {
	return "Hello World!"
}

func TestMain(t *testing.T) {
	got := Hello()
	want := "Hello World!"

	if got != want {
		t.Errorf("got %q, but want %q", got, want)
	}
}
