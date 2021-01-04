## 安装与环境搭建

### 程序下载

#### mac👨🏻‍💻

- 安装包下载地址: https://golang.google.cn/dl/

- 打开安装文件，默认安装到/usr/local/go/bin 目录，重新启动 terminal 让改动生效

- 在 terminal 中使用下面的命令来验证 go 安装成功 :

```
   $ go version
```

- 确认 go version 版本正确

#### windows

- 打开下载好的 MSI 文件，默认将 Go 安装到 C:\Go

- 验证成功安装 Go.

  - 点击开始菜单

  - 搜索 CMD

  - 在跳出的界面中输入:

```
    $ go version
```

- 确认 go version 版本正确

## Golang 官方工具

### 工具命令

### go get

Usage:

go get [-d] [-t] [-u] [build flags] [packages]

Examples:

安装一个工具的最新版本

```
$ go get golang.org/x/tools/cmd/goimports
```

升级一个指定的模块

```
$ go get -d golang.org/x/net
```

升级主模块所引入的包所在的模块

```
$ go get -d -u ./...
```

升级或者降级模块到指定版本

```
$ go get -d golang.org/x/text@v0.3.2
```

将模块升级到 master 分支上的提交版本

```
$ go get -d golang.org/x/text@master
```

删除一个模块依赖并且将需要该模块的其他模块降级到不需要它的版本

```
$ go get -d golang.org/x/text@none
```

`go get` 命令升级 go.mod 或者主模块中的模块依赖，然后 build 并且安装命令行所列出的包。
第一步是确定哪些模块需要升级。`go get` 命令接收 packages, package patterns, and module paths 作为参数列表。如果指定了包的具体参数，`go get` 会直接升级提供该包的模块。如果指定了包的通配符, `go get` 会扩展到满足通配符的全部包并且升级提供他们的模块。
如果参数指定了模块名而不是包名，并且模块的根目录没有包，go get 会升级模块但不会构建包。如果不提供参数 `go get` 的行为就像是 `go get` .(在当前目录的全部包)。不加参数的 `go get` 可以跟-u 标志一起使用来升级提供导入包的模块。

The first step is to determine which modules to update. go get accepts a list of packages, package patterns, and module paths as arguments. If a package argument is specified, go get updates the module that provides the package. If a package pattern is specified (for example, all or a path with a ... wildcard), go get expands the pattern to a set of packages, then updates the modules that provide the packages. If an argument names a module but not a package (for example, the module golang.org/x/net has no package in its root directory), go get will update the module but will not build a package. If no arguments are specified, go get acts as if . were specified (the package in the current directory); this may be used together with the -u flag to update modules that provide imported packages.

每一个参数都可以包括一个版本查询后缀，比如 `go get golang.org/x/text@v0.3.0.`。
一个版本查询后缀包含了@符号来指定特定的版本。这个后缀可能是指定了版本号(v0.3.0),或者指定了版本前缀(v0.3)，或者指定了分支名称(master),或者修改号(1234abcd),又或者是特殊的版本查询比如`latest`,`upgrade`,`patch`,`none`

Each argument may include a version query suffix indicating the desired version, as in go get golang.org/x/text@v0.3.0. A version query suffix consists of an @ symbol followed by a version query, which may indicate a specific version (v0.3.0), a version prefix (v0.3), a branch or tag name (master), a revision (1234abcd), or one of the special queries latest, upgrade, patch, or none. If no version is given, go get uses the @upgrade query.

只要 `go get` 正确把参数解析到一个指定版本的模块，`go get` 就会添加、改变、或者删除

Once go get has resolved its arguments to specific modules and versions, go get will add, change, or remove require directives in the main module's go.mod file to ensure the modules remain at the desired versions in the future. Note that required versions in go.mod files are minimum versions and may be increased automatically as new dependencies are added. See Minimal version selection (MVS) for details on how versions are selected and conflicts are resolved by module-aware commands.

其他模块可能会随着命令行添加、升级、降级一个模块而升级。比如假设 example.com/a 升级到了 v1.5.0 版本，这个版本需要 v1.2.0 版本的 example.com/b，那么 example.com/b 也会随着 a 的升级而升级。

