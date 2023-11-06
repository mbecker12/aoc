package aoc2022

import (
	"fmt"
	"math"
	"strconv"
	"strings"

	"aoc.mb/aocutils"
)

type monkey struct {
	name     string
	number   int
	operand  string
	ref1Name string
	ref2Name string
	ref1     *monkey
	ref2     *monkey
}

func getData(input []byte) []monkey {
	dataStrings := aocutils.SplitByteInput(input, "\n")
	operands := "+-*/"
	refNames := make([]string, 2)
	var matchingOperand string
	allMonkeys := make([]monkey, len(dataStrings)-1)
	for i, monkeyString := range dataStrings {
		var m monkey
		if monkeyString == "" {
			// ignore newline at the end
			continue
		}
		splitString := strings.Split(monkeyString, ":")
		name := strings.TrimSpace(splitString[0])

		if strings.ContainsAny(splitString[1], operands) {
			for _, op := range operands {
				if strings.Contains(splitString[1], string(op)) {
					refNames = strings.Split(splitString[1], string(op))
					matchingOperand = string(op)
				}
			}
			m = monkey{
				name:     name,
				number:   int(math.Inf(1)) - 1,
				operand:  matchingOperand,
				ref1Name: strings.TrimSpace(refNames[0]),
				ref2Name: strings.TrimSpace(refNames[1]),
			}
		} else {
			number, _ := strconv.Atoi(strings.TrimSpace(splitString[1]))
			m = monkey{
				name:   name,
				number: number,
			}
		}
		allMonkeys[i] = m
	}
	return allMonkeys
}

func (m *monkey) findReferences(monkeys []monkey) {
	// assume that if ref1Name is non-empty, so is ref2Name
	if m.ref1Name != "" {
		for i := 0; i < len(monkeys); i++ {
			if m.ref1Name == monkeys[i].name {
				m.ref1 = &monkeys[i]
				continue
			}
			if m.ref2Name == monkeys[i].name {
				m.ref2 = &monkeys[i]
				continue
			}
		}
	}
}

func (m *monkey) arithmetic(depth int) {
	/*
		Update m's number according to its references.
	*/
	fmt.Printf("Name: %s, Recursion depth: %d, ref1: %s, ref2: %s\n", m.name, depth, m.ref1Name, m.ref2Name)
	fmt.Println("Ref1:")
	fmt.Println(m.ref1)
	if m.ref1.ref1 != nil {
		m.ref1.arithmetic(depth + 1)
	}
	if m.ref2.ref1 != nil {
		m.ref2.arithmetic(depth + 1)
	}
	switch m.operand {
	case "+":
		m.number = m.ref1.number + m.ref2.number
	case "-":
		m.number = m.ref1.number - m.ref2.number
	case "*":
		m.number = m.ref1.number * m.ref2.number
	case "/":
		m.number = m.ref1.number / m.ref2.number
	}
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 21, Level %d\n", level)
	allMonkeys := getData(input)
	allRefMonkeys := make([]monkey, len(allMonkeys))

	// Big learning:
	// It's important to not iterate using 'range',
	// otherwise all objects get the same address
	for i := 0; i < len(allMonkeys); i++ {
		allMonkeys[i].findReferences(allMonkeys)
		allRefMonkeys[i] = allMonkeys[i]
	}

	switch level {
	case 1:
		var result int
		for i := 0; i < len(allRefMonkeys); i++ {
			if allRefMonkeys[i].name == "root" {
				allRefMonkeys[i].arithmetic(0)
				result = allRefMonkeys[i].number
			}
		}
		fmt.Printf("Root monkey contains result %d.\n\n", result)
		return result
	case 2:
		fmt.Printf("Level %d not implemented yet.\n\n", level)
		return -1
	default:
		fmt.Printf("Level %d not recognized\n\n", level)
		return -1
	}
}
