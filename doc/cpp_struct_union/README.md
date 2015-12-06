[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C++: struct, union

- [`struct`](#struct)
- [`union`](#union)

[↑ top](#c-character-string)
<br><br><br><br><hr>


#### `struct`

```cpp
#include <iostream>
#include <stdio.h>
#include <string.h>
using namespace std;

int main()
{
	struct data
	{
		string company; 
		char   name[6];
		int    number;
	};

	data dt;
	dt.company = "Google";
	string str = "Hello";
	strcpy(dt.name, str.c_str());
	dt.number = 100;

	cout << dt.company << endl; // Google
	cout << dt.name << endl;    // Hello
	cout << dt.number << endl;  // 100
}

```

[↑ top](#c-character-string)
<br><br><br><br><hr>


#### `union`

```cpp
#include <iostream>
using namespace std;

int main()
{
	// In union, all members share the same memory location.
	//
	// Unions can save us memory with different objects
	// sharing the memory location at different times.
	//
	union exampleUnion {
		int  num;
		char c;
		// data can be either int or char
	};
	// or }myUnion
	
	exampleUnion myUnion;
	myUnion.num = 10;
	myUnion.c = 'A';

	// myUnion.c overwrites num
	cout << myUnion.num << endl; // 65
	cout << myUnion.c << endl;   // A
}

```

[↑ top](#c-character-string)
<br><br><br><br><hr>
