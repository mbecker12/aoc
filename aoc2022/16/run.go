package aoc2022

import (
	"fmt"
	"math"
	"slices"

	"github.com/hmdsefi/gograph"
	"gonum.org/v1/gonum/mat"
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

func initTransitionMatrix(dim int) *mat.Dense {
	// prepare init value as +Infinity
	filler := make([]float64, dim*dim)
	for i := 0; i < dim*dim; i++ {
		filler[i] = math.Inf(1)
	}
	return mat.NewDense(dim, dim, filler)
}

func calcTransitionMatrix(graph gograph.Graph[string], dictLabelToIndex map[string]int, matrix *mat.Dense) *mat.Dense {
	var sourceIndex, destIndex int
	edges := graph.AllEdges()
	// init weights
	for _, edge := range edges {
		sourceIndex = dictLabelToIndex[edge.Source().Label()]
		destIndex = dictLabelToIndex[edge.Destination().Label()]
		matrix.Set(sourceIndex, destIndex, edge.Weight())
	}

	// find all connections
	dim := int(graph.Order())
	for k := 0; k < dim; k++ {
		for i := 0; i < dim; i++ {
			for j := 0; j < dim; j++ {
				if matrix.At(i, k)+matrix.At(k, j) < matrix.At(i, j) {
					matrix.Set(i, j, matrix.At(i, k)+matrix.At(k, j))
				}
			}
		}
	}

	return matrix
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 16, Level %d\n", level)
	graph := saveData(input)
	// currentVertex := graph.GetVertexByID("AA")
	dim := int(graph.Order())
	fmt.Println("dim", dim)
	dictLabelToIndex, _ := mapVertexToIndex(graph)
	matrix := initTransitionMatrix(dim)
	matrix = calcTransitionMatrix(graph, dictLabelToIndex, matrix)
	fmt.Println(matrix)

	// Printing out results
	for i := 0; i < dim; i++ {
		fmt.Println(fmt.Sprintf("Distance from node %v to:", i))
		for j := 0; j < dim; j++ {
			fmt.Println(fmt.Sprintf("  %v: %v", j, matrix.At(i, j)))
		}
	}
	fmt.Println(dictLabelToIndex)

	// // maxtime := 5
	// // openedValves := make(map[string]int)

	// // pressure := traverseGraph[string](currentVertex, openedValves, 0, 0, maxtime)
	// fmt.Println("pressure:", pressure)
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
