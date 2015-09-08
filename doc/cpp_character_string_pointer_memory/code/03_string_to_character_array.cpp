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
