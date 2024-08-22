package aoc2023

import (
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

func isInRange(x, start, end int) bool {
	return x >= start && x <= end
}

type Memo4 map[[4]int]bool
type Memo8 map[[8]int]bool

var rangeCache Memo4
var xyCache Memo8

func isStripInRange(a, b, c, d int, cache Memo4) bool {
	arr := [...]int{a, b, c, d}
	v, ok := cache[arr]
	if ok {
		return v
	}
	result := isInRange(a, c, d) || isInRange(b, c, d) || isInRange(c, a, b) || isInRange(d, a, b)
	cache[arr] = result
	return result
}

func (b1 *Brick) isXYOverlap(b2 Brick) bool {
	arr := [...]int{b1.front.x, b1.front.y, b1.back.x, b1.back.y,
		b2.front.x, b2.front.y, b2.back.x, b2.back.y}
	v, ok := xyCache[arr]
	if ok {
		return v
	}
	var result bool
	if isInRange(b1.front.x, b2.front.x, b2.back.x) {
		result = isStripInRange(b1.front.y, b1.back.y, b2.front.y, b2.back.y, rangeCache)
		xyCache[arr] = result
		return result
	} else if isInRange(b2.front.x, b1.front.x, b1.back.x) {
		result = isStripInRange(b2.front.y, b2.back.y, b1.front.y, b1.back.y, rangeCache)
		xyCache[arr] = result
		return result
	} else if isInRange(b1.front.y, b2.front.y, b2.back.y) {
		result = isStripInRange(b1.front.x, b1.back.x, b2.front.x, b2.back.x, rangeCache)
		xyCache[arr] = result
		return result
	} else if isInRange(b2.front.y, b1.front.y, b1.back.y) {
		result = isStripInRange(b2.front.x, b2.back.x, b1.front.x, b1.back.x, rangeCache)
		xyCache[arr] = result
		return result
	}
	result = false
	xyCache[arr] = result
	return result
}

func (b1 *Brick) IsBelow(b2 *Brick) bool {
	// check if b1 is exactly below b2

	// check if z coordinates could be touching each other
	if b1.front.z-b2.front.z == -1 ||
		b1.front.z-b2.back.z == -1 {
	} else {
		return false
	}
	return b1.isXYOverlap(*b2)

}

func (b1 *Brick) IsAbove(b2 *Brick) bool {
	// check if b1 is exactly above b2
	return b2.IsBelow(b1)
}

func (brick *Brick) canBrickAdvance(bricks []Brick) {
	var lowestBrickCoord int
	lowestBrickCoord = min(brick.front.z, brick.back.z)

	if lowestBrickCoord <= GroundLevel {
		brick.canAdvance = false
		return
	}
	for _, b := range bricks {
		if b.name == brick.name {
			continue
		}
		if b.IsBelow(brick) {
			brick.canAdvance = false
			return
		}
	}
	brick.canAdvance = true
}

func (brick *Brick) advanceBrick() {
	brick.front.z -= 1
	brick.back.z -= 1
}

func timestep(bricks []Brick, t int, verbose bool) {
	sort.Slice(bricks, func(i, j int) bool {
		return bricks[i].front.z < bricks[j].front.z
	})
	if verbose {
		fmt.Println("t:", t, "before:", bricks)
	}
	for i, b := range bricks {
		b.canBrickAdvance(bricks)
		bricks[i] = b
		if b.canAdvance {
			if verbose {
				fmt.Println("Brick ", i, "moving. Before:", b)
			}
			b.advanceBrick()
			if verbose {
				fmt.Println("Brick ", i, "moving. After :", b)
			}
			bricks[i] = b
		}
	}
	if verbose {
		fmt.Println("t:", t, "after:", bricks)
		fmt.Println("")
	}
}

func TimeLoop(bricks []Brick, maxtime int, verbose bool) []Brick {
	for t := 0; t < maxtime; t++ {
		timestep(bricks, t, verbose)
	}
	return bricks
}

func countDisintegrableBricks(bricks []Brick) int {
	supportedBy := make(map[Brick][]Brick)
	doubleSupportBricks := mapset.NewSet[string]()
	for _, b := range bricks {
		supportedBy[b] = []Brick{}
	}

	nAboveBrick := 0
	for _, b := range bricks {
		nAboveBrick = 0
		for _, c := range bricks {
			if b.name == c.name {
				continue
			}
			if c.IsAbove(&b) {
				supportedBy[c] = append(supportedBy[c], b)
				nAboveBrick += 1
			}
		}
		if nAboveBrick == 0 {
			doubleSupportBricks.Add(b.name)
		}
	}

	for _, v := range supportedBy {
		if len(v) > 1 {
			for _, x := range v {
				doubleSupportBricks.Add(x.name)
			}
		}
	}
	return doubleSupportBricks.Cardinality()
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 22, Level %d\n", level)
	rangeCache = make(map[[4]int]bool)
	xyCache = make(map[[8]int]bool)
	allBricks := GetData(input, level)
	verbose := true

	maxtime := 5

	switch level {
	case 1:
		t1 := time.Now()
		allBricks = TimeLoop(allBricks, maxtime, verbose)
		for _, brick := range allBricks {
			fmt.Println(brick)
		}
		nDisintegrable := countDisintegrableBricks(allBricks)
		t2 := time.Now()
		fmt.Println("nDisintegrable:", nDisintegrable)
		diff := t2.Sub(t1)
		fmt.Println(diff)
		return nDisintegrable
	case 2:
		fmt.Printf("Level %d not implemented\n\n", level)
		return -1
	default:
		fmt.Printf("Level %d not recognized\n\n", level)
		return -1
	}
}
