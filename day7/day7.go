package main

import (
	"fmt"
	"time"
)

func SolveTask1(fileName string) int {
	symbols := ReadFile(fileName)
	start := time.Now()
	splitted := 0
	for indLine, line := range symbols {
		if indLine >= len(symbols)-2 {
			break
		}
		for indCol, symbol := range line {
			if symbol == 'S' {
				symbols[indLine+1][indCol] = '|'
				continue
			}
			if symbol == '|' {
				if symbols[indLine+1][indCol] == '^' {
					symbols[indLine+1][indCol-1] = '|'
					symbols[indLine+1][indCol+1] = '|'
					splitted++
				} else {
					symbols[indLine+1][indCol] = '|'
				}
			}
		}
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P1: %v\n", duration.Seconds())
	return splitted
}

func SolveTask2(fileName string) int {
	symbols := ReadFile(fileName)
	start := time.Now()
	result := 0
	// Init Track for multiplies
	var multTrack [][]int
	for i := 0; i < len(symbols)-1; i++ {
		var multTrackLine []int
		for j := 0; j < len(symbols[0]); j++ {
			multTrackLine = append(multTrackLine, 0)
		}
		multTrack = append(multTrack, multTrackLine)
	}
	for indLine, line := range symbols {
		if indLine >= len(symbols)-2 {
			break
		}
		for indCol, symbol := range line {
			if symbol == 'S' {
				symbols[indLine+1][indCol] = '|'
				multTrack[indLine+1][indCol] = 1
				continue
			}
			if symbol == '|' {
				if symbols[indLine+1][indCol] == '^' {
					symbols[indLine+1][indCol-1] = '|'
					multTrack[indLine+1][indCol-1] += multTrack[indLine][indCol]
					symbols[indLine+1][indCol+1] = '|'
					multTrack[indLine+1][indCol+1] += multTrack[indLine][indCol]
				} else {
					symbols[indLine+1][indCol] = '|'
					multTrack[indLine+1][indCol] += multTrack[indLine][indCol]
				}
			}
		}
	}
	for i := 0; i < len(symbols[0]); i++ {
		result += multTrack[len(multTrack)-1][i]
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P2: %v\n", duration.Seconds())
	return result
}

func main() {
	t1_solution := SolveTask1("puzzle")
	t2_solution := SolveTask2("puzzle")
	fmt.Printf("Task 1: %d\nTask 2: %d\n", t1_solution, t2_solution)
}
