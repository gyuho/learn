#include <stdio.h>

void swapValue(int v1, int v2)
{
	int temp = v1;
	v1 = v2;
	v2 = temp;
}

void swapPointer(int * pt1, int * pt2)
{
	int temp = *pt1; // dereference
	*pt1 = *pt2;
	*pt2 = temp;
}

int main()
{
	int num1 = 1;
	int num2 = 2;
	swapValue(num1, num2);
	printf("num1 num2: %d %d\n", num1, num2);
	// 1 2

	swapPointer(&num1, &num2);
	printf("num1 num2: %d %d\n", num1, num2);
	// 2 1

	return 0;
}
