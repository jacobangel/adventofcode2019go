package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func getManhattanDistance(inputData string) int {
	wires := ConvertToCoords(inputData)
	fmt.Printf("%v\n", wires)
	wireA := wires[0]
	wireB := wires[1]
	distance := 100000000
	// This is where we would do a min heap or some shit like that.
	for i := 0; i < len(wireA)-1; i++ {
		// this is where we'd check for an intersection, and then
		// find the shortest path.
		line := [][2]int{wireA[i], wireA[i+1]}
		intersection, hasIntersect := getIntersects(line, wireB)
		if hasIntersect {
			localDistance := Abs(intersection[0]) + Abs(intersection[1])
			if localDistance < distance {
				distance = localDistance
			}
		}

	}

	return distance
}

func getIntersects(line [][2]int, wire [][2]int) ([2]int, bool) {
	if len(wire) < 2 {
		return [2]int{0, 0}, false
	}
	// you shoudl just frickin Coord type dude.
	intersectionsList := make([][2]int, 1)

	for i := 0; i < len(wire)-1; i++ {
		wireSection := [][2]int{wire[i], wire[i+1]}
		intersection, hasIntersect := intersect(line, wireSection)
		if hasIntersect {
			intersectionsList = append(intersectionsList, intersection)
		}
	}

	// fmt.Printf("%v\n", intersectionsList)

	return [2]int{0, 0}, false
}

const x = 0
const y = 1

func intersect(lineA [][2]int, lineB [][2]int) ([2]int, bool) {
	// we know they intersect if one's xs are > other xs
	// and
	if lineA[0][x] < lineB[0][x] && lineA[1][x] > lineB[1][x] &&
		lineA[0][y] > lineB[0][y] && lineA[1][y] < lineB[1][y] { // if this isn't a sign you should make a coord...
		return [2]int{lineA[1][x], lineB[1][y]}, true
	}
	if lineA[0][y] < lineB[0][y] && lineA[1][y] > lineB[1][y] &&
		lineA[0][x] > lineB[0][x] && lineA[1][x] < lineB[1][x] { // if this isn't a sign you should make a coord...
		return [2]int{lineB[1][x], lineA[1][y]}, true
	}

	return [2]int{}, false
}

func ConvertToCoords(inputData string) [][][2]int {
	fmt.Printf("The input is \n'%s' \n", inputData)

	wireList := strings.Split(inputData, "\n")
	wires := make([][][2]int, len(wireList))
	for wireIndex, line := range wireList {
		coordList := strings.Split(line, ",")
		lastCoord := [2]int{0, 0}
		for _, coord := range coordList {
			dir := string(coord[0])
			count, _ := strconv.Atoi(coord[1:])
			nextCoord := [2]int{lastCoord[0], lastCoord[1]}
			switch dir {
			case "U":
				nextCoord[1] -= count
			case "D":
				nextCoord[1] += count
			case "L":
				nextCoord[0] -= count
			case "R":
				nextCoord[0] += count
			}
			wires[wireIndex] = append(wires[wireIndex], nextCoord)
			//			fmt.Printf("%s: %s %d\n", coord, dir, count)
			lastCoord = nextCoord
		}
	}
	//fmt.Printf("Num coords %d\n%v", len(wires), wires)

	return wires
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
}

func day3() {
	day31()
	day32()
}
