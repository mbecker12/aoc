package main

import (
	"flag"
	"fmt"

	"aoc.mb/aoc2022"
	"aoc.mb/aocutils"
)

func main() {
	var year string = "2022"
	var level int
	var day string

	flag.IntVar(&level, "level", 1, "AoC Level")
	flag.StringVar(&day, "day", "1", "AoC Day")
	flag.Parse()

	fmt.Println(year, day, level)

	input := aocutils.DownloadAocInput(year, day)
	if input != nil {
		aocutils.WriteDaysDataToFile(year, day, input)
	}
	fmt.Println("Downloaded Input")
	fmt.Println()

	answer := aoc2022.Aoc(day, level, input)

	aocutils.SubmitAocResult(year, day, level, answer)
}
