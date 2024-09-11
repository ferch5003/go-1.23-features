package main

import (
	"fmt"
	"iter"
)

func BackwardSimple[T any](s []T) func(func() bool) {
	return func(yield func() bool) {
		for i := len(s) - 1; i >= 0; i-- {
			fmt.Println(s[i])
			if !yield() {
				return
			}
		}
	}
}

func Backward[T any](s []T) func(func(int, T) bool) {
	return func(yield func(int, T) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}

func BackwardSeq[T any](s []T) iter.Seq[T] {
	return func(yield func(T) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(s[i]) {
				return
			}
		}
	}
}

func BackwardSeq2[T any](s []T) iter.Seq2[int, T] {
	return func(yield func(int, T) bool) {
		for i := len(s) - 1; i >= 0; i-- {
			if !yield(i, s[i]) {
				return
			}
		}
	}
}