Other modules may be upgraded when a module named on the command line is added, upgraded, or downgraded if the new version of the named module requires other modules at higher versions. For example, suppose module example.com/a is upgraded to version v1.5.0, and that version requires module example.com/b at version v1.2.0. If module example.com/b is currently required at version v1.1.0, go get example.com/a@v1.5.0 will also upgrade example.com/b to v1.2.0.

如果使用@none 这个版本后缀，模块有可能会被删除。这是一种特殊的降级，依赖于该模块的其他模块也会被降级或者删除。一个模块依赖有可能被移除哪怕一个或者更多包被主模块中引入。这种情况下，下一次 build 命令有可能添加一个新的模块依赖。

A module requirement may be removed using the version suffix @none. This is a special kind of downgrade. Modules that depend on the removed module will be downgraded or removed as needed. A module requirement may be removed even if one or more of its packages are imported by packages in the main module. In this case, the next build command may add a new module requirement.

如果一个模块的不同版本同时需要使用，go 会报告错误。

If a module is needed at two different versions (specified explicitly in command line arguments or to satisfy upgrades and downgrades), go get will report an error.

在 `go get` 命令升级 go.mod 文件后，它会根据命令行的名称来作为构建包时的名称。可执行文件会被放到 GOBIN 环境变量的目录（默认是$GOPATH/bin 或者$HOME/go/bin)

After go get updates the go.mod file, it builds the packages named on the command line. Executables will be installed in the directory named by the GOBIN environment variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH environment variable is not set.

go get 支持下面的标记符号:

- -d  
   这个标记告诉 `go get` 不要构建或者安装包，当-d 被使用，`go get` 只会在 go.mod 文件中管理依赖。
- -u
  这个标记告诉 `go get` 只升级命令行中直接或者间接出现的包。每一个-u 选择的模块都会升级至它的最新版本，除非他已经被指定到了一个更高的版本(预发布版本)
  **-u=patch** 标记也告诉 `go get` 升级依赖，不过 `go get` 会把每个依赖升级到最新的 patch 版本。这跟@patch 版本查询有点类似。
- -t
  -t 标记告诉 go get 考虑对命令行出现的包进行测试的时候所需要的依赖。当-t 跟-u 一起使用，go get 会升级测试依赖到最新。

### go install

`go install`表示安装的意思，它先编译源代码得到可执行文件，然后将可执行文件移动到`GOPATH`的 bin 目录下。因为我们的环境变量中配置了`GOPATH`下的 bin 目录，所以我们就可以在任意地方直接执行可执行文件了。

### go build

go build 表示将源代码编译成可执行文件。

在 hello 目录下执行：

```
go build
```

或者在其他目录执行以下命令：

```
go build hello
```

go 编译器会去 GOPATH 的 src 目录下查找你要编译的 hello 项目

编译得到的可执行文件会保存在执行编译命令的当前目录下，如果是 windows 平台会在当前目录下找到 hello.exe 可执行文件。

可在终端直接执行该 hello.exe 文件：

```
c:\go\hello>hello.exe
Hello World!
```

我们还可以使用-o 参数来指定编译后得到的可执行文件的名字。

```
go build -o heiheihei.exe
```

#### 跨平台编译

默认我们`go build`的可执行文件都是当前操作系统可执行的文件，如果我想在 windows 下编译一个 linux 下可执行文件，那需要怎么做呢？

只需要指定目标操作系统的平台和处理器架构即可：

```bash
SET CGO_ENABLED=0  // 禁用CGO
SET GOOS=linux  // 目标平台是linux
SET GOARCH=amd64  // 目标处理器架构是amd64
```

_使用了 cgo 的代码是不支持跨平台编译的_

然后再执行`go build`命令，得到的就是能够在 Linux 平台运行的可执行文件了。

Mac 下编译 Linux 和 Windows 平台 64 位 可执行程序：

