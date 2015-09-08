#include <iostream>
using namespace std;

int find_stack_direction (int *num1)
{
	int num2 = 1;
	if (*num1 == 0)
	{
		num1 = &num2;
		return find_stack_direction (num1);
	}
	else
	{
		return ((&num2 > num1) ? 1 : -1);
	}
}

int main() {
	int num1 = 0;
	cout << find_stack_direction(&num1) << endl;
	// -1
	// (downwards)
}
