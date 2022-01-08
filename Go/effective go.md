## 介绍

Go 是一种新语言。尽管它借鉴了现有语言的思想，但它具有非同寻常的特性，使 Go 程序的特性与类似的程序不同。将 C ++或 Java 程序直接转换为 Go 不太可能产生令人满意的结果：Java 程序是用 Java 而不是 Go 编写的。另一方面，从 Go 角度考虑问题可能会产生一个成功但完全不同的程序。换句话说，要编写好语言，重要的是要了解其属性和习惯用法。了解 Go 编程中已建立的约定（例如命名，格式设置，程序构造等）也很重要，这样你编写的程序将易于其他 Go 程序员理解。

本文档提供了编写清晰，惯用的 Go 代码的技巧。它增强了[语言规范](https://golang.google.cn/ref/spec)，“ [Go](https://tour.golang.org/)之[旅”](https://tour.golang.org/)和“[如何编写 Go 代码”](https://golang.google.cn/doc/code.html)，您应该首先阅读所有这些内容。

### 例子

源码的作用不仅作为核心库，而且也是很好的如何使用语言的例子。此外，许多软件包都包含可运行的，自包含的可执行示例，您可以直接从[golang.org](https://golang.org/)网站上运行该示例 ，例如 [该](https://golang.org/pkg/strings/#example_Map)示例（如有必要，请单击“示例”一词以将其打开）。如果您对如何解决问题或可能如何实施解决方案有疑问，则库中的文档，代码和示例可以提供答案，想法和背景。

## 格式化

格式问题是最有争议但后果最不严重的问题。人们可以适应不同的格式样式，但是如果他们不必这样做会更好，如果每个人都遵循相同的样式，那么花在该主题上的时间就会更少。问题是如何在没有冗长的说明性样式指南的情况下接近这种乌托邦。

使用 Go，我们可以采用一种不寻常的方法，让机器处理大多数格式化问题。`gofmt`程序（也可作为`go fmt`，以软件包级别而不是源文件级别运行）读取 Go 程序，并以缩进和垂直对齐的标准样式发出源代码，并保留注释，并在必要时重新格式化注释。如果您想知道如何处理一些新的布局情况，请运行`gofmt`；如果答案似乎不正确，请整理一下程序（或提交有关的错误`gofmt`），请不要自行解决它。

作为例子，使用 Go 无需花时间对结构字段上的注释进行排列。 `Gofmt`将为您做到这一点。给出声明

```go
type T struct {
    name string // name of the object
    value int // its value
}
```

`gofmt` 会将将列对齐：

```go
type T struct {
    name    string // name of the object
    value   int    // its value
}
```

标准软件包中的所有 Go 代码都已使用格式化`gofmt`。

保留一些格式详细信息。非常简短：

- 缩进

  我们使用制表符进行缩进，`gofmt`并在默认情况下发出它们。仅在必要时使用空格。

- 线长

  Go 没有行长限制。如果一行代码太长了，则将其包裹起来并用额外的制表符缩进。

- 括弧

  Go 需要比 C 和 Java 更少的括号：控制结构（`if`， `for`，`switch`）的语法不需要括号。而且，运算符优先级层次更短更清晰

  像下面的代码就是它本身的意思，不像别的语言会有其它含义

  ```go
  x<<8 + y<<16
  ```

## 注释 Comment

Go 提供了 C 样式的`/* */`块注释和 C ++样式的`//`行注释。行注释是最常使用的；块注释主要显示为程序包注释，但在表达式中很有用，或者用于禁用大量代码。

该程序和 Web 服务器`godoc`处理 Go 源文件以提取有关软件包内容的文档。在顶级声明之前出现的注释（没有中间的换行符）将与声明一起被提取，以用作该项的解释性文本。这些注释的性质和样式决定了文档`godoc`生成的质量。

每个包都应在 package 子句前有一个*package 注释*，一个块注释。对于多文件包，包注释仅需要出现在一个文件中，任何一个都可以。包装评论应介绍包装，并提供与包装整体相关的信息。它会首先出现在`godoc`页面上，并应设置随后的详细文档。

```
/*
Package regexp implements a simple library for regular expressions.

The syntax of the regular expressions accepted is:

    regexp:
        concatenation { '|' concatenation }
    concatenation:
        { closure }
    closure:
        term [ '*' | '+' | '?' ]
    term:
        '^'
        '$'
        '.'
        character
        '[' [ '^' ] character-ranges ']'
        '(' regexp ')'
*/
package regexp
```

如果软件包很简单，则软件包注释可以简短。

```
// Package path implements utility routines for
// manipulating slash-separated filename paths.
```

注释不需要额外的格式，例如星号横幅。生成的输出甚至可能不会以固定宽度的字体显示，因此不必依赖对齐的间距`godoc`，例如`gofmt`，就可以了。注释是未解释的纯文本，因此 HTML 和其他注释（如`_this_`将*逐字复制）*不应该使用。`godoc`所做的一项调整是以固定宽度的字体显示缩进的文本，适用于程序片段。对于包注释 [`fmt`包](https://golang.google.cn/pkg/fmt/)使用此效果良好。

根据上下文的不同，`godoc`甚至可能不会重新格式化注释，因此请确保它们看起来直截了当：使用正确的拼写，标点和句子结构，折叠长行等。

在包中，顶级声明之前的任何注释都将用作该声明的*doc 注释*。程序中的每个导出（大写）名称都应带有文档注释。

Doc 注释最好作为完整的句子使用，从而可以进行各种各样的自动演示。第一句应该是单句摘要，以声明的名称开头。

```
//编译会解析一个正则表达式，如果成功，则返回
//可用于与文本匹配的Regexp。
func Compile（str字符串）（* Regexp，错误）{
```

如果每个文档注释都以其描述的项目名称开头，则可以使用[go](https://golang.google.cn/cmd/go/)工具的[doc](https://golang.google.cn/cmd/go/#hdr-Show_documentation_for_package_or_symbol) 子命令，并通过运行输出。想象一下，您忘记了名称“ Compile”，但在寻找正则表达式的解析函数，因此您运行了该命令， `grep`

```
$ go doc -all regexp | grep -i parse
```

如果包中的所有文档注释均以“此功能...”开头，则`grep` 不会帮助您记住该名称。但是，由于该软件包使用名称开头每个文档注释，因此您会看到类似这样的内容，该名称会回忆起您要查找的单词。

```
$ go doc -all regexp | grep -i parse
    编译将解析正则表达式，如果成功，则返回一个Regexp
    MustCompile类似于Compile，但如果无法解析该表达式，则会发生恐慌。
    解析。它简化了全局变量保存的安全初始化
$
```

Go 的声明语法允许对声明进行分组。单个文档注释可以引入一组相关的常量或变量。由于整个声明都已提出，因此这样的评论常常是敷衍了事。

```go
//未能解析表达式返回的错误代码。
var（
    ErrInternal = errors.New（“ regexp：内部错误”）
    ErrUnmatchedLpar = errors.New（“ regexp：不匹配的'（'”）
    ErrUnmatchedRpar = errors.New（“ regexp：不匹配'）'”）
    ...
）
```

分组还可以指示项目之间的关系，例如一组变量受互斥锁保护的事实。

```go
var（
    countLock sync.Mutex
    inputCount uint32
    outputCount uint32
    errorCount uint32
）
```

## 名字

名称在 Go 语言中与其他语言一样重要。它们甚至具有语义效果：包外部名称的可见性取决于其首字符是否为大写。因此，值得花一些时间讨论 Go 程序中的命名约定。

### 包名称

导入软件包时，软件包名称将成为其内容的访问器。

```go
import "bytes"
```

导入包后可以使用`bytes.Buffer`。如果每个使用该软件包的人都可以使用相同的名称来引用其内容，这将很有帮助，这意味着该软件包的名称应该很好：简短，简洁，令人回味。按照惯例，软件包使用小写的单字名称。不需要下划线或首字母大写。为了简便起见，Err 是错误的，因为每个使用您的软件包的人都会输入该名称。而且不用担心*先验*碰撞。包名称仅是导入的默认名称。它不必在所有源代码中都是唯一的，并且在发生冲突的极少数情况下，导入包可以选择其他名称以在本地使用。在任何情况下，混淆都是很少的，因为导入中的文件名决定了所使用的软件包。

另一个约定是，程序包名称是其源目录的基本名称。`src/encoding/base64` 导入的包名称为`"encoding/base64"`，但名称为`base64`，而不是`encoding_base64`或者 `encodingBase64`。

程序包的导入者将使用该名称来引用其内容，因此，程序包中导出的名称可以使用该事实来避免卡顿。（不要使用这种`import .`表示法，它可以简化必须在所测试的程序包之外运行的测试，但应避免这样做。）例如，`bufio`程序包中的缓冲读取器类型称为`Reader`，而不是`BufReader`，因为用户将其视为`bufio.Reader`，这是一个简洁明了的名称。此外，由于导入的实体始终使用其包名称来寻址，因此`bufio.Reader` 不会与冲突`io.Reader`。同样，通常会调用来创建新实例的函数（`ring.Ring`这是 Go 中*构造函数*的定义）`NewRing`，但是由于 `Ring`是程序包导出的唯一类型，由于调用了程序包`ring`，因此将其称为 just `New`，程序包的客户端将其视为`ring.New`。使用包结构可以帮助您选择好名字。

另一个简短的例子是`once.Do`； `once.Do(setup)方便了阅读，写成`once.DoOrWaitUntilDone(setup)`并不会更好。长名不会自动使代码更具可读性。有用的文档注释通常比加长名称更有价值。

### Getters

Go 不会自动为 getter 和 setter 提供支持。自己提供 getter 和 setter 并没有错，这样做通常是适当的，但是`Get`使用 getter 的名字既不是惯用的，也没有必要。如果您有一个名为`owner`（小写，未导出）的字段 ，则应调用 getter 方法`Owner`（大写，已导出），而不是`GetOwner`。使用大写名称进行导出提供了挂钩，以将字段与方法区分开。如果需要，可以使用 setter 函数`SetOwner`。这两个名字在实践中都读得很好：

```go
owner := obj.Owner()
if owner != user {
    obj.SetOwner(user)
}
```

### 接口名称

按照惯例，一个方法接口由该方法 name 加上后缀-er 或类似的修改命名构建的试剂名：`Reader`， `Writer`，`Formatter`， `CloseNotifier`等。

有许多这样的名称，兑现它们和它们捕获的函数名称很有用。 `Read`，`Write`，`Close`，`Flush`， `String`等有规范签名和意义。为避免混淆，除非您的方法具有相同的签名和含义，否则请不要给它们使用任何名称。相反，如果您的类型实现的方法的含义与熟知类型上的方法的含义相同，则为其赋予相同的名称和签名；调用您的 string-converter 方法`String`not `ToString`。

### 驼峰命名

最后，Go 中的约定是使用 `MixedCaps` 或`mixedCaps`而不使用下划线来编写多字名称。

## 分号

与 C 一样，Go 的形式语法使用分号来终止语句，但是与 C 中不同，这些分号不会出现在源代码中。相反，词法分析器使用一条简单规则在扫描时自动插入分号，因此输入文本几乎没有分号。

规则是这样的。如果换行符之前的最后一个标记是标识符（包括诸如`int`和的词`float64`），基本文字（例如数字或字符串常量）或标记之一

```
break continue fallthrough return ++ -- ) }
```

词法分析器总是在标记后插入分号。可以概括为：“如果换行符位于可以结束语句的标记之后，请插入分号”。

也可以在右括号之前省略分号，因此可以使用如下语句：

```
go func() { for { dst <- <-src } }()
```

不需要分号。惯用的 Go 程序仅在诸如`for`循环子句之类的地方使用分号 ，以分隔初始化程序，条件和延续元素。如果您以这种方式编写代码，则在一行上分隔多个语句也是必需的。

分号插入规则的一个后果是，你不能把一个控制结构（中左括号`if`，`for`，`switch`，或`select`）在下一行。如果这样做，将在分号之前插入一个分号，这可能会导致不想要的效果。这样写

```
if i < f() {
    g()
}
```

不像这样

```
if i < f()  // wrong!
{           // wrong!
    g()
}
```

## 控制结构

Go 的控制结构与 C 的控制结构相关，但在重要方面有所不同。没有`do`或没有`while`循环，只有略微概括 `for`； `switch`更灵活； `if`并`switch`接受可选的初始化语句，如`for`； `break`并`continue`声明接受可选的标签，以确定哪些中断或继续; 并且有新的控制结构，包括类型开关和多路通信多路复用器`select`。语法也略有不同：没有括号，并且主体必须始终用大括号分隔。

### If

在 Go 中，一个简单的`if`样子是这样的：

```go
if x > 0 {
    return y
}
```

强制括号鼓励`if`在多行上编写简单的语句。无论如何都是这样做的好风格，尤其是当主体包含诸如 a`return`或的控制语句时 `break`。

由于`if`并`switch`接受初始化语句，通常会看到用来设置局部变量的语句。

```go
if err := file.Chmod(0664); err != nil {
    log.Print(err)
    return err
}
```

在 Go library 中，你会发现，当一个`if`语句不流入下一条语句，也就是说，函数体以`break`，`continue`， `goto`，或`return`结束。不必要的 `else`被省略了。

```go
f, err := os.Open(name)
if err != nil {
    return err
}
codeUsing(f)
```

下面是一种常见情况的示例，在这种情况下，代码必须防范一系列错误情况。如果成功的控制流贯穿页面，代码的可读性会很好，从而消除了出现的错误情况。由于错误情况会以`return` 语句结尾，因此生成的代码不需要`else`语句。

```go
f, err := os.Open(name)
if err != nil {
    return err
}
d, err := f.Stat()
if err != nil {
    f.Close()
    return err
}
codeUsing(f, d)
```

### 重新声明和重新分配

旁白：上一节中的最后一个示例演示了`:=`简短声明表单如何工作的详细信息 。调用的声明`os.Open`为：

```
f，err：= os.Open（name）
```

该语句声明了两个变量`f`和`err`。几行后，对`f.Stat`read 的调用

```
d，err：= f.Stat（）
```

看起来好像在声明`d`和`err`。但是请注意，`err`这两个语句中都会出现。这种重复是合法的：`err`由第一个语句声明，但在第二个语句中仅仅是被重新分配了值。这意味着对`f.Stat`的调用将使用`err`上面声明的已有变量，并为其赋予一个新值。

在`:=`声明中，`v`即使已经声明了变量，也可能会出现该变量，条件是：

- 此声明与的现有声明在同一范围内`v` （如果`v`已经在外部范围中声明，则该声明将创建一个新变量§），
- 初始化中的对应值可分配给`v`
- 声明创建了至少一个其他变量。

这种不寻常的特性是纯粹的实用主义，`err`例如在长`if-else`链中易于使用单个值。您会看到它经常使用。

§这里值得一提的是，**在 Go 中，函数参数和返回值的作用域与函数主体相同**，即使它们在词法上出现在包围主体的括号之外。

### For

Go`for`循环类似于 C，但不相同。它统一了`for` ，`while`没有`do-while`。共有三种形式，其中只有一种具有分号。

```go
// Like a C for
for init; condition; post { }

// Like a C while
for condition { }

// Like a C for(;;)
for { }
```

简短的声明使在循环中可以轻松声明索引变量。

```go
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}
```

如果要遍历**数组**，切片，**字符串**或**映射**，或者从**通道**读取，则`range`子句可以管理该循环。

```go
for key, value := range oldMap {
    newMap[key] = value
}
```

如果只需要范围内的第一项（键或索引），请=直接去掉第二项：

```go
for key := range m {
    if key.expired() {
        delete(m, key)
    }
}
```

如果只需要范围（值）中的第二项，请使用*空白标识符*（下划线）来丢弃第一项：

```go
sum := 0
for _, value := range array {
    sum += value
}
```

如[后面的部分](https://golang.google.cn/doc/effective_go.html#blank)所述，空白标识符有许多用途。

对于字符串，`range` 它可以为您做更多的工作，通过解析 UTF-8 来分解单个 Unicode 代码点。错误的编码会占用一个字节并产生替换符文 U + FFFD。（名称（具有关联的内置类型）**`rune`**是单个 Unicode 的 Go 术语。有关 rune 的详细信息，请参见[语言规范](https://golang.google.cn/ref/spec#Rune_literals)。）

```go
for pos, char := range "日本\x80語" { // \x80 is an illegal UTF-8 encoding
    fmt.Printf("character %#U starts at byte position %d\n", char, pos)
}
```

Prints

```go
character U+65E5 '日' starts at byte position 0
character U+672C '本' starts at byte position 3
character U+FFFD '�' starts at byte position 6
character U+8A9E '語' starts at byte position 7
```

最后，Go 没有逗号运算符，`++`并且`--` 语句不是表达式。因此，如果您要在中运行多个变量，`for` 则应使用并行赋值。(i++`和`i--`在`Go`语言中是语句，不是表达式，因此不能赋值给另外的变量。此外没有`++i`和`--i)

```go
//反转a
对于i，j：= 0，len（a）-1; 我<j; i，j = i + 1，j-1 {
    a [i]，a [j] = a [j]，a [i]
}
```

### Switch

Go`switch`比 C 更通用。表达式不必是常数，甚至不必是整数，大小写从上到下进行评估，直到找到匹配项为止；如果`switch`没有表达式，则将其打开 `true`。它因此可能和习惯，写的 `if`- `else`- `if`-`else` 链作为`switch`。

```go
func unhex(c byte) byte {
    switch {
    case '0' <= c && c <= '9':
        return c - '0'
    case 'a' <= c && c <= 'f':
        return c - 'a' + 10
    case 'A' <= c && c <= 'F':
        return c - 'A' + 10
    }
    return 0
}
```

不会自动掉线，但不同的 case 可以用逗号分隔的列表显示。

```go
func shouldEscape(c byte) bool {
    switch c {
    case ' ', '?', '&', '=', '#', '+', '%':
        return true
    }
    return false
}
```

尽管它们在 Go 中不像其他一些类似 C 的语言那样普遍，但是可以使用`break`语句来提早终止`switch`。但是，有时需要跳出周围的循环而不是 Switch，在 Go 中，可以通过在循环上放置标签然后 break 到该标签来实现这一点。此示例显示了两种用法。

```go
Loop:
	for n := 0; n < len(src); n += size {
		switch {
		case src[n] < sizeOne:
			if validateOnly {
				break
			}
			size = 1
			update(src[n])

		case src[n] < sizeTwo:
			if n+1 >= len(src) {
				err = errShortInput
				break Loop
			}
			if validateOnly {
				break
			}
			size = 2
			update(src[n] + src[n+1]<<shift)
		}
	}
```

当然，该`continue`语句还接受可选标签，但仅适用于循环。

要结束本节，这是一个使用两个`switch`语句的字节片比较例程 ：

```go
// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
    for i := 0; i < len(a) && i < len(b); i++ {
        switch {
        case a[i] > b[i]:
            return 1
        case a[i] < b[i]:
            return -1
        }
    }
    switch {
    case len(a) > len(b):
        return 1
    case len(a) < len(b):
        return -1
    }
    return 0
}
```

### Type Switch

开关也可以用来发现接口变量的动态类型。这种*类型开关*使用括号内带有`type`关键字的类型声明的语法。如果开关在表达式中声明了变量，则该变量在每个子句中将具有相应的类型。在这种情况下重用名称也是符合习惯的，实际上是在每种情况下声明一个具有相同名称但类型不同的新变量。

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
    fmt.Printf("unexpected type %T\n", t)     // %T prints whatever type t has
case bool:
    fmt.Printf("boolean %t\n", t)             // t has type bool
case int:
    fmt.Printf("integer %d\n", t)             // t has type int
case *bool:
    fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
case *int:
    fmt.Printf("pointer to integer %d\n", *t) // t has type *int
}

```

