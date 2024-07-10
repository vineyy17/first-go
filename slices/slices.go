package main

import (
	"fmt"
)

func main() {
	var s []int                // s is a slice of int
	fmt.Println("len", len(s)) // len is "nil safe"

	if s == nil { // you can compare only a slice to nil
		fmt.Println("nil slice")
	}

	s2 := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("s2 = %#v\n", s2)

	s3 := s2[1:4] // slicing operation, half-open range
	fmt.Printf("s3 = %#v\n", s3)

	// fmt.Println(s2[:100]) // panic

	s3 = append(s3, 100)
	fmt.Printf("s3 (append) = %#v\n", s3)
	fmt.Printf("s2 (append) = %#v\n", s2) // s2 is changed as well!
	fmt.Printf("s2: len=%d, cap=%d\n", len(s2), cap(s2))
	fmt.Printf("s3: len=%d, cap=%d\n", len(s3), cap(s3))

	var s4 []int
	// s4 := make([]int, 0, 1_000) // single allocation
	for i := 0; i < 1_000; i++ {
		s4 = appendInt(s4, i)
	}
	fmt.Println("s4", len(s4), cap(s4))
}

func appendInt(s []int, v int) []int {
	i := len(s)
	if len(s) < cap(s) { // enough space in underlying array
		s = s[:len(s)+1]
	} else { // need to re-allocate and copy
		fmt.Printf("reallocate: %d->%d\n", len(s), 2*len(s)+1)
		s2 := make([]int, 2*len(s)+1)
		copy(s2, s)
		s = s2[:len(s)+1]
	}

	s[i] = v
	return s
}
