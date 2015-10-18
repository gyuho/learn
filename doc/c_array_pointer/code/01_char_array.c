#include <stdio.h>

int main(void)
{
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

	char str[]="Hello World!";
	printf("str length: %ld\n", sizeof(str)/sizeof(char));
	// str length: 13
	
	// we need null character to differentiate between these two:
	char charArray[]={'H', 'e', 'l', 'l', 'o'};
	printf("charArray: %s\n", charArray); // charArray: Hello

	char charArrayString[]={'H', 'e', 'l', 'l', 'o', '\0'};
	printf("charArrayString: %s\n", charArrayString); // charArrayString: Hello
	int idx=0;
	while (charArrayString[idx] != 0)
	{
		printf("%c", charArrayString[idx]);
		idx++;
	}
	printf("\n");
	// Hello

	charArrayString[3] = '\0';
	printf("charArrayString: %s\n", charArrayString); // charArrayString: Hel

	charArrayString[1] = 0;
	printf("charArrayString: %s\n", charArrayString); // charArrayString: H

    return 0;
}

