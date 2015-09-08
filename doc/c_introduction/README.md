[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# C: introduction

- [Reference](#reference)
- [Install](#install)
- [Hello World!](#hello-world)

[↑ top](#c-introduction)
<br><br><br><br>
<hr>







#### Reference

- [C Programming Language by Brian W. Kernighan and Dennis M. Ritchie](http://www.amazon.com/C-Programming-Language-2nd-Edition/dp/0131103628/ref=sr_1_1?ie=UTF8&qid=1394272687&sr=8-1&keywords=programming+c)
- [Programming in C by Stephen G. Kochan](http://www.amazon.com/Programming-C-4th-Developers-Library/dp/0321776410/ref=sr_1_3?s=books&ie=UTF8&qid=1394274646&sr=1-3&keywords=Programming+in+C)
- [C documentation](http://devdocs.io/c)
- [Stanford CS Essential C](http://cslibrary.stanford.edu/101)
- [Stanford CS Education Library](http://cslibrary.stanford.edu)
- [gcc](http://gcc.gnu.org/onlinedocs)
- [Learn C The Hard Way](http://c.learncodethehardway.org/book)
- [C11 Specification](http://www.open-std.org/JTC1/SC22/WG14/www/docs/n1570.pdf)
- [Everything You Need to Know to Write Good C Code](https://github.com/btrask/stronglink/blob/master/SUBSTANCE.md)

[↑ top](#c-introduction)
<br><br><br><br>
<hr>








#### Install

Please visit [here](https://gcc.gnu.org/).

[↑ top](#c-introduction)
<br><br><br><br>
<hr>







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

You can either:

- `cd code/` and `gcc hello.c` and `./a.out`
- `cd code/` and `gcc hello.c -o hello` and `./hello`

[↑ top](#c-introduction)
<br><br><br><br>
<hr>
