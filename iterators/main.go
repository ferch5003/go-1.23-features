package main

import (
	"fmt"
	"slices"
)

func main() {
	// CUSTOM
	fmt.Println("CUSTOM ITERATORS")
	fmt.Println()

	fmt.Println("SSIMPLE SPACE")
	sSimple := []int{1, 2, 3, 4, 5, 6, 7}
	for range BackwardSimple(sSimple) {
	}

	fmt.Println()

	fmt.Println("S SPACE")
	s := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range Backward(s) {
		fmt.Println("key: ", i, "value:", v)
	}

	fmt.Println()

	fmt.Println("S2 SPACE")
	s2 := []string{"a", "b", "c", "d", "f"}
	for v := range BackwardSeq(s2) {
		fmt.Println("value:", v)
	}

	fmt.Println()

	fmt.Println("S3 SPACE")
	s3 := []float64{1.1, 2.2, 3.3, 4.4, 5.5, 6.6, 7.7}
	for i, v := range BackwardSeq2(s3) {
		fmt.Println("key: ", i, "value:", v)
	}

	fmt.Println()

	// BUILT-IN
	fmt.Println("BUILT-IN ITERATORS")
	fmt.Println()

	fmt.Println("BS SPACE")
	bs := []int{1, 2, 3, 4, 5, 6, 7}
	for i, v := range slices.Backward(bs) {
		fmt.Println("key: ", i, "value:", v)
	}

	fmt.Println()
}
