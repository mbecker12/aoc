package aoc2023

import (
	aoc2023d22 "aoc.mb/aoc2023/22"
)

func Aoc(day string, level int, input []byte) int {
	var ret int
	switch day {
	case "22":
		ret = aoc2023d22.Run(input, level)
	}

	return ret
}
