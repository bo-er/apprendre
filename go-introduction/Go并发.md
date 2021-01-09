## Goroutine

goroutine 是由 Go 运行时管理的轻量级线程。

```go
go f(x,y,z)
```

会启动一个新的 Goroutine 并且执行

```go
f(x,y,z)
```

对 f,x,y 跟 z 的取值发生在当前 goroutine,而 f 函数的执行发生在新的 goroutine 中。
goroutine 在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync 包提供了这种能力，不过在 go 中并不会经常用到，因为还有其它的办法。

## 通道

通道是带有类型的管道，你可以通过通道操作符<-来发送或者接收值,箭头就是数据流的方向。

```go
ch <- v   //将v发送至通道ch
v := <-ch //从ch接收值并且赋予v
```

和映射与切片一样，通道在使用前必须创建:

```go
ch := make(chan int)
```

默认情况下，发送和接收操作在另一端准备好之前都会阻塞。这使得 goroutine 可以在没有显式的锁或者条件变量的情况下进行同步。

下面是一个对切片中的数进行求和，将任务分配给两个 goroutine 的例子。

```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
```

### 带缓冲的通道

通道是可以带缓冲的，将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的通道:

```go
ch := make(chan int,100)
```

仅当通道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。

### range 和 close

发送者可以通过 close 关闭一个通道来表示没有需要发送更多的值了。接收者可以通过为接收表达式分配第二个参数来测试通道是否被关闭。若没有值可以接收且通道已被关闭，那么在执行完

```go
v,ok := <- ch
```

之后 ok 会被设置为 false,ok 为 false 表示通道已经被关闭了

循环 for i := range c 会不断从通道接收值，直到它被关闭。

注意: 只有发送者才能关闭通道，接收者不能。向一个已经关闭的通道发送数据将引发程序 panic。

通道与文件不同，通常情况下不需要关闭他们。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，例如终止一个 range 循环。

例子:

```go
package main

import (
	"fmt"
)

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
```

打印结果为:

```
0
1
1
2
3
5
8
13
21
34
```

### select 语句

select 语句使一个 goroutine 可以等待多个通信操作。

select 会阻塞直到某个分支可以继续执行，这时就会执行该分支。当多个分支都准备好时会随机选择一个执行。

```go
package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

```

打印结果为:

```
0
1
1
2
3
5
8
13
21
34
quit
```

### default

当 select 中的其他分支都没有准备好时，default 分支就会执行。

为了在尝试发送或者接收的时候不发声阻塞，可以使用 default 分支:

```go
select {
case i := <-c:
    // 使用 i
default:
    // 从 c 中接收会阻塞时执行
}
```

例子:

```go
package main

import (
	"fmt"
	"time"
)

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
```

### sync.Mutex

我们已经看到通道非常适合在各个 Goroutine 之间进行通信

但是如果我们并不需要通信呢？比如说，如果我们只是想要保证每次只有一个 goroutine 能够访问一个共享的变量，从而避免冲突？

这里涉及的概念叫做`互斥(mutual exclusion)`,我们通常使用`互斥锁(Mutex)`这一数据结构来提供这种机制。go 标准库中提供了 sync.Mutex 互斥锁类型及其两个方法:

```
Lock
Unlock
```

我们可以通过在代码前调用 Lock 方法，在代码后调用 Unlock 方法来保证一段代码的互斥执行。我们也可以用 defer 语句来保证互斥锁一定会被解锁。

Mutex 的使用方法：

```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter 的并发使用是安全的。
type SafeCounter struct {
	v   map[string]int
	mux sync.Mutex
}

// Inc 增加给定 key 的计数器的值。
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	c.v[key]++
	c.mux.Unlock()
}

// Value 返回给定 key 的计数器的当前值。
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	// Lock 之后同一时刻只有一个 goroutine 能访问 c.v
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}
```

### 出现死锁的情况

- 直接读取空的通道

  不管是有缓存还是无缓存的通道，直接读取空通道都会死锁

  ```go
  package main

  func main() {
    ch := make(chan int)
    <-ch
  }
  ```

  解决办法: 采用 select case default,执行 default

  ```go
  package main

  import (
  "fmt"
  )

  func main() {
    ch := make(chan int, 3)
    ch <- 1
    select {
      case v := <-ch:
        fmt.Println(v)
      default:
        fmt.Println("chan has no data")
    }
  }
  ```

  将会打印: 1,如果不往通道添加 1 将打印"chan has no data"

- 无缓冲通道产生死锁

  往无缓冲通道中写入数据

  ```go
  package main

  import "fmt"

  func main() {
      ch := make(chan int)
      ch <- 1
      fmt.Println(<-ch)

  }
  ```

  一个解决办法是开启一个 goroutine 来往通道写入

  ```go
  package main

  import "fmt"

  func main() {
  ch := make(chan int)

  go func() {
      ch <- 1 // 开启子goroutine写入数据
  }()

  fmt.Println(<-ch) // 阻塞住，一旦ch有数据，则读取成功
  }

  ```

  另一个解决办法是采用有缓冲的通道

- 超过有缓冲通道容量产生死锁

  下面是一个简单的例子:

  ```go
  package main

  import "fmt"

  func main() {
      ch := make(chan int,1)
      ch <- 1
      ch <- 1
      fmt.Println(<-ch)

  }
  ```

  解决办法: 使用 select case default

  ```go
  package main

  import (
      "fmt"
      )

  func main() {
      ch := make(chan int, 3)

      for i:=1;i<5;i++{
      select {
          case ch<- i:
          fmt.Println(i)
          default:
          fmt.Println("装不下更多")
      }

      }
  }

  ```

- for range 产生死锁

  ```go
  package main

      import (
      "fmt"
      )

      func main() {
      ch := make(chan int, 3)

      ch <- 1
      ch <- 2
      ch <- 3

      // range 一直读取直到chan关闭，否则产生阻塞死锁
      for v := range ch {
          fmt.Println(v)
        }
      }
  ```

  执行上面的代码会得到这样的结果:
  由于 range 会一直读取到通道关闭，因此读取完全部数据后由于通道没有关闭，for range 尝试从没有内容的通道获取数据导致死锁。

  ```go
  1
  2
  3
  fatal error: all goroutines are asleep - deadlock!

  goroutine 1 [chan receive]:
  main.main()
      /tmp/sandbox815460815/prog.go:16 +0x114

  ```

  - 第一种解决办法是在 for range 中加上 if v == 3 {close(ch)}
  - 第二种解决办法是开启子 goroutine,主 goroutine sleep 几秒后退出

    ```go
    package main

    import (
    "fmt"
    "time"
    )

    func main() {
    ch := make(chan int, 3)

    ch <- 1
    ch <- 2
    ch <- 3

    go func(){
        for v := range ch {
        fmt.Println(v)
        }
    }()
    time.Sleep(5* time.Second)
    }
    ```

    实际上退出时间根本不需要 5 秒，而是在子 goroutine 出现死锁错误后就退出了。
