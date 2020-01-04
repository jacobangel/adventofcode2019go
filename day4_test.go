package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	assert.True(t, isValidPassword(111111))
	assert.False(t, isValidPassword(123456))
	assert.False(t, isValidPassword(234560))
}

func TestIsAscending(t *testing.T) {
	assert.True(t, isAscending(22345), "22345 should pass.")
	assert.False(t, isAscending(223450))
}

func TestHasDouble(t *testing.T) {
	assert.True(t, hasDouble(11))
	assert.False(t, hasDouble(123789))
	assert.False(t, hasDouble(12))
	assert.True(t, hasDouble(12378911))
	assert.True(t, hasDouble(1112378911))
}

func TestHasGroupedDouble(t *testing.T) {
	assert.True(t, hasEvenDouble(11))
	assert.False(t, hasEvenDouble(123789))
	assert.False(t, hasEvenDouble(12))
	assert.True(t, hasEvenDouble(112378911))
	assert.True(t, hasEvenDouble(1112378911))
	assert.False(t, hasEvenDouble(11123789))
	assert.True(t, hasEvenDouble(11112378911))
}

func TestIsValidGroupedPassword(t *testing.T) {
	assert.False(t, isValidGroupedPassword(111111))
	assert.True(t, isValidGroupedPassword(112233))
	assert.False(t, isValidGroupedPassword(123444))
	assert.True(t, isValidGroupedPassword(111122))
}