```bash
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Linux 下编译 Mac 和 Windows 平台 64 位可执行程序：

```bash
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build
```

Windows 下编译 Mac 平台 64 位可执行程序：

```bash
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build
```

### go clean

功能：用户删除项目的缓存文件或其他命令生成的文件。会删除以下文件（但不限于以下）

- 会删除编译 go 或命令源码文件而产生的文件，包括：“\_obj”和“\_test”目录，名称为“\_testmain.go”、“test.out”、“build.out”或“a.out”的文件，名称以“.5”、“.6”、“.8”、“.a”、“.o”或“.so”为后缀的文件。比如：执行 go build -work 会生成 WORK=C:\Users\my\AppData\Local\Temp\go-build296297870 ，go-build296297870 目录就是多生成的临时文件。

- 会删除当前目录下 `go build` 生成的 .exe 文件（假设有）。（即：删除当前代码包下生成的与包名同名或者与 Go 源码文件同名的可执行文件）

- 会删除 `go test` 命令并加入-c 标记时在当前代码包下生成的以包名加“.test”后缀为名的文件。

- go clean -i 命令：若在代码包中会删除 pkg 目录中的归档文件（.a 文件）；若在 main 包中会删除 bin 目录中安装的 .exe 可执行文件。

- go clean -n：会打印删除整个过程中用到的系统命令，但不会真正执行他们。

- go clean -n -x：在 -n 的基础之上真正执行命令，（与 go build -n -x 类似）。

- go clean -r：会递归删除当前路径包及其依赖包的一些目录与文件。

- go clean -i：删除 pkg 目录中因路径包生成的 .a 文件（即归档文件）。

- go clean -cache：删除因 `go build` 产生的全部缓存。

- go clean -testcache：让所有 `go build cache` 中的测试结果失效。

- go clean -modcache: 删除全部的 module 下载缓存，

### go doc

go doc 命令打印跟对象相关的文档注释以及参数。go doc 接受 0 参数或者 1 个参数或者两个参数。

- 零参数

  ```
  go doc
  ```

  它的作用是打印当前目录的包文档。

- 一个参数

  ```
    go doc <pkg>
    go doc <sym>[.<methodOrField>]
    go doc [<pkg>.]<sym>[.<methodOrField>]
    go doc [<pkg>.][<sym>.]<methodOrField>
  ```

  一个参数的写法是包名.待查找对象名。如果不提供包名则默认是当前目录所在包。如果 go doc 的参数名是大写开头则只在当前目录查找。

  比如:

  ```
  go doc api.device
  ```

- 两个参数

  ```
  go doc <pkg> <sym>[.<methodOrField>]
  ```

  在项目根目录可以执行下面的命令:

  ```
  go doc api device
  ```

  参数就是 device,包名是 api,如果不加包名 api 则需要 cd 到 api 的目录去执行查找。比如下面 cd 到 app 包执行不带包名的 go doc 命令:

  ```
  wujiabangs-MacBook-Pro:api steve$ go doc device
  package api // import "github.com/zsy-cn/sws-iot/admin/api"

  type Device struct {
          DeviceService devices.DeviceService
          OpcuaService  opcua.OpcUAService
  }
      Device 设备管理

  func (d *Device) AddUserChannel(c *gin.Context)
  func (d *Device) AddUserDevice(c *gin.Context)
  func (d *Device) AddUserDeviceToChannel(c *gin.Context)
  func (d *Device) AddUserProduct(c *gin.Context)
  func (d *Device) AddUserProductAttribute(c *gin.Context)
  func (d *Device) DeleteUserChannel(c *gin.Context)
  func (d *Device) DeleteUserDevice(c *gin.Context)
  func (d *Device) DeleteUserDeviceFormChannel(c *gin.Context)
  func (d *Device) DeleteUserProduct(c *gin.Context)
  func (d *Device) DeleteUserProductAttribute(c *gin.Context)
  func (d *Device) GetUserAllChannel(c *gin.Context)
  func (d *Device) GetUserAllDevices(c *gin.Context)
  func (d *Device) GetUserAllProducts(c *gin.Context)
  func (d *Device) GetUserChannelAllDevice(c *gin.Context)
  func (d *Device) GetUserDevice(c *gin.Context)
  func (d *Device) GetUserDeviceAllAttributes(c *gin.Context)
  func (d *Device) GetUserDeviceAllAttributesValue(c *gin.Context)
  func (d *Device) GetUserDeviceAttributeValue(c *gin.Context)
  func (d *Device) GetUserOPCUAChannelBrowse(c *gin.Context)
  func (d *Device) GetUserProduct(c *gin.Context)
  func (d *Device) GetUserProductAllAttributes(c *gin.Context)
  func (d *Device) StartScanUserDeviceAttribute(c *gin.Context)
  func (d *Device) StartScanUserProductAttribute(c *gin.Context)
  func (d *Device) StopScanUserDeviceAttribute(c *gin.Context)
  func (d *Device) StopScanUserProductAttribute(c *gin.Context)
  func (d *Device) UpdateDevice(c *gin.Context)
  func (d *Device) UpdateUserChannel(c *gin.Context)
  func (d *Device) UpdateUserDeviceAttributeScanTime(c *gin.Context)
  func (d *Device) UpdateUserProduct(c *gin.Context)
  ```

  可以看到上面不仅打印出了 Device 结构体的信息还有用到了该结构体的所有函数。

### gofmt

gofmt 的作用是格式化 go 程序

- 用法:

  ```
  gofmt [flags] [path ...]
  ```

- flags:

  ```
  -d
    不把重新格式化后的代码打印到输出窗口。
    如果一个文件的格式跟gofmt的不同，就把差别打印到标准输出。
    Do not print reformatted sources to standard output.
    If a file's formatting is different than gofmt's, print diffs
    to standard output.
  -e
      打印所有错误
    Print all (including spurious) errors.
  -l
    不要把重新格式化后的代码打印到输出。如果一个文件的格式跟gofmt的不同，就打印文
    件的名称到标准输出。
    Do not print reformatted sources to standard output.
    If a file's formatting is different from gofmt's, print its name
    to standard output.
  -r rule
    在重新格式化前让重写规则生效
    Apply the rewrite rule to the source before reformatting.
  -s
    在应用重写规则后（如果有），尝试简化代码
    Try to simplify code (after applying the rewrite rule, if any).
  -w
    不要把格式化后的代码打印到标准输出。
    如果文件的格式跟gofmt的不同，则用gofmt的版本覆盖这个文件。
    如果覆盖过程出现异常，原始文件将从备份恢复。
    Do not print reformatted sources to standard output.
    If a file's formatting is different from gofmt's, overwrite it
    with gofmt's version. If an error occurred during overwriting,
    the original file is restored from an automatic backup.
  ```

- debug:

  ```
  -cpuprofile filename
    Write cpu profile to the specified file.
  ```

- 格式:

  使用-r 命令时 gofmt 的重写规则必须是一个下面这种格式的字符串:

  ```
  pattern -> replacement
  ```

上面的 pattern 跟 replacement 都必须是合法的 Go 表达式，在 pattern 中，单字符的小写标识符是匹配任何子表达式的通配符。匹配上的表达式将被以 replacement 中相同的字符替换。

- 例子

  检查文件中不必要的括号:

  ```
  gofmt -r '(a) -> a' -l *.go
  ```

  去掉文件中不必要的括号:

  ```
  gofmt -r '(a) -> a' -w *.go
  ```

  将 slice 中显式的上限转化为隐式的

  ```
  gofmt -r 'α[β:len(α)] -> α[β:]' -w $GOROOT/src
  ```

### golint

linting 在计算机的世界表示自动对代码进行样式和程序性的检查。通常使用 lint 工具来实现这一目标，lint 工具是一个静态代码分析器。linting 的名称源于 Unix 系统中针对 C 语言的一项工具。

go 语言使用 Golint 来检查语法和格式错误,使用方式有两种:

- golint + 文件名
- golint + 文件夹

例子:

- 在项目根目录的终端输入（检查文件):

  ```
  golint ./admin/api/device.go
  ```

- 在项目根目录的终端输入(检查文件夹):

  ```
  golint ./admin/api
  ```

- golint 的常见问题:

  - don't use ALL_CAPS in Go names; use CamelCase

    不能使用下划线命名法，使用驼峰命名法

  - exported function Xxx should have comment or be unexported

    外部可见程序结构体、变量、函数都需要注释

  - var statJsonByte should be statJSONByte ; var taskId should be taskID

    通用名词要求大写
    iD/Id -> ID
    Http -> HTTP
    Json -> JSON
    Url -> URL
    Ip -> IP
    Sql -> SQL

  - don't use an underscore in package name ; don't use MixedCaps in package name ; xxXxx should be xxxxx

    包命名统一小写不使用驼峰和下划线

  - comment on exported type Repo should be of the form "Repo ..." (with optional leading article)

    注释第一个单词要求是注释程序主体的名称，注释可选不是必须的

  - type name will be used as user.UserModel by other packages, and that stutters ; consider calling this Model

    外部可见程序实体不建议再加包名前缀

  - if block ends with a return statement, so drop this else and outdent its block

    if 语句包含 return 时，后续代码不能包含在 else 里面

  - should replace errors.New(fmt.Sprintf(...)) with fmt.Errorf(...)

    errors.New(fmt.Sprintf(…)) 建议写成 fmt.Errorf(…)

  - receiver name should be a reflection of its identity; don't use generic names such as "this" or "self"

    receiver 名称不能为 this 或 self

  - error var SampleError should have name of the form ErrSample

    错误变量命名需以 Err/err 开头

  - should replace num += 1 with num++
    should replace num -= 1 with num--

    a+=1 应该改成 a++，a-=1 应该改成 a–-

### 其他 Go 命令

- env

  打印 go 的环境信息

- fix

  升级包以使用新的 API

- generate

  通过处理代码来生成 go 文件

- list

  列出项目的模块或者包

- run

  编译并且运行 go 程序

- test

  测试包

- tool

  运行指定的 go 工具

- version

  打印 go 的版本

- vet

  vet 用于检查 go 代码的可疑结构，它所提供的错误报告不保证全部正确，但是它能够捕捉编译器检查无法发现的问题。
  通常使用 vet 通过简单地 go vet 命令使用。不带参数则是检查当前目录，带参数比如:

  ```
  go vet my/project/
  ```

  就检查 my 文件夹下的 project 目录

## Go Mod 使用

### 包

每个 Go 程序都是由包构成的。

程序从 main 包开始运行。

按照约定，包名与导入路径的最后一个元素一致。例如，"math/rand" 包中的源码均以 package rand 语句开始。而 math 则是 module 的名称。

模块是一组包的集合，被定义在项目根目录的 go.mod 文件里。go.mod 记录了模块的模块路径，它也是引入包的时候使用的根路径。go 程序中的每一个依赖模块都需要路径以及版本号。

### import 导入

使用下面两种方式都可以导入包，推荐第一种分组的形式

```go
import (
	"fmt"
	"math"
)
```

```go
import "fmt"
import "math"
```

### 导出

go 语言中并没有类似 export 的关键字来导出函数或者变量。

在 Go 中如果一个名字以大写字母开头，那么它就是一个已导出的名字。当你 import 一个包的时候，只能引用这个包中已经导出的名字，也就是大写字母开头的名字。任何未导出的名字在这个包之外无法访问。

### 添加 GO Module

在$GOPATH/src 之外的任何位置创建一个文件夹，接着创建一个简单的 hello.go 程序,以及它的测试程序 hello_test.go（这是测试文件的命名规范），测试代码跟工作代码都属于 package hello 包因此它们在同一个文件夹下。

```go
package hello

