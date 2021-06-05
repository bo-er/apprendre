## 监听tcp并且处理进入的请求

可以看到两个方法一个是监听TCP连接，另外一个是处理TCP连接

```go
type server struct{
    tcpListener net.TCPListener
}

func (s *server) tcpListen(){
    for {
		conn, err := s.tcpListener.AcceptTCP()
		if err != nil {
			if neterr, ok := err.(net.Error); ok && !neterr.Temporary() {
				break
			}
			log.Printf("[ERR] Error accepting TCP connection: %s", err)
		}
		go s.handleConn(conn)
	}
}

func (s *server) handleConn(conn *net.TCPConn){

}

```