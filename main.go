package main

import (
	"flag"
	"fmt"
	"strconv"

	"aoc.mb/aoc2022"
	"aoc.mb/aocutils"
)

func main() {
	var year string = "2022"
	var day string = "1"
	// var level int = 1

	levelAddress := flag.String("level", "1", "AoC Level")
	flag.Parse()

	fmt.Println(year, day)

	input := aocutils.DownloadAocInput(year, day)
	if input != nil {
		aocutils.WriteDaysDataToFile(year, day, input)
	}
	fmt.Println("Downloaded Input")
	fmt.Println()

	level, _ := strconv.Atoi(*levelAddress)
	answer := aoc2022.Aoc(day, level, input)

	aocutils.SubmitAocResult(year, day, level, answer)
}
