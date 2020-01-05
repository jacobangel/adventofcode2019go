package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExample(t *testing.T) {
	assert.True(t, true, "True is true!")
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
