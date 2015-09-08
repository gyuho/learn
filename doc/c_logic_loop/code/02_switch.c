#include <stdio.h>

int main()
{
	while (1)
	{
		printf("Type an integer: ");
		int selected;
		scanf("%d", &selected);

		switch (selected)
		{
			case 0:
				{
					printf("selected 0\n");
					break;
				}
			case 1:
				{
					printf("selected 1\n");
					// continue; 
					// without this, it continues on 2
				}
			case 2:
				{
					printf("selected 2\n");
					continue;
				}
			case 3:
				{
					printf("selected 3\n");
					continue;
				}
			default:
				{
					printf("selected %d\n", selected);
				}
		}
		printf("Breaking out of loop!\n");
		break;
	}
}

/*
Type an integer: 3
selected 3
Type an integer: 2
selected 2
Type an integer: 1
selected 1
selected 2
Type an integer: 5
selected 5
Breaking out of loop!

Type an integer: 3
selected 3
Type an integer: 2
selected 2
Type an integer: 1
selected 1
selected 2
Type an integer: 0
selected 0
Breaking out of loop!
*/
