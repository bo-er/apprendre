## 指针

Go 拥有指针。指针保存了值的内存地址。

类型 \*T 是指向 T 类型值的指针。其零值为 nil。

```go
var p *int \\定义一个指针
```

& 操作符会生成一个指向其操作数的指针。

```go
i := 42
p = &i
```

`*` 操作符表示指针指向的底层值。

```go
fmt.Println(*p)  // 通过指针 p 读取 i
*p = 21         // 通过指针 p 设置 i
```

这也就是通常所说的“间接引用”或“重定向”。

与 C 不同，Go 没有指针运算。

```go
package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值


	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}
```

## 结构体(struct)

一个结构体（struct）就是一组字段（field）。Go 把 JAVA 类的字段跟方法拆开了，结构体就相当于 JAVA 中的字段。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

### 结构体字段

通过`.`来访问结构体字段，类似 javascript

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)
}
```

### 结构体指针

结构体字段也可以通过结构体指针来访问。

如果我们有一个指向结构体的指针 p，那么可以通过 `(*p).X` 来访问其字段 X。

不过这么写太啰嗦了，所以语言也允许我们使用隐式间接引用，直接写 `p.X` 就可以。

```go
package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v
	p.X = 1e9
	fmt.Println(v)
}


```

### 结构体语法

结构体文法通过直接列出字段的值来新分配一个结构体。

使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

特殊的前缀 & 返回一个指向结构体的指针。

```go
package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // 创建一个 Vertex 类型的结构体
	v2 = Vertex{X: 1}  // Y:0 被隐式地赋予
	v3 = Vertex{}      // X:0 Y:0
	p  = &Vertex{1, 2} // 创建一个 *Vertex 类型的结构体（指针）
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
```

### 结构体嵌套

如果结构体里嵌套了接口那么该结构体可以用实现了该接口的其他结构体来初始化，即使他们的属性不完全一致。

## 数组

类型 `[n]T` 表示拥有 `n` 个 `T` 类型的值的数组。

表达式

```go
var a [10]int
```

会将变量 `a` 声明为拥有 10 个整数的数组。

数组的**长度是其类型的一部分**，因此数组不能改变大小。这看起来是个限制，不过没关系，Go 提供了更加便利的方式 来使用数组。

```go

package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}
```

## 切片

每个数组的大小都是固定的。而切片则为数组元素提供动态大小，因此实践中切片比数组更为常用。

类型`[ ]T`表示一个元素类型为 T 的切片。

切片通过两个下标来界定，即一个上界和一个下界，二者以冒号分隔。

```go
a[low : high]
```

它会选择一个半开区间，包括第一个元素，但排除最后一个元素。

以下表达式创建了一个切片，它包含 `a` 中下标从 1 到 3 的元素：

```go
a[1:4]
```

切片的底层数据类型是数组，所以下面通过切取数组获取切片的做法是 OK 的:

```go
package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}

```

### 切片就像数组的引用

切片并不存储任何数据，它只是描述了底层数组中的一段。

更改切片的元素会修改其底层数组中对应的元素。

与它共享底层数组的切片都会观测到这些修改。

```go
package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
```

```go
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

  //定义切片时定义匿名的结构体
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}
```

### 切片的默认行为

在进行切片时，你可以利用它的默认行为来忽略上下界。

切片下界的**默认值是 0**，上界则是该**切片的长度**。

对于数组

```go
var a [10]int
```

来说，以下切片是等价的:

```go
a[0:10]
a[:10]
a[0:]
a[:]
```

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
```

### 切片的长度与容量

切片拥有**长度**跟**容量**两个属性

- 切片的长度就是它所包含的元素个数。
- 切片的容量是从它的第一个元素开始数，到其底层数组元素末尾的个数。
- 切片 s 的长度和容量可以通过表达式 `len(s)`和 `cap(s)`来获取
- 你可以通过重新切片来扩展一个切片，给它提供足够的容量。试着修改实例程序中的切片操作，向外扩展它的容量，看看会发生什么。

