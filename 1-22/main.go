package main

import "fmt"

func main() {
	rangeOver()
}

func rangeOver() {
	// SharingLoopVariablesBeforeFirstSolution
	fmt.Println("SharingLoopVariablesBeforeFirstSolution")
	SharingLoopVariablesBeforeFirstSolution()
	fmt.Println()

	// SharingLoopVariablesBeforeSecondSolution
	fmt.Println("SharingLoopVariablesBeforeSecondSolution")
	SharingLoopVariablesBeforeSecondSolution()
	fmt.Println()

	// SharingLoopVariables
	fmt.Println("SharingLoopVariables")
	SharingLoopVariables()
	fmt.Println()

	// RangeOverIntegers
	fmt.Println("RangeOverIntegers")
	RangeOverIntegers(10)
	fmt.Println()
}
