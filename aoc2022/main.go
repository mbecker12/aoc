package aoc2022

import (
	aoc2022d1 "aoc.mb/aoc2022/1"
	aoc2022d16 "aoc.mb/aoc2022/16"
	aoc2022d21 "aoc.mb/aoc2022/21"
	aoc2022d8 "aoc.mb/aoc2022/8"
)

func Aoc(day string, level int, input []byte) int {
	var ret int
	switch day {
	case "1":
		ret = aoc2022d1.Run(input, level)
	case "8":
		ret = aoc2022d8.Run(input, level)
	case "16":
		ret = aoc2022d16.Run(input, level)
	case "21":
		ret = aoc2022d21.Run(input, level)
	}

	return ret
}
