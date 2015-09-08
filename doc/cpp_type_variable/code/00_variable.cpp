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