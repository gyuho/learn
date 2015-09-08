#include <iostream>
#include <stdio.h>
using namespace std;

int main() {
	// In union, all members share the same memory location.
	//
	// Unions can save us memory with different objects
	// sharing the memory location at different times.
	//
	union Data {
		uint32_t i;
		char c[4];
		// so the data can be either
		// uint32_t or char
	};
	Data data;
	data.i = 0x0A0B0C0D;
	// data.c[4] = 0x0A0B0C0D; (X)
	printf("%X\n", data.c[0]);
	printf("%X\n", data.c[1]);
	printf("%X\n", data.c[2]);
	printf("%X\n", data.c[3]);
	if (data.c[0] == 10)
		cout <<  "Big-Endian" << endl;
	else
		cout << "Little-Endian" << endl;
}

/*
D
C
B
A
Little-Endian
*/
