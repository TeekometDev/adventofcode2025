package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
	"time"
)

type Coordinates struct {
	x int
	y int
}

type Area struct {
	cor1 Coordinates
	cor2 Coordinates
	area int
}

func SolveTask1(fileName string) int {
	file := ReadFile(fileName)
	var coordinates []Coordinates
	for _, line := range file {
		corS := strings.Split(line, ",")
		corXI, _ := strconv.Atoi(corS[0])
		corYI, _ := strconv.Atoi(corS[1])
		coordinates = append(coordinates, Coordinates{x: corXI, y: corYI})
	}
	res := 0
	for i := 0; i < len(coordinates)-1; i++ {
		for j := i + 1; j < len(coordinates); j++ {
			area := (int)(math.Abs((float64)(coordinates[i].x-coordinates[j].x))+1) * (int)(math.Abs((float64)(coordinates[i].y-coordinates[j].y))+1)
			if area > res {
				res = area
			}
		}
	}
	return res
}

func SolveTask2(fileName string) int {
	file := ReadFile(fileName)
	start := time.Now()
	var coordinates []Coordinates
	maxCoordinates := Coordinates{x: 0, y: 0}
	minCoordinates := Coordinates{x: 0, y: 0}
	validXValues := make(map[int]bool)
	validYValues := make(map[int]bool)
	for i, line := range file {
		corS := strings.Split(line, ",")
		corXI, _ := strconv.Atoi(corS[0])
		corYI, _ := strconv.Atoi(corS[1])
		if corXI > maxCoordinates.x {
			maxCoordinates.x = corXI
		}
		if corYI > maxCoordinates.y {
			maxCoordinates.y = corYI
		}
		if i == 0 {
			minCoordinates.x = corXI
			minCoordinates.y = corYI
		} else {
			if corXI < minCoordinates.x {
				minCoordinates.x = corXI
			}
			if corYI < minCoordinates.y {
				minCoordinates.y = corYI
			}
		}
		validXValues[corXI] = true
		validYValues[corYI] = true
		coordinates = append(coordinates, Coordinates{x: corXI, y: corYI})
	}
	edgeValues := make(map[Coordinates]bool)
	for i, currNode := range coordinates {
		var secondNode Coordinates
		if i < len(coordinates)-1 {
			secondNode = coordinates[i+1]
		} else {
			secondNode = coordinates[0]
		}
		if currNode.x > secondNode.x {
			for x := secondNode.x; x <= currNode.x; x++ {
				edgeValues[Coordinates{x: x, y: currNode.y}] = true
			}
		}
		if currNode.x < secondNode.x {
			for x := secondNode.x; x >= currNode.x; x-- {
				edgeValues[Coordinates{x: x, y: currNode.y}] = true
			}
		}
		if currNode.y > secondNode.y {
			for y := secondNode.y; y <= currNode.y; y++ {
				edgeValues[Coordinates{x: currNode.x, y: y}] = true
			}
		}
		if currNode.y < secondNode.y {
			for y := secondNode.y; y >= currNode.y; y-- {
				edgeValues[Coordinates{x: currNode.x, y: y}] = true
			}
		}
	}
	var areaList []Area
	for i := 0; i < len(coordinates)-1; i++ {
		for j := i + 1; j < len(coordinates); j++ {
			area := (int)(math.Abs((float64)(coordinates[i].x-coordinates[j].x))+1) * (int)(math.Abs((float64)(coordinates[i].y-coordinates[j].y))+1)
			areaList = append(areaList, Area{cor1: coordinates[i], cor2: coordinates[j], area: area})
		}
	}
	slices.SortFunc(areaList, compArea)
	for _, biggestArea := range areaList {
		xCors := []int{biggestArea.cor1.x, biggestArea.cor2.x}
		slices.Sort(xCors)
		yCors := []int{biggestArea.cor1.y, biggestArea.cor2.y}
		slices.Sort(yCors)
		stopThis := false
		for x := xCors[0] + 1; x < xCors[1]; x++ {
			_, x_ok := validXValues[x]
			if !x_ok {
				continue
			}
			for y := yCors[0] + 1; y < yCors[1]; y++ {
				_, y_ok := validYValues[y]
				if !y_ok {
					continue
				}
				testCor := Coordinates{x: x, y: y}
				_, ok := edgeValues[testCor]
				if ok {
					stopThis = true
					break
				}
			}
			if stopThis {
				break
			}
		}
		if stopThis {
			continue
		} else {
			end := time.Now()
			duration := end.Sub(start)
			fmt.Printf("Runtime P2: %v\n", duration.Seconds())
			return biggestArea.area
		}
	}
	return 0
}

func compArea(a, b Area) int {
	return cmp.Compare(b.area, a.area)
}

func main() {
	t1_solution := SolveTask1("puzzle")
	t2_solution := SolveTask2("puzzle")
	fmt.Printf("Task 1: %d\nTask 2: %d\n", t1_solution, t2_solution)
}
