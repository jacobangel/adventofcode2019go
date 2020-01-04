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

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func ConvertToCoords(inputData string) [][]Point {
	fmt.Printf("The input is \n'%s' \n", inputData)

	wireList := strings.Split(inputData, "\n")
	wires := make([][]Point, len(wireList))
	for wireIndex, line := range wireList {
		coordList := strings.Split(strings.TrimSpace(line), ",")
		lastCoord := Point{0, 0}
		wires[wireIndex] = append(wires[wireIndex], lastCoord)
		for _, coord := range coordList {
			dir := string(coord[0])
			count, _ := strconv.Atoi(coord[1:])
			nextCoord := Point{lastCoord.x, lastCoord.y}
			switch dir {
			case "U":
				nextCoord.y += count
			case "D":
				nextCoord.y -= count
			case "L":
				nextCoord.x -= count
			case "R":
				nextCoord.x += count
			}
			wires[wireIndex] = append(wires[wireIndex], nextCoord)
			//			fmt.Printf("%s: %s %d\n", coord, dir, count)
			lastCoord = nextCoord
		}
	}
	return wires
}

func getManhattanDistance(inputData string) int {
	wires := ConvertToCoords(inputData)
	wireA := wires[0]
	wireB := wires[1]
	fmt.Printf("Wire A: %v\nWire B: %v\n", wireA, wireB)
	distance := 100000000
	// This is where we would do a min heap or some shit like that.
	for i := 0; i < len(wireA)-1; i++ {
		// this is where we'd check for an intersection, and then
		// find the shortest path.
		line := []Point{wireA[i], wireA[i+1]}
		intersections, hasIntersect := getIntersects(line, wireB)
		if hasIntersect {
			for _, intersection := range intersections {
				localDistance := determineManhattanDistance(intersection)
				if localDistance < distance && localDistance != 0 {
					fmt.Printf("Found new winner: %v, LD: %d\n", intersection, localDistance)
					distance = localDistance
				}
			}
		}

	}

	return distance
}

func getWireDistance(inputData string) int {
	wires := ConvertToCoords(inputData)
	wireA := wires[0]
	wireB := wires[1]
	fmt.Printf("Wire A: %v\nWire B: %v\n", wireA, wireB)
	minDistA := getWireDists(wireA, wireB)
	minDistB := getWireDists(wireB, wireA)
	if minDistA > minDistB {
		return minDistB
	}
	return minDistA
}

func determineManhattanDistance(point Point) int {
	return Abs(point.x) + Abs(point.y)
}

func determineWireDistance(wire []Point) int {
	distance := 0
	for i := 0; i < len(wire)-1; i++ {
		distance += getLineLength(wire[i], wire[i+1])
	}
	return distance
}

func getLineLength(a Point, b Point) int {
	return integer(math.Round(math.Sqrt(float64((a.x-b.x)*(a.x-b.x) + (a.y-b.y)*(a.y-b.y))))
}

func getWireDists(wireA []Point, wireB []Point) int {
	distance := 100000000
	// This is where we would do a min heap or some shit like that.
	for i := 0; i < len(wireA)-1; i++ {
		// this is where we'd check for an intersection, and then
		// find the shortest path.
		line := []Point{wireA[i], wireA[i+1]}
		intersections, hasIntersect := getIntersects(line, wireB)
		if hasIntersect {
			for _, intersection := range intersections {
				if intersection.y == 0 && intersection.x == 0 {
					continue
				}
				wireToIntersection := append(wireA[0:i], intersection)

				localDistance := determineWireDistance(wireToIntersection)
				fmt.Printf("localDistance : %d, wire: %v\n", localDistance, wireToIntersection)
				if localDistance < distance && localDistance != 0 {
					fmt.Printf("Found new winner: %v, LD: %d\n", intersection, localDistance)
					distance = localDistance
				}
			}
		}
	}
	return distance
}

func getIntersects(line []Point, wire []Point) ([]Point, bool) {
	if len(wire) < 2 {
		fmt.Println("Wire is too short!")
		return []Point{Point{0, 0}}, false
	}
	// you shoudl just frickin Coord type dude.
	intersectionsList := make([]Point, 1)

	for i := 0; i < len(wire)-1; i++ {
		wireSection := []Point{wire[i], wire[i+1]}
		intersection, hasIntersect := getIntersection(line, wireSection)
		if hasIntersect {
			fmt.Printf("-> hasIntersect: %v, %v: %v\n", line, wireSection, intersection)
			intersectionsList = append(intersectionsList, intersection)
		}
	}

	if len(intersectionsList) > 0 {
		return intersectionsList, true
	}

	return []Point{Point{0, 0}}, false
}

const x = 0
const y = 1

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
	// if isEq(A, P) || isEq(A, Q) || isEq(B, P) || isEq(B, Q) {
	// 	return Point{0, 0}, false
	// }
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
}

func day32() {
	fmt.Println("Day 3, Part 2")
	data := loadFile("./day3-input.txt")
	distance := getWireDistance(data)
	fmt.Printf("The manhattan distance is %d\n", distance)
}

func day3() {
	day31()
	day32()
}
