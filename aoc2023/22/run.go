package aoc2023

import (
	"fmt"
	"sort"
	"time"

	mapset "github.com/deckarep/golang-set/v2"
)

type Memo4 map[[4]int]bool
type Memo3 map[[3]int]bool
type Memo8 map[[8]int]bool

var rangeCache Memo4
var xyCache Memo8
var oneDimCache Memo3
var Caching bool

func isInRange(x, start, end int) bool {
	var result bool
	arr := [...]int{x, start, end}
	if Caching {
		v, ok := oneDimCache[arr]
		if ok {
			return v
		}
	}
	result = x >= start && x <= end
	if Caching {
		oneDimCache[arr] = result
	}
	return result
}

func IsStripInRange(a, b, c, d int) bool {
	return isStripInRange(a, b, c, d)
}

func isStripInRange(a, b, c, d int) bool {
	arr := [...]int{a, b, c, d}
	if Caching {
		v, ok := rangeCache[arr]
		if ok {
			return v
		}
	}
	result := isInRange(a, c, d) || isInRange(b, c, d) || isInRange(c, a, b) || isInRange(d, a, b)
	if Caching {
		rangeCache[arr] = result
	}
	return result
}

func (b1 *Brick) isXYOverlap(b2 Brick) bool {
	arr := [...]int{b1.front.x, b1.front.y, b1.back.x, b1.back.y,
		b2.front.x, b2.front.y, b2.back.x, b2.back.y}
	if Caching {
		v, ok := xyCache[arr]
		if ok {
			return v
		}
	}
	var result bool
	if isInRange(b1.front.x, b2.front.x, b2.back.x) {
		result = isStripInRange(b1.front.y, b1.back.y, b2.front.y, b2.back.y)
		if Caching {
			xyCache[arr] = result
		}
		return result
	} else if isInRange(b2.front.x, b1.front.x, b1.back.x) {
		result = isStripInRange(b2.front.y, b2.back.y, b1.front.y, b1.back.y)
		if Caching {
			xyCache[arr] = result
		}
		return result
	} else if isInRange(b1.front.y, b2.front.y, b2.back.y) {
		result = isStripInRange(b1.front.x, b1.back.x, b2.front.x, b2.back.x)
		if Caching {
			xyCache[arr] = result
		}
		return result
	} else if isInRange(b2.front.y, b1.front.y, b1.back.y) {
		result = isStripInRange(b2.front.x, b2.back.x, b1.front.x, b1.back.x)
		if Caching {
			xyCache[arr] = result
		}
		return result
	}
	result = false
	if Caching {
		xyCache[arr] = result
	}
	return result
}

func (b1 *Brick) IsBelow(b2 *Brick) bool {
	// check if b1 is exactly below b2

	// check if z coordinates could be touching each other
	if b1.front.z-b2.front.z == -1 ||
		b1.front.z-b2.back.z == -1 ||
		b1.back.z-b2.front.z == -1 ||
		b1.back.z-b2.back.z == -1 {
		return b1.isXYOverlap(*b2)
	} else {
		return false
	}

}

func (b1 *Brick) IsAbove(b2 *Brick) bool {
	// check if b1 is exactly above b2
	return b2.IsBelow(b1)
}

func (brick *Brick) canBrickAdvance(bricks []Brick) {
	var lowestBrickCoord int
	lowestBrickCoord = brick.front.z

	if lowestBrickCoord <= GroundLevel {
		brick.canAdvance = false
		return
	}
	for _, b := range bricks {
		if b.name == brick.name {
			continue
		}
		if brick.front.z > b.back.z+1 {
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

func CountDisintegrableBricks(bricks []Brick) int {
	// A brick is disintegrable if there are
	// 1. No bricks directly above it
	// or
	// 2. The brick(s) above are held by at least one other brick
	// if the brick in question is disintegrated

	supportedBy := make(map[Brick][]Brick)            // <brick-id> : list of bricks supporting <brick-id>
	doubleSupportingBricks := mapset.NewSet[string]() // set of bricks that are at least doubly supporting
	uniqueSupportingBricks := mapset.NewSet[string]() // set of Bricks that are the lone supporting bricks to some
	for _, b := range bricks {                        // init map; init each <brick-id> with empty list of supporters
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
			doubleSupportingBricks.Add(b.name)
		}
	}

	for _, supportingBricks := range supportedBy {
		if len(supportingBricks) == 1 {
			uniqueSupportingBricks.Add(supportingBricks[0].name)
		}
		if len(supportingBricks) > 1 {
			for _, x := range supportingBricks {
				doubleSupportingBricks.Add(x.name)
			}
		}
	}
	doubleSupportingBricks = doubleSupportingBricks.Difference(uniqueSupportingBricks)
	return doubleSupportingBricks.Cardinality()
}

func GatherSupportedBricks(bricks []Brick) map[Brick][]Brick {
	// TODO
	// Gather which bricks are supported by a given brick
	supporting := make(map[Brick][]Brick) // <brick-id> : list of bricks that <brick-id> is supporting

	for _, b := range bricks { // init map; init each <brick-id> with empty list of supporters
		supporting[b] = []Brick{}
	}

	for _, b := range bricks {
		for _, c := range bricks {
			if b.name == c.name {
				continue
			}
			if c.IsBelow(&c) {
				supporting[c] = append(supporting[c], b)
			}
		}
	}
	return supporting
}

func CountChainReaction(brickPtr *Brick, supporting map[Brick][]Brick, counter int) int {
	ctr := 0
	if brickPtr != nil {
		counter += len(supporting[*brickPtr])
		for _, b := range supporting[*brickPtr] {
			counter += CountChainReaction(&b, supporting, counter)
		}
	}
	for _, v := range supporting {
		ctr += len(v)
		for _, b := range v {
			ctr += CountChainReaction(b, supporting, 0)
		}
	}
	return ctr
}

func prettyPrintSet(s mapset.Set[string]) {
	for x := range s.Iter() {
		fmt.Println(x)
	}
}

func prettyPrintMap(m map[Brick][]Brick) {
	for k, v := range m {
		fmt.Println(k, ":", v)
	}
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 22, Level %d\n", level)
	rangeCache = make(map[[4]int]bool)
	xyCache = make(map[[8]int]bool)
	oneDimCache = make(map[[3]int]bool)
	allBricks := GetData(input, level)
	verbose := false
	Caching = true

	maxtime := 400

	switch level {
	case 1:
		t1 := time.Now()
		allBricks = TimeLoop(allBricks, maxtime, verbose)
		// for _, brick := range allBricks {
		// 	fmt.Println(brick)
		// }
		nDisintegrable := CountDisintegrableBricks(allBricks)
		t2 := time.Now()
		fmt.Println("\nNumber of disintegrable bricks:", nDisintegrable)
		diff := t2.Sub(t1)
		fmt.Println("Execution took", diff)
		fmt.Println()
		// fmt.Println(oneDimCache)
		return nDisintegrable
	case 2:
		fmt.Printf("Level %d not implemented\n\n", level)
		return -1
	default:
		fmt.Printf("Level %d not recognized\n\n", level)
		return -1
	}
}
