package main

// Stays on stack: returns a copy of the value
func createOnStack() int {
	x := 42
	return x
}

// Escapes to heap: returns the pointer to the value
// The caller needs access to "y" after this function finishes.
// so it cannot be destroyed with the stack frame
func escapesToHeap() *int {
	y := 42
	return &y
}

func main() {
	a := createOnStack()
	b := escapesToHeap()

	// Print statements also cause variables to escape becoz fmt.Println
	// takes empty interfaces (any), which requires heap allocation for dynamic types
	_ = a
	_ = b
}