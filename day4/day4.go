package main

import (
	"fmt"
	"time"
)

func SolveTask1(fileName string) int {
	matrix := ReadFile(fileName)
	start := time.Now()
	result := 0
	for rowI, row := range matrix {
		for colI, element := range row {
			if element {
				if Task1Worker(&matrix, rowI, colI) {
					result++
				}
			}
		}
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P1: %v\n", duration.Seconds())
	return result
}

func SolveTask2(fileName string) int {
	currMatrix := ReadFile(fileName)
	start := time.Now()
	nextMatrix := currMatrix
	result := 0
	for {
		currRun := 0
		// Iterate over the current Matrix
		for rowI, row := range currMatrix {
			for colI, element := range row {
				if element {
					if Task1Worker(&currMatrix, rowI, colI) {
						result++
						currRun++
						nextMatrix[rowI][colI] = false
					}
				}
			}
		}
		if currRun == 0 {
			break
		}
		currMatrix = nextMatrix
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P2: %v\n", duration.Seconds())
	return result
}

func Task1Worker(matrix *[][]bool, row int, col int) bool {
	count := 0
	for rowI := row - 1; rowI <= row+1; rowI++ {
		if rowI < 0 || rowI >= len(*matrix) {
			continue
		}
		for colI := col - 1; colI <= col+1; colI++ {
			if colI < 0 || colI >= len((*matrix)[rowI]) {
				continue
			}
			if colI != col || rowI != row {
				if (*matrix)[rowI][colI] {
					count++
				}
			}

		}
	}
	return count < 4
}

func main() {
	res_T1 := SolveTask1("puzzle")
	res_T1MT := SolveTask1MT("puzzle")
	res_T2 := SolveTask2("puzzle")
	fmt.Printf("Task 1: %d\nTask 2: %d\n", res_T1, res_T2)
	fmt.Printf("MT - Task 1: %d\n", res_T1MT)
}
