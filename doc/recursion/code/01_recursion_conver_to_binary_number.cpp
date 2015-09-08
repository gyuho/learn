#include <iostream>
using namespace std;

long unsigned int toBinaryNumber(long unsigned int num) {
	if (num == 0)
		return 0;
	return (num % 2) + 10*toBinaryNumber(num/2);
}

int main()
{
	cout << toBinaryNumber(15) << endl; // 1111
}