```go
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
```

运行结果:

```go
len=6 cap=6 [2 3 5 7 11 13]
len=0 cap=6 []
len=4 cap=6 [2 3 5 7]
len=2 cap=4 [5 7]
```

### nil 切片

- 切片的零值是 nil

- nil 切片的长度和容量为 0 且没有底层数组

```go
package main

import "fmt"

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
```

- 用 make 创建切片
  切片可以用内建函数 make 来创建，这也是你创建动态数组的方式。

make 函数会分配一个元素为零值的数组并返回一个引用了它的切片。

```go
a := make([]int, 5)  // len(a)=5
```

要指定它的容量，需向 make 传入第三个参数。

```go
b := make([]int, 0, 5) // len(b)=0, cap(b)=5
b = b[:cap(b)] // len(b)=5, cap(b)=5
b = b[1:]      // len(b)=4, cap(b)=4
```

```go
package main

import "fmt"

func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

```

### 切片的切片

切片可以包含任何类型，甚至包括其它的切片

```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	// 创建一个井字板（经典游戏）
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// 两个玩家轮流打上 X 和 O
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}
```

### 向切片追加元素

为切片追加新的元素是种常用的操作，为此 Go 提供了内建的 append 函数。

```go
func append(s []T, vs ...T) []T
```

append 的第一个参数 s 是一个元素类型为 T 的切片，其余类型为 T 的值将会追加到该切片的末尾。
append 的结果是一个包含原切片所有元素加上新添加元素的切片，
当 s 的底层数组太小，不足以容纳所有给定的值时，它就会分配一个更大的数组。返回的切片会指向这个新分配的数组。

```go
package main

import "fmt"

func main() {
	var s []int
	printSlice(s)

	// 添加一个空切片
	s = append(s, 0)
	printSlice(s)

	// 这个切片会按需增长
	s = append(s, 1)
	printSlice(s)

	// 可以一次性添加多个元素
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}


```

### Range

for 循环的 range 形式可以遍历切片或者映射。
当使用 for 循环遍历切片的时候，每次迭代都会返回两个值。第一个值为当前元素的下标，第二个值为该下标所对应元素的一份副本。(go 语言始终是 copy by value,value 当然可以是指针,不过如果是指针 copy 的就是内存地址)

range 还可以通过\_来忽略下标。

```go
for i, _ := range pow
for _, value := range pow
```

若你只需要索引，忽略第二个变量即可。

```go
for i := range pow
```

```go
package main

import "fmt"

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}
}
```

```go
package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main() {
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

```

## 映射

映射将键映射到值
映射的零值为 nil。nil 映射既没有键，也不能添加键。
make 函数会返回给定类型的映射，并将其初始化备用。

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64 //两个float64变量
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
```

### 映射的语法

映射的语法跟结构体类似，不过必须有键名。

```go

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}

func main() {
	fmt.Println(m)
}
```

如果顶级类型只是一个类型名，你可以在语法的元素中省略它。

```go
package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
}

func main() {
	fmt.Println(m)
}
```

### 修改映射

在映射 m 中插入或者修改元素:

```go
m[key] =  elem
```

获取元素(单赋值):

```go
elem = m[key]
```

删除元素:

```go
delete(m,key)
```

通过双赋值检查某个键是否存在:

```go
elem,ok = m[key]

```

如果 key 在 m 中，ok 为 true;否则，ok 为 false
若 key 不在映射中，那么 elem 是该映射元素类型的零值
同样的，当从映射中获取某个不存在的键时，结果是映射的元素类型的零值。

注: 若 elem 或 ok 还未声明，你可以使用短变量声明:

```go
elem,ok := m[key]
```

```go
package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}
```

## 参考资料

- Go 切片：用法和本质 https://blog.go-zh.org/go-slices-usage-and-internals
