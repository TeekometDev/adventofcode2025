package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func SolveTask2MT(filename string) uint64 {
	lines := ReadFile(filename)
	start := time.Now()
	var solution uint64 = 0
	var mu sync.Mutex
	var wg sync.WaitGroup
	for _, line := range lines {
		wg.Add(1)
		go ResolveLineT2MT(line, &solution, &mu, &wg)
	}
	wg.Wait()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P2MT: %v\n", duration.Seconds())
	return solution
}

func ResolveLineT2MT(line string, sol *uint64, mu *sync.Mutex, wg *sync.WaitGroup) {
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
	mu.Lock()
	*sol += uint64(result)
	mu.Unlock()
	wg.Done()
}
