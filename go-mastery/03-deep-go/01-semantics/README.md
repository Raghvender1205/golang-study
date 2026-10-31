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

## Interfaces
Go interfaces are satisfied <b>implicitly</b> (structural typing) and under the hood, they are managed via a data structure known as `Fat Pointer`.

### Fat Pointer
Interface variable is not just a single pointer to a piece of memory. It is a <b>Fat Pointer</b> -- a 2-word (16 bytes on a 64-bit architecture) data structure. Go manages interfaces using two distinct internal C-style structs defined in the `runtime` package

1. `eface` (Empty Interface):- Used for `interface{}` or `any`. Since it has no methods, it only needs to know the type of data and pointer to the data itself.

2. `iface` (Non-Empty Interface):- Used for interfaces that defines methods. It contains a pointer to an interface table `itab` and a pointer to the concrete data.

The Internal Layout of `iface`

```
Go Interface Variable (iface)
  +-------------------+-------------------+
  |   tab (*itab)     |   data (unsafe)   |  <--- The "Fat Pointer" (2 words)
  +---------+---------+---------+---------+
            |                   |
            v                   v
     +------------+      +---------------+
     |  Interface |      | Concrete Value|
     |   Type     |      | (e.g., File)  |
     +------------+      +---------------+
     | Concrete   |
     |   Type     |
     +------------+
     | Method Ptrs| ---> [ Actual Func Code ] (Dynamic Dispatch)
     +------------+
```

- `itab`: This points to the Interface Table. The `itab` caches the exact layout of the interface type, the layout of the concrete type and most importantly a list of function pointers mapping the interface methods directly to the concrete type's implementation.
- `data`: A direct pointer to the allocated concrete value on the heap or stack.

[Here is the code](data-structures/interfaces/main.go)

This code uses Go's `unsafe` package to bypass type safety, look directly at the raw bits of an interface variable, expose the underlying fat pointer.

```sh
Interface Memory Location: 0x348cb9a40070
        -> itab Pointer (tab): 0x4d0b18
        -> Data Pointer (data): 0x348cb9a40080
Closing SQL Connection: prod-cluster-01
```

When we call `closer.Close()`, the Go runtime doesn't know what `closer` is at compile time. Instead, it executes an operation similar to this pseudo-assembly sequence.
- Fetch the `tab` pointer from the interface variable
- Resolve the function pointer address for `Close` out of the table offset.
- Call that function address, automatically passing the `data` pointer as the first argument (method receiver)

### 3. Value vs Pointer Receivers & The Method Table
A common compilation roadblock in Go is assigning a value to an interface when the interface methods are defined on a pointer reciever.

- If a method is defined on a pointer receiver (`func (s *SQLConnection) Close()`), that method is only added to the method table of `*SQLConnection`.
- If we attempt to assign a value type (`conn`) instead of pointer (`&conn`) to the interface, the compiler fails because the `itab` cannot be built, method
lacks a compatible pointer