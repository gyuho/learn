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
