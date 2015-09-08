#include <iostream>

int f(){
	f();
}

int main()
{
	f(); // stack overflows
	// Segmentation fault (core dumped)
}
