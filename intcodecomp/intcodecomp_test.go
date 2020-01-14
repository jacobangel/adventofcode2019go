package intcodecomp

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntCodeComp(t *testing.T) {
	assert.True(t, true)
}

func TestParseInstruction(t *testing.T) {
	instruction, paramOperations := parseInstruction(1002)
	assert.Equal(t, instruction, MULT)
	assert.Equal(t, paramOperations, []ParameterMode{POSITION, IMMEDIATE, POSITION})

	instruction, paramOperations = parseInstruction(1003)
	assert.Equal(t, instruction, STORE)
	assert.Equal(t, paramOperations, []ParameterMode{POSITION})

	instruction, paramOperations = parseInstruction(11001)
	assert.Equal(t, instruction, ADD)
	assert.Equal(t, paramOperations, []ParameterMode{POSITION, IMMEDIATE, IMMEDIATE})

	instruction, paramOperations = parseInstruction(1101)
	assert.Equal(t, instruction, ADD)
	assert.Equal(t, paramOperations, []ParameterMode{IMMEDIATE, IMMEDIATE, POSITION})
}

func TestInterpretProgram(t *testing.T) {
	data := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	data2 := []int{3, 21, 1008, 21, 8, 20, 1005, 20, 22, 107, 8, 21, 20, 1006, 20, 31, 1106, 0, 36, 98, 0, 0, 1002, 21, 125, 20, 4, 20, 1105, 1, 46, 104, 999, 1105, 1, 46, 1101, 1000, 1, 20, 4, 20, 1105, 1, 46, 98, 99}
	input1 := 0
	input2 := 8
	input3 := 18
	output1 := InterpretProgram(data, []int{input1})
	output2 := InterpretProgram(data2, []int{input2})
	output3 := InterpretProgram(data, []int{input3})
	expectedOutput1 := 999
	expectedOutput2 := 1000
	expectedOutput3 := 1001
	assert.Equal(t, expectedOutput1, output1)
	assert.Equal(t, expectedOutput2, output2)
	assert.Equal(t, expectedOutput3, output3)
}
