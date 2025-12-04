package main

import (
	"fmt"
	"time"
)

// Just testing a small optimization
func SolveTask2sameMatrix(fileName string) int {
	currMatrix := ReadFile(fileName)
	start := time.Now()
	result := 0
	iterationCount := 0
	for {
		iterationCount++
		currRun := 0
		// Iterate over the current Matrix
		for rowI, row := range currMatrix {
			for colI, element := range row {
				if element {
					if Task1Worker(&currMatrix, rowI, colI) {
						result++
						currRun++
						currMatrix[rowI][colI] = false
					}
				}
			}
		}
		if currRun == 0 {
			break
		}
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P2sm: %v, Iterations: %d\n", duration.Seconds(), iterationCount)
	return result
}
