package main

import "fmt"

func main() {
	// Initialize slice with Length 0, Capacity 3
	s := make([]int, 0, 3)
	fmt.Printf("Initial -> len: %d, cap: %d, ptr: %p\n", len(s), cap(s), s)

	// Append within capacity (no new allocation)
	s = append(s, 10, 20)
	fmt.Printf("Append 2 -> len: %d, cap: %d, ptr: %p\n", len(s), cap(s), s)

	// Append exceeding capacity (triggers reallocation)
	s = append(s, 30, 40)
	fmt.Printf("Append 2 more -> len: %d, cap: %d, ptr: %p\n", len(s), cap(s), s)
}