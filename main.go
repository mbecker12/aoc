package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"aoc.mb/aoc2022"
	"aoc.mb/aoc2023"
	"aoc.mb/aocutils"
)

func getChallengeData(year string, day string, useSmallData bool) ([]byte, error) {
	var challengeInput []byte
	if useSmallData {
		filename := aocutils.GetBasePath() + aocutils.GetDataFileName(year, day)
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

func cleanupSubmissionFiles(year, day string, level int) {
	fmt.Println("TODO: remove old submission files, keeping the latest one for each level in the given folder for")
	fmt.Printf("Year %s, Day %s, Level %d.", year, day, level)
}

func main() {
	var year string
	var level int
	var day string
	var dryrun, useSmallData, cleanupSubmissionData bool

	flag.IntVar(&level, "level", 1, "AoC Level")
	flag.StringVar(&day, "day", "1", "AoC Day")
	flag.StringVar(&year, "year", "2022", "AoC Year")
	flag.BoolVar(&dryrun, "dryrun", false, "Dryrun option")
	flag.BoolVar(&useSmallData, "smalldata", false, "Toggle using small example data if it exists in the folder")
	flag.BoolVar(&cleanupSubmissionData, "cleanup", false, "Toggle cleanup of old submission cache files. Will keep the latest submission file. Skips execution of the AoC problem.")
	flag.Parse()

	fmt.Println("Year", year, "\nDay", day, "\nLevel", level)

	if cleanupSubmissionData {
		fmt.Printf("Cleaning up submission data for Year %s, Day %s, Level %d. Skipping code execution.\n", year, day, level)
		cleanupSubmissionFiles(year, day, level)
		os.Exit(0)
	}

	challengeInput, err := getChallengeData(year, day, useSmallData)
	if err != nil {
		fmt.Println("Error detected while loading data.")
		os.Exit(-1)
	}
	if challengeInput != nil && !useSmallData {
		aocutils.WriteDaysDataToFile(year, day, challengeInput)
	}
	fmt.Println("Downloaded Input")
	fmt.Println()

	var answer int
	switch year {
	case "2022":
		answer = aoc2022.Aoc(day, level, challengeInput)
	case "2023":
		answer = aoc2023.Aoc(day, level, challengeInput)
	}

	aocutils.SubmitAocResult(year, day, level, answer, dryrun)
}
