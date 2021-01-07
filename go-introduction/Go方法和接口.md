## 方法

Go 语言没有类。不过你可以为类型定义方法。方法也使用 func 关键字（因为它就是函数，不过稍微特殊）。

方法就是一类带特殊的 **接收者** 参数的函数。

方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间。

```go
// A Mutex is a data type with two methods, Lock and Unlock.
type Mutex struct         { /* Mutex fields */ }
func (m *Mutex) Lock()    { /* Lock implementation */ }
func (m *Mutex) Unlock()  { /* Unlock implementation */ }

// NewMutex has the same composition as Mutex but its method set is empty.
type NewMutex Mutex

// The method set of PtrMutex's underlying type *Mutex remains unchanged,
// but the method set of PtrMutex is empty.
type PtrMutex *Mutex

// The method set of *PrintableMutex contains the methods
// Lock and Unlock bound to its embedded field Mutex.
type PrintableMutex struct {
	Mutex
}

// MyBlock is an interface type that has the same method set as Block.
type MyBlock Block
```

### 为结构体定义方法

在下面的例子中， Abs 方法拥有一个名为 v,类型为 Vertext 的接收者。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
```

### 方法即函数

记住: 方法只是个带接收者参数的函数。
现在这个 Abs 的写法就是个正常的函数，功能并没有什么变化。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(Abs(v))
}
```

### 为非结构体定义方法

你也可以为非结构体类型声明方法。
下面是一个带 Abs 方法的数值类型 MyFloat
你只能为在同一个包内定义的类型的接收者声明方法，而不能为其它包内定义的类型（包括 int 之类的内建类型）的接收者声明方法。

```go
package main

import (
	"fmt"
	"math"
)

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
```

### 指针接收者

你可以为指针接收者声明方法

这意味着对于某类型 T，接收者的类型可以用\*T 的语法。

下面的例子为\*Vertex 定义了 Scale 方法

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
```

指针接收者的方法可以修改接收者指向的值(就像 Scale 在这做的)。由于方法经常需要修改它的接收者，指针接收者比值接收者更常用。
将

```go
func (v *Vertex) Scale(f float64)
```

改为

```go
func (v Vertex) Scale(f float64)
```

会发现打印值不会改变。这是因为如果使用值接收者，那么 Scale 方法会对原始 Vertext 值的副本进行操作。（对于函数的其他参数也是如此。）Scale 方法必须使用指针接收者来更改 main 函数中声明的 Vertext 的值。

### 指针与函数

现在我们要把 Abs 和 Scale 方法重写为函数。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10)
	fmt.Println(Abs(v))
}

```

同样去掉\*将 Scale 函数变为

```go
func Scale(v Vertex, f float64){}
```

执行函数将会看到错误信息:

```go
./prog.go:23:8: cannot use &v (type *Vertex) as type Vertex in argument to Scale
```

可以看到为结构体定义的函数跟普通的函数还是有区别的，为结构体定义的指针接收者方法也可以通过值来调用。而定义好的参数为指针的函数，不能传入值作为参数。

### 方法与指针重定向

对比前面`指针接收者`跟`指针与函数`两个例子，你大概会注意到带指针参数的函数必须接受一个指针:

```go
var v Vertex
ScaleFunc(v,5)   // 编译错误!
ScaleFunc(&v,5)  //OK

```

而以指针为接收者的方法被调用时，接收者既能为值又能为指针：

```go

var v Vertex
v.Scale(5)    //OK
p := &v
p.Scale(10)   //OK
```

对于语句 v.Scale(5),即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。也就是说，由于 Scale 方法有一个指针接收者，为了方便起见，Go 会将语句 v.Scale(5)解释为(&v).Scale(5)。

例子:

```go
package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	p.Scale(3)
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
```

打印结果:

```
{60 80} &{96 72}
```

前面提到，带指针参数的函数必须接受一个指针。
同样的事情发生在相反的方向。接受一个值作为参数的函数必须接受一个指定类型的值。

```go

var v Vertex
fmt.Println(AbsFunc(v))    // OK
fmt.Println(AbsFunc(&v))   //编译错误!
```

而以值为接收者的方法被调用时，接收者既能为值又能为指针:

```go
var v Vertex
v.Scale(5)
p := &v
p.Scale(10)
```

