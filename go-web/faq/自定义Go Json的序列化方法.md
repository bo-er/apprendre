[自定义 Go Json 的序列化方法](https://colobu.com/2020/03/19/Custom-JSON-Marshalling-in-Go/)

我们知道，通过 tag,可以有条件地实现定制 Go JSON 序列化的方式，比如`json:",omitempty"`, 当字段的值为空的时候，我们可以在序列化后的数据中不包含这个值，而`json:"-"`可以直接不被 JSON 序列化,如果想被序列化 key`-`，可以设置 tag 为`json:"-,"`,加个逗号。

如果你为类型实现了`MarshalJSON() ([]byte, error)`和`UnmarshalJSON(b []byte) error`方法，那么这个类型在序列化反序列化时将采用你定制的方法。

这些都是我们常用的设置技巧。

如果临时想为一个 struct 增加一个字段的话，可以采用本译文的技巧，临时创建一个类型，通过嵌入原类型的方式来实现。他和[JSON and struct composition in Go](https://attilaolah.eu/2014/09/10/json-and-struct-composition-in-go/)一文中介绍的技巧还不一样(译文和 jsoniter-go 扩展可以阅读陶文的[Golang 中使用 JSON 的一些小技巧](https://zhuanlan.zhihu.com/p/27472716))。`JSON and struct composition in Go`一文中是通过嵌入的方式创建一个新的类型，你序列化和反序列化的时候需要使用这个新类型，而本译文中的方法是无痛改变原类型的`MarshalJSON`方式，采用`Alias`方式避免递归解析，确实是一种非常巧妙的方法。

以下是译文：

Go 的 `encoding/json`序列化`strcut`到 JSON 数据:

```go
package main

import (
	"encoding/json"
	"os"
	"time"
)

type MyUser struct {
	ID       int64     `json:"id"`
	Name     string    `json:"name"`
	LastSeen time.Time `json:"lastSeen"`
}

func main() {
	_ = json.NewEncoder(os.Stdout).Encode(
		&MyUser{1, "Ken", time.Now()},
	)
}
```

序列化的结果:

```
{"id":1,"name":"Ken","lastSeen":"2009-11-10T23:00:00Z"}
```

但是如果我们想改变一个字段的显示结果我们要怎么做呢？例如，我们想把`LastSeen`显示为 unix 时间戳。

最简单的方式是引入另外一个辅助 struct,在`MarshalJSON`中使用它进行正确的格式化：

```go
func (u *MyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID       int64  `json:"id"`
		Name     string `json:"name"`
		LastSeen int64  `json:"lastSeen"`
	}{
		ID:       u.ID,
		Name:     u.Name,
		LastSeen: u.LastSeen.Unix(),
	})
}
```

这样做当然没有问题，但是如果有很多字段的话就会很麻烦，如果我们能把原始 struct 嵌入到新的 struct 中，并让它继承所有不需要改变的字段就太好了:

```go
func (u *MyUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*MyUser
	}{
		LastSeen: u.LastSeen.Unix(),
		MyUser:   u,
	})
}
```

但是等等，问题是这个辅助 struct 也会继承原始 struct 的`MarshalJSON`方法，这会导致这个方法进入无限循环中，最后堆栈溢出。

解决办法就是为原始类型起一个别名，别名会有原始 struct 所有的字段，但是不会继承它的方法：

```go
func (u *MyUser) MarshalJSON() ([]byte, error) {
	type Alias MyUser
	return json.Marshal(&struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		LastSeen: u.LastSeen.Unix(),
		Alias:    (*Alias)(u),
	})
}
```

同样的技术也可以应用于`UnmarshalJSON`方法:

```go
func (u *MyUser) UnmarshalJSON(data []byte) error {
	type Alias MyUser
	aux := &struct {
		LastSeen int64 `json:"lastSeen"`
		*Alias
	}{
		Alias: (*Alias)(u),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}
	u.LastSeen = time.Unix(aux.LastSeen, 0)
	return nil
}
```
