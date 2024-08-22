package aoc2023

import (
	"fmt"
	"os"
	"strings"
	"testing"

	aoc2023 "aoc.mb/aoc2023/22"
	"aoc.mb/aocutils"
)

func getChallengeData(year string, day string, useSmallData bool) ([]byte, error) {
	var challengeInput []byte
	if useSmallData {
		filename := aocutils.GetDataFileName(year, day)
		filename = strings.Replace(filename, "data", "smalldata", 1)
		filename = aocutils.GetBasePath() + filename
		if !aocutils.FileExists(filename) {
			fmt.Printf("File %s does not exist\n", filename)
			return nil, os.ErrNotExist
		}
		challengeInput = aocutils.ReadDataFromFile(filename)
	} else {
		challengeInput = aocutils.DownloadAocInput(year, day)
	}
	return challengeInput, nil
}

func TestFallingBricksLevel1(t *testing.T) {
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

	filename := aocutils.GetDataFileName(year, day)
	filename = strings.Replace(filename, "data", "endcoords", 1)
	filename = aocutils.GetBasePath() + filename
	groundTruthData := aocutils.ReadDataFromFile(filename)
	finalBricks := aoc2023.GetData(groundTruthData, level)

	maxtime := 10
	allBricks = aoc2023.TimeLoop(allBricks, maxtime, false)

	for i, b := range allBricks {
		b.StopAdvance()
		if b != finalBricks[i] {
			fmt.Println("got     ", b)
			fmt.Println("expected", finalBricks[i])
			t.Errorf("Brick %d differs from ground truth.", i)
		}
	}
}

func constructTwoBricks(xf1, yf1, zf1, xb1, yb1, zb1,
	xf2, yf2, zf2, xb2, yb2, zb2 int) (aoc2023.Brick, aoc2023.Brick) {
	var front1, back1, front2, back2 aoc2023.Cube
	var brick1, brick2 aoc2023.Brick
	front1 = aoc2023.NewCube(xf1, yf1, zf1)
	back1 = aoc2023.NewCube(xb1, yb1, zb1)
	front2 = aoc2023.NewCube(xf2, yf2, zf2)
	back2 = aoc2023.NewCube(xb2, yb2, zb2)

	brick1 = aoc2023.NewBrick("0", front1, back1, false)
	brick2 = aoc2023.NewBrick("1", front2, back2, false)

	return brick1, brick2
}

func TestTouchingBricks(t *testing.T) {
	var brick1, brick2 aoc2023.Brick
	var isTouching bool

	brick1, brick2 = constructTwoBricks(0, 0, 1, 0, 1, 1, 0, 0, 2, 0, 0, 2)
	isTouching = brick1.IsBelow(&brick2)
	if !isTouching {
		t.Fail()
		fmt.Println(brick1)
		fmt.Println(brick2)
		fmt.Println()
	}
	isTouching = brick2.IsAbove(&brick1)
	if !isTouching {
		t.Fail()
		fmt.Println(brick1)
		fmt.Println(brick2)
		fmt.Println()
	}

	brick1, brick2 = constructTwoBricks(1, 0, 1, 1, 2, 1, 0, 0, 2, 2, 0, 2)
	isTouching = brick1.IsBelow(&brick2)
	if !isTouching {
		t.Fail()
		fmt.Println(brick1)
		fmt.Println(brick2)
		fmt.Println()
	}

	brick1, brick2 = constructTwoBricks(1, 0, 1, 1, 2, 1, 0, 2, 2, 2, 2, 2)
	isTouching = brick1.IsBelow(&brick2)
	if !isTouching {
		t.Fail()
		fmt.Println(brick1)
		fmt.Println(brick2)
		fmt.Println()
	}

	brick1, brick2 = constructTwoBricks(0, 1, 6, 2, 1, 6, 1, 1, 7, 1, 1, 8)
	isTouching = brick1.IsBelow(&brick2)
	if !isTouching {
		t.Fail()
		fmt.Println(brick1)
		fmt.Println(brick2)
		fmt.Println()
	}

	brick1, brick2 = constructTwoBricks(2, 0, 5, 2, 2, 5, 0, 1, 6, 2, 1, 6)
	isTouching = brick1.IsBelow(&brick2)
	if !isTouching {
		t.Fail()
		fmt.Println(brick1)
		fmt.Println(brick2)
		fmt.Println()
	}
}
