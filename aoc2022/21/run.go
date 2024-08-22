package aoc2022

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"

	"aoc.mb/aocutils"
)

var GroundLevel = 1

type monkey struct {
	name     string
	number   string
	operand  string
	ref1Name string
	ref2Name string
	ref1     *monkey
	ref2     *monkey
}

func getData(input []byte, level int) []monkey {
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
				number:   "empty",
				operand:  matchingOperand,
				ref1Name: strings.TrimSpace(refNames[0]),
				ref2Name: strings.TrimSpace(refNames[1]),
			}
		} else {
			number := strings.TrimSpace(splitString[1])
			m = monkey{
				name:   name,
				number: number,
			}
		}
		if m.name == "humn" && level == 2 {
			m.number = "x"
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
	if m.ref1.ref1 != nil {
		m.ref1.arithmetic(depth + 1)
	}
	if m.ref2.ref1 != nil {
		m.ref2.arithmetic(depth + 1)
	}
	m.number = "(" + m.ref1.number + m.operand + m.ref2.number + ")"
}

func SympySolve(result1, result2 string) string {
	cmd := exec.Command(
		aocutils.PythonBinPath,
		"/home/marvin/projects/aoc/aoc2022/21/solve.py",
		fmt.Sprintf("%s", result1),
		fmt.Sprintf("%s", result2),
	)

	var outbuf, errbuf strings.Builder
	cmd.Stdout = &outbuf
	cmd.Stderr = &errbuf
	err := cmd.Run()
	stderr := errbuf.String()
	stdout := outbuf.String()
	if err != nil {
		fmt.Println(stderr)
		fmt.Println("err:", err)
		panic(err)
	}
	return stdout
}

func convertSympyResult(sympySolution string) int {
	sympySolutiontrimmed := strings.TrimSpace(sympySolution)
	result, _ := strconv.Atoi(sympySolutiontrimmed)
	return result
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 21, Level %d\n", level)
	allMonkeys := getData(input, level)
	allRefMonkeys := make([]monkey, len(allMonkeys))
	var result string

	// Big learning:
	// It's important to not iterate using 'range',
	// otherwise all objects get the same address
	for i := 0; i < len(allMonkeys); i++ {
		allMonkeys[i].findReferences(allMonkeys)
		allRefMonkeys[i] = allMonkeys[i]
	}

	switch level {
	case 1:
		for i := 0; i < len(allRefMonkeys); i++ {
			if allRefMonkeys[i].name == "root" {
				allRefMonkeys[i].arithmetic(0)
				result = allRefMonkeys[i].number
				break
			}
		}
		sympySolution := SympySolve(result, "level1")
		resultLevel1 := convertSympyResult(sympySolution)
		fmt.Printf("Root monkey contains result %d.\n\n", resultLevel1)
		return 1
	case 2:
		var result1, result2 string
		for i := 0; i < len(allRefMonkeys); i++ {
			if allRefMonkeys[i].name == "root" {
				ref1Name := allRefMonkeys[i].ref1.name
				ref2Name := allRefMonkeys[i].ref2.name

				for j := 0; j < len(allRefMonkeys); j++ {
					if allRefMonkeys[j].name == ref1Name {
						allRefMonkeys[j].arithmetic(0)
						result1 = allRefMonkeys[j].number
						break
					}
				}
				for j := 0; j < len(allRefMonkeys); j++ {
					if allRefMonkeys[j].name == ref2Name {
						allRefMonkeys[j].arithmetic(0)
						result2 = allRefMonkeys[j].number
						break
					}
				}
				break
			}
		}
		sympySolution := SympySolve(result1, result2)
		resultLevel2 := convertSympyResult(sympySolution)

		fmt.Println("'humn' should take on the value", resultLevel2)

		return resultLevel2
	default:
		fmt.Printf("Level %d not recognized\n\n", level)
		return -1
	}
}
