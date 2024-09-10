package aoc2023

import (
	"strconv"
	"strings"

	"aoc.mb/aocutils"
	"golang.org/x/exp/constraints"
)

type Cube struct {
	x, y, z int
}

var GroundLevel int = 1

type Brick struct {
	name             string
	front, back      Cube
	vol              int
	canAdvance       bool
	xlen, ylen, zlen int
}

func NewCube(x, y, z int) Cube {
	return Cube{
		x: x,
		y: y,
		z: z,
	}
}

func (c Cube) GetX() int {
	return c.x
}
func (c Cube) GetY() int {
	return c.y
}
func (c Cube) GetZ() int {
	return c.z
}

func NewBrick(name string, front, back Cube, canAdvance bool) Brick {
	brick := Brick{
		name:       name,
		front:      front,
		back:       back,
		vol:        calcVolume(front, back),
		canAdvance: false,
	}
	brick.calcLen()
	return brick
}

func Abs[T constraints.Signed](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func calcVolume(p1, p2 Cube) int {
	vol := Abs(p1.x-p2.x) + Abs(p1.y-p2.y) + Abs(p1.z-p2.z) + 1
	return vol
}

func (b *Brick) calcLen() {
	b.xlen = Abs(b.front.x - b.back.x)
	b.ylen = Abs(b.front.y - b.back.y)
	b.zlen = Abs(b.front.z - b.back.z)
}

func (b *Brick) GetFront() Cube {
	return b.front
}

func (b *Brick) GetBack() Cube {
	return b.back
}

func (b *Brick) StopAdvance() {
	b.canAdvance = false
}

func GetData(input []byte, level int) []Brick {
	dataStrings := aocutils.SplitByteInput(input, "\n")
	allBricks := make([]Brick, len(dataStrings))

	for i, line := range dataStrings {
		if line == "" {
			// ignore newline at the end
			continue
		}
		var brick Brick
		cubeStrings := strings.Split(line, "~")
		cubes := make([]Cube, 2)
		for j, cubeString := range cubeStrings {
			coordinates := strings.Split(cubeString, ",")
			x, _ := strconv.Atoi(coordinates[0])
			y, _ := strconv.Atoi(coordinates[1])
			z, _ := strconv.Atoi(coordinates[2])
			cube := Cube{
				x: x,
				y: y,
				z: z,
			}
			cubes[j] = cube
		}
		brick = Brick{
			name:       strconv.Itoa(i),
			front:      cubes[0],
			back:       cubes[1],
			vol:        calcVolume(cubes[0], cubes[1]),
			canAdvance: false,
		}
		brick.calcLen()
		allBricks[i] = brick
	}
	return allBricks
}
