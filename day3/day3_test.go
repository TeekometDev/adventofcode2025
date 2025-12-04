package main

import "testing"

func TestTask1(t *testing.T) {
	got := SolveTask1("testfile")
	want := 357

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask1BasicWorkingExampleL1(t *testing.T) {
	got := ResolveLineT1("987654321111111")
	want := 98

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask1BasicWorkingExampleL2(t *testing.T) {
	got := ResolveLineT1("811111111111119")
	want := 89

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask1BasicWorkingExampleL3(t *testing.T) {
	got := ResolveLineT1("234234234234278")
	want := 78

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask1BasicWorkingExampleL4(t *testing.T) {
	got := ResolveLineT1("818181911112111")
	want := 92

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask2(t *testing.T) {
	got := SolveTask2("testfile")
	const want uint64 = 3121910778619

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask2MT(t *testing.T) {
	got := SolveTask2MT("testfile")
	const want uint64 = 3121910778619

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask2BasicWorkingExampleL1(t *testing.T) {
	got := ResolveLineT2("987654321111111")
	const want uint64 = 987654321111

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask2BasicWorkingExampleL2(t *testing.T) {
	got := ResolveLineT2("811111111111119")
	const want uint64 = 811111111119

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask2BasicWorkingExampleL3(t *testing.T) {
	got := ResolveLineT2("234234234234278")
	const want uint64 = 434234234278

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}

func TestTask2BasicWorkingExampleL4(t *testing.T) {
	got := ResolveLineT2("818181911112111")
	const want uint64 = 888911112111

	if got != want {
		t.Errorf("Expected %d and got %d", want, got)
	}
}
