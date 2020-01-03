package main

import (
	"fmt"
	"testing"
)

func Day3TestingStub(t *testing.T) {
	fmt.Println("We are printing something.")
	dat := "Hi"
	if dat != "" {
		t.Errorf("There was an error:\n '%s'", dat)
	}
}

const data1 = "R75,D30,R83,U83,L12,D49,R71,U7,L72\nU62,R66,U55,R34,D71,R55,D58,R83"
const data2 = "R98,U47,R26,D63,R33,U87,L62,D20,R33,U53,R51\nU98,R91,D20,R16,D67,R40,U7,R15,U6,R7"

func TestDistance(t *testing.T) {
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
	point1 := [][2]int{ 
		[2]int{5, 2},
		[2]int{5, -2},
	} 

	point2 := [][2]int{ 
		[2]int{2, 0},
		[2]int{12, 0},
	} 

	inter1, hasInter1 := intersect(point1, point2)
	if !hasInter1 {
		t.Errorf("Should have detected an intersection: %v", inter1)
	}
	if inter1[0] != 5 || inter1[1] != 0 {
		t.Errorf("We got the wrong coordinate: %v", inter1)
	} else {
		fmt.Printf("%v\n", inter1)
	}

	inter2, hasInter2 := intersect(point2, point1)
	if !hasInter2 {
		t.Errorf("Should have detected an intersection: %v", inter2)
	}
	if inter2[0] != 5 || inter2[1] != 0 {
		t.Errorf("We got the wrong coordinate: %v", inter2)
	} else {
		fmt.Printf("%v\n", inter2)
	}

}

func TestIntersect2(t *testing.T) {
	point1 := [][2]int{ 
		[2]int{2, 2},
		[2]int{2, -2},
	} 

	point2 := [][2]int{ 
		[2]int{0, 0},
		[2]int{5, 0},
	} 

	inter2, hasInter2 := intersect(point2, point1)
	if !hasInter2 {
		t.Errorf("Should have detected an intersection: %v", inter2)
	}
	if inter2[x] != 2 || inter2[y] != 0 {
		t.Errorf("We got the wrong coordinate: %v, expected [2, 0]", inter2)
	} else {
		fmt.Printf("%v\n", inter2)
	}

	inter1, hasInter1 := intersect(point1, point2)
	if !hasInter1 {
		t.Errorf("Should have detected an intersection: %v", inter1)
	}
	if inter1[0] != 2 || inter1[1] != 0 {
		t.Errorf("We got the wrong coordinate: %v", inter1)
	} else {
		fmt.Printf("%v\n", inter1)
	}


}

func TestNonIntersect(t *testing.T) {
	point1 := [][2]int{ 
		[2]int{10, 2},
		[2]int{10, -2},
	} 

	point2 := [][2]int{ 
		[2]int{2, 2},
		[2]int{2, -5},
	} 

	inter3, hasInter2 := intersect(point2, point1)
	if hasInter2 {
		t.Errorf("Should NOT have detected an intersection: %v", inter3)
	}
}