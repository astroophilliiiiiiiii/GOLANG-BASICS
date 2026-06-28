# 🚀 Go Basics

## BASIC CONCEPTS OF GO

### 1.) Syntactically Simpler
- Simple + clean syntax.
- Lesser keywords and lesser complexity than C++.

----------------------------------------------------------------------------

### 2.) Compiled Language
- Firstly it compiles then executes, that's why execution is faster.
- Faster than interpreted languages like Python.

----------------------------------------------------------------------------

### 3.) Crazy Concurrency Support (Goroutines)
- It has Goroutines (light-weight threads).
- Ek time pr boht saare tasks.
- Can efficiently handle thousands of requests simultaneously.

----------------------------------------------------------------------------

### 4.) Statically Typed
- Type of variable is fixed at the time of declaration and can't be changed.

```go
var x = 10   // Always int
```

----------------------------------------------------------------------------

### 5.) Performance

#### Execution
- Slightly slower than C++ (because of Garbage Collector).
- Faster than Java and Python.

#### Compilation
- Faster than C++.

##### Garbage Collector (GC)
- Go has a Garbage Collector (GC). GC automatically frees unused memory.
- C++ has **no** Garbage Collector. We manage memory ourselves.
- Because of GC, Go has a small runtime overhead, so C++ is usually slightly faster in execution.

----------------------------------------------------------------------------

### 6.) Go vs Java

Go is usually faster than Java, especially for:

- 🌐 Web servers
- 💻 CPU-intensive tasks
- ⚡ Applications that need very fast response time (low latency)

#### Why?
- ✅ Go is a compiled language.
- ✅ It has lightweight Goroutines for concurrency.
- ✅ It has efficient memory management.

------------------------------------------------------------------

# THREADS

Suppose there are 3 tasks:

- File Download
- Music Play
- Message Send

### Single Thread
- All 3 tasks execute one by one.

### Multiple Threads
- All tasks run simultaneously.

------------------------------------------------------------------

# Goroutine

**Goroutine = Go's lightweight thread.**

Just put the `go` keyword before a function call.

go printNumbers()
go downloadFile()

Both can run simultaneously.

--------------------------------------------------------------------

### Why Goroutines?

#### Normal Thread (C++/Java)
- More memory.
- More creation cost.

#### Goroutine
- Very less memory.
- Very less creation time.

That's why Go can create a very large number of Goroutines efficiently.

----------------------------------------------------------------------------

### Real-Life Example

Suppose you are building a chat application.

At the same time:
- User A sends a message.
- User B uploads an image.
- User C starts a video call.

Go uses Goroutines to handle each task concurrently, so the server can serve many users efficiently.

----------------------------------------------------------------------------

# Companies Using Go

- Docker → Built using Go.
- Kubernetes → Built using Go.

----------------------------------------------------------------------------
