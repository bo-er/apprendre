## Overhead of calling len() in Go

There are two cases:

- **Local slice:** length will be cached and there is no overhead
- **Global slice** or passed (by reference): length cannot be cached and there is overhead

### No overhead for local slices

For locally defined slices the length is cached, so there is no runtime overhead. You can see this in the assembly of the following program:

```golang
func generateSlice(x int) []int {
    return make([]int, x)
}

func main() {
    x := generateSlice(10)
    println(len(x))
}
```

Compiled with `go tool compile -S test.go` this yields, amongst other things, the following lines:

```golang
MOVQ    "".x+40(SP),BX
MOVQ    BX,(SP)
// ...
CALL    ,runtime.printint(SB)
```

What happens here is that the first line retrieves the length of `x` by getting the value located 40 bytes from the beginning of `x` and most importantly caches this value in `BX`, which is then used for every occurrence of `len(x)`. The reason for the offset is that an array has the following structure ([source](https://code.google.com/p/go/source/browse/src/cmd/gc/go.h?name=go1.3.3#831)):

```golang
typedef struct
{               // must not move anything
    uchar   array[8];   // pointer to data
    uchar   nel[4];     // number of elements
    uchar   cap[4];     // allocated number of elements
} Array;
```

`nel` is what is accessed by `len()`. You can see this in the [code generation](https://code.google.com/p/go/source/browse/src/cmd/6g/gsubr.c?name=go1.3.3#1302) as well.

### Global and referenced slices have overhead

For shared values caching of the length is not possible since the compiler has to assume that the slice changes between calls. Therefore the compiler has to write code that accesses the length attribute directly every time. Example:

```golang
func accessLocal() int {
    a := make([]int, 1000) // local
    count := 0
    for i := 0; i < len(a); i++ {
        count += len(a)
    }
    return count
}

var ag = make([]int, 1000) // pseudo-code

func accessGlobal() int {
    count := 0
    for i := 0; i < len(ag); i++ {
        count += len(ag)
    }
    return count
}
```

Comparing the assembly of both functions yields the crucial difference that as soon as the variable is global the access to the `nel` attribute is not cached anymore and there will be a runtime overhead:

```golang
// accessLocal
MOVQ    "".a+8048(SP),SI // cache length in SI
// ...
CMPQ    SI,AX            // i < len(a)
// ...
MOVQ    SI,BX
ADDQ    CX,BX
MOVQ    BX,CX            // count += len(a)

// accessGlobal
MOVQ    "".ag+8(SB),BX
CMPQ    BX,AX            // i < len(ag)
// ...
MOVQ    "".ag+8(SB),BX
ADDQ    CX,BX
MOVQ    BX,CX            // count += len(ag)
```