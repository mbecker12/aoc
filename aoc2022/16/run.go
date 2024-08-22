package aoc2022

import (
	"fmt"
	"slices"

	"github.com/hmdsefi/gograph"
)

func addPressure[T comparable](openedValves map[T]int, totalPressure int) int {
	for _, v := range openedValves {
		totalPressure += v
	}
	return totalPressure
}

func openValve[T comparable](valve *gograph.Vertex[T], openedValves map[T]int, timestep int) (map[T]int, int) {
	openedValvesCopy := make(map[T]int)
	for k, v := range openedValves {
		openedValvesCopy[k] = v
	}
	openedValvesCopy[valve.Label()] = int(valve.Weight())
	return openedValvesCopy, timestep + 1
}

func isValveOpen[T comparable](valve *gograph.Vertex[T], openedValves map[T]int) bool {
	_, isNeighborOpen := openedValves[valve.Label()]
	return isNeighborOpen
}

func traverseGraph[T comparable](
	currentVertex *gograph.Vertex[T],
	openedValves map[T]int,
	currentPressure int,
	timestep int,
	maxtime int,
) int {
	currentPressure = addPressure[T](openedValves, currentPressure)
	if timestep == maxtime {
		return currentPressure
	}

	neighbors := currentVertex.Neighbors()
	tmpResults := []int{0, 0, 0, 0, 0}

	for i, neighbor := range neighbors {
		if !isValveOpen(neighbor, openedValves) {
			openedValves, timestep = openValve(neighbor, openedValves, timestep)
		}
		tmpresult := traverseGraph[T](neighbor, openedValves, currentPressure, timestep+1, maxtime)
		tmpResults[i] = tmpresult
	}
	return slices.Max(tmpResults)
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 16, Level %d\n", level)
	graph := saveData(input)
	currentVertex := graph.GetVertexByID("AA")

	maxtime := 5
	openedValves := make(map[string]int)

	pressure := traverseGraph[string](currentVertex, openedValves, 0, 0, maxtime)
	fmt.Println("pressure:", pressure)
	switch level {
	case 1:
		fmt.Printf("Level %d not implemented yet\n", level)
		return -1
	case 2:
		fmt.Printf("Level %d not implemented yet\n", level)
		return -1
	default:
		fmt.Printf("Level %d not recognized\n\n", level)
		return -1
	}
}
