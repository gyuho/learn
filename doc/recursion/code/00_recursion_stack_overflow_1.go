package main

func f() {
	f()
}

func main() {
	f()
}

/*
runtime: goroutine stack exceeds 1000000000-byte limit
fatal error: stack overflow

runtime stack:
runtime.throw(0x465799)
	/usr/local/go/src/runtime/panic.go:491 +0xad
runtime.newstack()
	/usr/local/go/src/runtime/stack.c:784 +0x555
runtime.morestack()
	/usr/local/go/src/runtime/asm_amd64.s:324 +0x7e

goroutine 1 [stack growth]:

...

	/home/ubuntu/go/src/github.com/gyuho/learn/doc/recursion/code/00_recursion_stack_overflow_1.go:4 +0x1b fp=0xc228038520 sp=0xc228038518
...additional frames elided...

...
*/
