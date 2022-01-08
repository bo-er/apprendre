## for

- Go 只有一种循环结构：for 循环。

  基本的 for 循环由三部分组成，它们用分号隔开：

        初始化语句：在第一次迭代前执行
        条件表达式：在每次迭代前求值
        后置语句：在每次迭代的结尾执行

  初始化语句通常为一句短变量声明，该变量声明仅在 for 语句的作用域中可见。

  一旦条件表达式的布尔值为 false，循环迭代就会终止。

  注意：和 C、Java、JavaScript 之类的语言不同，Go 的 for 语句后面的三个构成部分外没有小括号， 大括号 { } 则是必须的。

  实例:

  ```
  package main

  import "fmt"

  func main() {
      sum := 0
      for i := 0; i < 10; i++ {
          sum += i
      }
      fmt.Println(sum)
  }
  ```

- for 循环的初始化语句跟后置语句都是可选的:

  ```
  package main

  import "fmt"

  func main() {
      sum := 1
      for ; sum < 1000; {
          sum += sum
      }
      fmt.Println(sum)
  }

  ```

- for 就是 go 语言中的 while,如果要把 for 当做 while 使用,去掉分号就可以:

  ```
  package main

  import "fmt"

  func main() {
      sum := 1
      for sum < 1000 {
          sum += sum
      }
      fmt.Println(sum)
  }

  ```

- 无限循环

  如果省略循环条件，该循环就不会结束，因此无限循环可以写得很紧凑。

  ```
  package main

  func main() {
      for {
      }
  }

  ```

## if

Go 的 if 语句跟 for 循环类似，表达式外无需小括号(),但是大括号是必须的。

```
package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
```

- if 的简短语句

  同 for 一样， if 语句可以在条件表达式前执行一个简单的语句。

  该语句声明的变量作用域仅在 if 之内。

  （在最后的 return 语句处使用 v 看看。）

  ```
  package main

  import (
      "fmt"
      "math"
  )

  func pow(x, n, lim float64) float64 {
      if v := math.Pow(x, n); v < lim {
          return v
      }
      return lim
  }

  func main() {
      fmt.Println(
          pow(3, 2, 10),
          pow(3, 3, 20),
      )
  }
  ```

- if 跟 else

  if 的简短语句声明的变量除了在 if 语句后的代码块可以访问，在 else 的代码块也可以访问。

  ```
  package main

  import (
      "fmt"
      "math"
  )

  func pow(x, n, lim float64) float64 {
      if v := math.Pow(x, n); v < lim {
          return v
      } else {
          fmt.Printf("%g >= %g\n", v, lim)
      }
      // 这里开始就不能使用 v 了
      return lim
  }

  func main() {
      fmt.Println(
          pow(3, 2, 10),
          pow(3, 3, 20),
      )
  }
  ```

## Switch

如果你的项目需要编写一连串`if-else`语句，那么为了简便使用 switch。它运行第一个值等于条件表达式的 case 语句。

Go 的 switch 语句类似于 C、C++、Java、JavaScript 和 PHP 中的，不过 Go 只运行选定的 case，而非之后所有的 case。 实际上，Go 自动提供了在这些语言中每个 case 后面所需的 break 语句。 除非以 fallthrough 语句结束，否则分支会自动终止。 Go 的另一点重要的不同在于 switch 的 case 无需为常量，且取值不必为整数。

```
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.\n", os)
	}
}
```

- switch 的求值顺序

  switch 的 case 语句从上到下顺次执行，直到匹配成功时停止。

  （例如，

  switch i {
  case 0:
  case f():
  }

  在 i==0 时 f 不会被调用。

  ```
  package main

  import (
      "fmt"
      "time"
  )

  func main() {
      fmt.Println("When's Saturday?")
      today := time.Now().Weekday()
      switch time.Saturday {
      case today + 0:
          fmt.Println("Today.")
      case today + 1:
          fmt.Println("Tomorrow.")
      case today + 2:
          fmt.Println("In two days.")
      default:
          fmt.Println("Too far away.")
      }
  }
  ```

- 没有条件的 switch

  没有条件的 switch 指 switch{},它跟 switch true 一样。

  这种形式能将一长串 if-then-else 写得更加清晰。

  ```
  package main

  import (
      "fmt"
      "time"
  )

  func main() {
      t := time.Now()
      switch {
      case t.Hour() < 12:
          fmt.Println("Good morning!")
      case t.Hour() < 17:
          fmt.Println("Good afternoon.")
      default:
          fmt.Println("Good evening.")
      }
  }
  ```

## defer

defer 语句会将函数推迟到外层函数返回之后执行。

推迟调用的函数其参数会立即求值，但直到外层函数返回前该函数都不会被调用。

```
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")
}
```

- defer 栈

  推迟的函数调用会被压入一个栈中。当外层函数返回时，被推迟的函数会按照后进先出的顺序调用。

  ```
  package main

  import "fmt"

  func main() {
      fmt.Println("counting")

      for i := 0; i < 10; i++ {
          defer fmt.Println(i)
      }

      fmt.Println("done")
  }
  ```

  上面的程序执行结果为:

  ```
  counting
  done
  9
  8
  7
  6
  5
  4
  3
  2
  1
  0
  ```

- 扩展阅读:

  👁 https://blog.go-zh.org/defer-panic-and-recover
