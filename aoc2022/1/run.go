package aoc2022

import (
	"fmt"
	"log"
	"sort"
	"strconv"

	"aoc.mb/aocutils"
)

type elf struct {
	calories []int
}

func (e *elf) SetCalories(d []int) {
	e.calories = d
}

func (e *elf) Calories() []int {
	return e.calories
}

func assembleData(groups []string) []elf {
	var elves = make([]elf, len(groups))
	fmt.Println("Initialized elves")
	n := len(groups)
	for i := 0; i < n; i++ {
		group := aocutils.SplitStringInput(groups[i], "\n")
		var tmp = make([]int, len(group))
		for j := 0; j < len(group); j++ {
			if group[j] == "" {
				continue
			}
			num, err := strconv.Atoi(group[j])
			if err != nil {
				log.Fatalln(err)
			}
			tmp[j] = num
		}
		elves[i].SetCalories(tmp)
	}
	return elves
}

func sum(data []int) int {
	s := 0
	for _, d := range data {
		s += d
	}
	return s
}

func findMax(elves []elf) (int, int) {
	max := 0
	maxIdx := -1

	for i, e := range elves {
		sum := sum(e.Calories())
		if sum >= max {
			max = sum
			maxIdx = i
		}
	}
	return max, maxIdx
}

func findTopNSum(elves []elf, n int) int {
	var allCalories = make([]int, len(elves))
	for i, e := range elves {
		allCalories[i] = sum(e.Calories())
	}

	sort.Ints(allCalories)

	m := len(allCalories)
	s := 0
	for i := m - 1; i >= m-n; i-- {
		s += allCalories[i]
	}
	return s
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 1, Level %d\n", level)
	groups := aocutils.SplitByteInput(input, "\n\n")
	fmt.Println("Split into groups")
	fmt.Println()

	elves := assembleData(groups)
	fmt.Printf("Saved %d elves to structs\n\n", len(elves))

	switch level {
	case 1:
		max, maxIdx := findMax(elves)
		fmt.Printf("Found max = %d at index %d\n\n", max, maxIdx)
		return max
	case 2:
		top3 := findTopNSum(elves, 3)
		fmt.Printf("Found sum of top 3 elves = %d\n\n", top3)

		return top3
	default:
		fmt.Printf("Level %d not recognized\n\n", level)
		return -1
	}
}
