/*
gcc 00_variable.c -o 00_variable;
./00_variable;
*/

#include <stdio.h>
#include <stdlib.h>

void inc()
{
	int av = 10;
	static int sv = 10;
	av += 5;
	sv += 5;
	printf("av = %d, sv = %d\n", av, sv);
}

// static storage duration
int A;       // uninitialized global variable; this system initializes with zero
int B = 9;   // initialized global variable
 
int main(void)
{
	int i;
	for (i = 0; i < 10; ++i)
		inc();

	printf("&A = %p\n", (void*)&A);
	printf("&B = %p\n", (void*)&B);
 
	// "automatic" storage duration:
	//	variable is allocated at the beginning of the enclosing code block
	//	and deallocated at the end. 
	int A = 1;   // hides global A
	printf("&A = %p\n", (void*)&A);
 
	// "static" storage duration:
	//	variable is allocated when the program begins
	//	and deallocated when the program ends.
	//	It keeps the state between function calls.
	static int B=1; // hides global B
	printf("&B = %p\n", (void*)&B);
 
	// "allocated" storage duration:
	//	variable is allocated and deallocated by dynamic memory allocation functions.
	int *pt = (int*)malloc(sizeof(int));   // start allocated storage duration
	printf("address of int in allocated memory = %p\n", (void*)pt);
	free(pt);                              // stop allocated storage duration 
 
	return 0;
}

/*
av = 15, sv = 15
av = 15, sv = 20
av = 15, sv = 25
av = 15, sv = 30
av = 15, sv = 35
av = 15, sv = 40
av = 15, sv = 45
av = 15, sv = 50
av = 15, sv = 55
av = 15, sv = 60
&A = 0x601060
&B = 0x601050
&A = 0x7ffceb15b210
&B = 0x601058
address of int in allocated memory = 0x1195010
*/

