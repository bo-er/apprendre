[原文地址](https://juejin.cn/post/6844904147339198472)

# Protobuf 语法全解析

Protocol Buffers（protobuf）是一种语言无关，平台无关，可扩展的用于序列化结构化数据的方式——类似 XML，但比 XML 更灵活，更高效。虽然平常工作中经常用到 protobuf，但很多时候只是停留在基本语法的使用上，很多高级特性和语法还掌握不全，在阅读一些开源 proto 库的时候，总会看到一些平常没有使用过的语法，影响理解。

本文基于 Go 语言，总结了所有的`proto3`常用和不常用的语法和示例，助你全面掌握 protobuf 语法，加深理解，扫清源码阅读障碍。

# Quick Start

使用 protobuf 语法编写`xxx.proto`文件，然后将其编译成可供特定语言识别和使用的代码文件，供程序调用，这是 protobuf 的基本工作原理。

以 Go 语言为例，使用官方提供的编译器会将`xxx.proto`文件编译成`xxx.pb.go`文件——一个普通的 go 代码文件。
要使用 protobuf，首先我们需要下载 protobuf 编译器——protoc，但 Go 语言并没有被编译器直接支持，而是通过插件的方式被编译器引用，所以同时我们还需要下载 Go 语言的编译插件：

1. 下载合适环境的编译器（`protoc-$VERSION-$PLATFORM.zip`）：[github.com/protocolbuf…](https://github.com/protocolbuffers/protobuf/releases)
2. 下载安装 Go 语言编译插件：`go install google.golang.org/protobuf/cmd/protoc-gen-go`
   安装完毕后，我们准备如下文件`$SRC_DIR/quick_start.proto`:

```proto
syntax = "proto3";

message SearchRequest {
  string query = 1;
  int32 page_number = 2;
  int32 result_per_page = 3;
}

```

执行编译器命令：`protoc --go_out=$DST_DIR $SRC_DIR/quick_start.proto`。 该命令将编译`$SRC_DIR/quick_start.proto`文件，并且将其基于 Go 语言的编译输出结果保存到文件`$DST_DIR/quick_start.qb.go`中：

```go

type SearchRequest struct {
	Query                string   `protobuf:"bytes,1,opt,name=query,proto3" json:"query,omitempty"`
	PageNumber           int32    `protobuf:"varint,2,opt,name=page_number,json=pageNumber,proto3" json:"page_number,omitempty"`
	ResultPerPage        int32    `protobuf:"varint,3,opt,name=result_per_page,json=resultPerPage,proto3" json:"result_per_page,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}


```

在程序中引入生成文件`quick_start.qb.go`所在的包，就可以用 protobuf 的方式对结构体进行序列化和反序列化。
序列化：

```go
req := &pb.SearchRequest{} //此处pb是 quick_start.qb.go 所在包的别名
// ...

// 序列化结构体，写入文件
out, err := proto.Marshal(req)
if err != nil {
        log.Fatalln("Failed to encode search request :", err)
}
if err := ioutil.WriteFile(fname, out, 0644); err != nil {
        log.Fatalln("Failed to write search request:", err)
}

```

反序列化：

```go
// 从文件读取消息，并将其反序列化成结构体
in, err := ioutil.ReadFile(fname)
if err != nil {
        log.Fatalln("Error reading file:", err)
}
book := &pb.SearchRequest{}
if err := proto.Unmarshal(in, book); err != nil {
        log.Fatalln("Failed to parse search request:", err)
}

```

# A Bit of Everything

quick start 示例中展示的是最基础的用法，下面我们通过一个包含所有`proto3`语法的示例，逐一讲解 protobuf 的各项语法和功能。
示例代码在这里可以找到：[a_bit_of_everything.proto](https://github.com/DrmagicE/proto-example/blob/master/a_bit_of_everything.proto)
在代码根目录下执行`protoc --go_out=plugins=grpc:. a_bit_of_everything.proto`生成`xxx.pb.go`文件。

## package

```proto
syntax = "proto3";
option go_package = "examplepb";  // 编译后的golang包名
package example.everything; // proto包名
...

```

在示例文件的起始位置会看到`go_package`和`package`两个关于包的声明，但这两个`package`表达的意义并不相同，`package example.everything;`表明的是当前`.proto`文件所在的包名，跟 Go 语言类似，在相同的包名下，不能定义相同名称的`message`，`enum`或是`service`。 `option go_package = "examplepb"` 则定义了一个文件级别的`option`，用于指定编译后的 golang 包名。

## import

```proto
...
import "google/protobuf/any.proto";
import "google/protobuf/descriptor.proto";
//import "other.proto";
...

import`用于引入其他的proto文件，当在当前文件中要使用其他proto文件的定义时，需要将其`import`进来，然后可以通过类似`packageName.MessageName`的方式来引用需要的内容，跟Go语言的`import`十分类似。执行编译`protoc`的时候，需要加上`-I`参数来指定`import`文件的路径，例如： `protoc -I $GOPATH/src --go_out=. a_bit_of_everything.proto
```

> 示例中引入的 any.proto 和 descriptor.proto 已经内置到 protoc 中，故编译本示例不需要加-I 参数

## 标量类型 （Scalar Value Types）

| proto 类型 | Go 类型 | 备注                                                        |
| ---------- | ------- | ----------------------------------------------------------- |
| double     | float64 |                                                             |
| float      | float   |                                                             |
| int32      | int32   | 编码负数值相对低效                                          |
| int64      | int64   | 编码负数值相对低效                                          |
| uint32     | uint32  |                                                             |
| uint64     | uint64  |                                                             |
| sint32     | int32   | 当值为负数时候，编码比 int32 更高效                         |
| sint64     | int64   | 当值为负数时候，编码比 int64 更高效                         |
| fixed32    | uint32  | 当值总是大于 2^28 时，编码比 uint32 更高效                  |
| fixed64    | uint64  | 当值总是大于 2^56 时，编码比 uint32 更高效                  |
| sfixed32   | int32   |                                                             |
| sfixed64   | int64   |                                                             |
| bool       | bool    |                                                             |
| string     | string  | 只能是 utf-8 编码或者 7-bit ASCII 文本，且长度不得大于 2^32 |
| bytes      | []byte  | 不大于 2^32 的任意长度字节序列                              |

## message 消息

```proto
// 普通的message
message SearchRequest {
    string query = 1;
    int32 page_number = 2;
    int32 result_per_page = 3;
}

```

`message`可以包含多个字段声明，每个字段声明需要包含字段类型，字段名称和一个唯一序号。字段类型可以是标量，枚举或是其他`message`类型。唯一序号用于标识该字段在消息二进制编码中位置。

> 还可以用`repeated`来修饰字段类型，详见下文`repeated`说明。

## 枚举类型

```proto
...
// 枚举 enum
enum Status {
    STATUS_UNSPECIFIED = 0;
    STATUS_OK  = 1;
    STATUS_FAIL= 2;
    STATUS_UNKNOWN = -1; // 不推荐有负数
}
...

```

通过`enum`关键字定义枚举类型，在 protobuf 中，枚举是一个 int32 类型。第一个枚举值必须从 0 开始，如果不希望在代码中使用 0 值，可以将第一个值用`XXX_UNSPECIFIED`作为占位符。由于 enum 类型实际上是用 protobuf 的 int32 类型的编码方式编码，故不推荐在枚举类型中使用负数。

> `XXX_UNSPECIFIED`只是一种代码规范。并不影响代码行为。

## 保留字段 (Reserved Fields) & 保留枚举值(Reserved Values)

```proto
// 保留字段
message ReservedMessage {
    reserved 2, 15, 9 to 11;
    reserved "foo", "bar";
    // string abc = 2;  // 编译报错
    // string foo = 3;  // 编译报错
}
// 保留枚举
enum ReservedEnum {
    reserved 2, 15, 9 to 11, 40 to max;
    reserved "FOO", "BAR";
    // FOO = 0; // 编译报错
    F = 0;
}

```

如果我们将某`message`中的字段删除了，后面更新可能会重新使用这些字段。当新旧两种 proto 定义都在线上运行时，编解码可能会发生错误。例如有新旧两个版本的`Foo`:

```proto
// old version
message Foo {
    string a = 1;
}

// new version
message Foo {
    int32 a = 1;
}

```

如果使用新版本的 proto 来解析旧版本的消息，就会发生错误，因为新版本 proto 会尝试将`a`解析成 int32，但实际上旧版本 proto 是按照 string 类型来对`a`进行编码的。protobuf 通过提供`reserved`关键字来避免新旧版本冲突的问题：

```proto
// new version
message Foo {
    reserved 1; // 标记第一个字段是保留的
    int32 a = 2; // 序号从2开始，就不会与旧版本的string类型a冲突了
}

```

## 嵌套

```proto
// nested 嵌套message
message SearchResponse {
    message Result {
        string url = 1 ;
        string title = 2;
    }
    enum Status {
        UNSPECIFIED = 0;
        OK  = 1;
        FAIL= 2;
    }
    Result results = 1;
    Status status = 2;
}

```

`message`允许多层嵌套，`message`和`enum`都可以嵌套。被嵌套的`message`和`enum`不仅可以在当前`message`中使用，也可以被其他`message`引用：

```proto
message OtherResponse {
    SearchResponse.Result result = 1;
    SearchResponse.Status status = 2;
}

```

## 复合类型

除标量类型外，protobuf 还提供了一些非标量类型，在本文中我把它们统称为复合类型。

> 复合类型并不是官方划分的类别。是本文为了便于理解而归纳总结的一个概念。

### repeated

```proto
// repeated
message RepeatedMessage {
    repeated SearchRequest requests = 1;
    repeated Status status = 2;
    repeated int32 number = 3;
}

```

`repeated`可以作用在`message`中的变量类型上。只有**标量类型**，**枚举类型**和**message 类型**可以被`repeated`修饰。`repeated`表示当前修饰变量可以被重复任意次（包括 0 次），其实就是表示当前修饰类型的一个变长数组，也就是 Go 语言中的`slice`：

```go
// repeated
type RepeatedMessage struct {
	Requests             []*SearchRequest `protobuf:"bytes,1,rep,name=requests,proto3" json:"requests,omitempty"`
	Status               []Status         `protobuf:"varint,2,rep,packed,name=status,proto3,enum=example.everything.Status" json:"status,omitempty"`
	Number               []int32          `protobuf:"varint,3,rep,packed,name=number,proto3" json:"number,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}


```

### map

```proto
message MapMessage{
    map<string, string> message = 1;
    map<string, SearchRequest> request = 2;
}

```

除了`slice`，当然还有`map`。其中 key 的类型可以是**除去`double`,`float`,`bytes`以外**的标量类型，value 的类型可以是任意标量类型，枚举类型和 message 类型。protobuf 的`map`编译成 Go 语言后也是用`map`来表示：

```go
...
// map
type MapMessage struct {
	Message              map[string]string         `protobuf:"bytes,1,rep,name=message,proto3" json:"message,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Request              map[string]*SearchRequest `protobuf:"bytes,2,rep,name=request,proto3" json:"request,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}
...

```

### any

```proto
...
import "google/protobuf/any.proto";
...
message AnyMessage {
    string message = 1;
    google.protobuf.Any details = 2;
}
...

```

`any`类型可以包含一个不需要指定类型的任意的序列化消息。要使用`any`类型，需要`import google/protobuf/any.proto`。`any`类型字段的 encode/decode 交由各语言的运行时各自实现，例如在 Go 语言中可以这样读写`any`类型的字段：

```go
...
import "github.com/golang/protobuf/ptypes"
...
func getSetAny() {
	fmt.Println("getSetAny")
	req := &examplepb.SearchRequest{
	    Query: "query",
	}
	// 将SearchRequest打包成Any类型
	a, err := ptypes.MarshalAny(req)
	if err != nil {
	    log.Println(err)
	    return
	}
	// 赋值
	anyMsg := &examplepb.AnyMessage{
	    Message: "any message",
	    Details: a,
	}

	req = &examplepb.SearchRequest{}
	// 从Any类型中还原proto消息
	err = ptypes.UnmarshalAny(anyMsg.Details, req)
	if err != nil {
	    log.Println(err)
	}
	fmt.Println("any:", req)
}

```

### one of

```proto
// one of
message OneOfMessage {
    oneof test_oneof {
        string m1 = 1;
        int32 m2 =2;
    }
}

```

如果某消息包含多个字段，但这些字段同一时间最多只允许一个被设置时，可以通过`oneof`来保证这样的行为。对`oneof`中任意一个字段设值，都会将其他字段清空。例如对上述的例子，`test_oneof`字段要么是 string 类型的 m1,要么是 int32 类型的 m2。在 Go 语言中读写`oneof`的示例如下：

```go
func getSetOneof() {
	fmt.Println("getSetOneof")
	oneof := &examplepb.OneOfMessage{
		// 同一时间只能设值一个值
		TestOneof: &examplepb.OneOfMessage_M1{
			M1: "this is m1",
		},
	}
	fmt.Println("m1:", oneof.GetM1())  // this is m1
	fmt.Println("m2:", oneof.GetM2()) // 0
}

```

## options & extensions

相信大部的 gopher 在平常使用 protobuf 的过程中都很少关注`options`，80%的开发工作也不需要直接用到`options`。但 options 是一个很有用的功能，其大大提高了 protobuf 的扩展性，我们有必要了解它。`options`其实是 protobuf 内置的一些`message`类型，其分为以下几个级别：

- 文件级别(file-level options)
- 消息级别(message-level options)
- 字段级别(field-level options)
- service 级别(service options)
- method 级别(method options)

protobuf 提供一些内置的`options`可供选择，也提供了通过`extend`关键字来扩展这些`options`，达到增加自定义`options`的目的。

> 在`proto2`语法中，`extend`可以作用于任何`message`，但在`proto3`语法中，`extend`仅能作用于这些定义`option`的`message`——仅用于自定义`option`。

`options`不会改变声明的整体含义（例如声明的是 int32 就是 int32，不会因为一个 option 改变了其声明类型），但可能会影响在特定情况下处理它的方式。例如我们可以使用内置的`deprecated option`将某字段标记为`deprecated`：

```proto
message Msg {
    string foo = 1;
    string bar = 2 [deprecated = true]; //标记为deprecated。
}

```

当我们需要编写自定义 protoc 插件时，可以通过自定义`options`为编译插件提供额外信息。举个例子，假设我要开发一个 proto 的校验插件，其生成`xxx.Validate()`方法来校验消息的合法性，我可以通过自定义`options`来提供生成代码的必要信息：

```proto
message Msg {
    // required是自定义options，表示foo字段必须非空
    string foo = 1; [required = true];
}

```

内置`options`的定义可以在[github.com/protocolbuf…](https://github.com/protocolbuffers/protobuf/blob/master/src/google/protobuf/descriptor.proto)找到，每种级别的`options`都对应一个`message`，分别是：

- FileOptions —— 文件级别
- MessageOptions —— 消息级别
- FieldOptions —— 字段级别
- ServiceOptions —— service 级别
- MethodOptions —— method 级别

以下将通过示例来逐一介绍这些级别的`options`，以及如何扩展这些`options`。

### 文件级别

```proto
...
option go_package = "examplepb";  // 编译后的golang包名
...
message extObj {
    string foo_string= 1;
    int64 bar_int=2;
}
// file options
extend google.protobuf.FileOptions {
    string file_opt_string = 1001;
    extObj file_opt_obj = 1002;
}
option (example.everything.file_opt_string) = "file_options";
option (example.everything.file_opt_obj) = {
    foo_string: "foo"
    bar_int:1
};

```

`go_package` 毫无疑问是 protobuf 内置提供的，用于指定编译后的 golang 包名。除了使用内置的外，可以通过`extend`字段来扩展内置的`FileOptions`，例如在上述例子中，我们新增了两个新的 option——string 类型的`file_opt_string`和 extObj 类型的`file_opt_obj`。并通过`option`关键字设置了两个文件级别的 options。在 Go 语言中，我们可以这样读取这些 options:

```go
func getFileOptions() {
	fmt.Println("file options:")
	msg := &examplepb.MessageOption{}
	md, _ := descriptor.MessageDescriptorProto(msg)
	stringOpt, _ := proto.GetExtension(md.Options, examplepb.E_FileOptString)
	objOpt, _ := proto.GetExtension(md.Options, examplepb.E_FileOptObj)
	fmt.Println("obj.foo_string:", objOpt.(*examplepb.ExtObj).FooString)
	fmt.Println("obj.bar_int", objOpt.(*examplepb.ExtObj).BarInt)
	fmt.Println("string:", *stringOpt.(*string))
}

```

打印结果：

```
file options:
	obj.foo_string: foo
	obj.bar_int 1
	string: file_options

```

### 消息级别

```proto
// message options
extend google.protobuf.MessageOptions {
    string msg_opt_string = 1001;
    extObj msg_opt_obj = 1002;
}
message MessageOption {
    option (example.everything.msg_opt_string) = "Hello world!";
    option (example.everything.msg_opt_obj) = {
        foo_string: "foo"
        bar_int:1
    };
    string foo = 1;
}

```

与文件级别大同小异，不再赘述。Go 语言读取示例：

```go
func getMessageOptions() {
	fmt.Println("message options:")
	msg := &examplepb.MessageOption{}
	_, md := descriptor.MessageDescriptorProto(msg)
	objOpt, _ := proto.GetExtension(md.Options, examplepb.E_MsgOptObj)
	stringOpt, _ := proto.GetExtension(md.Options, examplepb.E_MsgOptString)
	fmt.Println("obj.foo_string:", objOpt.(*examplepb.ExtObj).FooString)
	fmt.Println("obj.bar_int", objOpt.(*examplepb.ExtObj).BarInt)
	fmt.Println("string:", *stringOpt.(*string))
}


```

### 字段级别

```proto
// field options
extend google.protobuf.FieldOptions {
    string field_opt_string = 1001;
    extObj field_opt_obj = 1002;
}
message FieldOption {
    // 自定义的option
    string foo= 1 [(example.everything.field_opt_string) = "abc",(example.everything.field_opt_obj) = {
        foo_string: "foo"
        bar_int:1
    }];
    // protobuf内置的option
    string bar = 2 [deprecated = true];
}

```

字段级别的 option 定义方式不使用`option`关键字，格式为：用[]包裹的用逗号分隔的 k=v 形式的数组。在 Go 语言中，我们可以这样读取这些 option：

```go
func getFieldOptions() {
	fmt.Println("field options:")
	msg := &examplepb.FieldOption{}
	_, md := descriptor.MessageDescriptorProto(msg)
	stringOpt, _ := proto.GetExtension(md.Field[0].Options, examplepb.E_FieldOptString)
	objOpt, _ := proto.GetExtension(md.Field[0].Options, examplepb.E_FieldOptObj)
	fmt.Println("obj.foo_string:", objOpt.(*examplepb.ExtObj).FooString)
	fmt.Println("obj.bar_int", objOpt.(*examplepb.ExtObj).BarInt)
	fmt.Println("string:", *stringOpt.(*string))
}

```

> 应用项目参考：[github.com/mwitkow/go-…](https://github.com/mwitkow/go-proto-validators) go-proto-validators 是一个用于生成可以校验 proto 消息合法性的 proto 编译插件，其使用字段级别的 option 来定义校验规则。

### service 和 method 级别

```proto
// service & method options
extend google.protobuf.ServiceOptions {
    string srv_opt_string = 1001;
    extObj srv_opt_obj = 1002;
}
extend google.protobuf.MethodOptions {
    string method_opt_string = 1001;
    extObj method_opt_obj = 1002;
}
service ServiceOption {
    option (example.everything.srv_opt_string) = "foo";
    rpc Search (SearchRequest) returns (SearchResponse) {
        option (example.everything.method_opt_string) = "foo";
        option (example.everything.method_opt_obj) = {
            foo_string: "foo"
            bar_int: 1
        };
    };
}

```

service 和 method 级别的 option 也是通过`option`关键字来定义，与文件级别和消息级别 option 类似，不再赘述。Go 语言读取示例：

```go
func getServiceOptions() {
	fmt.Println("service options:")
	msg := &examplepb.MessageOption{}
	md, _ := descriptor.MessageDescriptorProto(msg)
	srv := md.Service[1] // ServiceOption
	stringOpt, _ := proto.GetExtension(srv.Options, examplepb.E_SrvOptString)
	fmt.Println("	string:", *stringOpt.(*string))
}
func getMethodOptions() {
	fmt.Println("method options:")
	msg := &examplepb.MessageOption{}
	md, _ := descriptor.MessageDescriptorProto(msg)
	srv := md.Service[1] // ServiceOption
	objOpt, _ := proto.GetExtension(srv.Method[0].Options, examplepb.E_MethodOptObj)
	stringOpt, _ := proto.GetExtension(srv.Method[0].Options, examplepb.E_MethodOptString)
	fmt.Println("	obj.foo_string:", objOpt.(*examplepb.ExtObj).FooString)
	fmt.Println("	obj.bar_int", objOpt.(*examplepb.ExtObj).BarInt)
	fmt.Println("	string:", *stringOpt.(*string))
}

```

> 应用项目参考：[github.com/grpc-ecosys…](https://github.com/grpc-ecosystem/grpc-gateway)
> grpc-gateway 通过为 rpc 的 method 自定义 option，来表达由 grpc 到 http 的转换关系，通过文件级别和 service 级别的 option 来控制生成 swagger 的行为。

# 参考

> [developers.google.cn/protocol-bu…](https://developers.google.cn/protocol-buffers/docs/proto3?hl=en) > [developers.google.cn/protocol-bu…](https://developers.google.cn/protocol-buffers/docs/reference/proto3-spec?hl=en) > [github.com/mwitkow/go-…](https://github.com/mwitkow/go-proto-validators) > [github.com/grpc-ecosys…](
