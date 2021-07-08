# What the heck is setresuid?

When I was reading some code from umc I noticed a system call: `syscall.setresuid()`
What the heck is it?
## Description


 >setresuid() sets the **real user ID**, the **effective user ID**, and the
    **saved set-user-ID** of the calling process.

> An unprivileged process may change its real UID, effective UID,
and saved set-user-ID, each to one of: the current real UID, the
current effective UID or the current saved set-user-ID.

A privileged process (on Linux, one having the CAP_SETUID
capability) may set its real UID, effective UID, and saved set-
user-ID to arbitrary values.

If one of the arguments equals -1, the corresponding value is not
changed.

Regardless of what changes are made to the real UID, effective
UID, and saved set-user-ID, the filesystem UID is always set to
the same value as the (possibly new) effective UID.

Completely analogously, setresgid() sets the real GID, effective
GID, and saved set-group-ID of the calling process (and always
modifies the filesystem GID to be the same as the effective GID),
with the same restrictions for unprivileged processes.


## 简单的背景

每一个进程都有自己的"进程凭证"，这个凭证包含了`PID`,`PPID`,`PGID`,`session ID`以及实际跟有效的用户，用户组ID: `RUID`,`EUID`,`RGID`,`EGID`

那么接下来的目标就是了解上面的概念是什么

## 查看UID 跟GID

```sh
grep $LOGNAME /etc/passwd
```
打印结果:

```
root:x:0:0:root:/root:/bin/bash

```
上面的结果中root就是用户名，UID跟GID都是0

## 查看RUID跟RGID

当我们使用ps aux | grep mysql的时候无论如何至少能够获取到一条进程信息，这个就是