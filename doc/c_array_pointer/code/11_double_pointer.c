#include <stdio.h>

/*
void myFunc(TYPE * arr) {}
void myFunc(TYPE arr[]) {}

void myFunc(TYPE ** arr) {}
void myFunc(TYPE * arr[]) {}
*/

void printStrings(int argc, char * argv[])
{
	int i;
	for (i=0; i<argc; i++)
		printf("%s ", argv[i]);
	printf("\n");
}

int main(void)
{
	char * str[3] = {
		"A", "B", "C"
	};

	printStrings(3, str);
	// A B C

	return 0;
}
