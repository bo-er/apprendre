# Google I/O 2012

## 举例

### 简单无聊的打印

```go
func boring(msg string){
    for i:=0;;i++{
        fmt.Println(msg,i)
        time.Sleep(time.Second)
    }
}

func main(){
    boring("boring")
}
```

打印结果:

```go
boring 0
boring 1
boring 2
boring 3
boring 4
boring 5
```

### 加了 go 关键字的打印

```go

func boring(msg string){
    for i:=0;;i++{
        fmt.Println(msg,i)
        time.Sleep(time.Second)
    }
}

func main(){
    go boring("boring")
}
```

打印结果有些让人惊讶:

```go

```

程序没有执行任何打印就退出了，如果在 go boring("boring")下一行加上 time.Sleep(time.Second)可以看到有打印结果:

```go
boring 0
boring 1
```

## Goroutines

goroutine 是一个独立执行的函数，由 go 声明触发。

它有自己的 call stack,会随着需要增长或者收缩。

goroutine 开销很小，实际使用可以有几百甚至上千个 goroutines

goroutine 不是线程

一个线程可以有上千个 goroutine

相反，goroutines 是按照需要动态多路复用到线程上，以保持所有 goroutines 的运行状态。

当然如果把 goroutine 看成是开销很小的线程，你也可以这么理解。

## channels

Go 中的 channels 提供了一个建立在两个 goroutine 之间的连接，允许他们相互通信。

```go
var c chan int
c = make(chan int)

//上面的两行写法等价于下面的一行写法
c := make(chan int)
```

给通道发送信号

```go

c <- 1
```

从通道接收信号

```go
value  = <-c
```

### 使用通道

改造上面的无聊函数:

```go
func boring(c chan string, msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		c <- "ready"
	}
}

func main() {
	c := make(chan string)
	go boring(c, "boring")
	for i := 0; i < 5; i++ {
		<-c
	}
}
```

打印结果:

```
boring 0
boring 1
boring 2
boring 3
boring 4
boring 5
```

## Synchronization

当 main 函数执行 <- c,它会等待一个值从 c 中被取出。

同理，当 boring 函数执行 c <- "ready"时，它也会等待一个`接收者`做好准备。

一个发送者跟一个接收者必须同时做好进行通信的准备，否则 go 将等待直到他们做好准备。

因此通道具有**通信**和**同步**两个属性。

### 有缓冲通道

Go 也可以创建有缓冲的通道

有缓冲通道移除了通道同步的属性

## 设计模式

### 生成器:返回一个通道的函数

go 中的通道也是一级值，就跟字符串和整数一样。

看下面的例子:

```go
func boring(msg string) <-chan string { //返回一个只接受的通道
	c := make(chan string)
	go func() { //在生成通道的函数内部启动goroutine
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c //将生成的通道返回给调用者
}

func main() {
	c := boring("I am boring!") //返回一个通道的函数
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
}
```

打印结果:

```
You say: "I am boring! 0"
You say: "I am boring! 1"
You say: "I am boring! 2"
You say: "I am boring! 3"
You say: "I am boring! 4"
```

如果生成多个通道:

```go
func main() {
	steve := boring("steve is boring!") //返回一个通道的函数
	eve := boring("eve is boring")
	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-steve)
		fmt.Printf("You say: %q\n", <-eve)
	}
}
```

打印结果是:

```
You say: "steve is boring! 0"
You say: "eve is boring 0"
You say: "steve is boring! 1"
You say: "eve is boring 1"
You say: "steve is boring! 2"
You say: "eve is boring 2"
You say: "steve is boring! 3"
You say: "eve is boring 3"
You say: "steve is boring! 4"
You say: "eve is boring 4"
```

之所以 steve 跟 eve 轮流出现是因为通道的同步属性，如果 eve 通道已经做好了发送数据的准备而 steve 通道没有那么 eve 通道将等待 steve 通道。如果不希望这种固定的执行模式，可以使用 fan-in 函数来让 go 从准备好的通道立即取出数据。

#### FYI

- fan-in : Fan-in is the number of devices that can be permitted to drive a single logic gate input.
- Fan-out is the number of devices a single logic gate output can drive.
- Multiplexing is the selection of one signal from many available signals to be routed to a single output.
- Demultiplexing is the reverse process of Multiplexing

### Fan-in 无序模式

下面的例子也是一个很好的示范，说明了 go 中的一切都是值，包括<- chan

```go
func boring(msg string) <-chan string { //返回一个只接受的通道
	c := make(chan string)
	go func() { //在生成通道的函数内部启动goroutine
		for i := 0; ; i++ {
			c <- fmt.Sprintf("%s %d", msg, i)
		}
	}()
	return c //将生成的通道返回给调用者
}

func fanIn(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for {c <- <-input1}}()
	go func() {
		for {c <- <-input2}}()
	return c
}

func main() {
	c := fanIn(boring("Steve"), boring("Eve"))
	for i := 0; i < 10; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("You are boring,I am leaving")
}
```

打印结果是无序的:

```
Eve 0
Eve 1
Steve 0
Steve 1
Eve 2
Steve 2
Eve 3
Eve 4
Steve 3
Eve 5
You are boring,I am leaving
```

### Fan-in 有序模式

往一个通道发送一个通道，让 goroutine 轮流执行

接收所有的消息，并且

首先定义一个 Message 结构，它包含一个用于回复的通道。

```go
type Message struct{
    str string
    wait chan bool
}
```

## SELECT

select 使得 go 的并发变成了语言的一部分

select 语句提供了另外一种解决多通道问题的方法。它跟 switch 很像，但是不同于 switch，select 的每一个 case 都是一次通信过程:

- 所有的通道都会被评估(evaluate)
- selection 在至少一次通信可以继续前会阻塞
- 如果多个 selection 可以执行，select 会伪随机地选择一个 case
- 如果有 default case 并且没有通道准备好将立即执行

### 使用 SELECT 的 FAN-IN

之前的通道生成器是这样实现的:

```go
func main(input1,input2 <- chan string) <- chan string{
    c := make(chan string)
    go func(){for{c <- <- input1}}()
    go func(){for{c <- <- input2}}()
    return c
}
```

上面的例子启动了两个 goroutien 来从两个通道拷贝输入，但是使用 select 之后我们只需要启动一个 goroutine,谁准备好了就从谁那里拷贝：

```go
func fanIn(input1,input2 <- chan string) <- chan string{
    c := make(chan string)
    go func(){
        for{
            select{
                case s := <- input1: c <-s
                case s := <- input2: c<-s
            }
        }
    }()
    return c
}

```

### 使用 Select 来设置超时

`time.After`函数返回了一个阻塞指定时间的通道。时间间隔过后，通道将传递一次当前时间。

```go

func main(){
    c := boring("Joe")
    for {
        select {
            case s := <- c:
                fmt.Println(s)
            case <- time.After(1 * time.Second):
                fmt.Println("You are too slow!")
                return
        }
    }
}
```

也可以在循环外设置超时:

```go
func main(){
    c := boring("Joe")
    timeout := time.After(5 *time.Second)
    for{
        select{
            case s := <-c:
                fmt.Println(s)
            case <- timeout:
                fmt.Println("You talk too much")
                return
        }
    }
}
```

### 退出通道(quit channel)

```go
quit := make(chan bool)
c := boring("Joe",quit)
for i := rand.Intn(10);i >= 0; i--{
    fmt.Println(<-c)
}
quit <- true

```

```go
select{
    case c <- fmt.Sprintf("%s:%d",msg,i):
        //什么也不做
    case <- quit:
        return
}
```
