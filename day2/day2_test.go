package main

import "testing"

func TestImport(t *testing.T) {
	got := ReadFile("testfile")
	want := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"

	if got != want {
		t.Error("File read is not fitting!")
	}
}

func TestT1_sync(t *testing.T) {
	got := SolveTask1_Sync("testfile")
	const want uint64 = 1227775554

	if got != want {
		t.Errorf("Got %d instead of %d", got, want)
	}
}

func TestT1_async(t *testing.T) {
	got := SolveTask1_Async("testfile")
	const want uint64 = 1227775554

	if got != want {
		t.Errorf("Got %d instead of %d", got, want)
	}
}

func TestT2(t *testing.T) {
	got := SolveTask2("testfile")
	const want uint64 = 4174379265

	if got != want {
		t.Errorf("Got %d instead of %d", got, want)
	}
}
