[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# C: type, variable, operator

- [type](#type)
- [variable](#variable)
- [operator](#operator)

[↑ top](#c-type-variable-operator)
<br><br><br><br>
<hr>


#### type

- [Basic Types and Operators](http://cslibrary.stanford.edu/101/EssentialC.pdf)

[↑ top](#c-type-variable-operator)
<br><br><br><br>
<hr>


#### variable

```c
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

```

In more detail:

> The C programming language manages memory **statically**, **automatically**, or
> **dynamically**. **Static-duration** variables are allocated in **main memory**, usually
> along with the executable code of the program, and **persist for the lifetime
> of the program**; **automatic-duration** variables are allocated on the **stack** and
> come and **go** as **functions** are called and **return**. For static-duration and
> automatic-duration variables, the size of the allocation must be compile-time
> constant (except in C99, which allowed variable-length automatic arrays). If
> the required size is not known until run-time (for example, if data of
> arbitrary size is being read from the user or from a disk file), then using
> fixed-size data objects is inadequate.
>
> The lifetime of allocated memory can also cause concern. Neither static- nor
> automatic-duration memory is adequate for all situations. Automatic-allocated
> data cannot persist across multiple function calls, while static data persists
> for the life of the program whether it is needed or not. In many situations the
> programmer requires greater flexibility in managing the lifetime of allocated
> memory.
>
> These limitations are avoided by using dynamic memory allocation in which
> memory is more explicitly (but more flexibly) managed, typically, by allocating
> it from the free store (informally called the "heap"), an area of memory
> structured for this purpose. In C, the library function `malloc` is used to
> allocate a block of memory on the **heap**. The program accesses this block of
> memory via a pointer that `malloc` returns. When the memory is no longer needed,
> the pointer is passed to `free` which deallocates the memory so that it can be
> used for other purposes.
>
> [*C dynamic memory allocation*](https://en.wikipedia.org/wiki/C_dynamic_memory_allocation)
> *by Wikipedia*

[↑ top](#c-type-variable-operator)
<br><br><br><br>
<hr>


#### operator

```c
#include <stdio.h>

int main(void)
{
	int num0;
	int num1;
	num0=7;
	num1=2;
	printf("%d ÷ %d: %f\n", num0, num1, (double)num0/num1);
	// 7 ÷ 2: 3.500000 
	printf("%d ÷ %d's quotient:  %d\n", num0, num1, num0/num1);
	// 7 ÷ 2's quotient:  3 
	printf("%d ÷ %d's remainder: %d\n\n", num0, num1, num0%num1);
	// 7 ÷ 2's remainder: 1

	printf("num0: %d\n", num0);
	printf("++num0: %d\n", ++num0);
	printf("num0: %d\n\n", num0);

	printf("num1: %d\n", num1);
	printf("num1++: %d\n", num1++);
	printf("num1: %d\n\n", num1);
	/*
		num0: 7
		++num0: 8
		num0: 8

		num1: 2
		num1++: 2
		num1: 3
	*/

	printf("AND: %d\n", (1<2) && (3<4)); // AND: 1
	printf("OR:  %d\n", (1>2) || (3>4)); // OR:  0

	char array[7] = "Hello";
	int i;
	for (i=0; i<7; i++)
		printf("%d: %c\n", i, array[i]);
	printf("\n");
	for (i=0; i<7; ++i)
		printf("%d: %c\n", i, array[i]);
	printf("\n");
	/*
		0: H
		1: e
		2: l
		3: l
		4: o
		5: 
		6: 

		0: H
		1: e
		2: l
		3: l
		4: o
		5: 
		6: 
	*/

	return 0;
}

```

[↑ top](#c-type-variable-operator)
<br><br><br><br>
<hr>

