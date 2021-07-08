### What are OSes? / Logistics

- What is an operating system
- Dual-mode operation; protection; exceptions
- process VMs
- Process = thread + address space
- the Unix/Monolithic kernel design
- xv6 and Unix
- System calls (in xv6)
- logistics

### Multiprogramming and Dual-Mode Operation

- System calls in xv6 (review)
- Other exception handling
- Context switches generally
- Context switches in xv6 (start)

### The Unix API 1

- Context switches in xv6 (finish)
- Process control blocks
- POSIX versus Unix
- Process creation and management: fork(), exec*(), wait()

### The Unix API 2

- Shells, generally
- Shell assignment
- Unix: everything is a file, stdio, stdout
- stdio.h versus system calls
- POSIX file API: pipe(), open(), read(), write(), close()

### The Unix API 3 /  Scheduling 1

- POSIX API: pipe() (finish)
- xv6 creating the first process
- threads versus processes
- queues of processes and schedulers    
  - aside: non-CPU scheduling
- alternating I/O and CPU bursts
- the process state machine

### Scheduling 2

- scheduling metrics: fairness, response time, throughput
- FCFS, RR
- priority scheduling
- SJF, SRTF

### Scheduling 3

- approximating SJF: multi-level feedback scheduling
- proportional share scheduling    
  - lottery scheduling
- Linux’s Completely Fair Scheduler
- (if time) real-time scheduling

### Threads / Synchronization 1: Locks 1

- pthreads API — pthread_create, pthread_join
- some pthreads examples
- bank account synchronization example / lost write
- race conditions
- building locks is tricky: too much milk
- mutual exclusion / critical sections
- locks
- aside: standard container rules
- disabling interrupts for locks

### Synchronization 2: Locks 2, Mutexes

- disabling interrupts for locks, continued
- load/store reordering
- cache coherency, MSI and snooping
- atomic operations: test-and-set, CAS
- xv6’s spinlock
- test-and-test-and-set
- false sharing
- avoiding busy-waits: mutexes

### Synchronization 3: Monitors and Semaphores

- barriers
- monitors    
  - producer/consumer with monitors
- intuition for using monitors
- monitor exercises (just started)

### Synchronization 4: Semaphores and Monitors con't / Reader+Writer Locks

- monitor exercises (continued)
- counting semaphores    
  - producer/consumer with counting semaphores
- relating monitors and counting semaphores
- reader/writers problem    
  - reader/writers with monitors
  - reader or writer priority
- POSIX rwlocks

### Synchronization 5: Reader+Writer con't / Deadlock

- reader/writers problem (con’t)
- deadlock    
  - definition: resources and conditions
  - examples
- deadlock prevention
- deadlock detection and resource acquisition graphs

### Alternatives to Threads / Virtual Memory 1

- (briefly) event-based programming/message passing
- review(???): what is virtual memory?
- xv6 paging (start)

### Virtual Memory 2: page table tricks

- xv6 paging (con’t)
- page table tricks    
  - allocate on demand
  - copy-on-write
- (if time) demand paging
- (if time) the page cache, high-level

### Virtual Memory 3: mmap, page cache (intro)

- memory mapped files, POSIX mmap    
  - `/proc/$$/maps`
- demand paging
- the page cache
- Linux process memory map data structures
- page replacement algorithms (ideal)    
  - ideal: Belady’s MIN
  - ideal possible: LRU
- working set model
- accessed bits  and simulating them
- dirty bits and simulating them

### Virtual 4 / I/O 0

- accessed bits
- page replacement algorithms (possible)    
  - second chance
  - SEQ
  - clock algorithm and variants
- non-LRU replacement    
  - handling scanning and readahead
  - (if time) fairness and page replacement
- start device files

### I/O 1 / Filesystems 1

- continue device files
- device drivers and interrupts    
  - “top” and “bottom” halves
  - device interface
  - devices as magic memory
  - aside on I/O space
  - direct memory access
- disk interface: sectors
- FAT: linked lists on disk (start)

### Filesystems 2

- FAT: linked lists on disk (con’t)
- hard disks
- SSDs
- the inode concept (start)

### Filesystems 3

- the inode concept (con’t)
- double- and triply-indirect blocks
- sparse files
- hard links and symbolic links
- locality: block groups
- extents
- trees on disk

### Filesystems 4

- ordering writes carefully
- fsck, etc.
- write-ahead logging and journaling filesystems
- mirroring
- snapshots

### Sockets / Distributed Systems 1

- mounting filesystems
- reasons for distribution
- DNS, IP addresses, port numbers, connections
- sockets    
  - getaddrinfo
  - POSIX API: socket, bind, listen, accept, connect
- supplemental references:    
  - [Wikipedia’s Berkeley Sockets article](https://en.wikipedia.org/wiki/Berkeley_sockets)
  - [echo-server example](https://www.cs.virginia.edu/~cr4bd/4414/F2019/files/echo-server.cc), [echo-client example](https://www.cs.virginia.edu/~cr4bd/4414/F2019/files/echo-client.cc)

### Distributed Systems 2: RPC / Failure

- remote procedure calls
- fail-stop
- two-phase commit (start)

### Distributed Systems 3: Failure / Network filesystems

- two-phase commit
- two-phase assignment
- very briefly: distributed consensus
- network filesystems
- stateful versus stateless
- caching in network filesystems    
  - open-to-close consistency
  - callbacks

### Access Control

- access control lists
- user IDs on Unix
- /bin/login
- set-user-ID
- (briefly) capabilities

### Virtual Machines

- trap and emulate
- reflecting exceptions
- software support for virtualized virtual memory
- (if time) hardware virtualization support