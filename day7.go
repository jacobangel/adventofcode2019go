package main

import (
	"./intcodecomp"
	"fmt"
)

type Amp struct {
	retVal int
	opcode intcodecomp.Opcode
	data   []int
	ip     int
}

/**
This is Heaps Algorithm. (https://stackoverflow.com/questions/30226438/generate-all-permutations-in-go)

The algorithm generates (n-1)! permutations of the first n-1 elements,
adjoining the last element to each of these.
This will generate all of the permutations that end with the last element.
If n is odd, swap the first and last element and
if n is even, then swap the ith element (i is the counter starting from 0) and the last element and
repeat the above algorithm till i is less than n.
In each iteration, the algorithm will produce all the permutations that end with the current last element.
*/
func getPermutations(items []int) [][]int {
	var helper func([]int, int)
	output := [][]int{}
	helper = func(list []int, n int) {
		if n == 1 {
			final := make([]int, len(list))
			copy(final, list)
			output = append(output, final)
		} else {
			for i := 0; i < n; i++ {
				helper(list, n-1)
				t := 0
				if n%2 == 0 {
					t = i
				}
				holder := list[t]
				list[t] = list[n-1]
				list[n-1] = holder
			}
		}
	}
	helper(items, len(items))
	return output
}

func FindMaxAmplificationSeq(programData []int) (int, []int) {
	perms := getPermutations([]int{0, 1, 2, 3, 4})

	maxValue := 0
	var maxPerm []int

	for _, perm := range perms {
		newValue := 0
		localdata := programData

		for _, code := range perm {
			newValue, _, _, _ = intcodecomp.InterpretProgram(localdata, []int{code, newValue}, 0)
		}
		if maxValue < newValue {
			maxPerm = perm
			maxValue = newValue
		}
	}

	fmt.Printf("data: %d\nperms: %d\n", programData, perms)
	return maxValue, maxPerm
}

func FindMaxAmplificationSeqLoop(programData []int) (int, []int) {
	perms := getPermutations([]int{5, 6, 7, 8, 9})

	maxValue := 0
	var maxPerm []int

	for _, perm := range perms {
		newValue := 0
		var newMemory []int
		var opcode intcodecomp.Opcode
		ip := 0
		amplifiers := make([]Amp, len(perm))
		// loop through once and set each code.
		// loop through once to set the int coes
		for index, code := range perm {
			localdata := programData
			localdata = make([]int, len(programData))
			copy(localdata, programData)
			_, opcode, newMemory, ip = intcodecomp.InterpretProgram(localdata, []int{code}, 0)
			amplifiers[index] = Amp{newValue, opcode, newMemory, ip}
		}
		for amplifiers[4].opcode != intcodecomp.STOP {
			for index, _ := range perm {
				if amplifiers[index].opcode == intcodecomp.STOP {
					fmt.Printf("Amplifier #%d is done!!!!!!\n", index)
					continue
				}
				if amplifiers[index].opcode == intcodecomp.STOP {
					continue
				}
				prevReg := index - 1
				if prevReg < 0 {
					prevReg = len(perm) - 1
				}
				// fmt.Printf("amt #%d: starting at ip %d, with value %d on data %v\n", index, amplifiers[index].ip, amplifiers[prevReg].retVal, amplifiers[index].data)
				// fmt.Printf("Data values are equal: %t", testEq(amplifiers[index].data, amplifiers[prevReg].data))
				lastVal := amplifiers[index].retVal

				newValue, opcode, newMemory, ip = intcodecomp.InterpretProgram(amplifiers[index].data, []int{amplifiers[prevReg].retVal}, amplifiers[index].ip)
				if opcode == intcodecomp.STOP {
					newValue = lastVal
				}
				amplifiers[index] = Amp{newValue, opcode, newMemory, ip}
			}
		}
		// fmt.Printf("val: %d, Amp: %v, perm: %v \n", newValue, amplifiers[4], perm)
		if maxValue < newValue {
			maxPerm = perm
			maxValue = newValue
		}
	}

	return maxValue, maxPerm
	// return 0, []int{}
}

func day71() {
	fmt.Println("Day 7.1")

	programData := LoadProgramData("./day7_input.txt")
	fmt.Println(programData)
	returnVal, seq := FindMaxAmplificationSeq(programData)
	fmt.Printf("The answer is: '%d' and %d\n", returnVal, seq)
}

func day72() {
	fmt.Println("Day 7.2")
	programData := LoadProgramData("./day7_input.txt")
	fmt.Println(programData)
	returnVal, seq := FindMaxAmplificationSeqLoop(programData)
	fmt.Printf("The answer is: '%d' and %d\n", returnVal, seq)
}

func day7() {
	fmt.Println("Day 7")
	fmt.Println("---------")
	day71()
	day72()
	fmt.Println()
}
