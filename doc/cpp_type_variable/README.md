[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C++: type, variable

- [type](#type)
- [variable](#variable)

[↑ top](#c-type-variable)
<br><br><br><br><hr>


#### type

- [C++ type system](https://msdn.microsoft.com/en-us/library/hh279663.aspx)

[↑ top](#c-type-variable)
<br><br><br><br><hr>


#### variable

```c++
/*
g++ -std=c++11 00_variable.cpp -o 00_variable;
./00_variable;
*/
#include <iostream>
using namespace std;

int main()
{
	// declaring variables:
	int a, b;

	a = 7;
	b = 2;
	a = a + 1;
	b++;
	auto c = a - b;
	// int c = a - b;
	cout << c << endl;
	// 5

	string st1("Hello");
	cout << st1 << endl;
	// Hello

	string st2 = "World";
	cout << st2 << endl;
	// World

	return 0;
}

```

<br>
According to [FAQ](https://isocpp.org/wiki/faq/newbie#main-returns-int)
, `return 0;` is not needed:

> In C++, main() need not contain an explicit return statement.
> In that case, the value returned is 0, meaning successful execution.

```c++
#include <iostream>

int main()
{
	std::cout << "This program returns the integer value 0\n";
}

```

[↑ top](#c-type-variable)
<br><br><br><br><hr>
