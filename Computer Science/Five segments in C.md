For those future visitors who may be interested in knowing about those memory segments, I am writing important points about 5 memory segments in C:

**Some heads up:**

1. Whenever a C program is executed some memory is allocated in the RAM for the program execution. This memory is used for storing the frequently executed code (binary data), program variables, etc. The below memory segments talks about the same:
2. Typically there are three types of variables:
   - Local variables (also called as automatic variables in C)
   - Global variables
   - Static variables
   - You can have global static or local static variables, but the above three are the parent types.

**5 Memory Segments in C:**

# 1. Code Segment

- The code segment, also referred as the text segment, is the area of memory which contains the frequently executed code.
- The code segment is often read-only to avoid risk of getting overridden by programming bugs like buffer-overflow, etc.
- The code segment does not contain program variables like local variable (_also called as automatic variables in C_), global variables, etc.
- Based on the C implementation, the code segment can also contain read-only string literals. For example, when you do `printf("Hello, world")` then string "Hello, world" gets created in the code/text segment. You can verify this using `size` command in Linux OS.
- [Further reading](https://en.wikipedia.org/wiki/Code_segment)

# Data Segment

The data segment is divided in the below two parts and typically lies below the heap area or in some implementations above the stack, but the data segment never lies between the heap and stack area.

## 2. Uninitialized data segment

- This segment is also known as **bss**.
- This is the portion of memory which contains:
  1. **Uninitialized global variables** **\*(including pointer variables)\***
  2. **Uninitialized constant global variables**.
  3. **Uninitialized local static variables**.
- Any global or static local variable which is not initialized will be stored in the uninitialized data segment
- For example: global variable `int globalVar;` or static local variable `static int localStatic;` will be stored in the uninitialized data segment.
- If you declare a global variable and initialize it as `0` or `NULL` then still it would go to uninitialized data segment or bss.
- [Further reading](https://en.wikipedia.org/wiki/.bss)

## 3. Initialized data segment

- This segment stores:
  1. **Initialized global variables** **\*(including pointer variables)\***
  2. **Initialized constant global variables**.
  3. **Initialized local static variables**.
- For example: global variable `int globalVar = 1;` or static local variable `static int localStatic = 1;` will be stored in initialized data segment.
- This segment can be **further classified into initialized read-only area and initialized read-write area**. _Initialized constant global variables will go in the initialized read-only area while variables whose values can be modified at runtime will go in the initialized read-write area_.
- **\*The size of this segment is determined by the size of the values in the program's source code, and does not change at run time\***.
- [Further reading](https://en.wikipedia.org/wiki/Data_segment)

# 4. Stack Segment

- Stack segment is used to store variables which are created inside functions (

  function could be main function or user-defined function

  ), variable like

  1. **Local variables** of the function **\*(including pointer variables)\***
  2. **Arguments passed to function**
  3. **Return address**

- Variables stored in the stack will be removed as soon as the function execution finishes.

- [Further reading](https://en.wikipedia.org/wiki/Stack-based_memory_allocation)

# 5. Heap Segment

- This segment is to support dynamic memory allocation. If the programmer wants to allocate some memory dynamically then in C it is done using the `malloc`, `calloc`, or `realloc` methods.
- For example, when `int* prt = malloc(sizeof(int) * 2)` then eight bytes will be allocated in heap and memory address of that location will be returned and stored in `ptr` variable. The `ptr` variable will be on either the stack or data segment depending on the way it is declared/used.
- [Further reading](https://en.wikipedia.org/wiki/Memory_management#HEAP)
