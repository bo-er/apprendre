## 简单的for循环例子


```go
func main() {
	a := []int{1, 2, 3, 4, 5}
	for i := 0; i < len(a); i++ {
	}
}

```
上面的代码执行 `go tool compile -S main.go` 得到的结果:

```
"".main STEXT nosplit size=14 args=0x0 locals=0x0 funcid=0x0
        0x0000 00000 (main.go:3)        TEXT    "".main(SB), NOSPLIT|ABIInternal, $0-0
        0x0000 00000 (main.go:3)        FUNCDATA        $0,  gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)  // FUNCDATA包含了gc所需要的信息
        0x0000 00000 (main.go:3)        FUNCDATA        $1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
        0x0000 00000 (main.go:3)        XORL    AX, AX  // AX XOR AX => AX, 即初始化AX
        0x0002 00002 (main.go:5)        JMP     7 // 无条件jump到 00007指令
        0x0004 00004 (main.go:5)        INCQ    AX //增加AX
        0x0007 00007 (main.go:5)        CMPQ    AX, $5 //比较AX跟值5
        0x000b 00011 (main.go:5)        JLT     4 // 如果AX比值5小则跳转到7
        0x000d 00013 (main.go:5)        RET       // 上面为false则返回
        0x0000 31 c0 eb 03 48 ff c0 48 83 f8 05 7c f7 c3        1...H..H...|..
go.cuinfo.packagename. SDWARFCUINFO dupok size=0
        0x0000 6d 61 69 6e                                      main
""..inittask SNOPTRDATA size=24
        0x0000 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00 00  ................
        0x0010 00 00 00 00 00 00 00 00                          ........
gclocals·33cdeccccebe80329f1fdbee7f5874cb SRODATA dupok size=8
        0x0000 01 00 00 00 00 00 00 00                          ........

```