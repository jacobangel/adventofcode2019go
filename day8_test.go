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