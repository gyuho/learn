[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Segmentation fault, debug

- [Reference](#reference)
- [segmentation fault](#segmentation-fault)
- [stack overflow](#stack-overflow)

[↑ top](#segmentation-fault-debug)
<br><br><br><br>
<hr>







#### Reference

- [Segmentation fault](https://en.wikipedia.org/wiki/Segmentation_fault)
- [Stack overflow](https://en.wikipedia.org/wiki/Stack_overflow)
- [Troubleshooting Segmentation Violations/Faults](http://web.mit.edu/10.001/Web/Tips/tips_on_segmentation.html)
- [Debugging Go Code with GDB](https://golang.org/doc/gdb)
- [Go has a debugger—and it's awesome!](https://blog.cloudflare.com/go-has-a-debugger-and-its-awesome/)
- [derekparker/delve](https://github.com/derekparker/delve)
- [mailgun/godebug](https://github.com/mailgun/godebug)

[↑ top](#segmentation-fault-debug)
<br><br><br><br>
<hr>






#### segmentation fault

> Segmentation faults have various causes, and are a common problem in programs
> written in the C programming language, where they arise primarily due to
> errors in use of pointers for virtual memory addressing, particularly illegal
> access. Another type of memory access error is a bus error, which also has
> various causes, but is today much rarer; these occur primarily due to
> incorrect physical memory addressing, or due to misaligned memory access –
> these are memory references that the hardware cannot address, rather than
> references that a process is not allowed to address.
>
> A segmentation fault occurs when a program attempts to access a memory
> location that it is not allowed to access, or attempts to access a memory
> location in a way that is not allowed (for example, attempting to write to a
> read-only location, or to overwrite part of the operating system).
>
> [*Segmentation fault*](https://en.wikipedia.org/wiki/Segmentation_fault) *by
> Wikipeida*

<br>
Then let's create some segmentation fault in Go:

```go
package main

func main() {
	var ptr *int = nil
	*ptr = 100
}

/*
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xb code=0x1 addr=0x0 pc=0x400c02]
*/

```

Let's debug this with `gdb`:

```bash
$ gdb ./00_segmentation_fault

(gdb) run
Starting program: /home/ubuntu/go/src/github.com/gyuho/learn/doc/segmentation_fault_debug/code/00_segmentation_fault 

Program received signal SIGSEGV, Segmentation fault.
main.main ()
    at /home/ubuntu/go/src/github.com/gyuho/learn/doc/segmentation_fault_debug/code/00_segmentation_fault.go:5
5		*ptr = 100
(gdb)

```

<br>
In C++:

```cpp
#include <iostream>
using namespace std;

int main()
{
	int *ptr = NULL;
	*ptr = 100; // Write to invalid memory address
	// Segmentation fault (core dumped)
}

```

Let's debug this with `gdb`:

```bash
$ gdb ./a.out

(gdb) run
Starting program: /home/ubuntu/go/src/github.com/gyuho/learn/doc/segmentation_fault_debug/code/a.out 
Traceback (most recent call last):
  File "/usr/share/gdb/auto-load/usr/lib/x86_64-linux-gnu/libstdc++.so.6.0.19-gdb.py", line 63, in <module>
    from libstdcxx.v6.printers import register_libstdcxx_printers
ImportError: No module named 'libstdcxx'

Program received signal SIGSEGV, Segmentation fault.
0x00000000004006dd in main ()
(gdb) 

```

[↑ top](#segmentation-fault-debug)
<br><br><br><br>
<hr>









#### stack overflow

> In software, a stack overflow occurs if the stack pointer exceeds the stack
> bound. The call stack may consist of a limited amount of address space, often
> determined at the start of the program. The size of the call stack depends on
> many factors, including the programming language, machine architecture,
> multi-threading, and amount of available memory. 
>
> [*Stack overflow*](https://en.wikipedia.org/wiki/Stack_overflow) *by Wikipedia*

<br>
Here's an example of stack overflow in Go:

```go
package main

func f() {
	g()
}

func g() {
	f()
}

func main() {
	f()
	/*
	   runtime: goroutine stack exceeds 1000000000-byte limit
	   fatal error: stack overflow

	   ...
	*/
}

```

<br>
And in C++:

```cpp
#include <iostream>

int f();
int g();

int f(){
	g();
}

int g() {
	f();  
}

int main()
{
	f(); // stack overflows
	// Segmentation fault (core dumped)
}

```

[↑ top](#segmentation-fault-debug)
<br><br><br><br>
<hr>
