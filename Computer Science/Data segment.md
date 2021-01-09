In [computing](https://en.wikipedia.org/wiki/Computing), a **data segment** (often denoted **.data**) is a portion of an [object file](https://en.wikipedia.org/wiki/Object_file) or the corresponding [address space](https://en.wikipedia.org/wiki/Address_space) of a program that contains initialized [static variables](https://en.wikipedia.org/wiki/Static_variable), that is, [global variables](https://en.wikipedia.org/wiki/Global_variable) and [static local variables](https://en.wikipedia.org/wiki/Static_local_variable). The size of this segment is determined by the size of the values in the program's source code, and does not change at [run time](<https://en.wikipedia.org/wiki/Run_time_(program_lifecycle_phase)>).

The data segment is read/write, since the values of variables can be altered at run time. This is in contrast to the _read-only data segment_ (_rodata segment_ or _.rodata_), which contains static constants rather than variables; it also contrasts to the [code segment](https://en.wikipedia.org/wiki/Code_segment), also known as the text segment, which is read-only on many architectures. Uninitialized data, both variables and constants, is instead in the [BSS segment](https://en.wikipedia.org/wiki/BSS_segment).

Historically, to be able to support memory address spaces larger than the native size of the internal address register would allow, early CPUs implemented a system of segmentation whereby they would store a small set of indexes to use as offsets to certain areas. The [Intel 8086](https://en.wikipedia.org/wiki/Intel_8086) family of CPUs provided four segments: the code segment, the data segment, the stack segment and the extra segment. Each segment was placed at a specific location in memory by the software being executed and all instructions that operated on the data within those segments were performed relative to the start of that segment. This allowed a 16-bit address register, which would normally be able to access 64 KB of memory space, to access 1 MB of memory space.

This segmenting of the memory space into discrete blocks with specific tasks carried over into the programming languages of the day and the concept is still widely in use within modern programming languages.

## Program memory

A computer program memory can be largely categorized into two sections: read-only and read/write. This distinction grew from early systems holding their main program in [read-only memory](https://en.wikipedia.org/wiki/Read-only_memory) such as [Mask ROM](https://en.wikipedia.org/wiki/Mask_ROM), [PROM](https://en.wikipedia.org/wiki/Programmable_read-only_memory) or [EEPROM](https://en.wikipedia.org/wiki/EEPROM). As systems became more complex and programs were loaded from other media into RAM instead of executing from ROM, the idea that some portions of the program's memory should not be modified was retained. These became the _.text_ and _.rodata_ segments of the program, and the remainder which could be written to divided into a number of other segments for specific tasks.

### Text

The **code segment**, also known as a **text segment** or simply as **text**, is where a portion of an [object file](https://en.wikipedia.org/wiki/Object_file) or the corresponding section of the program's [address space](https://en.wikipedia.org/wiki/Address_space) that contains [executable](https://en.wikipedia.org/wiki/Executable) instructions is stored and is generally read-only and fixed size.

### Data

[![img](https://upload.wikimedia.org/wikipedia/commons/thumb/5/50/Program_memory_layout.pdf/page1-149px-Program_memory_layout.pdf.jpg)](https://en.wikipedia.org/wiki/File:Program_memory_layout.pdf)

This shows the typical layout of a simple computer's program memory with the text, various data, and stack and heap sections.

The _.data_ segment contains any global or static variables which have a pre-defined value and can be modified. That is any variables that are not defined within a function (and thus can be accessed from anywhere) or are defined in a function but are defined as _static_ so they retain their address across subsequent calls. Examples, in C, include:

```
   int val = 3;
   char string[] = "Hello World";
```

The values for these variables are initially stored within the read-only memory (typically within _.text_) and are copied into the _.data_ segment during the start-up routine of the program.

Note that in the above example, if these variables had been declared from within a function, they would default to being stored in the local stack frame.

### BSS

Main article: [BSS segment](https://en.wikipedia.org/wiki/BSS_segment)

The BSS segment, also known as _uninitialized data_, is usually adjacent to the data segment. The BSS segment contains all global variables and static variables that are initialized to zero or do not have explicit initialization in source code. For instance, a variable defined as `static int i;` would be contained in the BSS segment.

### Heap

The heap area commonly begins at the end of the .bss and .data segments and grows to larger addresses from there. The heap area is managed by [malloc](https://en.wikipedia.org/wiki/Malloc), calloc, realloc, and free, which may use the [brk](https://en.wikipedia.org/wiki/Sbrk) and [sbrk](https://en.wikipedia.org/wiki/Sbrk) system calls to adjust its size (note that the use of brk/sbrk and a single "heap area" is not required to fulfill the contract of malloc/calloc/realloc/free; they may also be implemented using [mmap](https://en.wikipedia.org/wiki/Mmap)/munmap to reserve/unreserve potentially non-contiguous regions of virtual memory into the process' [virtual address space](https://en.wikipedia.org/wiki/Virtual_address_space)). The heap area is shared by all threads, shared libraries, and dynamically loaded modules in a process.

### Stack

Main article: [Call stack](https://en.wikipedia.org/wiki/Call_stack)

The stack area contains the program [stack](<https://en.wikipedia.org/wiki/Stack_(data_structure)>), a [LIFO](<https://en.wikipedia.org/wiki/LIFO_(computing)>) structure, typically located in the higher parts of memory. A "stack pointer" register tracks the top of the stack; it is adjusted each time a value is "pushed" onto the stack. The set of values pushed for one function call is termed a "stack frame". A stack frame consists at minimum of a return address. [Automatic variables](https://en.wikipedia.org/wiki/Automatic_variable) are also allocated on the stack.

The stack area traditionally adjoined the heap area and they grew towards each other; when the stack pointer met the heap pointer, free memory was exhausted. With large address spaces and virtual memory techniques they tend to be placed more freely, but they still typically grow in a converging direction. On the standard PC [x86 architecture](https://en.wikipedia.org/wiki/X86_architecture) the stack grows toward address zero, meaning that more recent items, deeper in the call chain, are at numerically lower addresses and closer to the heap. On some other architectures it grows the opposite direction.

## Interpreted languages

Some interpreted languages offer a similar facility to the data segment, notably [Perl](https://en.wikipedia.org/wiki/Perl)[[1\]](https://en.wikipedia.org/wiki/Data_segment#cite_note-1) and [Ruby](<https://en.wikipedia.org/wiki/Ruby_(programming_language)>).[[2\]](https://en.wikipedia.org/wiki/Data_segment#cite_note-2) In these languages, including the line `__DATA__` (Perl) or `__END__` (Ruby, old Perl) marks the end of the code segment and the start of the data segment. Only the contents prior to this line are executed, and the contents of the source file after this line are available as a file object: `PACKAGE::DATA` in Perl (e.g., `main::DATA`) and `DATA` in Ruby. This can be considered a form of [here document](https://en.wikipedia.org/wiki/Here_document) (a file literal).

## See also

- [Segmentation (memory)](<https://en.wikipedia.org/wiki/Segmentation_(memory)>)
- [Segmentation fault](https://en.wikipedia.org/wiki/Segmentation_fault)
- [Linker (computing)](<https://en.wikipedia.org/wiki/Linker_(computing)>)
- [Code segment](https://en.wikipedia.org/wiki/Code_segment)
- [.bss](https://en.wikipedia.org/wiki/.bss)
- [Uninitialized variable](https://en.wikipedia.org/wiki/Uninitialized_variable)
- [Stack (abstract data type)](<https://en.wikipedia.org/wiki/Stack_(abstract_data_type)>)
- [Process control block](https://en.wikipedia.org/wiki/Process_control_block)

## References

1. **[^](https://en.wikipedia.org/wiki/Data_segment#cite_ref-1)** [perldata: Special Literals](http://perldoc.perl.org/perldata.html#Special-Literals)
2. **[^](https://en.wikipedia.org/wiki/Data_segment#cite_ref-2)** Ruby: Object: [**END**](http://ruby-doc.org/docs/keywords/1.9/Object.html#method-i-__END__)

## External links

- ["C startup"](http://www.bravegnu.org/gnu-eprog/c-startup.html). _bravegnu.org_.
- ["mem_sequence.c - sequentially lists memory regions in a process"](https://web.archive.org/web/20090202113414/http://blog.ooz.ie/2008/09/0x03-notes-on-assembly-memory-from.html). Archived from [the original](http://blog.ooz.ie/2008/09/0x03-notes-on-assembly-memory-from.html) on 2009-02-02.
- van der Linden, Peter (1997). [*Expert C Programming: Deep C Secrets*](http://www.electroons.com/8051/ebooks/expert C programming.pdf) (PDF). Prentice Hall. pp. 119ff.
