## Go 中控制并发

- WaitGroup
- Context

WaitGroup 适用场景是多个 Goroutine 执行同一事情
例子:

```go

func main() {
var wg sync.WaitGroup
wg.Add(2)
go func() {
    fmt.Println("我完成了工作")
    wg.Done()
}()

go func() {
    fmt.Println("我也完成了工作")
    wg.Done()
}()

wg.Wait()

}

```

上面是任务全部执行完结束，如何主动通知让程序停止？

可以使用`channel` + `select`

```go

func main() {
	stop := make(chan bool)
	counter := 0
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("任务结束咯!")
				return
			default:
				counter++
				fmt.Printf("继续工作第%d个任务...\n", counter)
			}
		}
	}()

	time.Sleep(1 * time.Second)
	stop <- true
}
```

但是如果 goroutine 中启动了 goroutine，或者有多个 goroutine 怎么办？

### 使用 Context！

```go

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("任务结束咯!")
				
				return
			default:

				fmt.Println("继续工作...")
			}
		}
	}()

	time.Sleep(1 * time.Second)
	cancel()
}
```

一个 cancel 结束所有 go routine 的例子:

```go
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, "搬砖")
	go worker(ctx, "写代码")
	go worker(ctx, "造轮子")

	time.Sleep(5 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}

func worker(ctx context.Context, workname string) {

	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Printf("名为%s的工作停止!\n", workname)
				return
			default:
				fmt.Printf("正在进行%s的工作任务\n", workname)
				time.Sleep(1 * time.Second)
			}
		}
	}()

}

```
