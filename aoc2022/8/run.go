package aoc2022

import (
	"fmt"

	"aoc.mb/aocutils"
)

type coordinates struct {
	x int
	y int
}

func saveData(input []byte) [][]int {

	dataStrings := aocutils.SplitByteInput(input, "\n")
	nRows := len(dataStrings) - 1 // disregard empty line at the end
	nCols := len(dataStrings[0])

	dataInts := make([][]int, nRows)
	for i := 0; i < nRows; i++ {
		dataInts[i] = make([]int, nCols)
	}

	for i, row := range dataStrings {
		for j, d := range row {
			dataInts[i][j] = int(d - '0')
		}
	}
	return dataInts
}

func findVisibleInLine(line []int) []int {
	max := -1
	maxIdx := 0
	visibleIdx := []int{}

	for i, x := range line {
		if x < max && i < maxIdx {
			// x is visible
			visibleIdx = append(visibleIdx, i)
		}
		if x > max {
			// new max found
			max = x
			maxIdx = i

			visibleIdx = append(visibleIdx, maxIdx)
			if x == 9 {
				break
			}
		}
	}
	return visibleIdx
}

func findVisibilityInLine(line []int, startIdx int, startHeight int) int {
	// view from left
	length := len(line)
	leftVisibility := length - 1 - startIdx
	rightVisibility := startIdx
	for i := startIdx; i < length; i++ {
		if i == startIdx {
			continue
		}
		if line[i] >= startHeight {
			leftVisibility = i - startIdx
			break
		}
	}
	// view from right
	for i := startIdx; i > 0; i-- {
		if i == startIdx {
			continue
		}
		if line[i] >= startHeight {
			rightVisibility = startIdx - i
			break
		}
	}
	return leftVisibility * rightVisibility
}

func isCoordInList(allCoordinates []coordinates, coord coordinates) bool {
	tmp := false
	for _, c := range allCoordinates {
		if c.x == coord.x && c.y == coord.y {
			return true
		}
	}
	return tmp
}

func addCoordinates(allCoordinates []coordinates, newCoordinates []coordinates) []coordinates {
	for _, newCoord := range newCoordinates {
		if !isCoordInList(allCoordinates, newCoord) {
			allCoordinates = append(allCoordinates, newCoord)
		}
	}
	return allCoordinates
}

func findVisible(data [][]int) []coordinates {
	var visibleCoord []int
	nRows := len(data)
	nCols := len(data[0])

	// view from left
	var coordsLeft []coordinates
	for i := 0; i < nRows; i++ {
		visibleCoord = findVisibleInLine(data[i])
		for _, y := range visibleCoord {
			coordsLeft = append(coordsLeft, coordinates{i, y})
		}
	}

	// view from right
	var coordsRight []coordinates
	for i := 0; i < nRows; i++ {
		var auxData []int = make([]int, nCols)
		for a := nCols - 1; a >= 0; a-- {
			auxData[nCols-a-1] = data[i][a]
		}
		visibleCoord = findVisibleInLine(auxData)
		// fmt.Println(visibleCoord)
		for _, y := range visibleCoord {
			coordsRight = append(coordsRight, coordinates{i, nCols - y - 1})
		}
	}

	// view from top
	var coordsTop []coordinates
	for j := 0; j < nCols; j++ {
		var auxData []int = make([]int, nRows)
		for a := 0; a < nRows; a++ {
			auxData[a] = data[a][j]
		}
		visibleCoord = findVisibleInLine(auxData)
		// invert x and y
		for _, x := range visibleCoord {
			coordsTop = append(coordsTop, coordinates{x, j})
		}
	}
	// view from bottom
	var coordsBottom []coordinates
	for j := 0; j < nCols; j++ {
		var auxData []int = make([]int, nRows)
		for a := nRows - 1; a >= 0; a-- {
			auxData[nRows-a-1] = data[a][j]
		}
		visibleCoord = findVisibleInLine(auxData)
		// invert x and y
		for _, x := range visibleCoord {
			coordsBottom = append(coordsBottom, coordinates{nRows - x - 1, j})
		}
	}

	var allCoordinates []coordinates
	allCoordinates = addCoordinates(coordsLeft, coordsRight)
	allCoordinates = addCoordinates(allCoordinates, coordsTop)
	allCoordinates = addCoordinates(allCoordinates, coordsBottom)

	return allCoordinates
}

func findVisibility(data [][]int) int {
	var maxVisibility int
	nRows := len(data)
	nCols := len(data[0])
	var auxData []int = make([]int, nRows)

	for i := 0; i < nRows; i++ {
		for j := 0; j < nCols; j++ {
			horizontalData := data[i]
			horizontalIdx := j
			horizontalVal := data[i][j]

			for a := 0; a < nRows; a++ {
				auxData[a] = data[a][j]
			}
			verticalData := auxData
			verticalIdx := i
			verticalVal := auxData[i]

			// horizontal
			horizontalScore := findVisibilityInLine(horizontalData, horizontalIdx, horizontalVal)

			// vertical
			verticalScore := findVisibilityInLine(verticalData, verticalIdx, verticalVal)
			if horizontalScore*verticalScore > maxVisibility {
				maxVisibility = horizontalScore * verticalScore
			}
		}
	}
	return maxVisibility
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 8, Level %d\n", level)
	data := saveData(input)

	switch level {
	case 1:
		visible := findVisible(data)
		nVisible := len(visible)
		fmt.Printf("There are %d trees visible.\n\n", nVisible)
		return nVisible
	case 2:
		visibility := findVisibility(data)
		fmt.Printf("Max visibility score: %d\n\n", visibility)
		// os.Exit(1)
		return visibility
	default:
		fmt.Printf("Level %d not recognized\n\n", level)
		return -1
	}
}