对于语句 v.Scale(5),即便 v 是个值而非指针，带指针接收者的方法也能被直接调用。也就是说，由于 Scale 方法有一个指针接收者，为方便起见，Go 会将语句 v.Scale(5)解释为(&v).Scale(5)

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
	fmt.Println(AbsFunc(v))

	p := &Vertex{4, 3}
	fmt.Println(p.Abs())
	fmt.Println(AbsFunc(*p))
}

```

### 选择值或者指针作为接收者

使用指针接收者的原因有两点:

- 方法能够修改其接收者指向的值
- 避免在每次调用方法时复制该值。若值的类型为大型结构体的时，这样做会更加高效。
  在本例中，Scale 和 Abs 接收者的类型为\*Vertex,即便 Abs 并不需要修改其接收者。

通常来说，所有给定类型的方法都应该有值或者指针接收者，但并不应该二者混用。

```go
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Vertex{3, 4}
	fmt.Printf("Before scaling: %+v, Abs: %v\n", v, v.Abs())
	v.Scale(5)
	fmt.Printf("After scaling: %+v, Abs: %v\n", v, v.Abs())
}
```

## 接口

接口类型是由一组方法签名定义的集合。

接口类型的变量可以保存任何实现了这些方法的值。

下面的例子有一处错误:

```go
package main

import (
	"fmt"
	"math"
)

type Abser interface {
	Abs() float64
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat 实现了 Abser
	a = &v // a *Vertex 实现了 Abser

	// 下面一行，v 是一个 Vertex（而不是 *Vertex）
	// 所以没有实现 Abser。
	a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

```

可以看到 Vertex 指针接收者实现了 Abser 接口，而值没有实现 Abser 接口，因此下面的一行是错误的。

```go
fmt.Println(a.Abs())
```

这就是指针接收者的方法不能通过值来调用的一个例子。

### 接口与隐式实现

类型通过实现一个接口的所有方法来实现这个接口。既然无需专门的显示声明，也就没有 JAVA 中的"implements"关键字。隐式接口从接口的实现中解耦了定义，这样接口的实现可以出现在任何包中，无需提前准备。因此，也就无需在每一个实现上增加新的接口名称。这样同时也鼓励了明确的接口定义。

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// 此方法表示类型 T 实现了接口 I，但我们无需显式声明此事。
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
```

### 接口值

接口也是值。它们可以像其他值一样传递。

接口值可以用作函数的参数或者返回值。

在内部，接口值可以看做包含值和具体类型的元组：

```
(value,type)
```

接口值保存了一个具体底层类型的具体值。

接口值调用方法时会执行其底层类型的同名方法。

```go
package main

import (
	"fmt"
	"math"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type F float64

func (f F) M() {
	fmt.Println(f)
}

func main() {
	var i I

	i = &T{"Hello"}
	describe(i)
	i.M()

	i = F(math.Pi)
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### 底层值为 nil 的接口值

即便接口内的具体值为 nil，方法仍然会被 nil 接收者调用。

在一些语言中，这会触发一个空指针异常，但在 Go 中通常会写一些方法来优雅地处理它,下面的例子中的 M 方法就是一个典型的例子。

注意: 保存了 nil 具体值的接口其本身并不为 nil

```go
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I

	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

打印结果:

```
(<nil>, *main.T)
<nil>
(&{hello}, *main.T)
hello

```

### nil 接口值

nil 接口值既不保存值也不保存具体类型
为 nil 接口调用方法会产生运行时的错误，因为接口的元组内并未包含能够指明该调用哪个`具体`方法的类型。

```go
package main

import "fmt"

type I interface {
	M()
}

func main() {
	var i I
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
```

### 空接口

指定了零个方法的接口值被称为*空接口:*

```
interface{}
```

空接口可以保存任何类型的值。（因为每个类型都至少实现了零个方法)
空接口被用来处理未知类型的值。例如，fmt.Print 可接受类型为 interface{}的任意数量的参数。

```go
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)

	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

```

打印结果为:

```
(<nil>, <nil>)
(42, int)
(hello, string)
```

### 类型断言

`类型断言`提供了访问接口值底层具体值的方式。

```
t := i.(T)
```

该语句断言接口值 i 保存了具体类型 T,并将其底层类型为 T 的值赋予变量 t。
若 i 并未保存 T 类型的值，该语句就会触发一个 panic。

