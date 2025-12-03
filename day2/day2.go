package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func InvestigateRange_T1_async(rangeStr string, resultVal *uint64, waitGroup *sync.WaitGroup, mu *sync.Mutex) {
	rngList := strings.Split(rangeStr, "-")
	start, _ := strconv.Atoi(rngList[0])
	end, _ := strconv.Atoi(rngList[1])
	for curr := start; curr <= end; curr++ {
		currStr := strconv.Itoa(curr)
		if len(currStr)%2 == 0 {
			first := currStr[0:(len(currStr) / 2)]
			second := currStr[(len(currStr) / 2):]
			if first == second {
				mu.Lock()
				*resultVal += uint64(curr)
				mu.Unlock()
			}
		}
	}
	defer waitGroup.Done()
}

func InvestigateRange_T1(rangeStr string) uint64 {
	rngList := strings.Split(rangeStr, "-")
	start, _ := strconv.Atoi(rngList[0])
	end, _ := strconv.Atoi(rngList[1])
	var resultVal uint64 = 0
	for curr := start; curr <= end; curr++ {
		currStr := strconv.Itoa(curr)
		if len(currStr)%2 == 0 {
			first := currStr[0:(len(currStr) / 2)]
			second := currStr[(len(currStr) / 2):]
			if first == second {
				resultVal += uint64(curr)
			}
		}
	}
	return resultVal
}

func SolveTask1_Sync(fileName string) uint64 {
	content := ReadFile(fileName)
	start := time.Now()
	listOfRanges := strings.Split(content, ",")
	var result uint64 = 0
	for _, value := range listOfRanges {
		result += InvestigateRange_T1(value)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime w/o Multithreading: %v\n", duration.Seconds())
	return result
}

func SolveTask1_Async(fileName string) uint64 {
	content := ReadFile(fileName)
	start := time.Now()
	listOfRanges := strings.Split(content, ",")
	var result uint64 = 0
	var waitGroup sync.WaitGroup
	var mu sync.Mutex
	for _, value := range listOfRanges {
		waitGroup.Add(1)
		go InvestigateRange_T1_async(value, &result, &waitGroup, &mu)
	}
	waitGroup.Wait()
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime w Multithreading: %v\n", duration.Seconds())
	return result
}

func InvestigateRange_T2(rangeStr string, result *uint64, waitGroup *sync.WaitGroup, mu *sync.Mutex) {
	rngList := strings.Split(rangeStr, "-")
	start, _ := strconv.Atoi(rngList[0])
	end, _ := strconv.Atoi(rngList[1])
	for curr := start; curr <= end; curr++ {
		currStr := strconv.Itoa(curr)
		for currLen := 1; currLen <= (len(currStr) / 2); currLen++ {
			if analyzeString(currStr, currLen) {
				mu.Lock()
				*result += uint64(curr)
				mu.Unlock()
				break
			}
		}
	}
	defer waitGroup.Done()
}

func analyzeString(testString string, currLen int) bool {
	if (len(testString) % currLen) != 0 {
		return false
	}
	compare := testString[0:currLen]
	for currInvest := currLen; currInvest <= (len(testString) - currLen); currInvest += currLen {
		curr := testString[currInvest:(currInvest + currLen)]
		if compare != curr {
			return false
		}
	}
	return true
}

func SolveTask2(fileName string) uint64 {
	content := ReadFile(fileName)
	listOfRanges := strings.Split(content, ",")
	var result uint64 = 0
	var waitGroup sync.WaitGroup
	var mu sync.Mutex
	for _, value := range listOfRanges {
		waitGroup.Add(1)
		go InvestigateRange_T2(value, &result, &waitGroup, &mu)
	}
	waitGroup.Wait()
	return result
}

func main() {
	res_T1 := SolveTask1_Sync("puzzle")
	res_T1_async := SolveTask1_Async("puzzle")
	res_T2 := SolveTask2("puzzle")
	fmt.Printf("Task 1: %d\nTask 1 Async: %d\nTask 2: %d\n", res_T1, res_T1_async, res_T2)
}
