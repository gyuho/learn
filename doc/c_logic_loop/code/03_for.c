#include <stdio.h>
#include <stdlib.h>

#define SIZE 5

int main(void)
{
	unsigned i = 0, array [SIZE];

	for( ; i < SIZE; ++i)
		array [i] = rand();
	printf("Array Updated\n");

	for (i = 0; i < SIZE; ++i)
		printf("%d, ", array[i]);
	printf("\n");
 
	return EXIT_SUCCESS;
}

/*
Array Updated
1804289383, 846930886, 1681692777, 1714636915, 1957747793,
*/
