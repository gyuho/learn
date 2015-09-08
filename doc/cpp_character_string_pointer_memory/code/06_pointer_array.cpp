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