## Functions

### Multiple return values

Go 的不寻常功能之一是函数和方法可以返回多个值。这种形式可以用来改进 C 程序中的一些笨拙的习惯用法：带内错误返回，例如`-1`for`EOF` 和修改由地址传递的参数。

在 C 语言中，写错误是通过统计负数来表示的，错误代码会在易失性位置中被隐藏掉。在 Go 中，`Write` 可以返回一个计数*和*一个错误：“Yes, you wrote some bytes but not all of them because you filled the device”。`Write`软件包中文件的方法签名`os`为：

```go
func (file *File) Write(b []byte) (n int, err error)
```

就像文件说的，它返回写入的字节数和一个非空`error`，当`n` `!=` `len(b)`的时候。这是一种常见的样式。有关更多示例，请参见错误处理部分。

类似的方法避免了将指针传递给返回值以模拟引用参数的需要。这是一个简单的函数，可从字节片中的某个位置获取一个数字，然后返回该数字和下一个位置。

```go
func nextInt(b []byte, i int) (int, int) {
    for ; i < len(b) && !isDigit(b[i]); i++ {
    }
    x := 0
    for ; i < len(b) && isDigit(b[i]); i++ {
        x = x*10 + int(b[i]) - '0'
    }
    return x, i
}
```

您可以使用它来扫描输入切片 b`中的数字，如下所示：

```go
  for i := 0; i < len(b); {
        x, i = nextInt(b, i)
        fmt.Println(x)
    }
