umask在linux中用于减小`permission bits`,这个动作可以用`unmask`来描述。如果permission已经设为0则不做任何事情。

举个例子，如果你需要`unmask`linux系统的默认权限(文件是666,文件夹是777),那么你可以使用下面的命令:

```
umask 077
```
它的二进制形式是:

```
000 111 111
```

umask的计算方式:  

```
permission = file permission & (!umask)
```

因此如果umask是022那么对于默认创建的文件夹将以下面的方式计算:

```
文件夹权限 111 111 111
umask    000 010 010
!umask   111 101 101
计算结果   111 101 101

```
最后的计算结果是755