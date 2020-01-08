package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDay6(t *testing.T) {
	assert.False(t, false)
}

func TestTree(t *testing.T) {
	a := Node{"A", map[string]*Node{}}
	b := Node{"B", map[string]*Node{}}
	c := a.append("B")
	assert.Equal(t, a.value, "A")
	assert.Equal(t, a.children["B"].value, "B")
	assert.Equal(t, *a.children["B"], b)
	assert.Equal(t, b, c)
}

func TestFindNode(t *testing.T) {
	a := Node{"A", map[string]*Node{}}
	a.append("B")
	d := a.append("D")
	a.append("C")
	d.append("G")
	r := d.append("R")

	c := findNode(&a, "D")
	assert.Equal(t, "D", c.value, "D")
	assert.Equal(t, "G", c.children["G"].value)
	assert.Equal(t, "R", d.children["R"].value)
	assert.Equal(t, r, *findNode(&a, "R"))
}

func TestGraph(t *testing.T) {
	testGraph := Graph{map[string]*GraphNode{}}
	testGraph.addNode("a")
	testGraph.addNode("b")
	testGraph.addNode("x")
	testGraph.addNode("y")
	testGraph.addEdge("a", "y")
	testGraph.addEdge("a", "x")
	assert.True(t, testGraph.hasNode("a"))
	assert.True(t, testGraph.hasNode("b"))
	assert.False(t, testGraph.hasNode("c"))
	assert.True(t, testGraph.hasEdge("a", "y"))
	assert.True(t, testGraph.hasEdge("a", "x"))
	assert.False(t, testGraph.hasEdge("y", "a"))
	assert.Equal(t, 2, len(testGraph.edgeList["a"].edges))
	assert.Equal(t, 0, len(testGraph.edgeList["y"].edges))
}

func TestMinDist(t *testing.T) {
	programData := `COM)B
B)C
C)D
D)E
E)F
B)G
G)H
D)I
E)J
J)K
K)L
K)YOU
I)SAN`
	graph := strToGraph(programData)
	// fmt.Printf("Graph: %v\n", graph)
	count := getMinDistanceBetween(&graph, "YOU", "SAN")
	assert.Equal(t, count, 4)
}
