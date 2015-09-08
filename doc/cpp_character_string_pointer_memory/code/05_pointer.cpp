#include <iostream>
using namespace std;

int main()
{
	int val = 10;
	int* valPt = &val;
	*valPt = 100;
	cout << val << endl;
	// 100
}
