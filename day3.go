package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func ParseWireInput(inputData string) [][]Point {
	wireList := strings.Split(inputData, "\n")
	wires := make([][]Point, len(wireList))
	for wireIndex, line := range wireList {
		wires[wireIndex] = ConvertStrToCoords(line)
	}

	return wires
}

type Instruction struct {
	direction string
	step      int
}

func ConvertStrToCoords(line string) []Point {
	wire := make([]Point, 1)
	rawInstructionList := strings.Split(strings.TrimSpace(line), ",")
	lastCoord := Point{0, 0}
	wire = append(wire, lastCoord)
	for _, rawInstruction := range rawInstructionList {
		instruction := strToInstruction(rawInstruction)
		nextCoord := getNextPointFromInstruction(lastCoord, instruction)
		wire = append(wire, nextCoord)
		lastCoord = nextCoord
	}
	return wire
}

func strToInstruction(rawInstruction string) Instruction {
	dir := string(rawInstruction[0])
	step, _ := strconv.Atoi(rawInstruction[1:])
	return Instruction{dir, step}
}

func getNextPointFromInstruction(previousCoord Point, instr Instruction) Point {
	newCoord := Point{previousCoord.x, previousCoord.y}
	step := instr.step
	switch instr.direction {
	case "U":
		newCoord.y += step
	case "D":
		newCoord.y -= step
	case "L":
		newCoord.x -= step
	case "R":
		newCoord.x += step
	}
	return newCoord
}

func getManhattanDistance(inputData string) int {
	wires := ParseWireInput(inputData)
	wireA := wires[0]
	wireB := wires[1]
	// fmt.Printf("Wire A: %v\nWire B: %v\n", wireA, wireB)
	distance := 100000000
	// This is where we would do a min heap or some shit like that.
	for i := 0; i < len(wireA)-1; i++ {
		line := []Point{wireA[i], wireA[i+1]}
		intersections, hasIntersect := getIntersectionList(line, wireB)
		if hasIntersect {
			for _, intersection := range intersections {
				localDistance := determineManhattanDistance(intersection)
				if localDistance < distance && localDistance != 0 {
					// fmt.Printf("Found new winner: %v, LD: %d\n", intersection, localDistance)
					distance = localDistance
				}
			}
		}

	}

	return distance
}

func determineManhattanDistance(point Point) int {
	return Abs(point.x) + Abs(point.y)
}

func getWireDistance(inputData string) int {
	wires := ParseWireInput(inputData)
	wireA := wires[0]
	wireB := wires[1]
	return getShortestIntersections(wireA, wireB)
}

func calculateWireLength(wire []Point) int {
	distance := 0
	for i := 0; i < len(wire)-1; i++ {
		distance += getLineLength(wire[i], wire[i+1])
	}
	return distance
}

func getLineLength(a Point, b Point) int {
	baseDist := (a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y)
	return int(math.Round(math.Sqrt(float64(baseDist))))
}

/**
 * This entire function exists because I am a morally bankrupt programmer.
 * Basically it takes a line and find out where a point on it "fits", and
 * returns the length.
 */
func kludgeFindPointOnLine(a Point, wire []Point) (int, []Point) {
	for i := 0; i < len(wire)-1; i++ {
		wireSection := []Point{wire[i], wire[i+1]}
		hasIntersect := pointInSegment(wireSection, a)
		if hasIntersect {
			increment := 1
			if isEq(a, wire[i+1]) {
				increment = 2
			}
			wireToIntersection := make([]Point, i+1)
			copy(wireToIntersection, wire[0:i+increment])
			wireToIntersection = append(wireToIntersection, a)
			localDistance := calculateWireLength(wireToIntersection)
			return localDistance, wireToIntersection
		}
	}
	return 0, []Point{}
}