```

### Named result parameters

可以给 Go 函数的返回“参数”命名，并将其用作常规变量，就像传入的参数一样去使用它。命名后，函数开始时会将它们 **初始化为零值** 。如果函数执行不带参数的`return`语句，则将返回参数的当前值用作返回值。

名称不是强制性的，但它们可以使代码更短，更清晰：它们是文档。如果我们命名，`nextInt`则返回的结果显而易见`int` 。

```
func nextInt（b [] byte，pos int）（value，nextPos int）{
```

由于命名返回值已初始化并绑定到 return，因此这种写法简单明了。这是`io.ReadFull 很好的使用命名返回值的例子：

```
func ReadFull(r Reader, buf []byte) (n int, err error) {
    for len(buf) > 0 && err == nil {
        var nr int
        nr, err = r.Read(buf)
        n += nr
        buf = buf[nr:]
    }
    return
}
```

### Defer

Go 的`defer`语句将函数调用（ *延迟*函数）计划为在执行`defer`返回的函数之前立即运行。这是**处理异常情况**的一种不寻常但有效的方法，比如无论函数通过哪条路径返回都必须释放资源的情况。典型的例子是解锁互斥锁或关闭文件。

```go
// Contents returns the file's contents as a string.
func Contents(filename string) (string, error) {
    f, err := os.Open(filename)
    if err != nil {
        return "", err
    }
    defer f.Close()  // f.Close will run when we're finished.

    var result []byte
    buf := make([]byte, 100)
    for {
        n, err := f.Read(buf[0:])
        result = append(result, buf[0:n]...) // append is discussed later.
        if err != nil {
            if err == io.EOF {
                break
            }
            return "", err  // f will be closed if we return here.
        }
    }
    return string(result), nil // f will be closed if we return here.
}

```

推迟调用诸如之类的函数 `Close` 有两个优点。首先，它保证您永远不会忘记关闭文件，如果以后编辑函数以添加新的返回路径，则很容易犯此错误。其次，由于 defer 关闭代码位于打开 os.Open 附近，这比将其放置在函数的末尾要清晰得多。

延迟函数的参数（如果函数是方法，则包括接收方）在*延迟* 执行时（而不是在*调用*执行时）进行评估。除了避免担心变量在函数执行时会更改值之外，这还意味着单个延迟的调用站点可以延迟多个函数的执行。这是一个愚蠢的例子。

```go
for i := 0; i < 5; i++ {
    defer fmt.Printf("%d ", i)
}
```

延迟的功能按 LIFO 顺序执行，因此该代码将 `4 3 2 1 0` 在函数返回时被打印。一个更合理的示例是通过程序跟踪函数执行的简单方法。我们可以编写一些简单的跟踪例程，如下所示：

```go
func trace(s string)   { fmt.Println("entering:", s) }
func untrace(s string) { fmt.Println("leaving:", s) }

// Use them like this:
func a() {
    trace("a")
    defer untrace("a")
    // do something....
}
```

通过利用以下事实，我们可以做得更好：在`defer`执行时评估延迟函数的参数。跟踪例程可以将参数设置为取消跟踪例程。这个例子：

```go
func trace(s string) string {
    fmt.Println("entering:", s)
    return s
}

func un(s string) {
    fmt.Println("leaving:", s)
}

//defer un()内部的参数将在执行defer语句的时候立即评估
func a() {
    defer un(trace("a"))
    fmt.Println("in a")
}

func b() {
    defer un(trace("b"))
    fmt.Println("in b")
    a()
}

func main() {
    b()
}
```

打印结果:

```go
entering: b
in b
entering: a
in a
leaving: a
leaving: b
```

对于习惯于使用其他语言的块级资源管理的程序员来说，这 `defer` 似乎很奇怪，但是它最有趣，功能最强大的应用恰恰是因为它不是基于程序块的而是基于函数的。在上的部分中 `panic`，`recover`我们将看到其可能性的另一个示例。

## Data

### 用 `new` 来分配

Go 有两个分配原语，内置函数 `new`和`make`。它们执行不同的操作，并应用于不同的类型，这可能会造成混淆，但是规则很简单。让我们`new`先谈谈。这是一个**分配内存**的内置函数，但与其他语言中的同名函数不同，它不会*初始化*内存，**只会将其*清零***。也就是说， `new(T)`为类型为 `T`的新单元**分配零存储**并返回其地址，值为 value `*T`。在 Go 术语中，它返回一个指向新分配的 type `T`的零值指针。

由于返回的内存`new`为零，因此在设计数据结构时安排使用每种类型的零值而无需进一步初始化将很有帮助。这意味着数据结构的用户可以创建数据结构`new`并开始使用。例如，的文档`bytes.Buffer`指出“零值`Buffer`是准备使用的空缓冲区”。同样，`sync.Mutex`没有显式的构造函数或`Init`方法。而是将 a 的零值`sync.Mutex` 定义为未锁定的互斥锁。

零值即有用属性会暂时起作用。考虑此类型声明。

```go
type SyncedBuffer struct {
    lock    sync.Mutex
    buffer  bytes.Buffer
}
```

type 的值`SyncedBuffer`也可以在分配或声明后立即使用。在下一个代码段中，`p`和`v`都可以正常工作，而无需进一步安排。

```
p：= new（SyncedBuffer）//输入* SyncedBuffer
var v SyncedBuffer //类型SyncedBuffer
```

### 构造函数和复合表达式

有时初始的零值不够好，因此需要初始化构造函数，如本例中从 package 派生的`os`那样。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := new(File)
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0
    return f
}
```

上面有很多样板。我们可以使用*复合表达式*来简化它，每次对其求值时都会创建一个新实例。

```go
func NewFile(fd int, name string) *File {
    if fd < 0 {
        return nil
    }
    f := File{fd, name, nil, 0}
    return &f
}
```

没错，复合表达式(composite literals)是指:

在声明变量的时候赋值

```go
// Short syntax
myArray := [5]int{3, 3, 3, 3, 3}
// OR
// Long syntax
var myArray [5]int = [5]int{3, 3, 3, 3, 3}
```

注意，与 C 语言不同，完全可以返回局部变量的地址。函数返回后，与变量关联的存储将保留。实际上，采用复合表达式的地址会在每次对其求值时分配一个新实例，因此我们可以将后两行结合在一起。

```
 return &File{fd, name, nil, 0}
```

复合表达式的字段按顺序排列，并且必须全部存在。但是，通过将元素明确标记为*字段*`:`_值_ 对，初始化器可以按任何顺序出现，而缺失的则保留为各自的零值。因此我们可以说

```
   return &File{fd: fd, name: name}
