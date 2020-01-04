package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

func TestAssertExample(t *testing.T) {
	assert.True(t, true, "True is true!")
}

func Day3TestingStub(t *testing.T) {
	fmt.Println("We are printing something.")
	dat := "Hi"
	if dat != "" {
		t.Errorf("There was an error:\n '%s'", dat)
	}
}

const data0 = "R8,U5,L5,D3\nU7,R6,D4,L4"
const data1 = "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"
const data2 = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

func TestManhattanDistance(t *testing.T) {
	dist0 := getManhattanDistance(data0)
	if dist0 != 6 {
		t.Errorf("Expected 6, got %d instead", dist0)
	}

	dist1 := getManhattanDistance(data1)
	if dist1 != 159 {
		t.Errorf("Expected 159, got %d instead", dist1)
	}
	dist2 := getManhattanDistance(data2)
	if dist2 != 135 {
		t.Errorf("Expected 135, got %d instead", dist2)
	}
}

func TestIntersect(t *testing.T) {
	point1 := []Point{
		Point{5, 2},
		Point{5, -2},
	}

	point2 := []Point{
		Point{2, 0},
		Point{12, 0},
	}

	inter1, hasInter1 := getIntersection(point1, point2)
	if !hasInter1 {
		t.Errorf("Should have detected an intersection: %v", inter1)
	}
	if inter1.x != 5 || inter1.y != 0 {
		t.Errorf("We got the wrong coordinate: %v", inter1)
	}

	inter2, hasInter2 := getIntersection(point2, point1)
	if !hasInter2 {
		t.Errorf("Should have detected an intersection: %v", inter2)
	}
	if inter2.x != 5 || inter2.y != 0 {
		t.Errorf("We got the wrong coordinate: %v", inter2)
	}
}

func TestIntersect2(t *testing.T) {
	point1 := []Point{
		Point{2, 2},
		Point{2, -2},
	}

	point2 := []Point{
		Point{-1, 0},
		Point{5, 0},
	}

	inter2, hasInter2 := getIntersection(point2, point1)
	if !hasInter2 {
		t.Errorf("Should have detected an intersection: %v", inter2)
	}
	if inter2.x != 2 || inter2.y != 0 {
		t.Errorf("We got the wrong coordinate: %v, expected [2, 0]", inter2)
	}

	inter1, hasInter1 := getIntersection(point1, point2)
	if !hasInter1 {
		t.Errorf("Should have detected an intersection: %v", inter1)
	}
	if inter1.x != 2 || inter1.y != 0 {
		t.Errorf("We got the wrong coordinate: %v", inter1)
	}

}

func TestNonIntersect(t *testing.T) {
	point1 := []Point{
		Point{10, 2},
		Point{10, -2},
	}

	point2 := []Point{
		Point{2, 2},
		Point{2, -5},
	}

	inter3, hasInter2 := getIntersection(point2, point1)
	if hasInter2 {
		t.Errorf("Should NOT have detected an intersection: %v", inter3)
	}
}

func TestIntersectReal(t *testing.T) {
	point1 := []Point{
		Point{8, 5},
		Point{3, 5},
	}

	point2 := []Point{
		Point{0, 0},
		Point{0, 7},
	}

	inter1, hasInter1 := getIntersection(point1, point2)
	if hasInter1 {
		t.Errorf("Should not have detected an intersection: %v", inter1)
	}
}
func TestIntersectReal2(t *testing.T) {
	point1 := []Point{
		Point{158, 53},
		Point{146, 53},
	}

	point2 := []Point{
		Point{100, 117},
		Point{100, 46},
	}

	inter1, hasInter1 := getIntersection(point1, point2)
	if hasInter1 {
		t.Errorf("Should not have detected an intersection: %v", inter1)
	}
}

func TestWireDistance(t *testing.T) {
	dist0 := getWireDistance(data0)
	if dist0 != 30 {
		t.Errorf("Expected 30, got %d instead", dist0)
	}

	dist2 := getWireDistance(data2)
	if dist2 != 410 {
		t.Errorf("Expected 410, got %d instead", dist2)
	}

	dist1 := getWireDistance(data1)
	if dist1 != 610 {
		t.Errorf("Expected 610, got %d instead", dist1)
	}
}

func TestLineLength(t *testing.T) {
	dist := getLineLength(Point{0, 0}, Point{5, 0})
	if dist != 5 {
		t.Errorf("Expected 5, got %d instead", dist)
	}
}

func TestLineLengthRoundsValues(t *testing.T) {
	dist := getLineLength(Point{0, 0}, Point{5, 5})
	if dist != int(math.Round(math.Sqrt(50))) {
		t.Errorf("Expected 7, got %d instead", dist)
	}
}

func TestDetermineWireDistance(t *testing.T) {
	wires := ParseWireInput(data0)
	wire1 := wires[0]
	wire2 := wires[1]
	dist := determineWireDistance(wire1)
	expected := 21
	if dist != expected {
		t.Errorf("Expected %d, got %d instead", expected, dist)
	}
	dist = determineWireDistance(wire2)
	expected = 21
	if dist != expected {
		t.Errorf("Expected %d, got %d instead", expected, dist)
	}
}
