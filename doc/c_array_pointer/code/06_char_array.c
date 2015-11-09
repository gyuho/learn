#include <stdio.h>

int main()
{
	char s0[] = "abc";
	// *s0 = "Hello";   // (X)
	printf("%s\n", s0); // abc
	s0[0] = 'X';
	printf("%s\n", s0); // Xbc

	char * s1 = "abc";
	s1 = "Hello";
	printf("%s\n", s1); // Hello
	s1[0] = 'X';        // (X)
	printf("%s\n", s1); // (X)

	return 0;
}
