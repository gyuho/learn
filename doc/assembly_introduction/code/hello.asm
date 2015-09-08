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
