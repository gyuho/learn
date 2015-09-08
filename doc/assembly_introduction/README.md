[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# Assembly: introduction

- [Reference](#reference)
- [Install](#install)
- [Hello World!](#hello-world)
- [`nasm` syntax](#nasm-syntax)

[↑ top](#assembly-introduction)
<br><br><br><br>
<hr>





#### Reference

- [`nasm` documentation](http://www.nasm.us/xdoc/2.11.08/html/nasmdoc0.html)
- [`nasm` tutorial](http://cs.lmu.edu/~ray/notes/nasmtutorial/)
- [x86 Assembly Guide](http://www.cs.virginia.edu/~evans/cs216/guides/x86.html)
- [x86_64 NASM Assembly Quick Reference](https://www.cs.uaf.edu/2006/fall/cs301/support/x86_64/)
- [Basic differences between x86 Assembly and X86-64 Assembly](https://www.exploit-db.com/papers/13136/)
- [github.com/0xAX/asm](https://github.com/0xAX/asm)

[↑ top](#assembly-introduction)
<br><br><br><br>
<hr>





#### Install

My machine setting as of today is `Linux 64-bit`:

```bash
$ cat /proc/cpuinfo | grep "model name" | head -1;
model name	: Intel(R) Core(TM) i7-4750HQ CPU @ 2.00GHz

$ lsb_release -a
Distributor ID:	Ubuntu
Description:	Ubuntu 14.04.3 LTS
Release:	14.04
Codename:	trusty

$ uname -sm
Linux x86_64
```

<br>

> x86-64 (also known as x64, x86_64 and **_`AMD64`_**) is the 64-bit
> version of the x86 instruction set.
>
> [x86-64](https://en.wikipedia.org/wiki/X86-64) *by Wikipedia*

<br>

And we will use `nasm` to write assembly:

> What Is NASM?
>
> The Netwide Assembler, NASM, is an 80x86 and x86-64 assembler designed for
> portability and modularity. It supports a range of object file formats,
> including Linux and \*BSD a.out, ELF, COFF, Mach-O, Microsoft 16-bit OBJ, Win32
> and Win64. It will also output plain binary files. Its syntax is designed to be
> simple and easy to understand, similar to Intel's but less complex. It supports
> all currently known x86 architectural extensions, and has strong support for
> macros.
>
> [*nasm.us*](http://www.nasm.us/xdoc/2.11.08/html/nasmdoc1.html#section-1.1)

<br>

To install:

```bash
sudo apt-get install nasm;
```

[↑ top](#assembly-introduction)
<br><br><br><br>
<hr>





#### Hello World!

```asm
; -------------------------------------
; nasm -f elf64 -o hello.o hello.asm
; ld -o hello hello.o
; ./hello
; -------------------------------------

section .data   ; data section stores constants
	myMessage	db	"Hello World!"
	; 1. myMessage         is label
	; 2. db                is instruction
	; 3. "Hello World!"    is operand(object)
 
section .text   ; text section contains code
	global	_start
; beginning of exported symbols
_start:
	mov	rax, 1           ; copy 1 to rax (temporary register)
	mov	rdi, 1           ; copy 1 to rdi (1st argument)
	mov	rsi, myMessage   ; copy myMessage to rsi (2nd argument)
	mov	rdx, 13          ; copy 13 to rdx (3rd argument)
	syscall              ; invoke system call write
	mov	rax, 60          ; copy 60 to rax (temporary register)
	mov	rdi, 0           ; copy 0 to rdi (1st argument)
	syscall              ; invoke system call write

```

<br>
And then to build:
```
nasm -f elf64 -o hello.o hello.asm;
```
`nasm` command builds an object file `hello.o` from `hello.asm`
using the `elf64` object file format. `object file` is a file
that contains `object code`, mostly in machine code, and information
for linkers. 

<br>
And then to link:
```
ld -o hello hello.o;
```
`ld` is a linker that takes the object file `hello.o` generated
by the compiler and combines it into a single executable file.

<br>
And then to execute:
```
./hello
```

This prints `Hello World!`.


[↑ top](#assembly-introduction)
<br><br><br><br>
<hr>





#### `nasm` syntax

```asm
section .data   ; data section stores constants
	myMessage	db	"Hello World!"
	; 1. myMessage         is label
	; 2. db                is instruction
	; 3. "Hello World!"    is operand(object)
```

<br>
And:

```asm
section .text   ; text section contains code
```

<br>
And:

```asm
	global	_start
_start:
```

`global` directive exports symbols in the code to object code.
And `_start` marks the beginning of exported symbols.


<br>
And:

```asm
	mov	rax, 1           ; copy 1 to rax (temporary register)
	mov	rdi, 1           ; copy 1 to rdi (1st argument)
	mov	rsi, myMessage   ; copy myMessage to rsi (2nd argument)
	mov	rdx, 13          ; copy 13 to rdx (3rd argument)
```

`mov` instruction copies data into its first operand, from the second operand.
`move A, B` copies the value in B into A. 


<br>
[This](http://www.nasm.us/doc/nasmdo11.html) explains following are
**registers** in 64-bit mode. Note that registers in assembly language
are not case-sensitive:

- `RAX`, `RCX`, `RDX`, `RBX`, `RSP`, `RBP`, `RSI`, `RDI`, `R8-R15`

<br>
And **_`register`_** is:

> In computer architecture, a processor **register** is a small amount of **storage**
> available as part of a digital processor, such as a central processing unit
> (CPU). Such registers are typically addressed by mechanisms other than main
> memory and can be accessed faster. Almost all computers, load-store
> architecture or not, **load data from a larger memory into registers** where it
> is used for **arithmetic, manipulated or tested by machine instructions**.
> Manipulated data is then often stored back into main memory, either by the
> same instruction or a subsequent one. Modern processors use either static or
> dynamic RAM as main memory, with the latter usually accessed via one or more
> cache levels.
>
> Processor registers are normally at the **top of the memory hierarchy**, and
> provide the fastest way to access data. The term normally refers only to the
> group of registers that are directly encoded as part of an instruction, as
> defined by the instruction set.
>
> [*Processor register*](https://en.wikipedia.org/wiki/Processor_register) *by Wikipedia*

<br>
So:

- `move rax, 1` copies 1 into `rax` register.
- `mov	rdi, 1` copies 1 into `rdi` register.
- `mov	rsi, myMessage` copies `myMessage` into `rsi` register.
- `mov	rdx, 13` copies 13 into `rdx` register.

Then what are `rax`, `rdi`, `rsi`, `rdx`? They are different registers.
https://www.exploit-db.com/papers/13136 explains:

- `rax`: temporary register; with variable arguments.
- `rdi`: used to pass 1st argument to functions.
- `rsi`: used to pass 2nd argument to functions.
- `rdx`: used to pass 3rd argument to functions; 2nd return.

Then what function are we passing 1, `myMessage`, 13, etc. to?
It calls the operating system to do the write.


<br>
And:

```asm
	syscall              ; invoke system call write
```

`syscall` in `nasm` invokes the operating system call write.
I am running this with Linux, which is written in *C*.
Then the **system call write** invokes
[this](https://en.wikipedia.org/wiki/Write_(system_call)):

```c
ssize_t write(int fd, const void *buf, size_t nbytes);
```

- `fd`: file descriptor. `0` is for standard input, `1` for standard output,
  and `2` for standard error.
- `*buf`: pointer to a character array that stores data from the file pointed
  by `fd`.
- `nbytes`: number of bytes to write onto the file, from the character array.




<br>
And:

```asm
	mov	rax, 60          ; copy 60 to rax (temporary register)
	mov	rdi, 0           ; copy 0 to rdi (1st argument)
	syscall              ; invoke system call write
```

- `mov rax, 60`: copies `60` into `rax`, which means **exit syscall**.
- `mov rdi, 0`: copies `0` into `rdi`, which means **successful program
  exit**.

[↑ top](#assembly-introduction)
<br><br><br><br>
<hr>
