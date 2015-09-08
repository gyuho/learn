#include <stdio.h>

int main(void)
{
	int i = 2;
	if (i > 2) {
		printf("i > 2 is true\n");
	} else {
		printf("i > 2 is false\n");
	}

	i = 3;
	if (i == 3) printf("i == 3\n");

	if (i != 3) printf("i != 3 is true\n");
	else        printf("i != 3 is false\n");
}

/*
i > 2 is false
i == 3
i != 3 is false
*/
