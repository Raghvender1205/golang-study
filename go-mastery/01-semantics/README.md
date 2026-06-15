# Phase 1: Memory, Types and Semantics
How does go handle memory. Go is pass-by-value always

- Data structures under the hood: Understand that a slice is just a small struct (a pointer
to an array, length and capacity). If you pass a slice to a function, you are passing that struct by value
- Escape Analysis: How go compiler decides what lives on the stack versus the heap. It's critical for zero-allocation programming
- Interfaces: Grasp how implicit interface satisfaction works and the concept of "fat pointers" (`itab` and data pointer)

```
Slice Header Internal Representation
+-----------------+
| Data *uintptr   | ---> [ Underlying Array of Elements ]
+-----------------+
| Len  int        | ---> Current number of elements
+-----------------+
| Cap  int        | ---> Maximum capacity before reallocation
+-----------------+
```

## Slices: The "Fat Pointer" illusion
In Go, arrays are rigid; their size is part of their type (eg. `[3]int` is completely different type than `[4]int`). Slices (`[]int`) provide the flexibility but they are not arrays themselves. A slice is simply a lightweight descriptor -- a `struct` holding three values - that acts as a window into an underlying backing array.

[Here is the code](memory-layout/slices/main.go)

When we run this code, we notice that memory address (`ptr`) remains the same for the first append, but changes entirely on the second append. When a slice exceeds its capacity, Go allocates brand-new backing array, copies the old data and points the slice header to the new location

Before and after the allocation
```
State 1: make([]int, 0, 3)
+-----------------+       
| Data *uintptr   | --------> [ 0, 0, 0 ] (Pre-allocated backing array)
+-----------------+       
| Len: 0          |       
+-----------------+       
| Cap: 3          |       
+-----------------+       

State 2: After exceeding capacity (append 4 total items)
+-----------------+       
| Data *uintptr   | --------> [ 10, 20, 30, 40, 0, 0 ] (New Array, capacity doubled)
+-----------------+       
| Len: 4          |       
+-----------------+       
| Cap: 6          |       
+-----------------+       
                          [ 10, 20, 30 ] (Old Array - Marked for Garbage Collection)
```

## Escape Analysis: Stack vs Heap
Go doesn't make us manage memory manually like (`malloc` / `free`), but it does let us optimize <b>where</b> memory lives

- Stack: Fast, self cleaning, strictly scoped to a function
- Heap: Slower to allocate, requires Garbage Collector (GC) to clean up, which consumes CPU cycles.

Go compiler performs `Escape Analysis` to determine if a variable can safely live on the stack, or if it "escapes" the function's scope and must be allocated to heap.

[Here is the code](memory-layout/escape/main.go)

To see the compiler's decisions in real-time, run the build command with `-m` gcflag

```
go build -gcflags="-m" main.go
```

You will something like this
```sh
# command-line-arguments
./main.go:4:6: can inline createOnStack
./main.go:12:6: can inline escapesToHeap
./main.go:17:6: can inline main
./main.go:18:20: inlining call to createOnStack
./main.go:19:20: inlining call to escapesToHeap
./main.go:13:2: moved to heap: y
```