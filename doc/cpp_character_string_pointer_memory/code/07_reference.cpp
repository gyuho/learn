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
