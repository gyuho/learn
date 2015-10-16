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

	printf("AND: %d\n", (1<2) && (3<4)); // AND: 1
	printf("OR:  %d\n", (1>2) || (3>4)); // OR:  0

	return 0;
}

/*
num0: 7
++num0: 8
num0: 8

num1: 2
num1++: 2
num1: 3
*/

