package main

import (
	"cmp"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

type Node struct {
	x       int
	y       int
	z       int
	id      int
	network int
}

type Edge struct {
	node1Id  int
	node2Id  int
	distance float64
}

type Network struct {
	id    int
	nodes []int
}

func SolveTask1(fileName string, connectTimes int) int {
	file := ReadFile(fileName)
	nodeList := createNodes(&file)
	edgeList := calculateDistances(&nodeList)
	slices.SortFunc(edgeList, compDist)
	var networks []Network
	for i := 0; i < connectTimes; i++ {
		currEdge := edgeList[i]
		node1 := &nodeList[currEdge.node1Id]
		node2 := &nodeList[currEdge.node2Id]

		if node1.network == -1 && node2.network == -1 {
			// No Network
			newNetwork := Network{id: len(networks), nodes: []int{node1.id, node2.id}}
			networks = append(networks, newNetwork)
			node1.network = newNetwork.id
			node2.network = newNetwork.id
		} else if node1.network == node2.network {
			continue
		} else if node1.network > -1 && node2.network == -1 {
			networks[node1.network].nodes = append(networks[node1.network].nodes, node2.id)
			node2.network = node1.network
		} else if node1.network == -1 && node2.network > -1 {
			networks[node2.network].nodes = append(networks[node2.network].nodes, node1.id)
			node1.network = node2.network
		} else if node1.network > -1 && node2.network > -1 {
			// Combine Networks
			oldNetwork := node2.network
			for len(networks[oldNetwork].nodes) > 0 {
				changeNodeId := networks[oldNetwork].nodes[0]
				networks[oldNetwork].nodes = networks[oldNetwork].nodes[1:]
				nodeList[changeNodeId].network = node1.network
				networks[node1.network].nodes = append(networks[node1.network].nodes, changeNodeId)
			}
		}
	}
	slices.SortFunc(networks, compSizes)
	retVal := 1
	for i := 0; i < 3; i++ {
		retVal = retVal * len(networks[i].nodes)
	}
	return retVal
}

func SolveTask2(fileName string) int {
	file := ReadFile(fileName)
	nodeList := createNodes(&file)
	edgeList := calculateDistances(&nodeList)
	slices.SortFunc(edgeList, compDist)
	var networks []Network
	continueProcess := true
	i := 0
	result := 0
	for continueProcess {
		currEdge := edgeList[i]
		node1 := &nodeList[currEdge.node1Id]
		node2 := &nodeList[currEdge.node2Id]

		if node1.network == -1 && node2.network == -1 {
			// No Network
			newNetwork := Network{id: len(networks), nodes: []int{node1.id, node2.id}}
			networks = append(networks, newNetwork)
			node1.network = newNetwork.id
			node2.network = newNetwork.id
		} else if node1.network == node2.network {

		} else if node1.network > -1 && node2.network == -1 {
			networks[node1.network].nodes = append(networks[node1.network].nodes, node2.id)
			node2.network = node1.network
		} else if node1.network == -1 && node2.network > -1 {
			networks[node2.network].nodes = append(networks[node2.network].nodes, node1.id)
			node1.network = node2.network
		} else if node1.network > -1 && node2.network > -1 {
			// Combine Networks
			oldNetwork := node2.network
			for len(networks[oldNetwork].nodes) > 0 {
				changeNodeId := networks[oldNetwork].nodes[0]
				networks[oldNetwork].nodes = networks[oldNetwork].nodes[1:]
				nodeList[changeNodeId].network = node1.network
				networks[node1.network].nodes = append(networks[node1.network].nodes, changeNodeId)
			}
		}
		// Check if only one valid Network is left with all nodes
		for netInd := 0; netInd < len(networks); netInd++ {
			if len(networks[netInd].nodes) == len(nodeList) {
				result = node1.x * node2.x
				continueProcess = false
			}
		}
		i++
		if i >= len(edgeList) {
			continueProcess = false
		}
	}
	return result
}

func createNodes(file *[]string) []Node {
	var list []Node
	for ind, line := range *file {
		coordinates := strings.Split(line, ",")
		x, _ := strconv.Atoi(coordinates[0])
		y, _ := strconv.Atoi(coordinates[1])
		z, _ := strconv.Atoi(coordinates[2])
		list = append(list, Node{x: x, y: y, z: z, id: ind, network: -1})
	}
	return list
}

func calculateDistances(nodeList *[]Node) []Edge {
	var edges []Edge
	for i := 0; i < len(*nodeList)-1; i++ {
		node1 := (*nodeList)[i]
		for j := i + 1; j < len(*nodeList); j++ {
			node2 := (*nodeList)[j]
			dist := math.Pow((float64)(node2.x-node1.x), 2) + math.Pow((float64)(node2.y-node1.y), 2) + math.Pow((float64)(node2.z-node1.z), 2)
			edges = append(edges, Edge{node1Id: node1.id, node2Id: node2.id, distance: dist})
		}
	}
	return edges
}

func compDist(a, b Edge) int {
	return cmp.Compare(a.distance, b.distance)
}

func compSizes(a, b Network) int {
	return cmp.Compare(len(b.nodes), len(a.nodes))
}

func main() {
	t1_solution := SolveTask1("puzzle", 1000)
	t2_solution := SolveTask2("puzzle")
	fmt.Printf("Task 1: %d\nTask 2: %d\n", t1_solution, t2_solution)
}
