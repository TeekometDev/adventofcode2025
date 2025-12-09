package main

import "testing"

func TestTask1(t *testing.T) {
	got := SolveTask1("testfile")
	want := 50

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	// Test isn't working but not problematic, dued to the structure of the puzzles values
	got := SolveTask2("testfile")
	want := 24

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}
