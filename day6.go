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

type Node struct {
	value    string
	children map[string]*Node
}

func (n *Node) appendNode(value Node) Node {
	n.children[value.value] = &value
	return value
}

func (n *Node) append(value string) Node {
	newChild := Node{value, map[string]*Node{}}
	n.children[value] = &newChild
	return newChild
}

// func findItem(value string, root *Node) *Node {
// 	queue := make([]*Node, 0)
// 	queue = append(queue, root)
// 	for len(queue) != 0 {
// 		next := queue[0]
// 		queue := queue[1:]
// 		if next.value == value {
// 			return next
// 		}
// 		if len(next.children) != 0 {
// 			for _, child := range next.children {
// 				queue = append(queue, child)
// 			}
// 		}
// 	}
// 	return nil
// }

func NewNode(value string) Node {
	return Node{value, map[string]*Node{}}
}

func hasChild(node *Node, value string) bool {
	return findNode(node, value) != nil
}

func addNode(root *Node, value string, childNode *Node) {
	parent := findNode(root, value)
	if parent == nil {
		fmt.Errorf("This is technically an error")
	}
	parent.children[childNode.value] = childNode
}

func findNode(root *Node, value string) *Node {
	if root == nil {
		return nil
	}
	if root.value == value {
		return root
	}

	for _, child := range root.children {
		search := findNode(child, value)
		if search != nil {
			return search
		}
	}

	return nil
}

func strToTree(datum string) Node {
	data := strings.Fields(datum)
	fmt.Printf("DUDE: %d\n", len(data))
	var root *Node
	for _, relationship := range data {
		orbits := strings.Split(relationship, ")")
		parentValue := orbits[0]
		childValue := orbits[1]
		fmt.Printf("%s ~-> %s\n", parentValue, childValue)
		// find the parent. if we can't find the parent, then
		// we add it as the root
		// if we can find the parent, then we add the child
		if !hasChild(root, parentValue) {
			rootNode := NewNode(parentValue)
			root = &rootNode
		}
		child := NewNode(childValue)
		addNode(root, parentValue, &child)
	}

	return *root
}

func strToGraph(datum string) Graph {
	data := strings.Fields(datum)
	fmt.Printf("DUDE: %d\n", len(data))
	var root Graph = Graph{map[string]*GraphNode{}}
	for _, relationship := range data {
		orbits := strings.Split(relationship, ")")
		childValue := orbits[0]
		parentValue := orbits[1]
		root.addNode(childValue)
		root.addNode(parentValue)
		root.addEdge(parentValue, childValue)
		// fmt.Printf("%s ~-> %s\n", parentValue, childValue)
	}

	return root
}

func countOrbits(graph *Graph, key string, level int) int {
  if len(graph.edgeList[key].edges)	== 0 {
		return level
	}

	count := 0

	for childKey, hasEdge := range graph.edgeList[key].edges {
		if hasEdge {
			count += countOrbits(graph, childKey, level + 1)
		}
	}

	return count
}

func sumOrbits(graph *Graph, level int) int {
	fmt.Println(graph)
	count := 0
	for key, child := range graph.edgeList {
		fmt.Println(key, child)
		count += countOrbits(graph, key, 0)
	}
	return count
}

func printTree(node *Node, count int, level int) int {

	if node == nil {
		fmt.Printf("Just count %d\n", count)
		return count
	}

	for _, child := range node.children {
		fmt.Printf("New level: %d\n", level+1)
		count = printTree(child, count, level+1)
	}
	fmt.Printf("Adding %d\n", count+level)
	return count + level
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
	// start at the nodes, and walk up until they're seen!

	dist := 0
	seenRight := []string{pointA}		
	seenLeft := []string{pointB}		
	nodeRight := graph.findNode(pointA)
	nodeLeft  := graph.findNode(pointB)
	
	for nodeLeft != nil || nodeRight != nil {
		if nodeLeft != nil {
			for key, _ := range nodeLeft.edges {
				seenLeft = append(seenLeft, key)
				foundIndex := contains(seenRight, key)
				if foundIndex != -1 {
				  fmt.Printf("We had: %v, %v\n", seenRight, seenLeft)
					fmt.Printf("We found the top ancestor! %s\n", key)
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
					fmt.Printf("We found the top ancestor! %s\n", key)
				  fmt.Printf("We had: %v, %v\n", seenRight, seenLeft)
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
	// fmt.Printf("Graph: %v\n", graph)
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
