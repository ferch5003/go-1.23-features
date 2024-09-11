package main

import "fmt"

func SharingLoopVariablesBeforeFirstSolution() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		// Copying value of v and pass it to go function.
		go func(v string) {
			fmt.Println(v)
			done <- true
		}(v)
	}

	for _ = range values {
		<-done
	}
}

func SharingLoopVariablesBeforeSecondSolution() {
	done := make(chan bool)

	values := []string{"a", "b", "c"}
	for _, v := range values {
		// Creates new variable v and use the new variable on the go function.
		v := v
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}

func SharingLoopVariables() {
	done := make(chan bool)

	// Now you can pass the value to the go function immediately without creating new variables
	// or copy the value of v.
	values := []string{"a", "b", "c"}
	for _, v := range values {
		go func() {
			fmt.Println(v)
			done <- true
		}()
	}

	for _ = range values {
		<-done
	}
}

func RangeOverIntegers(n int) {
	for i := range n {
		fmt.Println("Iteration", i)
	}
}
