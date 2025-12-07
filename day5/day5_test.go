package main

import "testing"

func TestPart1(t *testing.T) {
	got := SolveTask1("testfile")
	want := 3

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func TestPart2(t *testing.T) {
	got := SolveTask2("testfile")
	want := 14

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}
