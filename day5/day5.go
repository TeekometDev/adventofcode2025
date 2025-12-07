package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func SolveTask1(filename string) int {
	lines := ReadFile(filename)
	var freshItemRanges [][]int
	result := 0
	idReg, _ := regexp.Compile("^[0-9]+$")
	rangeReg, _ := regexp.Compile("^[0-9]+-[0-9]+$")
	for _, line := range lines {
		if rangeReg.MatchString(line) {
			ranges := strings.Split(line, "-")
			start, _ := strconv.Atoi(ranges[0])
			end, _ := strconv.Atoi(ranges[1])
			freshItemRanges = append(freshItemRanges, []int{start, end})
		} else if idReg.MatchString(line) {
			key, _ := strconv.Atoi(line)
			if checkInRange(key, &freshItemRanges) {
				result++
			}
		}
	}
	return result
}

func SolveTask2(filename string) int {
	lines := ReadFile(filename)
	var freshItemRanges [][]int
	result := 0
	rangeReg, _ := regexp.Compile("^[0-9]+-[0-9]+$")
	for _, line := range lines {
		if rangeReg.MatchString(line) {
			ranges := strings.Split(line, "-")
			start, _ := strconv.Atoi(ranges[0])
			end, _ := strconv.Atoi(ranges[1])
			freshItemRanges = append(freshItemRanges, []int{start, end})
		}
	}
	reorderRanges(&freshItemRanges)
	for _, element := range freshItemRanges {
		result += element[1] - element[0] + 1
	}
	return result
}

func checkInRange(actNum int, ranges *[][]int) bool {
	for _, currRange := range *ranges {
		if actNum >= currRange[0] && actNum <= currRange[1] {
			return true
		}
	}
	return false
}

func reorderRanges(ranges *[][]int) {
	for currIndex, currRange := range *ranges {
		currStart := currRange[0]
		currEnd := currRange[1]
		for invIndex, invRange := range *ranges {
			if invIndex == currIndex {
				continue
			}
			invStart := invRange[0]
			invEnd := invRange[1]
			// Case Start in investigated Range
			if currStart > invStart && currStart < invEnd {
				if currEnd > invEnd {
					(*ranges)[invIndex][1] = currEnd
				}
				removeIndexFromRanges(currIndex, ranges)
				reorderRanges(ranges)
				return
			}
			// Case End in investigated Range
			if currEnd > invStart && currEnd < invEnd {
				if currStart < currEnd {
					(*ranges)[invIndex][0] = currStart
				}
				removeIndexFromRanges(currIndex, ranges)
				reorderRanges(ranges)
				return
			}
			// Case identical
			if currStart == invStart && currEnd == invEnd {
				removeIndexFromRanges(invIndex, ranges)
				reorderRanges(ranges)
				return
			}
			// Case combining
			if currStart == invEnd {
				(*ranges)[invIndex][1] = currEnd
				removeIndexFromRanges(currIndex, ranges)
				reorderRanges(ranges)
				return
			}
			if currEnd == invStart {
				(*ranges)[invIndex][0] = currStart
				removeIndexFromRanges(currIndex, ranges)
				reorderRanges(ranges)
				return
			}
		}
	}
}

func removeIndexFromRanges(index int, ranges *[][]int) {
	*ranges = append((*ranges)[:index], (*ranges)[index+1:]...)
}

func main() {
	t1_solution := SolveTask1("puzzle")
	t2_solution := SolveTask2("puzzle")
	fmt.Printf("Task 1: %d\nTask 2: %d\n", t1_solution, t2_solution)
}
