#include <stdio.h>

int main(void)
{
	int i;
	for ( i = 1; i < 101; i++ )
	{
		if ( i%15 == 0 )
		{
			printf("FizzBuzz\n");
		}
		else if ( i%3 == 0 )
		{
			printf("Fizz\n");
		}
		else if ( i%5 == 0 )
		{
			printf("Buzz\n");
		}
		else 
		{
			printf("%d\n", i);
		}
	}
	return 0;
}

// ...
// Buzz
// 41
// Fizz
// 43
// 44
// FizzBuzz
// 46
// ...
