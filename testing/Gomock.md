- 如果一个testCase里需要调用某个Mock方法，第一次需要返回1第二次需要返回2，可以这么做:

```
m.EXPECT().Version().Times(1).Return(byte(4))
m.EXPECT().Version().Times(1).Return(byte(5))
```
`指定调用次数`之后，就能实现调用同样的方法第一次返回byte(4)第二次返回byte(5)