package aoc2022

import (
	aoc2022d1 "aoc.mb/aoc2022/1"
)

func Aoc(day string, level int, input []byte) int {
	var ret int
	switch day {
	case "1":
		ret = aoc2022d1.Run(input, level)
	}
	return ret
}