[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# C++: character, string, pointer, memory

- [Reference](#reference)
- [character](#character)
- [`string`](#string)
- [reverse string](#reverse-string)
- [`string` to *character array*](#string-to-character-array)
- [STL `string`](#stl-string)
- [pointer](#pointer)
- [pointer, array](#pointer-array)
- [ampersand `&`](#ampersand-)
- [pointer, array, function](#pointer-array-function)

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### Reference

- [Pointers and Memory](http://cslibrary.stanford.edu/102/)

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### character

```cpp
#include <iostream>
#include <string.h>
using namespace std;

int main()
{
	// string literals are regular arrays
	//
	// \0 is a null character
	char bt0[] = {'H', 'e', 'l', 'l', 'o', '\0'};
	cout << "bt0 length: " << strlen(bt0) << endl;
	// bt0 length: 5

	char bt1[] = "Hello";
	char* bt2 = "Hello"; // deprecated conversion from string constant to ‘char*’
	cout << "bt2 length: " << strlen(bt2) << endl;
	// bt2 length: 5

	cout << (bt0 == bt1) << endl; // 0

	cout << bt0 << endl << bt1 << endl << bt2 << endl;
	// Hello
	// Hello
	// Hello

	int i = 0;
	while (bt0[i] != '\0'){
		cout << bt0[i];
		i++;
	}
	cout << endl; // Hello
	i = 0;
	while (bt1[i] != '\0'){
		cout << bt1[i];
		i++;
	}
	cout << endl; // Hello

	// Is character array mutable? Yes.
	bt0[0] = 'A';
	cout << bt0 << endl;
	// Aello

	typedef unsigned char BYTE;
	BYTE text[] = "text";
	cout << text << endl; // text
}

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### `string`

```cpp
#include <iostream>

int main()
{
	std::string st1 = "Hello";
	std::string st2 = "Hello";

	std::cout << (st1 == st2) << std::endl;
	// 1
	
	// Is string mutable?
	// No.
	std::cout << st1[0] << std::endl; // H
	
	// st1[0] = "A";
	// invalid conversion from 'const char*' to 'char'
	
	std::cout << st1 << std::endl;
	// Hello
}

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### reverse string

```cpp
#include <iostream>
#include <string>
#include <algorithm>
#include <string.h>
using namespace std;

void reverseInPlace(char* str);
char* reverseReturn(char* str);
string reverseRecursive(string str);

int main()
{
	string str = "Hello World!";
	reverse(str.begin(), str.end());
	cout << str << endl; // !dlroW olleH

	char st1[] = "Hello World!";
	reverseInPlace(st1);
	cout << st1 << endl; // !dlroW olleH

	char st2[] = "Hello World!";
	char* rs2 = reverseReturn(st2);
	cout << rs2 << endl; // !dlroW olleH
	delete [] rs2;

	string st3 = "Hello World!";
	cout << reverseRecursive(st3) << endl; // !dlroW olleH
}

void reverseInPlace(char* str) {
	// unsigned integer type
	// type able to represent the size of any object in bytes
	size_t size = strlen(str);
	if (size < 2) {
		return;
	}
	for ( size_t i = 0, j = size - 1; i < j; i++, j-- ) {
		char tempChar = str[i];
		str[i] = str[j];
		str[j] = tempChar;
	}
}

// DO NOT define with char s[]
char* reverseReturn(char* str)
{
	int length = strlen(str);

	// char bts[length];
	char* bts = (char*)malloc(length);
	// Dynamic allocation needs to be 
	// deallocated manually

	int i, j;
	for (i=0, j=length-1; i < j; ++i, --j)
	{
		bts[i] = str[j];
		bts[j] = str[i];
	}

	return bts;
}

string reverseRecursive(string str)
{
	if (str.length() == 1)
		return str;
	return reverseRecursive(str.substr(1, str.length())) + str.at(0);
}

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### `string` to character array

```cpp
#include <iostream>
#include <string.h>
using namespace std;

int main() 
{
	string str = "Hello";
	char *cstr = new char[str.length() + 1];
	strcpy(cstr, str.c_str());

	cout << cstr << endl;
	// Hello
	
	// Deallocate storage space of array
	delete [] cstr;
}

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### STL `string`

```cpp
#include <iostream>
#include <string>
using namespace std;

int main()
{
	string str = "Hello World!!";
	str.pop_back();
	cout << str << endl;
	// Hello World!

	cout << str.back() << endl;  // !
	cout << str.empty() << endl; // 0
}

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### pointer

Pointer is a variable to store the address value of data(variable).
`&num` returns the address of `num`. `int * num` defines a pointer type
variable `num`. `*` is also used to dereference *or access* the memory
that the pointer points to.

```cpp
#include <iostream>
using namespace std;

int main()
{
	int val = 10;
	int* valPt = &val;
	*valPt = 100;
	cout << val << endl;
	// 100
}

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### pointer, array

```cpp
#include <iostream>
using namespace std;

int main() 
{
	// array points to the first element
	int numbers[5];
	int* pt;
	pt = numbers;
	*pt = 10;

	*(pt + 3) = 40;
	int* tp = numbers + 4;
	*tp = 50;

	pt++; *pt = 20;
	pt++; *pt = 30;

	for (int i=0; i<5; ++i)
		cout << numbers[i] << ", ";
	cout << endl;
	// 10, 20, 30, 40, 50, 
}

```

Note that the name of `array` is defined as pointer.
**And you cannot update the address of original `array`**:

```cpp
#include <stdio.h>

int main() 
{
	int arr[3] = {0, 1, 2};
	printf("Name of array: %p\n", arr);
	printf("&arr[0]: %p\n", &arr[0]);
	printf("&arr[1]: %p\n", &arr[1]);
	printf("&arr[2]: %p\n", &arr[2]);
	/*
		Name of array: 0x7ffc539d5a60
		&arr[0]: 0x7ffc539d5a60
		&arr[1]: 0x7ffc539d5a64
		&arr[2]: 0x7ffc539d5a68
	*/

	// int num = 0;
	// int* pnum = &num;
	// arr = pnum;

	// arr = &arr[1];
}

/*
07_pointer_array.cpp:13:6: error: incompatible types in assignment of ‘int*’ to ‘int [3]’
  arr = pnum;
      ^
07_pointer_array.cpp:15:6: error: incompatible types in assignment of ‘int*’ to ‘int [3]’
  arr = &arr[1];
      ^
*/

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### ampersand `&`

```cpp
#include <iostream>
using namespace std;

void updateByValue( int n )
{
	n = 1;
}


// & in C++
// 1. Declare a variable as a reference.
// 2. Take the address of a variable.

// compiler automatically takes the 
// address(reference) of argument.
void updateByRef( int& n )
{
	n = 2;
}

// need to pass the pointer 
// explicitly
void updateByPtr( int *n )
{
	*n = 3;
}

int main()
{
	int i=5;

	updateByValue( i ); 
	cout << i << endl; // 5

	updateByRef( i );
	cout << i << endl; // 2
 
	updateByPtr( &i );
  	cout << i << endl; // 3
}

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>


#### pointer, array, function

```cpp
#include <iostream>
using namespace std;

// (X) C++ does not allow to pass an entire
// array as an argument to a function.
// However, You can pass a pointer to an array.
int getSize1(int nums[])
{
	size_t size = sizeof(nums) / sizeof(nums[0]);
	return size;
}

// (X) C++ does not allow to pass an entire
// array as an argument to a function.
// However, You can pass a pointer to an array.
int getSize2(int* nums)
{
	size_t size = sizeof(nums) / sizeof(nums[0]);
	return size;
}

void printArray(int arr[], int size)
{
	int i;
	for (i=0; i < size; i++)
		printf("%d ", arr[i]);
	printf("\n");
}

int main()
{
	int nums[] = {0, 1, 3};
	cout << "*nums: " << *nums << endl;
	cout << "*(nums+2): " << *(nums+2) << endl;
	size_t size1 = sizeof(nums) / sizeof(nums[0]);
	size_t size2 = sizeof(nums) / sizeof(*nums);
	cout << "size1: " << size1 << endl;
	cout << "size2: " << size2 << endl;
	cout << "getSize1: " << getSize1(nums) << endl;
	cout << "getSize2: " << getSize2(nums) << endl;
	printArray(nums, 3); // 0 1 3
}

/*
*nums: 0
*(nums+2): 3
size1: 3
size2: 3
getSize1: 2
getSize2: 2
*/

```

[↑ top](#c-character-string-pointer-memory)
<br><br><br><br>
<hr>
