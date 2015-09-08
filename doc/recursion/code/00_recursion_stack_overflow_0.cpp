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
