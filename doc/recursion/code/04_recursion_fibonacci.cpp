#include <iostream>
using namespace std;

long unsigned int fib(long unsigned int num) {
	if (num == 0)
		return 0;
	else if (num == 1)
		return 1;
	else
		return fib(num-1) + fib(num-2);
}

int main()
{
	for (int i=0; i<15; ++i)
		cout << fib(i) << ", ";
	cout << endl;
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 
}
