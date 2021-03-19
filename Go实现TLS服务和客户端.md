# [来自鸟窝](https://colobu.com/2016/06/07/simple-golang-tls-examples/)

传输层安全协议（Transport Layer Security，缩写：TLS），及其前身安全套接层（Secure Sockets Layer，缩写：SSL）是一种安全协议，目的是为互联网通信提供安全及数据完整性保障。

SSL包含记录层（Record  Layer）和传输层，记录层协议确定了传输层数据的封装格式。传输层安全协议使用X.509认证，之后利用非对称加密演算来对通信方做身份认证，之后交换对称密钥作为会谈密钥（Session key）。这个会谈密钥是用来将通信两方交换的数据做加密，保证两个应用间通信的保密性和可靠性，使客户与服务器应用之间的通信不被攻击者窃听。

本文并没有提供一个TLS的深度教程，而是提供了两个Go应用TLS的简单例子，用来演示使用Go语言快速开发安全网络传输的程序。



### TLS历史

> 1994年早期，NetScape公司设计了SSL协议（Secure Sockets Layer）的1.0版，但是未发布。
> 1994年11月，NetScape公司发布SSL 2.0版，很快发现有严重漏洞。
> 1996年11月，SSL 3.0版问世，得到大规模应用。
> 1999年1月，互联网标准化组织ISOC接替NetScape公司，发布了SSL的升级版[TLS 1.0版](https://www.ietf.org/rfc/rfc2246.txt)。
> 2006年4月和2008年8月，TLS进行了两次升级，分别为[TLS 1.1](https://tools.ietf.org/html/rfc4346)版和[TLS 1.2](https://tools.ietf.org/html/rfc5246)版。最新的变动是2011年TLS 1.2的修订版。
> 现在正在制定 [tls 1.3](https://github.com/tlswg/tls13-spec)。

### 证书生成

首先我们创建私钥和证书。

#### 服务器端的证书生成

使用了"服务端证书"可以确保服务器不是假冒的。

1、 生成服务器端的私钥

```
openssl genrsa -out server.key 2048
```

2、 生成服务器端证书

```
openssl req -new -x509 -key server.key -out server.pem -days 3650
```

or

```
go run $GOROOT/src/crypto/tls/generate_cert.go --host localhost
```

#### 客户端的证书生成

除了"服务端证书"，在某些场合中还会涉及到"客户端证书"。所谓的"客户端证书"就是用来证明客户端访问者的身份。
比如在某些金融公司的内网，你的电脑上必须部署"客户端证书"，才能打开重要服务器的页面。
我会在后面的例子中演示"客户端证书"的使用。

3、 生成客户端的私钥

```
openssl genrsa -out client.key 2048
```

4、 生成客户端的证书

```
openssl req -new -x509 -key client.key -out client.pem -days 3650
```

或者使用下面的脚本：

```
#!/bin/bash
# call this script with an email address (valid or not).
# like:
# ./makecert.sh demo@random.com
mkdir certs
rm certs/*
echo "make server cert"
openssl req -new -nodes -x509 -out certs/server.pem -keyout certs/server.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1"
echo "make client cert"
openssl req -new -nodes -x509 -out certs/client.pem -keyout certs/client.key -days 3650 -subj "/C=DE/ST=NRW/L=Earth/O=Random Company/OU=IT/CN=www.random.com/emailAddress=$1"
```

### Golang 例子

Go [Package tls](https://golang.org/pkg/crypto/tls/)部分实现了 tls 1.2的功能，可以满足我们日常的应用。[Package crypto/x509](https://golang.org/pkg/crypto/x509/)提供了证书管理的相关操作。

#### 服务器证书的使用

本节代码提供了服务器使用证书的例子。

下面的代码是服务器的例子：

```go
package main

import (
	"bufio"
	"crypto/tls"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	config := &tls.Config{Certificates: []tls.Certificate{cert}}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
```

首先从上面我们创建的服务器私钥和pem文件中得到证书`cert`，并且生成一个tls.Config对象。这个对象有多个字段可以设置，本例中我们使用它的默认值。
然后用`tls.Listen`开始监听客户端的连接，accept后得到一个net.Conn，后续处理和普通的TCP程序一样。

然后，我们看看客户端是如何实现的：

```go
package main

import (
	"crypto/tls"
	"log"
)

func main() {
	conf := &tls.Config{
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}
```

`InsecureSkipVerify`用来控制客户端是否证书和服务器主机名。如果设置为true,则不会校验证书以及证书中的主机名和服务器主机名是否一致。
因为在我们的例子中使用自签名的证书，所以设置它为true,仅仅用于测试目的。

可以看到，整个的程序编写和普通的TCP程序的编写差不太多，只不过初始需要做一些TLS的配置。

你可以`go run server.go`和`go run client.go`测试这个例子。

#### 客户端证书的使用

在有的情况下，需要双向认证，服务器也需要验证客户端的真实性。在这种情况下，我们需要服务器和客户端进行一点额外的配置。

服务器端：

```go
package main

import (
	"bufio"
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
	"net"
)

func main() {
	cert, err := tls.LoadX509KeyPair("server.pem", "server.key")
	if err != nil {
		log.Println(err)
		return
	}

	certBytes, err := ioutil.ReadFile("client.pem")
	if err != nil {
		panic("Unable to read cert.pem")
	}
	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("failed to parse root certificate")
	}

	config := &tls.Config{
		Certificates: []tls.Certificate{cert},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		ClientCAs:    clientCertPool,
	}
	ln, err := tls.Listen("tcp", ":443", config)
	if err != nil {
		log.Println(err)
		return
	}
	defer ln.Close()

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	r := bufio.NewReader(conn)
	for {
		msg, err := r.ReadString('\n')
		if err != nil {
			log.Println(err)
			return
		}

		println(msg)

		n, err := conn.Write([]byte("world\n"))
		if err != nil {
			log.Println(n, err)
			return
		}
	}
}
```

因为需要验证客户端，我们需要额外配置下面两个字段：

```
ClientAuth:   tls.RequireAndVerifyClientCert,
ClientCAs:    clientCertPool,
```

然后客户端也配置这个`clientCertPool`:

```go
package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

func main() {
	cert, err := tls.LoadX509KeyPair("client.pem", "client.key")
	if err != nil {
		log.Println(err)
		return
	}

	certBytes, err := ioutil.ReadFile("client.pem")
	if err != nil {
		panic("Unable to read cert.pem")
	}

	clientCertPool := x509.NewCertPool()
	ok := clientCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("failed to parse root certificate")
	}

	conf := &tls.Config{
		RootCAs:            clientCertPool,
		Certificates:       []tls.Certificate{cert},
		InsecureSkipVerify: true,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:443", conf)
	if err != nil {
		log.Println(err)
		return
	}
	defer conn.Close()

	n, err := conn.Write([]byte("hello\n"))
	if err != nil {
		log.Println(n, err)
		return
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
		return
	}

	println(string(buf[:n]))
}
```

运行这两个代码`go run server2.go`和`go run client2.go`,可以看到两者可以正常的通讯，如果用前面的客户端`go run client.go`，不能正常通讯，因为前面的客户端并没有提供客户端证书。

> **更正** 使用自定义的CA的例子可以参考 https://github.com/golang/net/tree/master/http2/h2demo

```
Make CA:
$ openssl genrsa -out rootCA.key 2048
$ openssl req -x509 -new -nodes -key rootCA.key -days 1024 -out rootCA.pem
... install that to Firefox

Make cert:
$ openssl genrsa -out server.key 2048
$ openssl req -new -key server.key -out server.csr
$ openssl x509 -req -in server.csr -CA rootCA.pem -CAkey rootCA.key -CAcreateserial -out server.crt -days 500
```