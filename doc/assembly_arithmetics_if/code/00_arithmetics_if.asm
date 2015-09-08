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
