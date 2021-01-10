All computing environments must deal with [memory management](https://en.wikipedia.org/wiki/Memory_management). This article discusses some memory management concepts used by the [Go programming language](https://golang.org/). This article is written for programmers familiar with basic memory management concepts but unfamiliar with Go memory management in particular.

# Stack, Heap, and Fixed Size Segments

For the purposes of this article, there are three ways to allocate memory: the stack, the heap, and fixed size segments.

## Stack

[The stack](https://en.wikipedia.org/wiki/Stack-based_memory_allocation) has a top that moves up and down. Space is allocated on the stack by moving the top up (i.e. pushing items on the stack) and space is deallocated by moving the top down (i.e. popping items off the stack). The top is an address that can be incremented and decremented with fast arithmetic operations.

Typically, a functions parameters and local variables are allocated on the stack.

Each [goroutine](https://golang.org/doc/effective_go.html#goroutines) has its own stack; thus, no [synchronization](<https://en.wikipedia.org/wiki/Synchronization_(computer_science)>) (e.g., locking) is necessary.

Goroutine stacks are allocated on the heap. If the stack needs to grow beyond the amount allocated for it, then heap operations (allocate new, copy old to new, free old) will occur.

## Heap

Unlike the stack, [the heap](https://en.wikipedia.org/wiki/Memory_management#DYNAMIC) does not have a single partition of allocated and free regions. Rather, there is a set of of free regions. A data structure must be used to implement this set of free regions. When an item is allocated, it is removed from the free regions. When an item is freed, it is added back to the set of free regions.

Unlike the stack, the heap is not owned by one goroutine, so manipulating the set of free regions in the heap requires synchronization (e.g., locking).

## Fixed Sized Segments

Memory can also be allocated in one of the fixed sized segments, such as the [data segment](https://en.wikipedia.org/wiki/Data_segment) and [code segment](https://en.wikipedia.org/wiki/Code_segment). Fixed sized segments are defined at compile time and do not change size at runtime. Read-write fixed size segments (e.g., the data segment) contain global variables while read-only segments (e.g., code segment and rodata segment) contain constant values and instructions.[1](https://dougrichardson.us/2016/01/23/go-memory-allocations.html#fn:2)

# What Goes Where?

[The Go Programming Language Specification](https://golang.org/ref/spec) does not define where items will be allocated. For example, a variable defined as `var x int` could be allocated on the stack or the heap and still follow the language spec. Likewise, the integer pointed to by _p_ in `p := new(int)` could be allocated on the stack or the heap.

However, certain requirements will exclude some choices of memory in certain conditions. For instance:

- The size of the data segment cannot change at run time, and therefore cannot be used for data structures that change size.
- The lifetime of items in the stack are ordered by their position on the stack. If the top of the stack is address _X_ then everything above _X_ will be deallocated while everything below _X_ will remain allocated. Memory allocated by a function can escape that function if referenced by an item outside the scope of the function and therefore cannot be allocated on the stack (because it’s still being referenced), and neither can it be allocated in the data segment (because the data segment cannot grow at runtime), thus it must be allocated on the heap – although inlining can remove some of these heap allocations.

## Escape Analysis

[Escape analysis](https://en.wikipedia.org/wiki/Escape_analysis) is used to determine whether an item can be allocated on the stack. It determines if an item created in a function (e.g., a local variable) can escape out of that function or to other goroutines. For example, in the following function, x escapes from the function that defines it:

```go
package escapeanalysis

func Foo() *int {
    var x int
    return &x
}
```

Items that escape must be allocated on the heap. Thus _x_ would be allocated on the heap.(https://dougrichardson.us/2016/01/23/go-memory-allocations.html#fn:1)

The exact escape analysis algorithm can change between Go versions. However, you can use `go tool compile -m` to print optimization decisions, which include the escape analysis. For example, on the previous program with Go version 1.5.2, you get the following output:

```
escape.go:3: can inline Foo
escape.go:4: moved to heap: x
escape.go:5: &x escapes to heap
```

# Garbage Collector

Go uses [garbage collection](<https://en.wikipedia.org/wiki/Garbage_collection_(computer_science)>) for memory management. The Go garbage collector occasionally has to [stop the world](https://en.wikipedia.org/wiki/Tracing_garbage_collection#Stop-the-world_vs._incremental_vs._concurrent) to complete the collection task. Since [Go version 1.5](https://golang.org/doc/go1.5#gc), the collector is designed so that the _stop the world_ task will take no more than 10 milliseconds out of every 50 milliseconds of execution time.

The garbage collector has to be aware of both heap and stack allocated items. This is easy to see if you consider a heap allocated item, _H_, referenced by a stack allocated item, _S_. Clearly, the garbage collector cannot free _H_ until _S_ is freed and so the garbage collector must be aware of lifetime of _S_, the stack allocated item.

# Performance

If your process is [CPU bound](https://en.wikipedia.org/wiki/CPU-bound), use [runtime/pprof](https://golang.org/pkg/runtime/pprof/) package and `go tool pprof` to profile your program. If you see symbols like growslice and newobject taking up a lot of time, optimizing memory allocations may improve performance.

Assuming you’ve determined optimizing memory use would improve performance of your program, then reduce the number of allocations – especially heap allocations.

1. Reuse memory you’ve already allocated.
2. Restructure your code so the compiler can make stack allocations instead of heap allocations. Use `go tool compile -m` to help you identify escaped variables that will be heap allocated and then rewrite your code so that they can be stack allocated.
3. Restructure your CPU bound code to pre-allocate memory in a few big chunks rather than continuously allocating small chunks.

# References

- [Go 1.4+ Garbage Collection (GC) Plan and Roadmap – 2014-8-6](https://docs.google.com/document/d/16Y4IsnNRCN43Mx0NZc5YXZLovrHvvLhK_h0KN8woTO4/edit)
- [Go Escape Analysis Flaws – 2015-2-10](https://docs.google.com/document/d/1CxgUBPlx9iJzkz9JWkb6tIpTe5q32QDmz8l0BouG0Cw/preview)
- [Golang Escape Analysis – 2015-11-11](https://web.archive.org/web/20170930011137/http://blog.rocana.com/golang-escape-analysis)
- [Profiling Go Programs – 2011-6-24](https://blog.golang.org/profiling-go-programs)
- [The Go Programming Language Specification](https://golang.org/ref/spec)

# Footnotes

1. Thank you to K. Richard Pixley (former developer of GNU ld, GNU as, GNU BFD) for pointing out that data can also be stored in the code segment. [↩](https://dougrichardson.us/2016/01/23/go-memory-allocations.html#fnref:2)
2. Other optimizations (like inlining) could allow the compiler to allocate seemingly escaped variables on the stack. [↩](https://dougrichardson.us/2016/01/23/go-memory-allocations.html#fnref:1)
