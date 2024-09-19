package main

import (
	"testing"
	"unique"
)

// Strings
var (
	s1 = "This is the same string"
	s2 = "This is the same string"
)

// Slices
var (
	sl1 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	sl2 = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
)

// Structs
type Unique struct {
	Value string
}

var (
	st1 = Unique{s1}
	st2 = Unique{s2}
)

func BenchmarkString_UsingComparison(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = s1 == s2
	}
}

func BenchmarkString_UsingUnique(b *testing.B) {
	u1 := unique.Make(s1)
	u2 := unique.Make(s2)

	for i := 0; i < b.N; i++ {
		_ = u1 == u2
	}
}

func BenchmarkSliceInt_UsingComparison(b *testing.B) {
	for i := 0; i < b.N; i++ {
		if len(sl1) != len(sl2) {
		}
		for i := range sl1 {
			_ = sl1[i] == sl2[i]
		}
	}
}

func BenchmarkSliceInt_UsingUnique(b *testing.B) {
	var u1 []unique.Handle[int]
	for _, v := range sl1 {
		u1 = append(u1, unique.Make(v))
	}

	var u2 []unique.Handle[int]
	for _, v := range sl2 {
		u2 = append(u2, unique.Make(v))
	}

	for i := 0; i < b.N; i++ {
		if len(u1) != len(u2) {
		}
		for i := range u1 {
			_ = u1[i] == u2[i]
		}
	}
}

func BenchmarkStruct_UsingComparison(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = st1 == st2
	}
}

func BenchmarkStruct_UsingUnique(b *testing.B) {
	u1 := unique.Make(st1)
	u2 := unique.Make(st2)

	for i := 0; i < b.N; i++ {
		_ = u1 == u2
	}
}