```

在一个有限的情况下，如果一个复合表达式完全不包含任何字段，它将为该类型创建一个零值。表达式`new(File)`和`&File{}`是等效的。

也可以为**数组**，**切片**和**映射**创建复合表达式，其中字段标签为索引或映射键。在这些例子中，无论值是`Enone`， `Eio`和`Einval`，只要它们是不同的就行。

```go
a := [...]string   {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
s := []string      {Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
```

### Allocation with `make`

回到分配。内置函数`make(T, `_args_`)`的用途不同于`new(T)`。它仅创建**切片**，**映射**和**通道**，并返回**类型 T**（**不是\*T**）的*初始化* （**不是零值**）值。区别的原因是，这三种类型用到了在使用之前必须初始化的数据结构的引用。例如，切片是一个三项描述符，其中包含指向数据（处于数组内部），长度和容量的指针，在初始化这三项之前，切片是 nil。对于切片，映射和通道， make 初始化内部数据结构,并准备要使用的值。例如:

```go
make（[] int，10，100）
```

分配一个空间为 100 的 int 数组，然后创建一个长度为 10 且容量为 100（数组的容量)的切片结构指向该数组的前 10 个元素。（创建切片时，可以省略容量；有关更多信息，请参见切片部分。）相反，`new([]int)`返回指向新分配的，零切片结构的指针，即指向`nil`切片值的指针。

这些示例说明了`new`和 之间的区别`make`。

```go
var p *[]int = new([]int)       // allocates slice structure; *p == nil; 几乎不会用到
var v  []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

// Unnecessarily complex:
var p *[]int = new([]int)
*p = make([]int, 100, 100)

// 复合习惯的:
v := make([]int, 100)
```

请记住，这 `make` 仅适用于**映射**，**切片**和**通道**，不返回指针。要获得显式指针可以使用 **new** ，或者显式地获取变量的地址。

### 数组

数组在计划内存的详细布局时很有用，有时可以帮助避免分配，但是数组主要是切片的根基，切片是下一节的主题。为奠定该主题的基础，以下是有关数组的几句话。

在 Go 和 C 中，数组的工作方式之间存在主要差异。在 Go 中，

- 数组是值的集合。将一个数组分配给另一个数组将复制所有元素。
- **特别是，如果将数组传递给函数，它将接收该数组的*副本*，而不是指向它的指针**。
- **数组的大小是其类型的一部分。类型`[10]int` 和`[20]int`是不同的**。

数组是值这个属性既有用又昂贵。如果您想要类 C 的行为和效率，可以将指针传递给数组。

```go
func Sum(a *[3]float64) (sum float64) {
    for _, v := range *a {
        sum += v
    }
    return
}

array := [...]float64{7.0, 8.5, 9.1}
x := Sum(&array)  // Note the explicit address-of operator
```

但是，即使这种方式也不是 Go 所常用的。请改用切片。

### 切片

切片包装数组可为数据序列提供更通用，更强大和更方便的接口。除了具有明确维数的元素（例如转换矩阵）外，Go 中的大多数数组编程都是使用切片而不是简单数组完成的。

切片包含对基础数组的引用，如果将一个切片分配给另一个切片，则两个切片均引用同一数组。如果函数采用 slice 参数，则对 slice 的元素所做的更改将对调用者可见，这类似于将指针传递给基础数组。因此`Read` 函数可以接受一个切片参数，而不是一个指针和一个计数; 切片内的长度明确了要读取的数据上限。这是 os 包中 File 类型 `Read`方法的签名 ：

```go
func (f *File) Read(buf []byte) (n int, err error)
```

该方法返回读取的字节数和错误值（如果有）。读入所述第一 32 个字节的较大的缓冲区的 `buf`，切(slice)一下缓冲 buf。

```go
n, err := f.Read(buf[0:32])
```

这种切片是普通且有效的。实际上，如果不考虑效率，以下代码段也能读取缓冲区的前 32 个字节。

```go
   var n int
    var err error
    for i := 0; i < 32; i++ {
        nbytes, e := f.Read(buf[i:i+1])  // Read one byte.
        n += nbytes
        if nbytes == 0 || e != nil {
            err = e
            break
        }
    }
```

切片的长度可以更改，只要它仍然满足基础数组的长度限制；只需将其分配给自身的一部分即可。切片的*容量*（可通过内置函数 cap`获取）cap可以报告切片可以使用的最大长度。下面是一个将数据追加到切片的功能。如果数据超出容量，则会重新分配片。返回结果切片。该函数使用`len`和`cap`应用于`nil`切片时也可以，并返回 0 这一点。

```go
func Append(slice, data []byte) []byte {
    l := len(slice)
    if l + len(data) > cap(slice) {  // reallocate
        // Allocate double what's needed, for future growth.
        newSlice := make([]byte, (l+len(data))*2)
        // The copy function is predeclared and works for any slice type.
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0:l+len(data)]
    copy(slice[l:], data)
    return slice
}
```

之后必须返回分片，因为尽管`Append` 可以修改的元素`slice`，但分片本身（保存指针，长度和容量的运行时数据结构）是按值传递的。

添加到切片的想法非常有用，它被`append`内置函数捕获 。但是，要了解该功能的设计，我们需要更多信息，因此我们将在稍后再讨论。

### 二维切片

Go 的数组和切片是一维的。要创建等效于 2D 数组或切片的数组，必须定义一个数组的数组或切片的切片，如下所示：

```go
type Transform [3] [3] float64 //一个3x3数组，实际上是一个数组的数组。
type LinesOfText [] [] byte //字节切片的一部分。
```

由于切片的长度是可变的，因此可能使每个内部切片的长度不同。这可能是常见的情况，例如在我们的`LinesOfText` 示例中：每行都有独立的长度。

```go
text := LinesOfText{
	[]byte("Now is the time"),
	[]byte("for all good gophers"),
	[]byte("to bring some fun to the party."),
}
```

有时有必要分配 2D 切片，例如，在处理像素的扫描线时可能会出现这种情况。有两种方法可以实现此目的。一种是独立分配每个分片；另一种是分配单个数组，并将单个切片指向该数组。使用哪种取决于您的应用程序。如果切片可能增大或缩小，则应独立分配它们，以免覆盖下一行；如果不是，则使用单一分配构造对象可能会更有效。作为参考，以下是这两种方法的示意图。首先，一次一行：

```go
//分配顶级切片。
picture：= make（[] [] uint8，YSize）// y的每单位一行。
//循环遍历行，为每行分配切片。
for i：=range picture{
	picture[i] = make（[]uint8，XSize）
}
```

现在作为一种分配，分成几行：

```go
// Allocate the top-level slice, the same as before.
picture := make([][]uint8, YSize) // One row per unit of y.
// Allocate one large slice to hold all the pixels.
pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
// Loop over the rows, slicing each row from the front of the remaining pixels slice.
for i := range picture {
	picture[i], pixels = pixels[:XSize], pixels[XSize:]
}
```

### Maps

映射是一种方便且功能强大的内置数据结构，该结构将一种类型的值（_键_）与另一种类型的值（*元素*或*值*）相关联。键可以是定义了相等运算符的任何类型，例如整数，浮点数和复数，字符串，指针，接口（只要动态类型支持相等），结构和数组。切片不能用作映射键，因为未在其上定义相等性。像切片一样，映射保留对基础数据结构的引用。如果将地图传递给更改地图内容的函数，则更改将在调用方中可见。

可以使用带有冒号分隔的键/值对的常规复合文字语法来构建映射，因此在初始化过程中轻松构建它们。

```go
var timeZone = map[string]int{
    "UTC":  0*60*60,
    "EST": -5*60*60,
    "CST": -6*60*60,
    "MST": -7*60*60,
    "PST": -8*60*60,
}
```

语法上分配和获取映射值的方式类似于对数组和切片执行相同的操作，只是索引不必为整数。

```
offset：= timeZone ["EST"]
```

尝试使用映射中不存在的键来获取映射值时，将为映射中的条目类型返回零值。例如，如果映射包含整数，则查找不存在的键将返回`0`。集合可以实现为具有值类型的映射`bool`。将映射项设置`true`为将值放入集合中，然后通过简单的索引对其进行测试。

```go
attended := map[string]bool{
    "Ann": true,
    "Joe": true,
    ...
}

if attended[person] { // will be false if person is not in the map
    fmt.Println(person, "was at the meeting")
}
```

有时您需要从零值中区分出缺失的条目。是否有条目`"UTC"` 或为 0，因为它根本不在映射中？您可以采用多种分配形式进行区分。

```
var seconds int
var ok bool
seconds, ok = timeZone[tz]
```

由于明显的原因，这被称为“comma ok”的表达习惯。在此示例中，如果`tz`存在，`seconds` 将进行适当设置并`ok`为 true；如果不是， `seconds`则将其设置为零，并且`ok`为 false。这是一个将其与良好的错误报告结合在一起的函数：

```go
func offset(tz string) int {
    if seconds, ok := timeZone[tz]; ok {
        return seconds
    }
    log.Println("unknown time zone:", tz)
    return 0
}
```

要在映射中检查 key 是否存在而又不需要获取到 value 的值，就使用空白符号\_

```
_, present := timeZone[tz]
```

要删除 map entry，请使用`delete` 内置函数，其内置参数是 map 和要删除的键。即使 map 上已经没有键，也可以这样做。

```
delete(timeZone, "PDT")  // Now on Standard Time
```

### Printing

Go 中的格式化打印使用类似于 C`printf` 家族的样式，但功能更丰富，更通用。该函数住在`fmt` 包装和有大写的名字：`fmt.Printf`，`fmt.Fprintf`， `fmt.Sprintf`等。字符串函数（`Sprintf`等）返回字符串，而不是填充提供的缓冲区。

您不需要提供格式字符串。对于每一个`Printf`， `Fprintf`和`Sprintf`有另一种双功能，如`Print`和`Println`。这些函数不采用格式字符串，而是为每个参数生成默认格式。这些`Println`版本还会在参数之间插入一个空格，并在输出中添加一个换行符，而`Print`仅当双方的操作数都不是字符串时，这些版本才会添加空格。在此示例中，每行产生相同的输出。

```go
fmt.Printf("Hello %d\n", 23)
fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
fmt.Println("Hello", 23)
fmt.Println(fmt.Sprint("Hello ", 23))
```

格式化的打印功能`fmt.Fprint` 和它的朋友们将实现该`io.Writer`接口的任何对象作为第一个参数。变量`os.Stdout` 和`os.Stderr`熟悉的实例。

在这里，事情开始与 C 背道而驰。首先，诸如这样的数字格式不带有标志性或大小标志。相反，打印例程使用参数的类型来决定这些属性。

```go
var x uint64 = 1<<64 - 1
fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))
```

Prints

```go
18446744073709551615 ffffffffffffffff; -1 -1
```

如果只需要默认转换（例如，十进制表示整数），则可以使用包罗万象的格式`%v`（表示“值”）；结果跟使用`Print`，和`Println`一样。此外，该格式可以打印*任何*值，甚至可以打印数组，切片，结构和映射。这是上一节中定义的时区映射的打印语句。

```go
fmt.Printf("%v\n", timeZone)  // or just fmt.Println(timeZone)
```

输出：

```go
map[CST：-21600 EST：-18000 MST：-25200 PST：-28800 UTC：0]
```

对于映射，`Printf`和它的朋友们按字母顺序对输出进行字典排序。

打印结构时，修改后的格式**`%+v`**会打印出结构名: 值的格式，对于任何值，可选格式**`%#v`**都会以完整的 Go 语法打印该值。

```go
type T struct {
    a int
    b float64
    c string
}
t := &T{ 7, -2.35, "abc\tdef" }
fmt.Printf("%v\n", t)
fmt.Printf("%+v\n", t)
fmt.Printf("%#v\n", t)
fmt.Printf("%#v\n", timeZone)
```

Prints

```go
&{7 -2.35 abc   def}
&{a:7 b:-2.35 c:abc     def}
&main.T{a:7, b:-2.35, c:"abc\tdef"}
map[string]int{"CST":-21600, "EST":-18000, "MST":-25200, "PST":-28800, "UTC":0}
```

（请注意，“＆”号。）将引号字符串格式`%q`应用于类型为`string`或时也可以使用`[]byte`。`%#q`如果可能，替代格式将使用反引号代替。（该`%q`格式也适用于整数和符文，生成单引号符文常量。）此外，该方法`%x`适用于字符串，字节数组和字节片以及整数，生成长十六进制字符串，且格式为空格（`% x`），在字节之间放置空格。

另一种方便的格式是`%T`，它打印值的*类型*。

```go
fmt.Printf("%T\n", timeZone)
```

打印

```
map [string] int
```

如果要控制自定义类型的默认格式，只需给该类型定义一个`String()方法。对于我们的简单类型`T`，可能看起来像这样。

```
func (t *T) String() string {
    return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}
fmt.Printf("%v\n", t)
```

以以下格式打印

```
7 / -2.35 /“ abc \ tdef”
```

（如果您需要打印类型*值*`T`以及指向的指针`T`，则 for 的接收器`String`必须为值类型；此示例使用了指针，因为这对于结构类型更有效且更惯用。有关[指针与值接收器的联系](https://golang.google.cn/doc/effective_go.html#pointers_vs_values)，请参见下文更多信息。）

我们的`String`方法之所以能够调用，`Sprintf`是因为打印 routine 是完全可重入的，并且可以通过这种方式包装。但是，关于此方法，有一个重要的细节要理解：不要在构造 String 方法的时候直接调用 Sprint 方法并且尝试直接打印。如果`Sprintf` 调用尝试将接收方直接打印为字符串，则可能会发生这种情况，因为直接打印变量又会再次调用 String( 。如本例所示，这是一个常见且容易犯的错误。

```go
type MyString string

func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
}
```

它也很容易修复：将参数转换为基本字符串类型，打印时不调用 String()方法。

```go
type MyString string
func (m MyString) String() string {
    return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}
```

在[初始化部分，](https://golang.google.cn/doc/effective_go.html#initialization)我们将看到另一种避免这种递归的技术。

另一种打印技术是将打印 routine 的参数直接传递给另一个此类 routine。`Printf`的方法签名使用类型`...interface{}` 作为其最终参数，以指定可以在格式之后显示任意数量的参数（任意类型）。

```go
func Printf(format string, v ...interface{}) (n int, err error) {
```

在函数内`Printf`，其`v`作用类似于类型的变量， `[]interface{}`但如果将其传递给另一个可变参数函数，则其作用类似于常规参数列表。这是`log.Println`我们上面使用的功能的实现。它直接将其参数传递给 `fmt.Sprintln`实际格式。

```go
// Println以fmt.Println的方式打印到标准记录器。
func Println（v ... interface {}）{
    std.Output（2，fmt.Sprintln（v ...））//输出带有参数（int，string）
}
```

我们在嵌套调用中写完`...`之后`v`，`Sprintln`以告诉编译器将其`v`视为参数列表。否则，它将`v`作为单个切片参数传递 。

打印比我们这里讨论的还要多。有关详细信息，请参见`godoc`软件包的文档`fmt`。

顺便说一句，`...`参数可以是特定类型，例如`...int` 对于选择最小整数列表的 min 函数而言：

```go
func Min(a ...int) int {
    min := int(^uint(0) >> 1)  // largest int
    for _, i := range a {
        if i < min {
            min = i
        }
    }
    return min
}
```

### Append

现在，我们缺少了解释`append`内置功能设计所需的内容。`append` 的签名与上面的自定义`Append`函数不同。示意图如下：

```
func append(slice []T, elements ...T) []T
```

其中*T*是任何给定类型的占位符。实际上，您无法在 Go 中编写`T` 由调用者确定类型的函数。这`append`就是内置的原因：它需要编译器的支持。

`append`所做的工作是将元素附加到切片的末尾并返回结果。需要返回结果，因为与我们手写的一样`Append`，底层数组可能会更改。这个简单的例子

```go
x := []int{1,2,3}
x = append(x, 4, 5, 6)
fmt.Println(x)
```

打印`[1 2 3 4 5 6]`。所以`append`工作有点像`Printf`，收集任意数量的参数。

但是，如果我们想做我们想做的事`Append`并将一个切片附加到另一个切片上怎么办？很简单：使用**`...`**！下面的代码段产生与上面相同的输出。

```go
x：= [] int {1,2,3}
y：= [] int {4,5,6}
x =append（x，y ...）
fmt.Println（x）
```

如果没有**`...`**，上面的代码就不会编译，因为类型是错误的。`y`不是 type `int`。

## 初始化

尽管从表面上看，它与 C 或 C ++中的初始化没有太大区别，但是 Go 中的初始化功能更强大。可以在初始化期间构建复杂的结构，并且正确处理了初始化对象之间（甚至不同包之间）的排序问题。

### 常数

Go 中的常量就是常量。即使在函数中定义为局部变量时，也可以在编译时创建它们，并且只能是 numbers，characters（runes），字符串或布尔值。由于编译时的限制，定义它们的表达式必须是可由编译器评估的常量表达式。例如， `1<<3`是一个常量表达式，而 `math.Sin(math.Pi/4)`不是因为函数调用`math.Sin`需要在运行时发生。

在 Go 中，使用枚举器创建枚举常量`iota` 。由于`iota`可以作为表达式的一部分，并且表达式可以隐式重复，因此可以轻松地构建复杂的值集。

```go
type ByteSize float64

const (
    _           = iota // ignore first value by assigning to blank identifier
    KB ByteSize = 1 << (10 * iota)
    MB
    GB
    TB
    PB
    EB
    ZB
    YB
)
```

将`String`方法附加到任何用户定义的类型的能力使得任意值都可以自动格式化自身以进行打印。尽管您会看到它最常用于结构，但该技术对于标量类型（例如`ByteSize`这样的浮点类型）也很有用

```go
func (b ByteSize) String() string {
    switch {
    case b >= YB:
        return fmt.Sprintf("%.2fYB", b/YB)
    case b >= ZB:
        return fmt.Sprintf("%.2fZB", b/ZB)
    case b >= EB:
        return fmt.Sprintf("%.2fEB", b/EB)
    case b >= PB:
        return fmt.Sprintf("%.2fPB", b/PB)
    case b >= TB:
        return fmt.Sprintf("%.2fTB", b/TB)
    case b >= GB:
        return fmt.Sprintf("%.2fGB", b/GB)
    case b >= MB:
        return fmt.Sprintf("%.2fMB", b/MB)
    case b >= KB:
        return fmt.Sprintf("%.2fKB", b/KB)
    }
    return fmt.Sprintf("%.2fB", b)
}
```

表达式`YB`打印为`1.00YB`，而`ByteSize(1e13)`打印为`9.09TB`。

这里的使用`Sprintf` 来实现`ByteSize`的`String`方法是安全的（避免循环调用）不是因为使用了转换，而是因为它调用`Sprintf`时用了`%f`，这不是一个字符串格式：`Sprintf`需要一个字符串的时候才会调用 String 方法，而`%f` 想要一个浮点值。

### 变量

变量可以像常量一样被初始化，但是初始化器可以是在运行时计算的通用表达式。

```go
var (
    home   = os.Getenv("HOME")
    user   = os.Getenv("USER")
    gopath = os.Getenv("GOPATH")
)
```

### 初始化 init 函数

最后，每个源文件都可以定义自己的 niladic`init`函数来设置所需的任何状态。（实际上，每个文件可以具有多个 `init`函数。）init`在包中的所有变量声明初始化完成后才调用。而且包中声明的变量只会在所有导入的包都已初始化完成后进行初始化。

除了不能表示为声明的初始化外，`init`函数的常见用法是在实际执行开始之前验证或修复程序状态的正确性。

```go
func init() {
    if user == "" {
        log.Fatal("$USER not set")
    }
    if home == "" {
        home = "/home/" + user
    }
    if gopath == "" {
        gopath = home + "/go"
    }
    // gopath may be overridden by --gopath flag on command line.
    flag.StringVar(&gopath, "gopath", gopath, "override default GOPATH")
}
```

## 方法

### 指针 VS 值

如我们所见`ByteSize`，可以为任何命名类型（指针或接口除外）定义方法；接收者不必是结构。

在上面的切片讨论中，我们编写了一个`Append` 函数。我们可以将其定义为切片方法。为此，我们首先声明一个可以绑定该方法的命名类型（named type），然后使该方法的接收者成为该类型的值。

```go
type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
    // Body exactly the same as the Append function defined above.
}
```

这仍然需要方法返回更新的切片。我们可以通过重新定义方法采取消除这种笨拙的做法，让方法使用指向 ByteSlice 的指针作为方法接收者，因此该方法可以直接覆盖调用者的切片。

```go
func (p *ByteSlice) Append(data []byte) {
    slice := *p
    // Body as above, without the return.
    *p = slice
}
```

实际上，我们可以做得更好。如果我们修改函数，使其看起来像是标准`Write`方法，就像这样，

```go
func (p *ByteSlice) Write(data []byte) (n int, err error) {
    slice := *p
    // Again as above.
    *p = slice
    return len(data), nil
}
```

然后该类型`*ByteSlice`满足标准接口 `io.Writer`，这很方便。例如，我们可以打印成一张。

```go
    var b ByteSlice
    fmt.Fprintf（＆b，“这一小时有％d天\ n”，7）
```

我们传递 ByteSlice`的地址因为只有指针`\*ByteSlice`满足`io.Writer`。有关方法接收者使用指针还是值的规则是，可以在指针和值上调用值接收者方法，但是只能在指针上调用指针接收者方法。

之所以出现此规则，是因为指针方法可以修改接收者。在值上调用它们将导致该方法接收该值的副本，因此任何修改都将被丢弃。（在值上调用指针接收者方法并不会修改这个值，而指针接收者方法又需要能改变调用者的值所以产生冲突）

因此，该语言不允许出现此错误。但是，有一个方便的例外。**当值是可寻址的时**，GO 语言将通过自动插入地址运算符来处理在值上调用指针方法的常见情况。在我们的示例中，变量`b`是可寻址的，因此我们可以`Write`使用 just 调用其方法`b.Write`。编译器会将其重写`(&b).Write`为我们。

顺便说一句，在字节切片上使用`Write`的想法对于实现`bytes.Buffer`至关重要。

## 接口及其他类型

### 接口

Go 中的接口提供了一种指定对象行为的方法：如果可以做到*这一点*，则可以在*此处*使用它 。我们已经看过几个简单的例子。定制打印函数可以用一种`String`方法来实现，而`Fprintf`可以用一种`Write`方法生成任何东西的输出。只有一个或两个方法的接口在 Go 代码中很常见，并且通常会使用从该方法派生的名称（例如实现`Write`名为`io.Writer` 的方法）。

一个类型可以实现多个接口。例如，一个集合可以通过在包中的例程进行排序`sort`，如果它实现了 `sort.Interface`，其中包含`Len()`， `Less(i, j int) bool`以及`Swap(i, j int)`，它也可以有一个自定义的格式。在这个人为的例子中，`Sequence`两者都满足。

```go
type Sequence []int

// Methods required by sort.Interface.
func (s Sequence) Len() int {
    return len(s)
}
func (s Sequence) Less(i, j int) bool {
    return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
    copy := make(Sequence, 0, len(s))
    return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String() string {
    s = s.Copy() // Make a copy; don't overwrite argument.
    sort.Sort(s)
    str := "["
    for i, elem := range s { // Loop is O(N²); will fix that in next example.
        if i > 0 {
            str += " "
        }
        str += fmt.Sprint(elem)
    }
    return str + "]"
}
```

### Conversions

`Sequence`的`String`方法是重新创建`Sprint`已经对切片进行的工作。（它的复杂度为 O（N²），这很差。）如果在调用之前将`Sequence`转换为`[]int 格式，我们可以分担任务（并加快速度）。

```go
func (s Sequence) String() string {
    s = s.Copy()
    sort.Sort(s)
    return fmt.Sprint([]int(s))
}
```

此方法是用于`Sprintf`从`String`方法安全调用的转换技术的另一个示例 。因为如果忽略类型名称，这两个类型（`Sequence`和`[]int`）是相同的，因此在它们之间进行转换是合法的。转换不会创建新值，而只是暂时地充当现有值具有新类型的行为。（还有其他一些合法的转换，例如从整数到浮点的转换，它们确实创建了一个新值。）

在 Go 程序中，习惯用法是转换表达式的类型以访问不同的方法集。例如，我们可以使用现有类型`sort.IntSlice`将整个示例简化为：

```go
type Sequence []int

// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
    s = s.Copy()
    sort.IntSlice(s).Sort()
    return fmt.Sprint([]int(s))
}
```

现在，而不是`Sequence`实现多个接口（排序和打印），我们使用一个数据项的转换为多种类型的能力（`Sequence`，`sort.IntSlice` 和`[]int`），每个做这项工作的某些部分。在实践中，这种情况较不常见，但可以有效。

### 接口转换和类型断言

[类型开关](https://golang.google.cn/doc/effective_go.html#type_switch)是一种转换形式：它们采用一个接口，并且对于开关中的每种情况，在某种意义上都将其转换为该情况的类型。这是下面的代码展示了`fmt.Printf`如何使用类型开关将值转换为字符串的简化版本。如果已经是字符串，则我们希望接口保留实际的字符串值，而如果它具有 `String`方法，则需要调用该方法后的结果。

```go
type Stringer interface {
    String() string
}

var value interface{} // Value provided by caller.
switch str := value.(type) {
case string:
    return str
case Stringer:
    return str.String()
}
```

第一种情况找到了具体的价值。第二个将接口转换为另一个接口。这样混合类型就很好了。

如果我们只关心一种类型该怎么办？如果我们知道该值包含一个`string` 而我们只想提取它？一个单例类型开关可以，但*类型断言也可以*。类型断言采用接口值并从中提取指定的显式类型的值。该语法是从打开类型开关的子句中借用的，但具有显式类型而不是`type`关键字：

```go
value.(typeName)
```

结果是具有静态类型 typeName 的新值``。该类型必须是保存在接口的具体类型，或者是值可以转换为的第二种接口类型。为了提取我们知道在值中的字符串，我们可以这样写：

```go
str := value.(string)
```

但是，如果事实证明该值不包含字符串，则程序将因运行时错误而崩溃。为了防止这种情况，请使用“逗号，好”惯用法来安全地测试该值是否为字符串：

```go
str, ok := value.(string)
if ok {
    fmt.Printf("string value is: %q\n", str)
} else {
    fmt.Printf("value is not a string\n")
}
```

如果类型断言失败，`str`则该类型断言将仍然存在并且为字符串类型，但是它将具有零值（一个空字符串）。

为了说明该功能，这里有一个`if`-`else` 语句，该语句等效于打开此部分的类型开关。

```go
if str, ok := value.(string); ok {
    return str
} else if str, ok := value.(Stringer); ok {
    return str.String()
}
```

### 概论

如果类型仅存在于实现接口，并且永远不会有超出该接口的导出方法，则无需导出类型本身。仅导出接口即可清楚地知道该值除了接口中描述的内容外没有其他有趣的行为。它还避免了需要在通用方法的每个实例上重复文档。

在这种情况下，构造函数应返回接口值而不是实现类型。作为一个例子，在散列库中 crc32.NewIEEE`和`adler32.New` 返回接口类型`hash.Hash32`。在 Go 程序中将 CRC-32 算法替换为 Adler-32，仅需要更改构造函数调用即可；其余代码不受算法更改的影响。

一种类似的方法允许将各个`crypto`包中的流密码算法与它们链接在一起的分组密码分开。`Block`数据`crypto/cipher`包中的接口指定了分组密码的行为，该密码提供了单个数据块的加密。然后，类似于该`bufio`包，实现该接口的密码包可用于构造该`Stream`接口表示的流式密码，而无需了解块加密的详细信息。

该 `crypto/cipher`接口是这样的：

```go
type Block interface {
    BlockSize() int
    Encrypt(dst, src []byte)
    Decrypt(dst, src []byte)
}

type Stream interface {
    XORKeyStream(dst, src []byte)
}
```

这是计数器模式（CTR）流的定义，它将块密码转换为流密码。注意，分组密码的详细信息已被抽象掉：

```go
// NewCTR returns a Stream that encrypts/decrypts using the given Block in
// counter mode. The length of iv must be the same as the Block's block size.
func NewCTR(block Block, iv []byte) Stream
```

`NewCTR`不仅适用于一种特定的加密算法和数据源，而且适用于`Block`接口的任何实现以及任何 `Stream`。因为它们返回接口值，所以用其他加密模式替换 CTR 加密是本地化的更改。构造函数调用必须进行编辑，但是由于周围的代码必须仅将结果视为 a `Stream`，因此不会注意到差异。

### 接口和方法

由于几乎所有内容都可以添加方法，因此几乎所有内容都可以满足接口。有一个说明性示例在 http`包中 ，它定义了`Handler`接口。任何实现`Handler`的对象都可以处理 HTTP 请求。

```go
type Handler interface{
    ServeHTTP（ResponseWriter,* Request）
}
```

`ResponseWriter`本身是一个接口，提供对将响应返回给客户端所需的方法的访问。这些方法包括标准`Write`方法，因此 `http.ResponseWriter`可以在任何用到了`io.Writer` 的地方使用。 `Request`是一个包含来自客户端的请求的已解析的结构。

为简便起见，让我们忽略 POST，并假设 HTTP 请求始终是 GET；简化不会影响处理程序的设置方式。这是处理程序的简单实现，用于计算访问页面的次数。

```go
// Simple counter server.
type Counter struct {
    n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ctr.n++
    fmt.Fprintf(w, "counter = %d\n", ctr.n)
}
```

（注意我们的主题，注意如何`Fprintf`打印到 `http.ResponseWriter`。）在真实的服务器中，访问`ctr.n`需要防止并发访问。请参阅`sync`和`atomic`软件包以获取建议。

供参考，这里是如何将这样的服务器附加到 URL 树上的节点。

```go
import "net/http"
...
ctr := new(Counter)
http.Handle("/counter", ctr)
```

但是为什么要给`Counter`构造一个结构？整数就足够了。（接收方必须是一个指针，这样增量才能对调用方可见。）

```go
// Simpler counter server.
type Counter int

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    *ctr++
    fmt.Fprintf(w, "counter = %d\n", *ctr)
}
```

如果您的程序有一些内部状态需要在网页被访问时接收通知怎么办？将通道绑定到网页。

```go
// A channel that sends a notification on each visit.
// (Probably want the channel to be buffered.)
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
    ch <- req
    fmt.Fprint(w, "notification sent")
}
```

最后，假设我们要介绍`/args`调用服务器二进制文件时使用的参数。编写函数以打印参数很容易。

```go
func ArgServer（）{
    fmt.Println（os.Args）
}
```

我们如何将其变成 HTTP 服务器？我们可以让 ArgServer 变成某个类型的方法，忽视它的值，但是有一种更简洁的方法。由于我们可以为除指针和接口之外的任何类型定义方法，因此我们可以为函数编写方法。该`http`软件包包含以下代码：

```go

// The HandlerFunc type is an adapter to allow the use of
// ordinary functions as HTTP handlers.  If f is a function
// with the appropriate signature, HandlerFunc(f) is a
// Handler object that calls f.
type HandlerFunc func(ResponseWriter, *Request)

// ServeHTTP calls f(w, req).
//为函数编写方法
func (f HandlerFunc) ServeHTTP(w ResponseWriter, req *Request) {
    f(w, req)
}
```

`HandlerFunc`是具有方法的类型`ServeHTTP`，因此该类型的值可以处理 HTTP 请求。看一下该方法的实现：接收者是一个函数`f`，并且该方法调用`f`。这可能看起来很奇怪，但与接收方是通道和在该通道上发送方法没有什么不同。

为了`ArgServer`成为 HTTP 服务器，我们首先将其修改为具有正确的签名。

```go
// Argument server.
func ArgServer(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, os.Args)
}
```

`ArgServer`现在有相同的签名`HandlerFunc`，因此它可以被转换成该类型来访问它的方法，就像我们转换`Sequence`到`IntSlice` 访问`IntSlice.Sort`。设置它的代码很简洁：

```go
http.Handle("/args", http.HandlerFunc(ArgServer))
```

当有人访问该页面时`/args`，安装在该页面上的处理程序具有值`ArgServer` 和类型`HandlerFunc`。HTTP 服务器将以接收者身份调用该`ServeHTTP` 类型的方法，该方法`ArgServer`将依次调用 `ArgServer`（通过`f(w, req)` 内部调用`HandlerFunc.ServeHTTP`）。然后将显示参数。

在本节中，我们由结构，整数，通道和函数组成了 HTTP 服务器（给结构、整数、通道、函数定义了 ServeHTTP)，这是因为接口只是方法的集合，几乎可以给任何类型定义方法。

## 空白标识符

在[`for` `range`循环](https://golang.google.cn/doc/effective_go.html#for) 和[map](https://golang.google.cn/doc/effective_go.html#maps)的上下文中，我们已经多次提到了空白标识符 。可以使用任何类型的任何值来分配或声明空白标识符，并且可以无害地丢弃该值。这有点像写入 Unix`/dev/null`文件：它表示只写值，用作需要变量但实际值无关的占位符。它的用途超出了我们已经看到的范围。

### 多重分配中的空白标识符

在`for` `range`循环中使用空白标识符是一般情况的一种特殊情况：多重分配。

如果赋值在左侧需要多个值，但是程序不会使用其中一个值，则赋值左侧的空白标识符可以避免创建虚拟变量的需要，并明确说明：**该值将被丢弃**。例如，当函数调用返回一个值和一个错误但只有错误是重要的时候，丢弃不相关的值。

```go
if _, err := os.Stat(path); os.IsNotExist(err) {
	fmt.Printf("%s does not exist\n", path)
}
```

有时，您会看到丢弃该错误值以忽略该错误的代码。这是可怕的做法。始终检查错误返回；提供它们是有原因的。

```
// Bad! This code will crash if path does not exist.
fi, _ := os.Stat(path)
if fi.IsDir() {
    fmt.Printf("%s is a directory\n", path)
}
```

### 未使用的导入和变量

导入包或声明变量而不使用它是错误的。未使用的导入会使程序臃肿，并且编译缓慢，而已初始化但未使用的变量至少会浪费计算量，并且可能表明有一个较大的错误。但是，当程序正在积极开发中时，经常会出现未使用的导入和变量，并且为了继续进行编译而删除它们，而稍后又需要它们，可能会很烦人。空白标识符提供了一种解决方法。

这个半编写的程序有两个未使用的导入（`fmt`和`io`）和一个未使用的变量（`fd`），因此不会编译，但是很高兴看到到目前为止的代码是否正确。

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
}
```

要让编译器对未使用的 imports 包保持沉默，请使用空白标识符来引用导入包中的符号。同样，将未使用的变量 fd 分配给空白标识符将使未使用的变量错误消失。该版本的程序可以编译。

```go
package main

import (
    "fmt"
    "io"
    "log"
    "os"
)

var _ = fmt.Printf // For debugging; delete when done.
var _ io.Reader    // For debugging; delete when done.

func main() {
    fd, err := os.Open("test.go")
    if err != nil {
        log.Fatal(err)
    }
    // TODO: use fd.
    _ = fd
}
```

按照惯例，全局声明应在导入之后立即赋值给\_空白符号并进行注释，以使其易于查找，并提醒以后进行清理。

### 为了使用副作用而导入

上一个示例中未使用的导入最终应当使用或删除（如`fmt`或`io`在上一个示例中）：空白分配将代码标识为正在开发的代码。但是有时仅出于副作用导入软件包是有用的，而不需要任何显式使用。例如， 程序包 net/http/pprof 在其`init`功能期间注册提供调试信息的 HTTP 处理程序。它具有导出的 API，但是大多数客户端只需要注册处理程序，即可通过网页访问数据。要仅出于副作用导入软件包，请将软件包重命名为空白标识符：

```
import _ "net/http/pprof"
```

这种导入形式清楚地表明，将 net/http/pprof 导入包中就是为了使用它的副作用，因为没有其他可能的用法：在此文件中，它没有名称。（如果这样做，并且我们没有使用该名称，则编译器将拒绝该程序。）

### 接口检查

正如我们在上面关于[接口](https://golang.google.cn/doc/effective_go.html#interfaces_and_types)的讨论中所看到的，类型不必显式声明它实现了接口。相反，类型仅通过实现接口的方法来实现接口。实际上，大多数接口转换都是静态的，因此在编译时进行检查。例如，除非\*os.File 实现 io.Reader 接口，否则 io.Reader 将不会编译 。

但是，某些接口检查确实在运行时发生。`encoding/json` 包中有一个实例，它定义了一个`Marshaler` 接口。当 JSON 编码器收到实现该接口的值时，编码器将调用该值的 marshaling 方法将其转换为 JSON，而不是执行标准转换。编码器在运行时使用以下[类型断言](https://golang.google.cn/doc/effective_go.html#interface_conversions)检查此属性：

```go
m, ok := val.(json.Marshaler)
```

如果仅需要询问某个类型是否实现了一个接口，而不实际使用该接口本身（也许作为错误检查的一部分），请使用空白标识符忽略类型声明的值：

```go
//类型推断
if _, ok := val.(json.Marshaler); ok {
    fmt.Printf("value %v of type %T implements json.Marshaler\n", val, val)
}
```

这种情况出现的一个地方是，有必要在实现该类型的包中保证它实际上满足接口的情况。如果某个类型（例如） `json.RawMessage`需要自定义 JSON 表示形式，则应实现 `json.Marshaler`，但是没有静态转换会导致编译器自动对此进行验证。如果类型意外地不满足该接口，则 JSON 编码器仍将起作用，但将不使用自定义实现。为了确保实现正确，可以在包中使用使用空白标识符的全局声明：

```go
var _ json.Marshaler = (*RawMessage)(nil)
```

在这个声明中，涉及将`*RawMessage`转换为`Marshaler` 需要`*RawMessage`实现`Marshaler`，并且该属性将在编译时被检查。如果`json.Marshaler`接口发生更改，则此软件包将不再编译，我们将注意到需要对其进行更新。

在此结构中出现空白标识符表示该声明仅存在于类型检查中，而不用于创建变量。但是，请不要对满足接口的每种类型执行此操作。按照惯例，只有在代码中不存在静态转换的情况下才使用此类声明，这种情况很少见。

## 嵌入

Go 没有提供典型的类型驱动的子类概念，但是它确实具有通过**将类型嵌入结构或接口**中来“借用”实现的各个部分的能力。

接口嵌入非常简单。我们之前提到过`io.Reader`and`io.Writer`接口；这是他们的定义。

```go
type Reader interface {
    Read(p []byte) (n int, err error)
}

type Writer interface {
    Write(p []byte) (n int, err error)
}
```

该`io`程序包还导出其他几个接口，这些接口指定可以实现多种此类方法的对象。例如，有`io.ReadWriter`一个同时包含`Read`和的接口`Write`。我们可以`io.ReadWriter`通过显式列出这两种方法来指定，但是嵌入这两种接口以形成新的接口更容易，更令人回味，如下所示：

```go
// ReadWriter is the interface that combines the Reader and Writer interfaces.
type ReadWriter interface {
    Reader
    Writer
}
```

这只是说，是什么样子：一个`ReadWriter`可以做一`Reader`做*和*一个什么`Writer` 呢; 它是嵌入式接口的并集。只有接口可以嵌入接口中。

相同的基本思想适用于结构，但意义更深远。所述`bufio`封装具有两个结构类型， `bufio.Reader以及bufio.Writer`，每一个都实现了 io 包中的类似接口。并且`bufio`还实现了缓冲的读取/写入接口，该操作通过使用嵌入将读取器和写入器组合为一个结构来完成：它列出了结构中的类型，但未提供字段名称。

```go
// ReadWriter stores pointers to a Reader and a Writer.
// It implements io.ReadWriter.
type ReadWriter struct {
    *Reader  // *bufio.Reader
    *Writer  // *bufio.Writer
}
```

嵌入的元素是指向结构的指针，当然，必须先将其初始化为指向有效结构，然后才能使用它们。该`ReadWriter`结构可以写成

```go
type ReadWriter struct {
    reader *Reader
    writer *Writer
}
```

但是为了提升字段的方法并满足`io`接口要求，我们还需要提供转发方法，如下所示：

```go
func (rw *ReadWriter) Read(p []byte) (n int, err error) {
    return rw.reader.Read(p)
}
```

通过直接嵌入结构，我们避免了这种记录。嵌入式类型的方法随着类型一起进来了，这意味着`bufio.ReadWriter` 不仅具有`bufio.Reader`和`bufio.Writer的方法，同时也满足了所有三个接口： `io.Reader`，`io.Writer`，和 `io.ReadWriter`。

嵌入和子类有一个重要的区别。当我们嵌入一个类型时，该类型的方法成为外部类型的方法，但是当调用它们时，该方法的接收者是内部类型，而不是外部类型。在我们的示例中，调用`bufio.ReadWriter`的 Read 方法时，其效果与上面写出的转发方法完全相同；接收者是 ReadWriter 的 reader 字段，而不是 `ReadWriter`本身。

嵌入也可以很方便。此示例显示了一个嵌入的字段以及一个常规的命名字段。

```go
type Job struct {
    Command string
    *log.Logger
}
```

该`Job`类型现在有`Print`，`Printf`，`Println` 和其他方法`*log.Logger`。`Logger` 当然，我们可以给一个字段名，但是没有必要这样做。现在，一旦初始化，我们就可以使用 Job 来打印日志：

```go
job.Println("starting now...")
```

该`Logger`是`Job`结构的一个普通字段，所以我们可以用通常的方法在构造函数中进行`Job`的初始化，这样：

```go
func NewJob(command string, logger *log.Logger) *Job {
    return &Job{command, logger}
}
```

或使用复合表达式，

```go
job := &Job{command, log.New(os.Stderr, "Job: ", log.Ldate)}
```

如果我们需要直接引用一个嵌入式字段，则该字段的类型名称（忽略包限定符）将用作字段名称，就像在 struct`Read`方法中一样`ReadWriter`。在这里，如果我们需要访问 `*log.Logger`一个的`Job`变量`job`，我们会写`job.Logger`，如果我们想要改进的方法，这将是有益的`Logger`。

```
func（job * Job）Printf（format string，args ... interface {}）{
    job.Logger.Printf（“％q：％s”，job.Command，fmt.Sprintf（format，args ...））
}
```

嵌入类型引入了名称冲突的问题，但是解决它们的规则很简单。首先，字段或方法`X`将其他任何对象`X`隐藏在该类型的更深层嵌套的部分中。如果`log.Logger`包含称为的字段或方法`Command`，则 Job 的`Command`字段``将占主导地位。内层 Command 将不可见

其次，如果相同的名称出现在相同的嵌套级别，则通常是错误的。`log.Logger`如果`Job`结构包含另一个称为 Logger 的字段或方法，则嵌入将是错误的。但是，如果在类型定义之外 Logger 从未出现，则可以。这种特点保护提供了一些保护，以防止外部嵌入的类型发生更改。如果添加的字段与另一个子类型中的另一个字段发生冲突并且这两个字段都不曾使用过，则没有问题。

## 并发

### 通过交流分享

并发编程是一个很大的话题，这里仅留有一些特定于 Go 的亮点。

需要实现对共享变量的正确访问使得在许多环境中进行并行编程变得很困难。Go 鼓励采用一种不同的方法，在这种方法中，共享值在通道之间传递，并且实际上，决不由单独的执行线程主动共享。在任何给定时间，只有一个 goroutine 可以访问该值。根据设计，不会发生数据竞争。为了鼓励这种思维方式，我们将其简化为一个口号：

> 不要通过共享内存进行通信；而是通过通信共享内存。

这种方法可能太过分了。例如，最好通过将互斥锁放在整数变量周围来完成引用计数。但是作为一种高级方法，使用通道来控制访问权限使编写清晰，正确的程序变得更加容易。

考虑该模型的一种方法是考虑一个 CPU 上运行的典型单线程程序。它不需要同步原始变量。现在运行另一个这样的实例；它也不需要同步。现在让这两个程序交流；如果通信就是同步器，则仍然不需要其他同步手段。例如，Unix 管道非常适合此模型。尽管 Go 的并发方法源自 Hoare 的通信顺序过程（CSP），但它也可以被视为类型安全泛化后的 Unix 管道。

### Goroutines

之所以称为*goroutine，*是因为现有的术语（线程，协程，进程等）传达了不准确的含义。goroutine 有一个简单的模型：它是在同一地址空间中与其他 goroutine 同时执行的函数。它是轻量级的，仅稍微跟分配堆栈空间比消耗资源。而且堆栈从小开始，因此消耗资源小，并且可以通过根据需要分配（和释放）堆内存来增长。

Goroutine 被多路复用到多个 OS 线程上，因此，如果一个阻塞，例如在等待 I / O 时，其他将继续运行。他们的设计隐藏了线程创建和管理的许多复杂性。

给函数或方法调用加上前缀 gogoroutine 中运行该调用。调用完成后，goroutine 会静默退出。（效果类似于 Unix Shell 的&表示法在后台运行命令。）

```
go list.Sort（）//同时运行list.Sort; 不要等待。
```

匿名函数在 goroutine 调用中可以派上用场。

```go
func Announce(message string, delay time.Duration) {
    go func() {
        time.Sleep(delay)
        fmt.Println(message)
    }()  // Note the parentheses - must call the function.
}
```

在 Go 中，匿名函数是闭包：实现可确保函数所引用的变量只要处于活动状态就可以保留。

这些示例不太实用，因为这些函数无法发出完成信号。为此，我们需要通道。

### 通道

与映射一样，通道也用`make`分配，并且结果值是对基础数据结构的引用。如果提供了可选的整数参数，则它将设置通道的缓冲区大小。对于**无缓冲**（同步通道），默认值为零。

```
ci：= make（chan int）//无缓冲的整数通道
cj：= make（chan int，0）//无缓冲的整数通道
cs：= make（chan * os.File，100）//指向文件的指针的缓冲通道
```

无缓冲通道将通信（值的交换）与同步相结合，从而确保两个计算（goroutines）处于已知状态。

使用通道有很多不错的习惯用法。我们从下面的例子开始。在上一节中，我们在后台启动了排序。通道可以允许启动 goroutine 等待排序完成。

```go
c := make(chan int)  // Allocate a channel.
// Start the sort in a goroutine; when it completes, signal on the channel.
go func() {
    list.Sort()
    c <- 1  // Send a signal; value does not matter.
}()
doSomethingForAWhile()
<-c   // Wait for sort to finish; discard sent value.
```

接收器始终阻塞，直到有数据要接收为止。如果通道未缓冲，则发送方将阻塞，直到接收方收到该值为止。如果通道具有缓冲区，则发送方仅在将值成功复制到缓冲区之前阻塞；如果缓冲区已满，则意味着等待直到某些接收者接受到一个值。

可以像信号灯一样使用缓冲的通道，例如以限制吞吐量。在此下面的示例中，传入的请求被传递到`handle`，后者将值发送到通道，处理该请求，然后从通道接收一个值，以为下一个使用者准备“信号量”。通道缓冲区的容量将同时呼叫的数量限制为`process`。

```go
var sem = make(chan int, MaxOutstanding)

func handle(r *Request) {
    sem <- 1    // Wait for active queue to drain.
    process(r)  // May take a long time.
    <-sem       // Done; enable next request to run.
}

func Serve(queue chan *Request) {
    for {
        req := <-queue
        go handle(req)  // Don't wait for handle to finish.
    }
}
```

一旦`MaxOutstanding`处理程序执行`process`完毕，任何其他处理程序都将阻止尝试发送到已填充的通道缓冲区，直到现有处理程序之一完成并从缓冲区接收消息为止。

但是，这种设计有一个问题：即使只有 MaxOutstanding 个请求可以随时运行，`Serve` 也会为每个传入请求创建一个新的 goroutine 。如此一来，如果请求太快，程序可能会消耗无限的资源。我们可以通过更改`Serve`以控制 goroutine 的创建来解决该缺陷。这是一个显而易见的解决方案，但是请注意，它有一个错误，我们将在随后修复：

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func() {
            process(req) // Buggy; see explanation below.
            <-sem
        }()
    }
}
```

错误是在**Go`for`循环中，循环变量在每次迭代中都会重复使用**，因此该`req` 变量在所有 goroutine 中共享。那不是我们想要的。我们需要确保`req`每个 goroutine 都是唯一的。这是一种实现方法，将 go 的值`req`作为参数传递给 goroutine 中的闭包：

```go
func Serve(queue chan *Request) {
    for req := range queue {
        sem <- 1
        go func(req *Request) {
            process(req)
            <-sem
        }(req)
    }
}
```

将此版本与先前版本进行比较，以了解闭包是如何声明并且允许的。另一个解决方案是仅创建一个具有相同名称的新变量，如下例所示：

```go
func Serve(queue chan *Request) {
    for req := range queue {
        req := req // Create new instance of req for the goroutine.
        sem <- 1
        go func() {
            process(req)
            <-sem
        }()
    }
}
```

写起来似乎很奇怪

```go
req := req
```

但这在 Go 中是合法且惯用的。您将获得具有相同名称的变量的新版本，有意在本地隐藏循环变量，但每个 goroutine 均具有唯一性。

回到编写服务器的一般问题，另一种很好管理资源的方法是启动固定数量的 handle goroutine,它们从请求通道读取数据。goroutine 的数量将同时调用的数量限制为`process`。此`Serve`函数还接受一个将告知其退出的通道；启动 goroutines 后，它将阻止从该通道接收。

```go
func handle(queue chan *Request) {
    for r := range queue {
        process(r)
    }
}

func Serve(clientRequests chan *Request, quit chan bool) {
    // Start handlers
    for i := 0; i < MaxOutstanding; i++ {
        go handle(clientRequests)
    }
    <-quit  // Wait to be told to exit.
}
```

### 通道的通道

Go 的最重要属性之一是通道是第一级的值，可以像分配其他值一样进行分配和传递。此属性的常见用法是实现安全的并行多路分解。

在上一节的示例中，`handle`是请求的理想处理程序，但是我们没有定义其处理的类型。如果该类型包括用于回复的渠道，则每个客户端可以提供自己的答案路径。这是 type 的示意图定义`Request`。

```go
type Request struct {
    args        []int
    f           func([]int) int
    resultChan  chan int
}
```

客户端提供一`Map` 个函数及其参数，以及请求对象内部的一个接收答案的通道。

```go
func sum(a []int) (s int) {
    for _, v := range a {
        s += v
    }
    return
}

request := &Request{[]int{3, 4, 5}, sum, make(chan int)}
// Send request
clientRequests <- request
// Wait for response.
fmt.Printf("answer: %d\n", <-request.resultChan)
```

在服务器端，处理程序功能是唯一更改的东西。

```
func handle（queue chan * Request）{
    对于req：=范围队列{
        req.resultChan <-req.f（req.args）
    }
}
```

要使它变得现实，显然还有很多工作要做，但是此代码是一个用于速率受限，并行，无阻塞 RPC 系统的框架，并且看不到互斥量。

### 并行化

这些想法的另一个应用是使多个 CPU 内核之间的计算并行化。如果可以将计算分解为可以独立执行的独立部分，则可以并行化计算，并在每个部分完成时发出信号。

假设我们要对一个项目的向量执行高消耗的操作，并且对每个项目的操作的值都是独立的，如本理想示例所示。

```go
type Vector []float64

// Apply the operation to v[i], v[i+1] ... up to v[n-1].
func (v Vector) DoSome(i, n int, u Vector, c chan int) {
    for ; i < n; i++ {
        v[i] += u.Op(v[i])
    }
    c <- 1    // signal that this piece is done
}

```

我们以循环方式独立启动各个部分，每个 CPU 一个。他们可以按任何顺序完成，但这无关紧要。在启动所有 goroutine 之后，我们通过排空通道来计数完成信号。

```go
const numCPU = 4 // number of CPU cores

func (v Vector) DoAll(u Vector) {
    c := make(chan int, numCPU)  // Buffering optional but sensible.
    for i := 0; i < numCPU; i++ {
        go v.DoSome(i*len(v)/numCPU, (i+1)*len(v)/numCPU, u, c)
    }
    // Drain the channel.
    for i := 0; i < numCPU; i++ {
        <-c    // wait for one task to complete
    }
    // All done.
}
```

可以为运行时询问哪个值合适，而不是为 numCPU 创建一个常量值。该函数`runtime.NumCPU` 返回计算机中硬件 CPU 内核的数量，因此我们可以编写

```
var numCPU = runtime.NumCPU（）
```

还有一个功能 `runtime.GOMAXPROCS`，可以报告（或设置）Go 程序可以同时运行的用户指定的内核数。它的默认值为，`runtime.NumCPU`但可以通过设置类似名称的 shell 环境变量或使用正数调用该函数来覆盖它。用零调用它只是查询值。因此，如果我们想满足用户的资源请求，我们应该写

```
var numCPU = runtime.GOMAXPROCS（0）
```

确保不要混淆并发的思想（将程序构造为独立执行的组件）和并行性，并发执行并行计算以提高多个 CPU 的效率。尽管 Go 的并发特性可以使一些问题易于并行计算，但 Go 是一种并发语言，而不是并行语言，并且并非所有并行化问题都适合 Go 的模型。有关区别的讨论，请参见此[博客文章中](https://blog.golang.org/2013/01/concurrency-is-not-parallelism.html)引用的演讲 。

### 缓冲区泄漏

并发编程工具甚至可以使非并发思想更容易表达。下面是从 RPC 包中抽象出来的示例。客户端 goroutine 循环从某个来源（可能是网络）接收数据。为了避免分配和释放缓冲区，它会保留一个空闲列表，并使用一个缓冲通道来表示它。如果通道为空，则会分配一个新的缓冲区。消息缓冲区准备就绪后，它将被发送到`serverChan`的服务器 。

```go
var freeList = make（chan * Buffer，100）
var serverChan = make（chan * Buffer）

func client（）{
    for{
        var b *Buffer
        //获取一个缓冲区（如果有）；如果没有则分配一个新的缓冲区。
        select {
        case b = <-freeList：
            // 拿到一个; 无事可做。
        default：
            //没有一个免费的，因此分配一个新的。
            b = new（Buffer）
        }
        load（b）//从网上读取下一条消息。
        serverChan <-b //发送到服务器。
    }
}
```

服务器循环从客户端接收每个消息，对其进行处理，然后将缓冲区返回到空闲列表。

```go
func server（）{
    for{
        b：= <-serverChan //等待工作。
        process（b）
        //如果有空间，请重新使用缓冲区。
        select {
        case freeList <-b：
            //在空闲列表上缓冲 无事可做。
        default：
            //免费列表已满，只需继续。
        }
    }
}
```

客户端尝试从中`freeList`中获取一个 Buffer；如果没有可用的，它将分配一个新的。除非列表已满，否则服务器将 buffer 放回到空闲列表中，在这种空闲列表已满的情况下，缓冲区会被垃圾收集器回收。（`default`语句中的子句在`select` 没有其他`case`准备就绪时执行，这意味着`selects`永不阻塞。）此实现仅依靠几行内容就建立了一个无泄漏存储桶列表，并依赖于缓冲通道和垃圾收集器进行记录。

## Errors

Library routines 必须经常向调用者返回某种错误指示。如前所述，Go 的多值返回可以轻松地在正常返回值旁边返回详细的错误描述。使用此功能提供详细的错误信息是一种很好的方式。例如，正如我们将看到的那样，`os.Open`不仅会`nil`在失败时返回一个指针，还会返回一个描述错误原因的错误值。

按照惯例，错误的类型为`error`，是一个简单的内置接口。

```go
type error interface{
    Error（）string
}
```

Library 作者可以自由地用一个更丰富的模型来实现此接口，从而不仅可以看到错误，而且可以提供一些上下文。如前所述，除了通常的`*os.File` 返回值外，`os.Open`还返回错误值。如果文件成功打开，则错误将为`nil`，但是如果出现问题，它将包含一个 `os.PathError`：

```go
// PathError records an error and the operation and
// file path that caused it.
type PathError struct {
    Op string    // "open", "unlink", etc.
    Path string  // The associated file.
    Err error    // Returned by the system call.
}

func (e *PathError) Error() string {
    return e.Op + " " + e.Path + ": " + e.Err.Error()
}

```

`PathError`的会`Error`产生一个像这样的字符串：

```
open /etc/passwx: no such file or directory
```

这种错误包括了有问题的文件名，操作以及所触发的操作系统错误，即使在导致错误的调用远未打印的情况下也有用。它比普通的“没有这样的文件或目录”提供更多信息。

在可行的情况下，错误字符串应标识其来源，例如通过使用前缀来命名产生错误的操作或程序包。例如，在 package 中 `image`，由于未知格式导致的解码错误的字符串表示形式是“ image：unknown format”。

关心精确错误详细信息的调用者可以使用类型切换或类型断言来查找特定错误并提取详细信息。对`PathErrors` 来说可能包括检查内部`Err` 字段是否存在可恢复的故障。

```go
for try := 0; try < 2; try++ {
    file, err = os.Create(filename)
    if err == nil {
        return
    }
    if e, ok := err.(*os.PathError); ok && e.Err == syscall.ENOSPC {
        deleteTempFiles()  // Recover some space.
        continue
    }
    return
}
```

这里第二条语句的`if`是另一种[类型断言](https://golang.google.cn/doc/effective_go.html#interface_conversions)。如果失败，`ok`则为 false，`e` 为`nil`。如果成功， 则为`ok`为true，表示错误的类型为`*os.PathError`，然后`e`也是一样，我们可以检查该错误的更多信息。

### Panic

向调用者报告错误的通常方法是返回一个 `error`作为额外的返回值。规范 `Read`方法是一个众所周知的实例。它返回一个字节数和一个`error`。但是，如果错误无法恢复怎么办？有时程序根本无法继续。

为此，有一个内置函数`panic` 实际上会创建一个运行时错误，该错误将使程序停止运行（但请参阅下一节）。该函数采用一个任意类型的参数（通常是字符串），以便在程序死亡时打印出来。这也是一种指示发生了不可能的事情的方法，例如退出无限循环。

```go
// A toy implementation of cube root using Newton's method.
func CubeRoot(x float64) float64 {
    z := x/3   // Arbitrary initial value
    for i := 0; i < 1e6; i++ {
        prevz := z
        z -= (z*z*z-x) / (3*z*z)
        if veryClose(z, prevz) {
            return z
        }
    }
    // A million iterations has not converged; something is wrong.
    panic(fmt.Sprintf("CubeRoot(%g) did not converge", x))
}

```

这只是一个示例，但实际的库函数应避免使用`panic`。如果问题可以掩盖或解决，最好还是让事情继续运行而不是取消整个程序。一个可能的反例是在初始化期间：如果该库确实无法进行设置，那么恐慌是可以理解的。

```
var user = os.Getenv（“ USER”）

func init（）{
    如果用户==“” {
        panic（“ $ USER没有价值”）
    }
}
```

### 恢复

当`panic`被调用时（包括对运行时错误的隐式调用，例如，对切片进行索引编制索引或失败类型声明），它将立即停止当前函数的执行并开始展开 goroutine 的堆栈，并在此过程中运行所有延迟函数。如果解散到达 goroutine 栈的顶部，程序将终止。但是，可以使用内置函数`recover`来重新获得对 goroutine 的控制并恢复正常执行。

调用将`recover`停止展开并返回传递给的参数`panic`。因为在展开时运行的唯一代码是在延迟函数内部，`recover` 所以仅在延迟函数内部有用。

一种应用`recover`是关闭服务器内部失败的 goroutine，而不会杀死其他正在执行的 goroutine。

```
func服务器（workChan <-chan *工作）{
    工作：=范围workChan {
        安全去做（工作）
    }
}

func safeDo（工作*工作）{
    延迟func（）{
        如果err：= recovery（）; err！= nil {
            log.Println（“工作失败：”，错误）
        }
    }（）
    做工作）
}
```

在此示例中，如果出现`do(work)`紧急情况，将记录结果，并且 goroutine 将干净地退出而不会打扰其他程序。延迟的关闭过程中无需执行任何其他操作；调用`recover`完全处理条件。

因为`recover`总是返回，`nil`除非直接从延迟函数调用，因此延迟代码可以调用本身使用的库例程，`panic`而`recover`不会失败。例如，in 中的 deferred 函数`safelyDo`可能在调用之前先调用日志记录函数`recover`，并且该日志记录代码将不受恐慌状态的影响。

有了我们的恢复模式，该`do` 函数（及其调用的任何函数）都可以通过调用干净地摆脱任何不良情况`panic`。我们可以使用该想法来简化复杂软件中的错误处理。让我们看一下`regexp`软件包的理想版本，该版本通过调用`panic`本地错误类型来报告解析错误。这是`Error`，`error`方法和`Compile`函数的定义。

```
// Error是解析错误的类型；它满足错误界面。
类型错误字符串
func（e Error）Error（）字符串{
    返回字符串（e）
}

//错误是* Regexp的一种方法，通过以下方法报告解析错误
//出现错误时惊慌失措。
func（regexp * Regexp）错误（错误字符串）{
    恐慌（错误（错误））
}

//编译返回正则表达式的解析表示形式。
func Compile（str字符串）（regexp * Regexp，错误错误）{
    regexp = new（正则表达式）
    //如果存在解析错误，doParse将会恐慌。
    延迟func（）{
        如果e：= recovery（）; e！= nil {
            regexp = nil //清除返回值。
            err = e。（Error）//如果不是解析错误，将重新出现紧急情况。
        }
    }（）
    返回regexp.doParse（str），nil
}
```

如果出现`doParse`紧急情况，恢复块会将返回值设置为—`nil`延迟函数可以修改命名的返回值。然后`err`，它将通过断言其具有本地类型来检查问题是否为解析错误`Error`。如果不是这样，则类型声明将失败，从而导致运行时错误，该错误将继续展开堆栈，就像没有任何中断一样。此检查意味着，如果发生意外情况（例如索引超出范围），即使我们正在使用`panic`并`recover`处理解析错误，代码也将失败。

有了错误处理，该`error`方法（因为它是绑定到类型的方法，所以它很好，甚至很自然，因为它具有与内置`error`类型相同的名称），可以很容易地报告解析错误，而不必担心展开解析堆栈用手：

```
如果pos == 0 {
    re.error（“'*'在表达式开始时是非法的”）
}
```

尽管此模式很有用，但应仅在包内使用。 `Parse`将内部`panic`调用转化为 `error`价值；它不会`panics` 向其客户公开。这是遵循的好规则。

顺便说一句，如果发生实际错误，此重新恐慌习惯用法会更改恐慌值。但是，原始故障和新故障都将显示在崩溃报告中，因此问题的根本原因仍然可见。因此，这种简单的重新恐慌方法通常就足够了-毕竟是崩溃。但是，如果您只想显示原始值，则可以编写更多代码来过滤意外的问题并使用原始错误重新恐慌。留给读者练习。

## Web 服务器

让我们完成一个完整的 Go 程序，一个 Web 服务器。这实际上是一种 Web 重新服务器。Google 提供了一项服务，`chart.apis.google.com` 可以将数据自动格式化为图表和图形。但是，很难以交互方式使用它，因为您需要将数据作为查询放入 URL。这里的程序为一种数据形式提供了一个更好的接口：给定一小段文本，它会在图表服务器上调用以产生 QR 码，即编码文本的盒子矩阵。该图像可以用手机的摄像头捕获，并解释为例如 URL，从而省去了在手机的小键盘上键入 URL 的麻烦。

这是完整的程序。解释如下。

```
包主

导入（
    “旗”
    “ html /模板”
    “日志”
    “ net / http”
）

var addr = flag.String（“ addr”，“：1718”，“ http服务地址”）// Q = 17，R = 18

var templ = template.Must（template.New（“ qr”）。Parse（templateStr））

func main（）{
    flag.Parse（）
    http.Handle（“ /”，http.HandlerFunc（QR））
    错误：= http.ListenAndServe（* addr，nil）
    如果err！= nil {
        log.Fatal（“ ListenAndServe：”，err）
    }
}

func QR（w http.ResponseWriter，req * http.Request）{
    templ.Execute（w，req.FormValue（“ s”））
}

const templateStr =`
<html>
<头>
<title> QR链接生成器</ title>
</ head>
<身体>
{{if。}}
<img src =“ http://chart.apis.google.com/chart?chs=300x300&cht=qr&choe=UTF-8&chl= {{。}}” />
<br>
{{。}}
<br>
<br>
{{结束}}
<form action =“ /” name = f method =“ GET”>
    <input maxLength = 1024 size = 70 name = s value =“” title =“文本到QR编码”>
    <input type =提交值=“ Show QR” name = qr>
</ form>
</ body>
</ html>
`
```

最多的部分`main`应该易于遵循。一个标志为我们的服务器设置默认的 HTTP 端口。模板变量`templ`是有趣的地方。它构建了一个 HTML 模板，该模板将由服务器执行以显示页面。稍后了解更多。

该`main`函数解析标志，并使用我们上面讨论的机制将函数绑定`QR`到服务器的根路径。然后`http.ListenAndServe`被称为启动服务器；服务器运行时会阻塞。

`QR`只会接收包含表单数据的请求，并以名为的表单值对数据执行模板`s`。

模板包`html/template`功能强大；该程序仅涉及其功能。本质上，它通过替换从传递给的数据项派生的元素`templ.Execute`（在本例中为表单值）来即时重写 HTML 文本。在模板文本（`templateStr`）中，用双括号分隔的段表示模板动作。仅当当前数据项的值（点）为非空时，from`{{if .}}` 才`{{end}}`执行`.`。即，当字符串为空时，该模板部分被抑制。

这两个摘要`{{.}}`表示要在网页上显示提供给模板的数据（查询字符串）。HTML 模板包会自动提供适当的转义符，因此可以安全地显示文本。

模板字符串的其余部分只是页面加载时显示的 HTML。如果解释太快，请参阅 模板包的[文档](https://golang.google.cn/pkg/html/template/)以进行更全面的讨论。

在那里，您可以找到：一个有用的 Web 服务器，其中包含几行代码以及一些数据驱动的 HTML 文本。Go 的功能强大到足以在几行中完成很多事情。
