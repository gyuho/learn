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
