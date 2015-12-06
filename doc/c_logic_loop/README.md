[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C: logic, loop

- [logic](#logic)
- [`if`](#if)
- [`switch`](#switch)
- [`for`](#for)
- [`while`](#while)
- [fizzbuzz](#fizzbuzz)

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### logic

```c
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

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### `if` 

```c
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

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### `switch`

```c
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

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### `for`

```c
#include <stdio.h>
#include <stdlib.h>

#define SIZE 5

int main(void)
{
	unsigned i = 0, array [SIZE];

	for( ; i < SIZE; ++i)
		array [i] = rand();
	printf("Array Updated\n");

	for (i = 0; i < SIZE; ++i)
		printf("%d, ", array[i]);
	printf("\n");
 
	return EXIT_SUCCESS;
}

/*
Array Updated
1804289383, 846930886, 1681692777, 1714636915, 1957747793,
*/

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### `while`

```c
#include <stdio.h>
#include <stdlib.h>

#define SIZE 5

int main(void)
{
	unsigned i = 0, array [SIZE];

	do
	{
		array [i] = rand();
		++i;
	} while (i < SIZE);

	printf("Array Updated!\n");

	i = 0;
	do
	{
		printf("%d ", array[i]);
		++i;
	} while (i < SIZE);
	printf("\n");


	int a = 10;
	while( a < 20 )
	{
		printf("a: %d\n", a);
		a++;
	}

	return EXIT_SUCCESS;
}

/*
Array Updated!
1804289383 846930886 1681692777 1714636915 1957747793 
a: 10
a: 11
a: 12
a: 13
a: 14
a: 15
a: 16
a: 17
a: 18
a: 19
*/

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### fizzbuzz

> Write a program that prints the numbers from 1 to 100. But for **multiples of
> three print “Fizz”** instead of the number and for the **multiples of**
> **_five_** **print “Buzz”. For numbers which are multiples of both**
> **_three_** and **_five_** **print “FizzBuzz”**.
>
> [**Fizz Buzz Test**](http://c2.com/cgi/wiki?FizzBuzzTest)


```c
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
 
```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>
