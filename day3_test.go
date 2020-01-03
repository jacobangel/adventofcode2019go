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
