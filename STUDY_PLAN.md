# Go study plan

This is a self-paced path starting after basic Go syntax and moving toward the advanced subjects in this repository. The learner writes every program; the teacher explains concepts, sets exercises, reviews attempts, and gives hints.

Follow the matching lesson directories in [go-mastery](go-mastery/README.md). Write each exercise in its lesson directory and keep your notes there too.

## How each lesson works

1. **Concept:** Explain one small idea in plain language, with a diagram when useful.
2. **Prediction:** Before running code, predict what a short example will do and explain why.
3. **Practice:** Write the solution yourself. Start with a small exercise, then change one requirement.
4. **Review:** Share your code, output, and questions. We review correctness, readability, and the Go concepts involved.
5. **Checkpoint:** Explain the idea back in your own words and complete a similar task without a walkthrough.

I will not write the exercise code for you. If you get stuck, I will give progressively stronger hints. I will show a complete solution only if you explicitly ask for one. We move on when you can apply the topic, not after a fixed number of days.

## Part 1: Data and program design

### 1. Arrays, slices, and maps

- Indexing, `range`, `len`, `append`, slice length and capacity, map lookup, and missing keys.
- Understand normal usage before studying backing arrays and runtime details.
- **Practice:** Count word frequencies and manage a list of tasks in memory.
- **Checkpoint:** Choose between a slice and a map for a new problem and explain why.

### 2. Structs, methods, and pointers

- Group data with structs; add behavior with methods; understand addresses and pointer receivers.
- Compare copying a value with changing the original value.
- **Practice:** Model a small library or inventory and add operations to update items.
- **Checkpoint:** Explain when a method should use a value or pointer receiver.

### 3. Errors and program boundaries

- Return errors, check them, add context, and distinguish expected failures from programmer mistakes.
- Introduce `defer` for cleanup. Learn `panic` only after ordinary error handling is clear.
- **Practice:** Validate input and handle invalid choices without crashing.
- **Checkpoint:** Make each possible failure understandable to a caller or user.

### 4. Files, text, and JSON

- Read and write files, handle paths and permissions, close resources, and work with text and JSON.
- Introduce readers and writers through practical use.
- **Practice:** Save and load your task list; report useful errors for a missing or malformed file.
- **Checkpoint:** Explain what could fail at each file operation and how your program responds.

### 5. Packages, modules, and tests

- Organize code into packages, use `go.mod`, export names deliberately, and write table-driven tests.
- Learn the standard tools: `go test`, `go fmt`, and `go vet`.
- **Practice:** Split one earlier program into packages and test its core logic.
- **Checkpoint:** Run the project and tests from a clean checkout.

## Part 2: Building useful Go programs

### 6. Interfaces and composition

- Define small interfaces from the behavior a caller needs; understand implicit satisfaction.
- Use interfaces where they simplify testing or allow multiple implementations.
- **Practice:** Give one program two interchangeable ways to store or display data.
- **Checkpoint:** Explain the concrete type, the interface, and which code depends on each.

### 7. HTTP and context

- Make HTTP requests and build a small HTTP server; handle request and response bodies and errors.
- Use `context` for cancellation and deadlines.
- **Practice:** Expose your task list through a small HTTP API.
- **Checkpoint:** Explain how a request ends, including cancellation and cleanup.

### 8. Concurrency

- Goroutines, channels, `select`, `sync.WaitGroup`, mutexes, and race detection.
- Learn when concurrency helps and how to avoid leaks and data races.
- **Practice:** Process a bounded set of jobs concurrently, then compare it with a sequential version.
- **Checkpoint:** Explain who owns shared data and how every goroutine stops.

## Part 3: Deep Go, using the existing repository as reference

Study these after the earlier checkpoints. For each topic, first predict behavior, then run or modify an example, then explain the result.

1. Slice backing arrays, aliasing, capacity growth, and `append`.
2. Allocation and escape analysis; inspect compiler output without assuming that a pointer always means heap allocation.
3. Interface method sets, nil interface values, and runtime representation.
4. Benchmarks and profiling before optimization.
5. Generics and iterators when they solve a concrete problem.
6. Runtime scheduling, garbage collection, and CGO as optional specialist topics.

## Progress record

For each completed topic, keep a short note with: the concept in your own words, the exercise you wrote, one mistake you fixed, and one question to revisit. The plan is complete when you can build a small Go program from a blank directory, test it, handle failures, and explain the key design choices without copying an example.
