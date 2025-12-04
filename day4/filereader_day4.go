package main

import (
	"bufio"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ReadFile(fileName string) [][]bool {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	var matrix [][]bool
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		var line []bool
		for _, symbol := range text {
			if symbol == '.' {
				line = append(line, false)
			} else if symbol == '@' {
				line = append(line, true)
			}
		}
		matrix = append(matrix, line)
	}
	return matrix
}