func Hello() string {
    return "Hello, world."
}
```

```go
package hello

import "testing"

func TestHello(t *testing.T) {
    want := "Hello, world."
    if got := Hello(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}
```

want 是目标值，got 是实际值，如果实际值不等于目标值则利用 t.Errorf 打印错误

到这个时候我们确实有一个 package 包但是并没有 module,因为没有 go.mod 文件

terminal 执行

```go
go mod init github.com/bo-er/hello
```

可以发现根目录下多了一个 go.mod 文件,它看起来是这样的:

```go
module github.com/bo-er/hello

go 1.15
```

go.mod 只在根目录出现，如果 hello 文件夹下需要增加一个包，那么这个包的地址就是

```go
github.com/bo-er/hello/包名
```

### 给项目添加依赖

改写 Hello.go,引入 rsc.io/quote 包，调用其 Hello 方法

```
package hello

import "rsc.io/quote"

func Hello() string {
    return quote.Hello()
}
```

执行 go test

```
wujiabangs-MacBook-Pro:hello steve$ go test
go: finding module for package rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: found rsc.io/quote in rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
PASS
ok      github.com/bo-er/helloworld/hello       0.350s
```

执行 go test 会自动查找代码中用到的 module，如果遇到包被 import 但是包所在的 module 却不在 go.mod 里面，go 命令会自动查找这个包所在的 module 然后把它添加到 go.mod,如果不指定版本则使用 latest 版本。

执行完命令可以看到 go.mod 文件多了一行:

```go
module github.com/bo-er/helloworld

go 1.15

require rsc.io/quote v1.5.2
```

并且多了一个 go.sum 文件:

go 命令使用 go.sum 文件来确保未来对这些模块的下载(比如别人 clone 了你的项目或者将项目部署到另外一台设备)跟第一次下载获取到的数据一致。所以为了防止依赖被不小心修改或者是被恶意篡改，请把 go.mod 和 go.sum 加到 git 中跟随项目一起提交。

```go
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c h1:qgOY6WgZOaTkIIMiVjBQcw93ERBE4m30iBm00nkL0i8=
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c/go.mod h1:NqM8EUOU14njkJ3fqMW+pc6Ldnwhi/IjpwHt7yyuwOQ=
rsc.io/quote v1.5.2 h1:w5fcysjrx7yqtD/aO+QwRjYZOKnaM9Uh2b40tElTs3Y=
rsc.io/quote v1.5.2/go.mod h1:LzX7hefJvL54yjefDEDHNONDjII0t9xZLPXsUe+TKr0=
rsc.io/sampler v1.3.0 h1:7uVkIFmeBqHfdjD+gZwtXXI+RODJ2Wc4O7MPEh/QiW4=
rsc.io/sampler v1.3.0/go.mod h1:T1hPZKmBbMNahiBKFy5HrXp6adAjACjK9JXDnKaTXpA=

```

可以看到 go.sum 里的内容格式要么是:

```go
<module> <version>/go.mod <hash>
```

要么是:

```
<module> <version> <hash>
```

其中 module 是依赖的路径，version 是依赖的版本号。hash 是以`h1:`开头的字符串，表示生成 checksum 的算法是第一版的 hash 算法（sha256）。checksum 就是“校验和”，用于保证每一个直接与间接依赖的完整性与准确性。

直接依赖就是当前项目直接引入的依赖，间接依赖就是引入的依赖包自己又用到了别的依赖。go.mod 只记录直接依赖。

但是如果你所引入的模块没有列入间接依赖,或者你所引入的模块没有 go.mod 文件，那么当前项目的 go.mod 将会列出这个依赖并且加上一个**//indirect**后缀。

如果想要查看所有直接、间接依赖使用下面的命令

```go
go list -m all
```

将会显示:

```go
github.com/bo-er/helloworld
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/sampler v1.3.0
```

golang.org/x/text`version`v0.0.0-20170915032832-14c0d48ead0c 的版本号 v0.0.0-20170915032832-14c0d48ead0c 是一个假的版本号，是 go 给未打标签的提交所使用的版本语法。

如果是添加本地依赖怎么做？如果本地有两个项目一个项目引入了另外一个项目的 module，并且被引入的项目没有上传到 github 之类的仓库上，那么要使用**replace**,比如替换本地的未上传也没有版本信息的 module（假设 uuid 项目跟 hellowrold 项目在同一个文件夹) ：

```go
module helloworld

go 1.15


require "uuid" v0.0.0
replace "uuid" => "../uuid"
```

它也可以用来替换国内无法访问的库，将 golang.org 上的包用 github 上的包替换，比如:

```go
replace

replace (
	golang.org/x/crypto v0.0.0-20180820150726-614d502a4dac => github.com/golang/crypto v0.0.0-20180820150726-614d502a4dac
	golang.org/x/net v0.0.0-20180821023952-922f4815f713 => github.com/golang/net v0.0.0-20180826012351-8a410e7b638d
	golang.org/x/text v0.3.0 => github.com/golang/text v0.3.0
)

```

### 给项目升级依赖

使用了 go module，版本通过语义版本标签表示。一个语义版本有三个部分：主要版本号，次要版本号，补丁版本号。比如 v0.1.2 表示主要版本是 0，次要版本是 1，补丁版本号是 2。

对于不同的主要版本 go module 使用不同的 path，比如一个是 `rsc.io/quote`, 另一个是 `rsc.io/quote/v2`, 还有一个是 `rsc.io/quote/v3`, 以此类推。这给了模块提供者升级版本的能力。

因此如果主要版本升级，模块提供者要修改模块名为例如 rsc.io/quote/v2,模块使用者也要同时修改自己所依赖的版本，将原来的

```go
import "rsc.io/quote"
```

修改为

```go
import "rsc.io/quote/v3"
```

如果不是主要版本升级就使用 go get 的系列命令:

```go
# 安装最新版本
$ go get golang.org/x/tools/cmd/goimports

# 升级指定模块
$ go get -d golang.org/x/net

# Upgrade modules that provide packages imported by packages in the main module.
# 更新主模块中引入的包的模块(不包括这些包自己所引入的模块)
$ go get -d -u ./...

# 升级或者降级到指定版本
$ go get -d golang.org/x/text@v0.3.2

# 更新（不是升级）到模块的master分支上提交的版本
$ go get -d golang.org/x/text@master

# Remove a dependency on a module and downgrade modules that require it
# 移除一个模块的依赖并且让依赖他的模块降级到一个不再依赖他的版本
$ go get -d golang.org/x/text@none
```

### 移除未使用的依赖

如果仅仅在.go 文件中去掉一个 import,执行下面的命令还是会看到去掉的模块，因此它并没有被真的删除。

```go
go list -m all
```

go build 或者 go test 可以轻易知道缺失的 module 然后将其添加到项目里，但是这两个命令并不知道什么时候 module 可以被安全的删除。这个要交给下面的命令。

```go
go mod tidy
```

go mod tidy 确保了 go.mod 文件中的模块信息跟代码中的模块匹配。它会添加缺失的依赖，删除代码中不使用的依赖。

### go mod vendor

执行下面的命令将在 main 模块的根目录下生成一个名为 vendor 的目录。

当 vendor 被启用时，`go`命令会从 vendor 目录获取依赖，而不是将依赖下载到 module 缓存然后从缓存中获取。

`go mod vendor`命令也会创建一个`vendor/modules.txt`文件，该文件包含已经 vendor 的包的名单以及每个包拷贝自哪里的地址。

另外`go mod vendor`命令会移除 vendor 文件夹如果它在创新建构之前已经存在。

下面的命令会打印 vendor 后的模块跟包的名称到标准错误输出

```go
go mod vendor -v
```

### 其他常用命令:

```go
go mod download    下载所有依赖的module到本地cache（默认为$GOPATH/pkg/mod目录）
go mod edit        编辑go.mod文件
go mod graph       打印模块依赖图
go mod init        初始化当前文件夹, 创建go.mod文件
go mod tidy        增加缺少的module，删除无用的module
go mod verify      检查在module cache中存放的主模块自从下载后没有被修改过。
go mod why         解释为什么需要依赖
```

### go mod 命令使用

#### go mod edit

格式化

因为我们可以手动修改 go.mod 文件，所以有些时候需要格式化该文件。Go 提供了一下命令：

```bash
go mod edit -fmt
```

添加依赖项

```bash
go mod edit -require=golang.org/x/text
```

移除依赖项

如果只是想修改`go.mod`文件中的内容，那么可以运行`go mod edit -droprequire=package path`，比如要在`go.mod`中移除`golang.org/x/text`包，可以使用如下命令：

```bash
go mod edit -droprequire=golang.org/x/text
```

关于`go mod edit`的更多用法可以通过`go help mod edit`查看。

## 开发工具介绍

开发工具: VS CODE

下载地址: https://code.visualstudio.com/Download

VSCODE 默认是英文，如果需要修改语言在 Extensions 中搜索 chinese 点击 install

1. 配置自动保存

2. 配置用户代码段片段 Ctrl/Command+Shift+P 在上方的弹出框中输入 configure user snippets 在弹出框中选择 Go，选择 go 语言后会打开一个 go.json 文件,可以添加一项配置，这样

```json
{
  // Place your snippets for go here. Each snippet is defined under a snippet name and has a prefix, body and
  // description. The prefix is what is used to trigger the snippet and the body will be expanded and inserted. Possible variables are:
  // $1, $2 for tab stops, $0 for the final cursor position, and ${1:label}, ${2:another} for placeholders. Placeholders with the
  // same ids are connected.
  // Example:
  // "Print to console": {
  // 	"prefix": "log",
  // 	"body": [
  // 		"console.log('$1');",
  // 		"$2"
  // 	],
  // 	"description": "Log output to console"
  // }
  "println": {
    "prefix": "pt",
    "body": "fmt.Println($0)",
    "description": "打印"
  }
}
```

### 安装工具包

- 安装 go 插件

  VS CODE 界面的最左侧有一竖条图标，点击最下面的 Extensions（扩展)图标搜索 go,
  然后点击 install(安装)

默认 GoPROXY 配置是：`GOPROXY=https://proxy.golang.org,direct`，由于国内访问不到`https://proxy.golang.org`，所以我们需要换一个 PROXY，这里推荐使用`https://goproxy.io`或`https://goproxy.cn`

可以执行下面的命令修改 GOPROXY：

```bash
go env -w GOPROXY=https://goproxy.cn,direct
```

配置好了 go proxy 再下载官方工具链，否则会下载失败

Windows 平台按下`Ctrl+Shift+P`，Mac 平台按`Command+Shift+P`，这个时候 VS Code 界面会弹出一个输入框

输入框中输入>go:install, VS CODE 将自动搜索指令，选择 Go:Install/Update Tools。然后勾选全部进行安装。

- 验证工具链安装成功
  最简单的做法是打开 terminal 然后执行比如`golint`如果没有任何提示(在 linux 中没有任何提示就是最好的提示 😂)就表示安装成功了。

## 发布程序

### 发布初始版本

对于第一个 go 程序很明显你需要打 v0 标签，也就是初始、不稳定的版本

由于 v0 不保证它的稳定性，因此所有刚开始的项目都必须使用 v0 标签。发布的过程是这样的:

1. 运行 go mod tidy,将包中使用到的 module 添加到 go.mod 并且下载下来，删除不再使用的 module

2. 运行 go test 确保一切 ok

3. 使用 git tag 来给项目打版本标签

4. 将新的 tag 推送到原始仓库

   ```go
   $ go mod tidy
   $ go test ./...
   ok      example.com/hello       0.015s
   $ git add go.mod go.sum hello.go hello_test.go
   $ git commit -m "hello: changes for v0.1.0"
   $ git tag v0.1.0
   $ git push origin v0.1.0
   $
   ```

   这样一来其他项目就可以依赖于 v0.1.0 版本的 example/hello 模块。你可以执行

   ```go
   go list -m example.com/hello@v0.1.0
   ```

   来确保这一点，如果你执行了命令并且使用了 goproxy,需要等待几分钟让代理缓存这个版本的 module

### 发布稳定版本

一旦你确定自己的模块所提供的 API 是稳定的，你可以发布发布 v1.0.0 版本。并且步骤跟上面的一样。

- 拓展阅读

  👁https://golang.org/ 官方文档

  ✓https://www.liwenzhou.com/李文周的博客

- 参考链接

  https://golang.org/dl/

  https://golang.org/ref/mod
