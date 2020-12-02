### 基本用法：

```go
package main

import (
    "fmt"
    "net/http"
)

type HelloHandler struct{}
func (h HelloHandler) ServeHTTP (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello Handler!")
}

func hello (w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello!")
}

func main() {
    server := http.Server{
        Addr: "127.0.0.1:8080",
    }
    helloHandler := HelloHandler{}
    http.Handle("/hello1", helloHandler)
    http.HandleFunc("/hello2", hello)
    server.ListenAndServe()
}
```

### http.Handel

首先，简单分析一下 `http.Handle(pattern string, handler Handler)`，`http.Handle(pattern string, handler Handler)` 接收两个参数，一个是路由匹配的字符串，另外一个是 `Handler` 类型的值：

```
func Handle(pattern string, handler Handler) { 
	//实际上是调用了下面的函数
	DefaultServeMux.Handle(pattern, handler) 
}
```

### Handler

实际上它是一个接口

```
type Handler interface {
	ServeHTTP(ResponseWriter, *Request)
}
```

### http.HandleFunc

该方法接收两个参数，一个是路由匹配的字符串，另外一个是函数func(ResponseWriter, *Request)` 类型的函数：

```
func HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    DefaultServeMux.HandleFunc(pattern, handler)
}

//然后继续调用DefaultServeMux.HandleFunc(pattern, handler)
func (mux *ServeMux) HandleFunc(pattern string, handler func(ResponseWriter, *Request)) {
    mux.Handle(pattern, HandlerFunc(handler))
}
//而实际上HandlerFunc是这样定义的：
type HandlerFunc func(ResponseWriter, *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
    f(w, r)
}
```

也就是说其实HandlerFunc就是一个Handler