package main

import (
	"fmt"
	"sync"
	"time"
)

func SolveTask1MT(fileName string) int {
	matrix := ReadFile(fileName)
	start := time.Now()
	result := 0
	var wg sync.WaitGroup
	var mu sync.Mutex
	for rowI, row := range matrix {
		for colI, element := range row {
			if element {
				wg.Add(1)
				go Task1WorkerMT(&matrix, rowI, colI, &result, &wg, &mu)
			}
		}
	}
	wg.Wait()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P1MT: %v\n", duration.Seconds())
	return result
}

func Task1WorkerMT(matrix *[][]bool, row int, col int, counter *int, waitGroup *sync.WaitGroup, mu *sync.Mutex) {
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
	if count < 4 {
		mu.Lock()
		*counter++
		mu.Unlock()
	}
	waitGroup.Done()
}
