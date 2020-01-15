package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsDay7(t *testing.T) {
	assert.False(t, false)
}

func TestThrustAmplification(t *testing.T) {
	data1 := []int{3, 15, 3, 16, 1002, 16, 10, 16, 1, 16, 15, 15, 4, 15, 99, 0, 0}
	data2 := []int{3, 23, 3, 24, 1002, 24, 10, 24, 1002, 23, -1, 23, 101, 5, 23, 23, 1, 24, 23, 23, 4, 23, 99, 0, 0}
	data3 := []int{3, 31, 3, 32, 1002, 32, 10, 32, 1001, 31, -2, 31, 1007, 31, 0, 33, 1002, 33, 7, 33, 1, 33, 31, 31, 1, 32, 31, 31, 4, 31, 99, 0, 0, 0}

	expectedPhase1 := 43210
	expectedSequence1 := []int{4, 3, 2, 1, 0}

	expectedPhase2 := 54321
	expectedSequence2 := []int{0, 1, 2, 3, 4}

	expectedPhase3 := 65210
	expectedSequence3 := []int{1, 0, 4, 3, 2}

	phase1, sequence1 := FindMaxAmplificationSeq(data1)
	phase2, sequence2 := FindMaxAmplificationSeq(data2)
	phase3, sequence3 := FindMaxAmplificationSeq(data3)

	assert.Equal(t, expectedPhase1, phase1)
	assert.Equal(t, expectedPhase2, phase2)
	assert.Equal(t, expectedPhase3, phase3)

	assert.Equal(t, expectedSequence1, sequence1)
	assert.Equal(t, expectedSequence2, sequence2)
	assert.Equal(t, expectedSequence3, sequence3)
}

func TestThrustAmplificationLoop(t *testing.T) {
	data1 := []int{3, 26, 1001, 26, -4, 26, 3, 27, 1002, 27, 2, 27, 1, 27, 26, 27, 4, 27, 1001, 28, -1, 28, 1005, 28, 6, 99, 0, 0, 5}
	data2 := []int{3, 52, 1001, 52, -5, 52, 3, 53, 1, 52, 56, 54, 1007, 54, 5, 55, 1005, 55, 26, 1001, 54, -5, 54, 1105, 1, 12, 1, 53, 54, 53, 1008, 54, 0, 55, 1001, 55, 1, 55, 2, 53, 55, 53, 4, 53, 1001, 56, -1, 56, 1005, 56, 6, 99, 0, 0, 0, 0, 10}

	expectedPhase1 := 139629729
	expectedSequence1 := []int{9, 8, 7, 6, 5}

	expectedPhase2 := 18216
	expectedSequence2 := []int{9, 7, 8, 5, 6}

	phase1, sequence1 := FindMaxAmplificationSeqLoop(data1)
	phase2, sequence2 := FindMaxAmplificationSeqLoop(data2)

	assert.Equal(t, expectedPhase1, phase1)
	assert.Equal(t, expectedPhase2, phase2)

	assert.Equal(t, expectedSequence1, sequence1)
	assert.Equal(t, expectedSequence2, sequence2)
}
func TestGetPerms(t *testing.T) {
	a := getPermutations([]int{0, 1})
	expected := [][]int{
		[]int{0, 1},
		[]int{1, 0},
	}
	assert.Equal(t, expected, a)
}
