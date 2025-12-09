package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func ReadFile(fileName string) [][]rune {
	start := time.Now()
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	var returnArray [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		var line []rune
		for _, element := range text {
			line = append(line, element)
		}
		returnArray = append(returnArray, line)
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Printf("File Runtime: %v\n", duration.Seconds())
	return returnArray
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
