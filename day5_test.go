package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDay5(t *testing.T) {
	assert.False(t, false)
}


func testTakeDigits(t *testing.T) {
	get1, remainder1 := takeDigits(1000, 2);
	assert.Equal(t, get1, 0, "2 digits from 1000 should be 00")
	assert.Equal(t, remainder1, 10)

	get1, remainder1 = takeDigits(1011, 2);
	assert.Equal(t, get1, 11, "2 digits from 1000 should be 00")
	assert.Equal(t, remainder1, 10)

	get1, remainder1 = takeDigits(-1011, 2);
	assert.Equal(t, get1, -1011, "2 digits from 1000 should be 00")
	assert.Equal(t, remainder1, -1011)

	get1, remainder1 = takeDigits(1, 2);
	assert.Equal(t, get1, 1, "2 digits from 1000 should be 00")
	assert.Equal(t, remainder1, 0)
}


func testParseInstruction(t *testing.T) {
	instruction, paramOperations := parseInstruction(1002);
	assert.Equal(t, instruction, 2)
	assert.Equal(t, paramOperations, []ParameterMode{POSITION, IMMEDIATE, POSITION})

	instruction, paramOperations = parseInstruction(1003);
	assert.Equal(t, instruction, 3)
	assert.Equal(t, paramOperations, []ParameterMode{POSITION})

	instruction, paramOperations = parseInstruction(11001);
	assert.Equal(t, instruction, 1)
	assert.Equal(t, paramOperations, []ParameterMode{IMMEDIATE, IMMEDIATE, POSITION})
}