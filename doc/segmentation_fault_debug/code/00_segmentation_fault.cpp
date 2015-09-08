#include <iostream>
using namespace std;

int main()
{
	int *ptr = NULL;
	*ptr = 100; // Write to invalid memory address
	// Segmentation fault (core dumped)
}

/*
(gdb) run

Program received signal SIGSEGV, Segmentation fault.
0x00000000004006dd in main ()
*/
