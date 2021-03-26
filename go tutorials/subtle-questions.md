- uint 不能直接相减，结果是负数会变成一个很大的 uint。
- channel 一定记得 close。
- goroutine 记得 return 或者中断，不然容易造成 goroutine 占用大量 CPU。
- 从 slice 创建 slice 的时候，注意原 slice 的操作可能导致底层数组变化。如果你要创建一个很长的 slice，尽量创建成一个 slice 里存引用，这样可以分批释放，避免 gc 在低配机器上 stop the world
- 面试的时候尽量了解协程，线程，进程的区别。
- 明白 channel 是通过注册相关 goroutine id 实现消息通知的。
- slice 底层是数组，保存了 len，capacity 和对数组的引用。
- 如果了解协程的模型，就知道所谓抢占式 goroutine 调用是什么意思。- 尽量了解互斥锁，读写锁，死锁等一些数据竞争的概念，debug 的时候可能会有用。
- 尽量了解 golang 的内存模型，知道多小才是小对象，为什么小对象多了会造成 gc 压力。

## Test related

- go test 出现**build failed**

go test 与其他的指定源码文件进行编译或运行的命令程序一样（参考：go run 和 go build），会为指定的源码文件生成一个虚拟代码包——“command-line-arguments”.执行 go test 命令时加入这个测试文件需要引用的源码文件，在命令行后方的文件都会被加载到 command-line-arguments 中进行编译。

比如开发 opcua 插件的时候需要用到最外层的 mocks 包里的某个 mock 文件,一开始执行下面的命令会出现**build fail**:

```
go test -cover  ./... -coverprofile=coverage.out -v -coverpkg ./...
```

解决办法是把 Mock 文件目录加入命令:

```
go test -cover  ./... -coverprofile=coverage.out -v ../../mocks/  -coverpkg ./...
```

- go vet 检查`shadows declaration`

shadow declaration 是指类似下面的现象:

```go
func main() {
	x := 1
	println(x)		// 1
	{
		println(x)	// 1
		x := 2
		println(x)	// 2	// 新的 x 变量的作用域只在代码块内部
	}
	println(x)		// 1
}
```

官方的`go vet`可以检查这种问题:

```
go vet ./...
```
