package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"

	"../aocutil"
)

func main() {
	file, err := os.Open("input07.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	lines := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("A: %v\n", solveA(lines))
	fmt.Printf("B: %v\n", solveB(lines))
}

// solveA solves the first half
func solveA(input []string) int {
	nodeMap := nodeMapMaker(input)
	shinyBagHolders := []string{}
	for _, e := range nodeMap["shiny gold"].inbound {
		shinyBagHolders = append(shinyBagHolders, e.color)
	}
	for i := 0; i < len(shinyBagHolders); i++ {
		edges := nodeMap[shinyBagHolders[i]].inbound
		for _, e := range edges {
			if !aocutil.ContainsString(shinyBagHolders, e.color) {
				shinyBagHolders = append(shinyBagHolders, e.color)
			}
		}
	}
	return len(shinyBagHolders)
}

// solveB solves the second half
func solveB(input []string) int {
	nodeMap := nodeMapMaker(input)
	return countContainers("shiny gold", nodeMap)
}

type edge struct {
	color string
	count int
}

type adjacent struct {
	inbound  []*edge
	outbound []*edge
}

func lineParser(line string) (string, []*edge) {
	edges := []*edge{}
	lineParts := strings.Split(line, " contain ")
	color := strings.Replace(lineParts[0], " bags", "", 1)
	containers := lineParts[1]
	re := regexp.MustCompile(`(?P<count>\d+) (?P<color>[a-z\s]+?) bag`)
	for _, container := range strings.Split(containers, ", ") {
		matchMap := aocutil.GetRegexpMap(re, container)
		if matchMap["count"] != "" {
			newEdge := &edge{matchMap["color"], aocutil.MustAtoi(matchMap["count"])}
			edges = append(edges, newEdge)
		}
	}
	return color, edges
}

func nodeMapMaker(input []string) map[string]adjacent {
	nodeMap := make(map[string]adjacent)
	for _, line := range input {
		color, edges := lineParser(line)
		// if color not in map, add it
		// if it is, set outbound edges
		if v, ok := nodeMap[color]; ok {
			v.outbound = edges
			nodeMap[color] = v
		} else {
			nodeMap[color] = adjacent{nil, edges}
		}
		// set inbound adjacencies
		for _, e := range edges {
			originColor := e.color
			if v, ok := nodeMap[originColor]; ok {
				v.inbound = append(v.inbound, &edge{color, 1})
				nodeMap[originColor] = v
			} else {
				nodeMap[originColor] = adjacent{[]*edge{{color, 1}}, nil}
			}
		}
	}
	return nodeMap
}

func countContainers(c string, m map[string]adjacent) int {
	outboundEdges := m[c].outbound
	if len(outboundEdges) == 0 {
		// nothing in this bag
		return 0
	}
	sum := 0
	for _, e := range outboundEdges {
		// count this bag + everything it contains
		sum += e.count + e.count*countContainers(e.color, m)
	}
	return sum

}
