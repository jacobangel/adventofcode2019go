package main

import (
	"fmt"
	"strings"
)

type GraphNode struct {
	edges map[string]bool
	value string
}

type Graph struct {
	edgeList map[string]*GraphNode
}

func (g *Graph) hasEdge(haystack string, needle string) bool {
	if g.hasNode(haystack) {
		node := g.edgeList[haystack]
		hasIt, valueExists := node.edges[needle]
		return valueExists && hasIt
	}

	return false
}

func (g *Graph) findNode(value string) *GraphNode {
	if g.hasNode(value) {
		return g.edgeList[value]
	}
	return nil
}

func (g *Graph) hasNode(value string) bool {
	_, ok := g.edgeList[value]
	return ok
}

func (g *Graph) addNode(value string) {
	if !g.hasNode(value) {
		g.edgeList[value] = &GraphNode{map[string]bool{}, value}
	}
}

func (g *Graph) addEdge(from string, to string) {
	g.addNode(from)
	g.addNode(to)
	g.edgeList[from].edges[to] = true
}

func strToGraph(datum string) Graph {
	data := strings.Fields(datum)
	var root Graph = Graph{map[string]*GraphNode{}}
	for _, relationship := range data {
		orbits := strings.Split(relationship, ")")
		childValue := orbits[0]
		parentValue := orbits[1]
		root.addNode(childValue)
		root.addNode(parentValue)
		root.addEdge(parentValue, childValue)
	}

	return root
}

func countOrbits(graph *Graph, key string, level int) int {
	if len(graph.edgeList[key].edges) == 0 {
		return level
	}

	count := 0

	for childKey, hasEdge := range graph.edgeList[key].edges {
		if hasEdge {
			count += countOrbits(graph, childKey, level+1)
		}
	}

	return count
}

func sumOrbits(graph *Graph, level int) int {
	fmt.Println(graph)
	count := 0
	for key := range graph.edgeList {
		count += countOrbits(graph, key, 0)
	}
	return count
}

func contains(hay []string, needle string) int {
	for index, value := range hay {
		if value == needle {
			return index
		}
	}
	return -1
}

func getMinDistanceBetween(graph *Graph, pointA string, pointB string) int {
	dist := 0
	seenRight := []string{pointA}
	seenLeft := []string{pointB}
	nodeRight := graph.findNode(pointA)
	nodeLeft := graph.findNode(pointB)

	for nodeLeft != nil || nodeRight != nil {
		if nodeLeft != nil {
			for key, _ := range nodeLeft.edges {
				seenLeft = append(seenLeft, key)
				foundIndex := contains(seenRight, key)
				if foundIndex != -1 {
					return len(seenLeft) + foundIndex - 3
					// this -3, because you don't count themselves
				}
				nodeLeft = graph.findNode(key)
			}
		}
		if nodeRight != nil {
			for key, _ := range nodeRight.edges {
				foundIndex := contains(seenLeft, key)
				seenRight = append(seenRight, key)
				if foundIndex != -1 {
					return len(seenRight) + foundIndex - 3
				}
				nodeRight = graph.findNode(key)
			}
		}
	}

	return dist
}

func day61() {

	fmt.Println("Day 6.1")
	programData := loadFile("./day6_input.txt")
	fmt.Println(programData)
	graph := strToGraph(programData)
	fmt.Printf("Graph: %v\n", graph)
	count := sumOrbits(&graph, 0)
	fmt.Printf("The answer is: '%d'\n", count)
}

func day62() {
	fmt.Println("Day 6.2")
	programData := loadFile("./day6_input.txt")
	graph := strToGraph(programData)
	count := getMinDistanceBetween(&graph, "YOU", "SAN")
	fmt.Printf("The answer is: '%d'\n", count)
}

func day6() {
	fmt.Println("Day 6")
	fmt.Println("---------")
	day61()
	day62()
	fmt.Println()
}
