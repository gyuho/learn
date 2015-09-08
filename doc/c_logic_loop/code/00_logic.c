#include <stdio.h>

// #include <stdbool.h>
//
// or
typedef int bool;
#define true  1
#define false 0

int main(void)
{
	bool b1 = true;
	bool b2 = false;
	printf("b1 && b2: %d\n", b1 && b2);
	printf("b1 || b2: %d\n", b1 || b2);

	int a = 10, b = 15, c = 20;
	printf("a < b : %d\n", (a < b));
	printf("b > c : %d\n", (b > c));

    return 0;
}

/*
b1 && b2: 0
b1 || b2: 1
a < b : 1
b > c : 0
*/
