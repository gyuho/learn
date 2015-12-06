[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C++: introduction

- [Reference](#reference)
- [Install](#install)
- [Hello World!](#hello-world)
- [argument](#argument)

[↑ top](#c-introduction)
<br><br><br><br><hr>


#### Reference

- [Standard C++](https://isocpp.org/)
- [cppreference.com](http://en.cppreference.com/w/)
- [cplusplus.com/reference](http://www.cplusplus.com/reference/)
- [C++ Core Guidelines](https://github.com/isocpp/CppCoreGuidelines)
- [C++ Language Reference](https://msdn.microsoft.com/en-us/library/3bstk3k5.aspx)
- [C/C++ Language and Standard Libraries](https://msdn.microsoft.com/en-us/library/hh875057.aspx)
- [cplusplus.com](http://www.cplusplus.com)
- [C++ documentation](http://devdocs.io/cpp)
- [gcc](http://gcc.gnu.org/onlinedocs)
- [Stanford CS 106B](http://cs.stanford.edu/people/eroberts/courses/cs106b)
- [Stanford CS 106L](http://www.stanford.edu/class/cs106l/course_reader.html)
- [Stanford CS Education Library](http://cslibrary.stanford.edu)
- [C++ FAQ](https://isocpp.org/faq)
- [C++ Primer Plus by Stephen Prata](http://www.amazon.com/Primer-Plus-Edition-Developers-Library/dp/0321776402/ref=sr_1_3?ie=UTF8&qid=1394277384&sr=8-3&keywords=C%2B%2B)

[↑ top](#c-introduction)
<br><br><br><br><hr>


#### Install

Please visit [here](https://gcc.gnu.org/).

[↑ top](#c-introduction)
<br><br><br><br><hr>


#### Hello World!

```c++
#include <stdio.h>
#include <iostream>
using namespace std;

int main(int argc, const char *argv[])
{
	printf("%s %d\n", "Hello World", 10);
    cout << "Hello World!";
    return 0;
}
// Hello World 10
// Hello World!
```

You can either:

- `cd code/` and `g++ -std=c++11 hello.cpp` and `./a.out`
- `cd code/` and `g++ -std=c++11 hello.cpp -o hello` and `./hello`

<br>

`int argc, const char *argv[]` is used to get arguments
from command line. If you do not need to process it,
just use `int main()` or `int main(void)`.

[↑ top](#c-introduction)
<br><br><br><br><hr>


#### argument

```cpp
#include <iostream>
using namespace std;

int main(int argc, char* argv[])
{
    // Check the number of parameters
    if (argc < 2) {
        cerr << "Usage: " << argv[0] << " NAME" << endl;
        return 1;
    }
    // Print the user's name:
    cout << argv[0] << "says hello, " << argv[1] << "!" << endl;
}

/*
$ ./a.out  Gyu-Ho

Then

./a.outsays hello, Gyu-Ho!
*/

```

[↑ top](#c-introduction)
<br><br><br><br><hr>
