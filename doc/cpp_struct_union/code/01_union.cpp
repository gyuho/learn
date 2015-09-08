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
