package aoc2023

import (
	"fmt"
	"os"
	"strings"
	"testing"

	aoc2023 "aoc.mb/aoc2023/22"
	"aoc.mb/aocutils"
)

// const Caching = false

func BenchmarkIsBelow(b *testing.B) {
	var brick1, brick2 aoc2023.Brick
	// rangeCache := make(map[[4]int]bool)
	// xyCache := make(map[[8]int]bool)

	brick1, brick2 = constructTwoBricks(0, 0, 1, 0, 1, 1, 0, 0, 2, 0, 0, 2)
	for i := 0; i < b.N; i++ {
		brick1.IsBelow(&brick2)
	}
}

func BenchmarkIsStripInRange(b *testing.B) {
	for i := 0; i < b.N; i++ {
		aoc2023.IsStripInRange(0, 1, 2, 3)
	}
}

func BenchmarkTimeStep(b *testing.B) {
	year := "2023"
	day := "22"
	challengeInput, err := getChallengeData(year, day, true)
	if err != nil {
		fmt.Println("Error detected while loading data.")
		os.Exit(-1)
	}
	if challengeInput != nil {
		aocutils.WriteDaysDataToFile(year, day, challengeInput)
	}
	fmt.Println("Downloaded Input")
	fmt.Println()

	level := 1
	fmt.Printf("Day 22, Level %d\n", level)
	allBricks := aoc2023.GetData(challengeInput, level)

	maxtime := 10
	for i := 0; i < b.N; i++ {
		aoc2023.TimeLoop(allBricks, maxtime, false)
	}
}

func BenchmarkDisintegrable(b *testing.B) {
	year := "2023"
	day := "22"
	level := 1
	filename := aocutils.GetDataFileName(year, day)
	filename = strings.Replace(filename, "data", "endcoords", 1)
	filename = aocutils.GetBasePath() + filename
	groundTruthData := aocutils.ReadDataFromFile(filename)

	finalBricks := aoc2023.GetData(groundTruthData, level)
	for i := 0; i < b.N; i++ {
		aoc2023.CountDisintegrableBricks(finalBricks)
	}

}