为了`判断`一个接口值是否保存了一个特定的类型，类型断言可以返回两个值：其底层值以及一个报告断言是否成功的布尔值。

```
t,ok := i.(T)
```

若 i 保存了一个 T,那么 T 将会是其底层值，而 ok 为 true
否则,ok 将为 false 而 t 将为 T 类型的零值，程序并不会产生恐慌。

```go
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // 报错(panic)
	fmt.Println(f)
}
```

打印结果为:

```


hello
hello true
0 false
panic: interface conversion: interface {} is string, not float64

goroutine 1 [running]:
main.main()
	/tmp/sandbox005824404/prog.go:17 +0x1fe
```

### 类型选择

`类型选择`是一种按顺序从几个类型断言中选择分支的结构。

类型选择与一般的 switch 语句相似，不过类型选择中的 case 为类型（而非值）。

它们针对给定接口值所存储的值的类型进行比较。

```go
switch v := i.(type) {
    case T:
        // v 的类型为 T
    case S:
        // v 的类型为 S
    default:
        //没有匹配,v与i的类型相同
}
```

类型选择中的声明与类型断言`i.(T)`的语法相同，只是具体类型 T 被替换成了关键字`type`

此选择语句判断接口值 i 保存的值类型是 T 还是 S。在 T 或 S 的情况下，变量 v 会分别按 T 或 S 类型保存 i 拥有的值。但是如果没有匹配，switch 到了默认 case，V 的值与接口类型跟 i 相同。

比如 `switch v := true.(type)`,v 的值是 true 然后类型是 bool 跟 true 保持一致。

```go
package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
```

打印结果:

```
Twice 21 is 42
"hello" is 5 bytes long
I don't know about type bool!
```

### Stringer

fmt 包中定义的 Stringer 是最普遍的接口之一

```go
type Stringer interface{
    String() string
}
```

Stringer 是一个可以用字符串描述自己的类型。fmt 包（还有很多包）都通过此接口来打印值。

```go
package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}

```

## 错误

Go 语言使用 error 值来表示错误状态

与 fmt.Stringer 类似，error 类型是一个内建接口:

```go
type error interface{
    Error() string
}
```

跟 fmt.Stringer 一样，fmt 包在打印值的时候会寻找 error 接口

通常函数会返回一个 error 值，调用该函数的代码应该判断这个错误是否等于 nil 来进行错误处理。

```go
i, err := strconv.Atoi("42")
if err != nil {
    fmt.Printf("couldn't convert number: %v\n", err)
    return
}
fmt.Println("Converted integer:", i)
```

error 为 nil 时表示成功;非 nil 的 error 表示失败。

```go
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

//由于MyError实现了Error方法，因此MyError的指针接收者实现了error接口，因此
//下面的run函数返回一个&MyError就是返回了error
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

//
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
打印结果:
```

at 2009-11-10 23:00:00 +0000 UTC m=+0.000000001, it didn't work

```

```

如果上面将 func (e \*MyError) Error() string{}改为 func (e MyError) Error() string{}，将会出现下面的错误:

```go
./prog.go:19:16: cannot use MyError literal (type MyError) as type error in return argument:
	MyError does not implement error (Error method has pointer receiver)
```

由于 MyError 的值接收者没有实现 Error 接口，因此 MyError 的值无法作为错误返回。

### Reader

io 包指定了 io.Reader 接口，它表示从数据流的末尾进行读取。
Go 标准库包含了该接口的许多实现，包括文件、网络连接、压缩和加密等等。
io.Reader 接口有一个 Read 方法:

```go
func (T) Read(b []byte) (n int, err error)
```

Read 用数据填充给定的字节切片并返回填充的字节数和错误值。在遇到数据流的结尾时，它会返回一个 io.EOF 错误。

下面的代码创建了一个 strings.Reader 并且以每次 8 字节的速度读取它的输出

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```

打印结果：

```
n = 8 err = <nil> b = [72 101 108 108 111 44 32 82]
b[:n] = "Hello, R"
n = 6 err = <nil> b = [101 97 100 101 114 33 32 82]
b[:n] = "eader!"
n = 0 err = EOF b = [101 97 100 101 114 33 32 82]
b[:n] = ""
```
