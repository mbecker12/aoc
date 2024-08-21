package aoc2022

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"aoc.mb/aocutils"
	"github.com/hmdsefi/gograph"
)

type valve struct {
	name     string
	flowrate int
	tunnels  []string
}

func extractValve(line string, r *regexp.Regexp) valve {
	var matches []string
	var valvename string
	var flowrate int
	var tunnel1 string
	var tunnel2 string
	var tunnel3 string
	var tunnel4 string
	var tunnel5 string

	matches = r.FindStringSubmatch(line)
	valvename = matches[r.SubexpIndex("valvename")]
	flowrate, _ = strconv.Atoi(matches[r.SubexpIndex("flowrate")])
	tunnel1 = matches[r.SubexpIndex("tunnel1")]
	tunnel1, _ = strings.CutSuffix(tunnel1, ",")

	tunnel2 = matches[r.SubexpIndex("tunnel2")]
	tunnel2, _ = strings.CutSuffix(tunnel2, ",")

	tunnel3 = matches[r.SubexpIndex("tunnel3")]
	tunnel3, _ = strings.CutSuffix(tunnel3, ",")

	tunnel4 = matches[r.SubexpIndex("tunnel4")]
	tunnel4, _ = strings.CutSuffix(tunnel4, ",")

	tunnel5 = matches[r.SubexpIndex("tunnel5")]
	tunnel5, _ = strings.CutSuffix(tunnel5, ",")

	// fmt.Printf("%s: %d.\n#%s#%s#%s#%s#%s\n\n", valvename, flowrate, tunnel1, tunnel2, tunnel3, tunnel4, tunnel5)

	tunnels := make([]string, 5)
	if tunnel1 != "" {
		tunnels = append(tunnels, tunnel1)
	}
	if tunnel2 != "" {
		tunnels = append(tunnels, tunnel2)
	}
	if tunnel3 != "" {
		tunnels = append(tunnels, tunnel3)
	}
	if tunnel4 != "" {
		tunnels = append(tunnels, tunnel4)
	}
	if tunnel5 != "" {
		tunnels = append(tunnels, tunnel5)
	}

	return valve{
		valvename,
		flowrate,
		tunnels,
	}

}

func saveData(input []byte) gograph.Graph[string] {
	lines := aocutils.SplitByteInput(input, "\n")
	r := regexp.MustCompile(
		`^(?P<generic>Valve) (?P<valvename>[A-Z]{2})(?P<moretext>\s[a-zA-Z]+\s[a-zA-Z]+\s[a-zA-Z]+=)(?P<flowrate>\d+);\s[a-zA-Z]+\s[a-zA-Z]+\s[a-zA-Z]+\s[a-zA-Z]+\s(?P<tunnel1>[A-Z]{2},?) ?(?P<tunnel2>[A-Z]{2},?)? ?(?P<tunnel3>[A-Z]{2},?)? ?(?P<tunnel4>[A-Z]{2},?)? ?(?P<tunnel5>[A-Z]{2},?)? ?`)

	var valves = make(map[string]valve)
	// valves = make([]valve, len(lines)-1)
	graph := gograph.New[string](gograph.Directed())

	// init graph vertices
	for _, line := range lines {
		if line == "" {
			// ignore newline at the end
			continue
		}

		_valve := extractValve(line, r)
		// fmt.Println(_valve)
		valves[_valve.name] = _valve
		graph.AddVertexByLabel(_valve.name, gograph.WithVertexWeight(float64(_valve.flowrate)))
	}

	// add directed edges for tunnels
	vertices := graph.GetAllVertices()
	for _, vertex := range vertices {
		currentvalve := valves[vertex.Label()]
		for _, tunnel := range currentvalve.tunnels {
			tunnelVertex := graph.GetVertexByID(tunnel)
			graph.AddEdge(vertex, tunnelVertex, gograph.WithEdgeWeight(1))
		}
	}
	return graph
}

func addPressure[T comparable](openedValves map[T]int, totalPressure int) int {
	for _, v := range openedValves {
		totalPressure += v
	}
	return totalPressure
}

func traverseGraph[T comparable](g gograph.Graph[T], currentVertex *gograph.Vertex[T], openedValves map[T]int, currentPressure int, timestep int, maxtime int) int {
	if timestep == maxtime {
		return addPressure[T](openedValves, currentPressure)
	}
	// currentVertex := g.GetVertexByID(currentVertexName)
	neighbors := currentVertex.Neighbors()

	maxWeight := 0
	var bestNeighbor *gograph.Vertex[T]
	for _, neighbor := range neighbors {
		if int(neighbor.Weight()) > maxWeight {

			for openValve, _ := range openedValves {
				// skip if valve has already been opened
				if openValve == neighbor.Label() {
					continue
				}
			}
			maxWeight = int(neighbor.Weight())
			bestNeighbor = neighbor
		}
	}
	// traverse to next vertex
	openedValves[bestNeighbor.Label()] = int(bestNeighbor.Weight())
	currentPressure = addPressure[T](openedValves, currentPressure)
	return traverseGraph[T](g, bestNeighbor, openedValves, currentPressure, timestep+1, maxtime)
}

func Run(input []byte, level int) int {
	fmt.Printf("Day 16, Level %d\n", level)
	// graph := saveData(input)
	// currentVertex := graph.GetVertexByID("AA")

	// maxtime := 30

	// traverseGraph[string](graph, currentVertex, openedValves, 0, 0, maxtime)
	// fmt.Println(graph)
	// switch level {
	// case 1:
	// 	fmt.Printf("Level %d not implemented yet\n", level)
	// 	return -1
	// case 2:
	// 	fmt.Printf("Level %d not implemented yet\n", level)
	// 	return -1
	// default:
	// 	fmt.Printf("Level %d not recognized\n\n", level)
	// 	return -1
	// }
	return -1
}
