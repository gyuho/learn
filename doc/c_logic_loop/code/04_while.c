#include <stdio.h>
#include <stdlib.h>

#define SIZE 5

int main(void)
{
	unsigned i = 0, array [SIZE];

	do
	{
		array [i] = rand();
		++i;
	} while (i < SIZE);

	printf("Array Updated!\n");

	i = 0;
	do
	{
		printf("%d ", array[i]);
		++i;
	} while (i < SIZE);
	printf("\n");


	int a = 10;
	while( a < 20 )
	{
		printf("a: %d\n", a);
		a++;
	}

	return EXIT_SUCCESS;
}

/*
Array Updated!
1804289383 846930886 1681692777 1714636915 1957747793 
a: 10
a: 11
a: 12
a: 13
a: 14
a: 15
a: 16
a: 17
a: 18
a: 19
*/
