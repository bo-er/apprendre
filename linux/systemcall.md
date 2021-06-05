## Where it begins

操作系统可能会有很多种，比如最开始 AT&T 开发的 systemV(five)跟伯克利的 BSD，到现在的各种操作系统比如 OpenBSD, FreeBSD. 不同的操作系统，我们可以根据他们所提供的 System call 方法来区别他们。System call 是操作系统代码中的可以被进程使用`特殊指令`调用的程序。

每一个进程应该都运行在自己的"内存盒子"里，如果它要执行 system call 那么就打破了那个边界。为了实现这一点，在系统调用中需要指定 system call 的内存地址。这些地址存放到了 `System call table` 里。

每一个进程都有一个`stack space`,当进程调用 syscall 的时候 stack space 会产生一块`stack frame`用于存放`syscall frame`。这么设计的理由是

1. 运行 system call 在进程的 context 中执行，避免了 context 切换。这样我们就不需要切换当前进程的`memory tables`了。否则实现会更加困难。

2. 允许 system call 被中断，暂停，恢复。当中断一个进程的时候，基本上不需要关心它是在运行`user code`还是`kernel code`

## Process States

一个进程可以处于`running`的状态，表示它正在被 CPU 处理，或者它可以`waiting`,意味着它等待调度器把它放回`running`的状态。正在运行的进程可以被阻塞。

进程阻塞的可能情况有好几种，最常见的是执行某些 system call,比如读取文件的 system call 就有可能阻塞进程。由于存储文件的硬件跟 CPU 相比速度非常缓慢，当操作系统收到读取文件的 system call 的时候会把进程阻塞，等到文件读取完毕再唤醒进程。

## System Call Types

- processes
- files
- networking sockets
- signals
- inter-process communication

  比如说 network socket

- terminals
- threads
- I/O devices

在调用 system call 的时候，与其当做对操作系统发出命令不如理解为对操作系统发送请求，这意味着调用的进行可能会收不到回应。

## Processes

- address space
- user ids
- file descriptors
- environment
- current and root directory

实际上在进程的 address space 中还专门存储了初始化跟未初始化的数据。他们都是用来存放全局变量的，一个用于存放未初始化的全局变量，另外一个存储初始化后的全局变量。当操作系统加载并且执行可执行文件时，可执行文件决定了这两者需要多大的空间

### Address Space Sections

地址空间会专门有一个 Code Section 用于存放代码，kernel section 则类似存放的是 kernel code.

stack section 刚开始的时候是空的，它会自动的变化。
heap section 则需要分配后才能使用，需要执行分配内存的 System Call. 基本上就是告诉操作系统: "你好，这是我的地址空间，请把它绑定到一个实际的内存空间"。如果你的代码尝试访问一块没有被分配内存的区域，就会导致内存错误。

### Mmap

它的作用是把一些数量的内存页放到地址空间里，并且将这些内存页导向实际内存地址。

- mmap

  "memory map" pages to the process address space
  
  ```
  address = mmap(5000)
  ```
- munmap

  "memory unmap" pages from the process address space

  ```
  // 通过地址取消内存分配
  munmap(address)

  ```

如果内存空间不够的话,mmap的操作会失败。这也是为什么虽然进程结束分配的内存会还给系统，但是仍然要进行内存回收一样，如果运行的是大型或者常驻型的程序如果不返还内存可能会出现内存不够分配的情况。

### Process Creation

fork system call:
当执行fork system call的时候，父母进程的调用结果为孩子进程的**pid**,而孩子进程的调用结果将返回0.
```
if fork() == 0:
    ... // new (child) process
else:
    ... // original (parent) process    
```
unix 系统创建新进程的做法是**复制自己**，这意味着复制**地址空间**，复制其他跟进程相关的比如**用户ID**。复制的说法可能会给人一种创建新进程会带来很大的开销的感觉，这是不靠谱的感觉。对于旧的Unix系统可能是这样，但是对于新版Unix系统，创建新进程只需要复制父母进程的**memory table**

只要fork的进程只读不写，那么我们不需要做任何额外的操作。