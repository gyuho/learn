#include <iostream>
using namespace std;

int main()
{
	bool b1 = true;
	bool b2 = false;
	cout << b1 && b2; // 1
	cout << endl;
	cout << b1 || b2; // 1
	cout << endl;

	int a = 10, b = 15, c = 20;
	cout << "a < b is "
		<< (a < b)
		<< endl;
	// a < b is 1
	
	cout << "b > c is "
		<< (b > c)
		<< endl;
	// b > c is 0
}
