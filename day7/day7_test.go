package main

import "testing"

func TestTask1(t *testing.T) {
	got := SolveTask1("testfile")
	want := 21

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}

func TestTask2(t *testing.T) {
	got := SolveTask2("testfile")
	want := 40

	if got != want {
		t.Errorf("Got %d and wanted %d", got, want)
	}
}
