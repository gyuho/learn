[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# C: array, pointer

- [Reference](#reference)
- [array](#array)
- [`char`](#char)
- [pointer](#pointer)

[↑ top](#c-array-pointer)
<br><br><br><br>
<hr>







#### Reference

- [C11 Specification](http://www.open-std.org/JTC1/SC22/WG14/www/docs/n1570.pdf)
- [Pointers and Memory](http://cslibrary.stanford.edu/102/)

[↑ top](#c-array-pointer)
<br><br><br><br>
<hr>






#### array

Let's say you have thousands of values to store in your program. One way, you
could:

```c
int num1, num2, num3, num4, num5;
int num6, num7, num8, num9, num10;
...
int num1000, num1001, num1002, num1003, num1004;
...
```

This is too verbose and error-prone. You should use array:

![array](img/array.png)

```c
#include <stdio.h>

int main(void)
{
	int arr0[2];
	arr0[0] = 10;
	arr0[1] = 20;
	int sum0=0, i;
	for (i=0; i<2; i++)
		sum0 += arr0[i];
	printf("sum0: %d\n", sum0);
	// sum0: 30

	int len=50;
	double arr1[len];
	int j;
	for (j=0; j<len; j++)
		arr1[j] = (double)j;
	double sum1=0;
	int k;
	for (k=0; k<len; k++)
		sum1 += arr1[k];
	printf("sum1: %f\n", sum1);
	// sum1: 1225.000000

	int arr2[3]={0, 1, 2};
	int i2;
	for (i2=0; i2<3; i2++)
		printf("%d ", arr2[i2]);
	printf("\n");
	// 0 1 2
	
	int arr3[]={0, 1, 2}; // automatically sized to 3
	int i3;
	for (i3=0; i3<3; i3++)
		printf("%d ", arr3[i3]);
	printf("\n");
	// 0 1 2

	int arr4[5]={1}; // automatically fill up with 0
	int i4;
	for (i4=0; i4<5; i4++)
		printf("%d ", arr4[i4]);
	printf("\n");
	// 1 0 0 0 0 

	printf("arr4 sizeof: %ld\n", sizeof(arr4));             // arr4 sizeof: 20
	printf("int sizeof: %ld\n", sizeof(int));               // int sizeof: 4
	printf("arr4 length: %ld\n", sizeof(arr4)/sizeof(int)); // arr4 length: 5

	printf("\ntype 3 integers: ");
	int arr[3];
	scanf("%d", &arr[0]);
	scanf("%d", &arr[1]);
	scanf("%d", &arr[2]);
	int ia;
	for (ia=0; ia<3; ia++)
		printf("%d ", arr[ia]);
	printf("\n");
	// type 3 integers: 100 200 300
	// 100 200 300

    return 0;
}

```

[↑ top](#c-array-pointer)
<br><br><br><br>
<hr>







#### `char`

C represents string literals with double quotes:

```c
printf("Hello World!\n")
```

And you can save this string in character array type. In C, string is an array
of characters, and the **`null character \0`** is automatically inserted at the
end of all character arrays. C needs `\0` to differentiate between two types of
array of characters, *one that cannot be printed with `%s`, the other that can
be printed with `%s`*:

![char_array](img/char_array.png)

```c
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

```

[↑ top](#c-array-pointer)
<br><br><br><br>
<hr>








#### pointer

Pointer is a variable to store the address value of data(variable).
`&num` returns the address of `num`. `int * num` defines a pointer type
variable `num`. `*` is also used to dereference *or access* the memory
that the pointer points to.

[↑ top](#c-array-pointer)
<br><br><br><br>
<hr>

