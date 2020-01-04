package main

import (
	"testing"
)

func TestDay2(t *testing.T) {
	if false {
		t.Fail()
	}
}

func TestFileLoading(t *testing.T) {
	dat := loadFile("")
	if dat != "" {
		t.Errorf("The file did not come back empty:\n '%s'", dat)
	}
	dat = loadFile("./day2-input.txt")
	if string(dat[len(dat)-1]) == "\x00" {
		t.Errorf("The file's return should not contain a null character")
	}
}

func TestArraySlicing(t *testing.T) {
	sliceOfData := ConvertToSlice("1,2,3,4")
	localSlice := []int{1, 2, 3, 4}
	if !testEq(sliceOfData, localSlice) {
		t.Errorf("The slices were not equal: %v, %v", sliceOfData, localSlice)
	}

	sliceWithNull := ConvertToSlice("1,2,3,4,6")
	if len(sliceWithNull) != 5 {
		t.Errorf("There was an error converting this thing over...")
	}
}

func testEq(a []int, b []int) bool {

	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestRunProgram1(t *testing.T) {
	// test if we can run a program with data;

	testProgram1 := []int{1, 0, 0, 0, 99}
	RunProgram(testProgram1)
	if testProgram1[0] != 2 {
		t.Errorf("We did not find a 2 in the right spot: '%d'", testProgram1)
	}
}

func TestRunProgram2(t *testing.T) {
	// test if we can run a program with data;
	testProgram2 := []int{2, 3, 0, 3, 99}
	RunProgram(testProgram2)
	if testProgram2[3] != 6 {
		t.Errorf("We didn't find 6 in the 3 index: '%d'", testProgram2)
	}
}

func TestRunProgram3(t *testing.T) {
	// test if we can run a program with data;
	testProgram3 := []int{2, 4, 4, 5, 99, 0}
	RunProgram(testProgram3)
	if testProgram3[5] != 9801 {
		t.Errorf("We did not find a 9801 in the right spot: '%d'", testProgram3)
	}
}

func TestRunProgram4(t *testing.T) {
	testProgram4 := []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	RunProgram(testProgram4)
	if testProgram4[0] != 30 && testProgram4[4] != 2 {
		t.Errorf("We did not find 30 and 2 in the right spot: '%d'", testProgram4)
	}
}
