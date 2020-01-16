package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDay8(t *testing.T) {
	assert.False(t, false)
}

func TestParseImage(t *testing.T) {
	layers := parseImage("123456789012", 3, 2)
	assert.Equal(t, [][]int{
			[]int{1,2,3,4,5,6},
			[]int{7,8,9,0,1,2},
	}, layers)
}

func TestRenderImage(t *testing.T) {
	message := parseImage("0222112222120000", 2, 2)
	message2 := parseImage("022222110222221121000000", 3, 2)
	// 022
	// 222
	
	// 110
	// 222
	
	// 221
	// 121

	// 000
	// 000"
	message = renderImage(message, 2, 2)
	message2 = renderImage(message2, 3, 2)
	assert.Equal(t, [][]int{
		[]int{0, 1},
		[]int{1, 0},
	}, message)

	assert.Equal(t, [][]int{
		[]int{0, 1, 0},
		[]int{1, 0, 1},
	}, message2)
	printedImage := printImage(message, 2, 2)
	printedImage2 := printImage(message2, 3, 2)
	expectedImage := "01\n10"
	expectedImage2 := "010\n101"
	assert.Equal(t, expectedImage, printedImage)
	assert.Equal(t, expectedImage2, printedImage2)
}