package main

import (
	"fmt"
	"strconv"
	"time"
)

func SolveTask1(filename string) int {
	lines := ReadFile(filename)
	start := time.Now()
	solution := 0
	for _, line := range lines {
		solution += ResolveLineT1(line)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P1: %v\n", duration.Seconds())
	return solution
}

func ResolveLineT1(line string) int {
	first := line[0]
	second := line[1]
	for index := 2; index < len(line); index++ {
		symbol := line[index]
		if second > first {
			first = second
			second = symbol
		} else if symbol > second {
			second = symbol
		}
	}
	firstStr := string(first)
	secondStr := string(second)
	sol, _ := strconv.Atoi(firstStr + secondStr)
	return sol
}

func SolveTask2(filename string) uint64 {
	lines := ReadFile(filename)
	start := time.Now()
	var solution uint64 = 0
	for _, line := range lines {
		solution += ResolveLineT2(line)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P2: %v\n", duration.Seconds())
	return solution
}

func ResolveLineT2(line string) uint64 {
	var symbolList []byte
	// Init 12-digit Symbol-List
	for i := 0; i < 12; i++ {
		symbolList = append(symbolList, line[i])
	}
	for i := 12; i < len(line); i++ {
		ShiftSymbols(&symbolList, line[i])
	}
	resultStr := ""
	for i := 0; i < len(symbolList); i++ {
		resultStr += string(symbolList[i])
	}
	result, _ := strconv.Atoi(resultStr)
	return uint64(result)
}

func ShiftSymbols(symbolList *[]byte, newSymbol byte) {
	hasDone := false
	for index := 0; index < len(*symbolList)-1; index++ {
		first := (*symbolList)[index]
		second := (*symbolList)[index+1]
		// Shift the digits
		if hasDone || second > first {
			(*symbolList)[index] = (*symbolList)[index+1]
			hasDone = true
		}
	}
	if !hasDone {
		if newSymbol > (*symbolList)[len(*symbolList)-1] {
			(*symbolList)[len(*symbolList)-1] = newSymbol
		}
	} else {
		(*symbolList)[len(*symbolList)-1] = newSymbol
	}
}

func main() {
	res_T1 := SolveTask1("puzzle")
	res_T2 := SolveTask2("puzzle")
	res_T2MT := SolveTask2MT("puzzle")
	fmt.Printf("Task 1: %d\nTask 2: %d\n", res_T1, res_T2)
	fmt.Printf("Task 2MT: %d\n", res_T2MT)
}
