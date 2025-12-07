package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

func SolveTask1(fileName string) int {
	lines := ReadFile(fileName)
	start := time.Now()
	result := 0
	var numbers [][]int
	var operands []string
	for i, line := range lines {
		symbols := strings.Fields(line)
		if i == len(lines)-1 {
			// last line
			operands = symbols
		} else {
			var tempArr []int
			for _, num := range symbols {
				newNum, _ := strconv.Atoi(num)
				tempArr = append(tempArr, newNum)
			}
			numbers = append(numbers, tempArr)
		}
	}
	for j, char := range operands {
		tempNum := numbers[0][j]
		for i := 1; i < len(numbers); i++ {
			if char == "*" {
				tempNum = tempNum * numbers[i][j]
			} else if char == "+" {
				tempNum += numbers[i][j]
			}
		}
		result += tempNum
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("Runtime P1: %v\n", duration.Seconds())
	return result
}

func SolveTask2(fileName string) int {
	lines := ReadFile(fileName)
	start := time.Now()
	result := 0
	var numbers [][]string
	highestColIndex := 0
	for _, line := range lines {
		var tempLine []string
		for i, sym := range line {
			tempLine = append(tempLine, string(sym))
			if i > highestColIndex {
				highestColIndex = i
			}
		}
		numbers = append(numbers, tempLine)
	}
	for ind := highestColIndex; ind >= 0; ind-- {
		op := numbers[len(numbers)-1][ind]
		if op == " " {
			continue
		}
		if op == "*" {
			tempNum := 1
			for runInd := highestColIndex; runInd >= ind; runInd-- {
				numString := ""
				for numIn := 0; numIn < len(numbers)-1; numIn++ {
					currSym := numbers[numIn][runInd]
					if currSym != " " {
						numString += currSym
					}
				}
				calcNum, _ := strconv.Atoi(numString)
				tempNum = tempNum * calcNum
			}
			result += tempNum
			highestColIndex = ind - 2
		}
		if op == "+" {
			tempNum := 0
			for runInd := highestColIndex; runInd >= ind; runInd-- {
				numString := ""
				for numIn := 0; numIn < len(numbers)-1; numIn++ {
					currSym := numbers[numIn][runInd]
					if currSym != " " {
						numString += currSym
					}
				}
				calcNum, _ := strconv.Atoi(numString)
				tempNum = tempNum + calcNum
			}
			result += tempNum
			highestColIndex = ind - 2
		}
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
