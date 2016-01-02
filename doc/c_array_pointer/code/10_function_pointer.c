#include <stdio.h>

void printAdd(int v1, int v2)
{
	printf("%d + %d = %d\n", v1, v2, v1 + v2);
}

void printString(char * str)
{
	printf("%s\n", str);
}

int main()
{
	int num1 = 1, num2 = 2;
	void (*fpAdd)(int, int) = printAdd;
	fpAdd(num1, num2);
	// 1 + 2 = 3

	char * str = "Hello World!";
	void (*fpStr)(char *) = printString;
	fpStr(str);
	// Hello World!

	return 0;
}