func getShortestIntersections(wireA []Point, wireB []Point) int {
	distance := 100000000
	// This is where we would do a min heap or some shit like that.
	for i := 0; i < len(wireA)-1; i++ {
		// this is where we'd check for an intersection, and then
		// find the shortest path.
		line := []Point{wireA[i], wireA[i+1]}
		intersections, hasIntersect := getIntersectionList(line, wireB)
		if hasIntersect {
			for _, intersection := range intersections {
				if intersection.y == 0 && intersection.x == 0 {
					continue
				}
				localDistance, _ := kludgeFindPointOnLine(intersection, wireA)
				otherDistance, _ := kludgeFindPointOnLine(intersection, wireB)
				combinedDist := localDistance + otherDistance

				if combinedDist < distance && combinedDist != 0 {
					distance = combinedDist
				}
			}
		}
	}
	return distance
}

func getIntersectionList(line []Point, wire []Point) ([]Point, bool) {
	if len(wire) < 2 {
		fmt.Println("Wire is too short!")
		return []Point{Point{0, 0}}, false
	}
	intersectionsList := make([]Point, 1)

	for i := 0; i < len(wire)-1; i++ {
		wireSection := []Point{wire[i], wire[i+1]}
		intersection, hasIntersect := getIntersection(line, wireSection)
		if hasIntersect {
			intersectionsList = append(intersectionsList, intersection)
		}
	}

	if len(intersectionsList) > 0 {
		return intersectionsList, true
	}

	return []Point{Point{0, 0}}, false
}

func getLineParts(A Point, B Point) (int, int, int) {
	a1 := B.y - A.y
	b1 := A.x - B.x
	c1 := a1*A.x + b1*A.y
	return a1, b1, c1
}
func isCenter(point Point) bool {
	return point.x == 0 && point.y == 0
}

func isEq(a Point, b Point) bool {
	return a.x == b.x && a.y == b.y
}

// it would be esasier at this piont to make a two item grid
// and simpley
func getIntersection(lineA []Point, lineB []Point) (Point, bool) {
	//a1x + b1y = c1
	//a2x + b2y = c2
	A := lineA[0]
	B := lineA[1]
	P := lineB[0]
	Q := lineB[1]
	if (isCenter(A) || isCenter(B)) || (isCenter(P) || isCenter(Q)) {
		return Point{0, 0}, false
	}
	if isEq(A, P) || isEq(A, Q) || isEq(B, P) || isEq(B, Q) {
		return Point{0, 0}, false
	}
	a1, b1, c1 := getLineParts(A, B)
	a2, b2, c2 := getLineParts(P, Q)
	// (a1b2 – a2b1) x = c1b2 – c2b1
	determinant := a1*b2 - a2*b1
	if determinant == 0 {
		return Point{-1, -1}, false
	}
	x := (b2*c1 - b1*c2) / determinant
	y := (a1*c2 - a2*c1) / determinant
	foundPoint := Point{x, y}
	isTrueIntersect := pointInSegment(lineA, foundPoint) && pointInSegment(lineB, foundPoint)

	return foundPoint, isTrueIntersect
}

func pointInSegment(segment []Point, point Point) bool {
	x := point.x
	y := point.y
	P := segment[0]
	Q := segment[1]
	return (x <= P.x || x <= Q.x) &&
		(x >= Q.x || x >= P.x) &&
		(y <= P.y || y <= Q.y) &&
		(y >= P.y || y >= Q.y)
}

func day31() {
	fmt.Println("Day 3, Part 1")
	// read in the values
	data := loadFile("./day3-input.txt")
	distance := getManhattanDistance(data)
	fmt.Printf("The manhattan distance is %d\n", distance)
	fmt.Println("The correct answer was: '870'")
}

func day32() {
	fmt.Println("Day 3, Part 2")
	data := loadFile("./day3-input.txt")
	distance := getWireDistance(data)
	fmt.Printf("The wire step distance is %d\n", distance)
	fmt.Println("The correct answer was: '13698'")
}

func day3() {
	fmt.Println("Day 3")
	fmt.Println("---------")
	day31()
	day32()
	fmt.Println()
}
