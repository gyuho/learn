[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C: introduction, printf, scanf

- [Reference](#reference)
- [Install](#install)
- [Hello World!](#hello-world)
- [`printf`](#printf)
- [`scanf`](#scanf)

[↑ top](#c-introduction-printf-scanf)
<br><br><br><br><hr>


#### Reference

- [Learn C and Build Your Own Lisp](http://www.buildyourownlisp.com/)
- [C Programming Language by Brian W. Kernighan and Dennis M. Ritchie](http://www.amazon.com/C-Programming-Language-2nd-Edition/dp/0131103628/ref=sr_1_1?ie=UTF8&qid=1394272687&sr=8-1&keywords=programming+c)
- [Programming in C by Stephen G. Kochan](http://www.amazon.com/Programming-C-4th-Developers-Library/dp/0321776410/ref=sr_1_3?s=books&ie=UTF8&qid=1394274646&sr=1-3&keywords=Programming+in+C)
- [C documentation](http://devdocs.io/c)
- [Stanford CS Essential C](http://cslibrary.stanford.edu/101)
- [Stanford CS Education Library](http://cslibrary.stanford.edu)
- [gcc](http://gcc.gnu.org/onlinedocs)
- [Learn C The Hard Way](http://c.learncodethehardway.org/book)
- [C11 Specification](http://www.open-std.org/JTC1/SC22/WG14/www/docs/n1570.pdf)
- [Everything You Need to Know to Write Good C Code](https://github.com/btrask/stronglink/blob/master/SUBSTANCE.md)

[↑ top](#c-introduction-printf-scanf)
<br><br><br><br><hr>


#### Install

Please visit [here](https://gcc.gnu.org/).

[↑ top](#c-introduction-printf-scanf)
<br><br><br><br><hr>


#### Hello World!

```c
#include <stdio.h>

int main(void)
{
    printf("Hello World!\n");
    // Hello World!
    return 0;
}
```

To compile and execute the program:

- `cd code/` and `gcc hello.c` and `./a.out`
- `cd code/` and `gcc hello.c -o hello` and `./hello`

<br>

In `int main(void)`, `int` is the return type and `void` is an argument type,
and `main` is the name of the function. And the function `printf` is defined in
a standard library `stdio.h`, which is called a `header file` to be imported
with `#include`. At the end, it returns `0` that in general means normal
execution without any error in Linux. You can return something other than 0 to
tell there's something wrong.

[↑ top](#c-introduction-printf-scanf)
<br><br><br><br><hr>


#### `printf`

```c
#include <stdio.h> // to use printf

int main(void)
{
	// In C, special characters can be single-quoted.
	char newLine = '\n';
    printf("\"%s\": %d%c", "Hello", 10, newLine);
	// "Hello": 10

    printf("%10sWorld\n", "Hello");
	//     HelloWorld
    
    int num9=9, num17=17;
    printf("Octet: %o %#o\n", num9, num9); // Octet: 11 011 
    printf("Hex: %x %#x\n", num17, num17); // Hex: 11 0x11

    double fnum=0.123456789;
    printf("Float: %f %e\n", fnum, fnum); // Float: 0.123457 1.234568e-01
    printf("Float: %.3f\n", fnum);        // Float: 0.123

    return 0;
}

```

[↑ top](#c-introduction-printf-scanf)
<br><br><br><br><hr>


#### `scanf`

```c
#include <stdio.h>

int main(void)
{
	int num0, num1, num2;
	printf("Type 3 integers: ");
	scanf("%o, %d, %x", &num0, &num1, &num2);

	printf("\nWe got:\n");
	printf("Octet in Decimal  : %d\n", num0); // 
	printf("Decimal in Decimal: %d\n", num1); // 
	printf("Hex in Decimal    : %d\n", num2); // 

	float fnum;
	printf("\nType 1 float: ");
	scanf("%f", &fnum);
	printf("Float: %f\n", fnum);

    return 0;
}

/*
Type 3 integers: 10, 10, 10

We got:
Octet in Decimal  : 8
Decimal in Decimal: 10
Hex in Decimal    : 16

Type 1 float: 1.23e-5
Float: 0.000012
*/

```

[↑ top](#c-introduction-printf-scanf)
<br><br><br><br><hr>

