package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"aoc.mb/aoc2022"
	"aoc.mb/aocutils"
)

func getChallengeData(year string, day string, useSmallData bool) ([]byte, error) {
	var challengeInput []byte
	if useSmallData {
		filename := aocutils.GetDataFileName(year, day)
		filename = strings.Replace(filename, "data", "smalldata", 1)
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

func main() {
	var year string = "2022"
	var level int
	var day string
	var dryrun bool
	var useSmallData bool

	flag.IntVar(&level, "level", 1, "AoC Level")
	flag.StringVar(&day, "day", "1", "AoC Day")
	flag.BoolVar(&dryrun, "dryrun", false, "Dryrun option")
	flag.BoolVar(&useSmallData, "smalldata", false, "Toggle uising small example data if it exists in the folder")
	flag.Parse()

	fmt.Println(year, day, level)

	challengeInput, err := getChallengeData(year, day, useSmallData)
	if err != nil {
		fmt.Println("Error detected while loading data.")
		os.Exit(-1)
	}
	if challengeInput != nil {
		aocutils.WriteDaysDataToFile(year, day, challengeInput)
	}
	fmt.Println("Downloaded Input")
	fmt.Println()

	answer := aoc2022.Aoc(day, level, challengeInput)

	aocutils.SubmitAocResult(year, day, level, answer, dryrun)
}
