[*back to contents*](https://github.com/gyuho/learn#contents)
<br>


# Assembly: arithmetics, if

- [Reference](#reference)
- [arithmetics, if](#arithmetics-if)

[↑ top](#assembly-arithmetics-if)
<br><br><br><br>
<hr>





#### Reference

- [`nasm` documentation](http://www.nasm.us/xdoc/2.11.08/html/nasmdoc0.html)
- [`nasm` tutorial](http://cs.lmu.edu/~ray/notes/nasmtutorial/)
- [x86 Assembly Guide](http://www.cs.virginia.edu/~evans/cs216/guides/x86.html)
- [x86_64 NASM Assembly Quick Reference](https://www.cs.uaf.edu/2006/fall/cs301/support/x86_64/)
- [Basic differences between x86 Assembly and X86-64 Assembly](https://www.exploit-db.com/papers/13136/)
- [github.com/0xAX/asm](https://github.com/0xAX/asm)

[↑ top](#assembly-arithmetics-if)
<br><br><br><br>
<hr>



#### arithmetics, if

Let's review some of the most frequent **instructions**:

- `mov A, B`: copy `B` into `A`
- `and A, B`: copy `A and B` into `A`
- `or A, B`: copy `A or B` into `A`
- `add A, B`: copy `A + B` into `A`
- `sub A, B`: copy `A - B` into `A`
- `inc A`: copy `A + 1` into `A`
- `dec A`: copy `A - 1` into `A`
- `syscall`: invokes an operating system write
- `db`: declares bytes to be in memory while program running

<br>

```asm
; -------------------------------------------------------------
; nasm -f elf64 -o 00_arithmetics_if.o 00_arithmetics_if.asm
; ld -o 00_arithmetics_if 00_arithmetics_if.o
; ./00_arithmetics_if
; -------------------------------------------------------------

section .data
    ; define constants
    num1:   equ 100
    num2:   equ 50
    ; initialize message
    myMessage:    db "Correct"
 
section .text
 
    global _start
_start:

    ; copy num1's value to rax (temporary register)
    mov rax, num1

    ; copy num2's value to rbx
    mov rbx, num2

    ; put rax + rbx into rax
    add rax, rbx

    ; compare rax and 150
    cmp rax, 150

    ; go to .exit label if rax and 150 are not equal
    jne .exit
    
    ; go to .correctSum label if rax and 150 are equal
    jmp .correctSum

; Print message that sum is correct
.correctSum:
    ; write syscall
    mov     rax, 1
    ; file descritor, standard output
    mov     rdi, 1
    ; message address
    mov     rsi, myMessage
    ; length of message
    mov     rdx, 15
    ; call write syscall
    syscall
    ; exit from program
    jmp .exit
 
; exit procedure
.exit:
    ; exit syscall
    mov    rax, 60
    ; exit code
    mov    rdi, 0
    ; call exit syscall
    syscall

```

And run this commands:

```bash
nasm -f elf64 -o 00_arithmetics_if.o 00_arithmetics_if.asm
ld -o 00_arithmetics_if 00_arithmetics_if.o
./00_arithmetics_if
```

This will prints out the message `Correct`.


[↑ top](#assembly-arithmetics-if)
<br><br><br><br>
<hr>
