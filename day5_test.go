package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsDay5(t *testing.T) {
	assert.False(t, false)
}


func TestTakeDigits(t *testing.T) {
	get1, remainder1 := takeDigits(1000, 2);
	assert.Equal(t, get1, 0, "2 digits from 1000 should be 00")
	assert.Equal(t, remainder1, 10)

	get1, remainder1 = takeDigits(1011, 2);
	assert.Equal(t, get1, 11, "2 digits from 1011 should be 11")
	assert.Equal(t, remainder1, 10)

	get1, remainder1 = takeDigits(-1011, 2);
	assert.Equal(t, get1, -1011, "2 digits from -1011 should be -1011")
	assert.Equal(t, remainder1, 0)

	get1, remainder1 = takeDigits(1, 2);
	assert.Equal(t, get1, 1, "2 digits from 1 should be 1")
	assert.Equal(t, remainder1, 0)

	get1, remainder1 = takeDigits(1101, 2);
	assert.Equal(t, get1, 1)

	get1, remainder1 = takeDigits(99, 2);
	assert.Equal(t, get1, 99)

	get1, remainder1 = takeDigits(990001, 5);
	assert.Equal(t, 90001, get1)
}


func TestParseInstruction(t *testing.T) {
	instruction, paramOperations := parseInstruction(1002);
	assert.Equal(t, instruction, MULT)
	assert.Equal(t, paramOperations, []ParameterMode{POSITION, IMMEDIATE, POSITION})

	instruction, paramOperations = parseInstruction(1003);
	assert.Equal(t, instruction, STORE)
	assert.Equal(t, paramOperations, []ParameterMode{POSITION})

	instruction, paramOperations = parseInstruction(11001);
	assert.Equal(t, instruction, ADD)
	assert.Equal(t, paramOperations, []ParameterMode{POSITION, IMMEDIATE, IMMEDIATE })

	instruction, paramOperations = parseInstruction(1101);
	assert.Equal(t, instruction, ADD)
	assert.Equal(t, paramOperations, []ParameterMode{IMMEDIATE, IMMEDIATE, POSITION })
}