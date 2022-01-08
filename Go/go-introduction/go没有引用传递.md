# Go 没有引用传递

[My post on pointers](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back) provoked a lot of debate about maps and pass by reference semantics. This post is a response to those debates.

由于 Go 没有`引用变量`,所以 go 没有引用传递的语法。

# 什么是引用变量？

在 C++这样的语言中，你可以给已经存在的变量声明一个`alias`别名。后面声明的这个别名变量称为引用变量。

```
#include <stdio.h>

int main() {
        int a = 10;
        int &b = a;
        int &c = b;

        printf("%p %p %p\n", &a, &b, &c); // 0x7ffe114f0b14 0x7ffe114f0b14 0x7ffe114f0b14
        return 0;
}
```

You can see that `a`, `b`, and `c` all refer to the same memory location. A write to `a` will alter the contents of `b` and `c`. This is useful when you want to declare reference variables in different scopes–namely function calls.

# Go does not have reference variables

Unlike C++, each variable defined in a Go program occupies a unique memory location.

```
package main

import "fmt"

func main() {
        var a, b, c int
        fmt.Println(&a, &b, &c) // 0x1040a124 0x1040a128 0x1040a12c
}
```

It is not possible to create a Go program where two variables share the same storage location in memory. It is possible to create two variables whose contents _point_ to the same storage location, but that is not the same thing as two variables who share the same storage location.

```
package main

import "fmt"

func main() {
        var a int
        var b, c = &a, &a
        fmt.Println(b, c)   // 0x1040a124 0x1040a124
        fmt.Println(&b, &c) // 0x1040c108 0x1040c110
}
```

In this example, `b` and c hold the same value–the address of `a`–however, `b` and `c` themselves are stored in unique locations. Updating the contents of `b` would have no effect on `c`.

# But maps and channels are references, right?

Wrong. Maps and channels are not references. If they were this program would print `false`.

```
package main

import "fmt"

func fn(m map[int]int) {
        m = make(map[int]int)
}

func main() {
        var m map[int]int
        fn(m)
        fmt.Println(m == nil)
}
```

If the map `m` was a C++ style reference variable, the `m` declared in `main` and the `m` declared in `fn` would occupy the same storage location in memory. But, because the assignment to `m` inside `fn` has no effect on the value of `m` in main, we can see that maps are not reference variables.

# Conclusion

Go does not have pass-by-reference semantics because Go does not have reference variables.

Next: [If a map isn’t a reference variable, what is it?](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)

### Related Posts:

1. [If a map isn’t a reference variable, what is it?](https://dave.cheney.net/2017/04/30/if-a-map-isnt-a-reference-variable-what-is-it)
2. [Declaration scopes in Go](https://dave.cheney.net/2016/12/15/declaration-scopes-in-go)
3. [Understand Go pointers in less than 800 words or your money back](https://dave.cheney.net/2017/04/26/understand-go-pointers-in-less-than-800-words-or-your-money-back)
4. [Slices from the ground up](https://dave.cheney.net/2018/07/12/slices-from-the-ground-up)
