#include <iostream>
#include <stdio.h>
#include <string>
#include <bitset>
using namespace std;

long unsigned int toLu(string bstr) {
	// auto size = bstr.length();
	// to_ulong converts to unsigned long integer

	// Templates are evaluated when compiled
	const int size = 100;
	long unsigned int num = bitset<size>(bstr).to_ulong();
	printf("%10s (decimal %lu)\n", bstr.c_str(), num);
	return num;
}

int main()
{
	cout << bitset<100>("101").to_ulong() << endl;
	// 5

	cout << toLu("101") << endl;
	// 5
	//       101 (decimal 5)

	bitset<4> mybits; // mybits: 0000
	cout << "mybits.set(): " << mybits.set() << endl; // mybits: 1111
	cout << "mybits.set(2,0): " << mybits.set(2,0) << endl;
	cout << "mybits.set(2): " << mybits.set(2) << endl;
	cout << "mybits.flip(): " << mybits.flip() << endl;
	cout << "mybits.flip(1): " << mybits.flip(1) << endl;
	cout << "mybits.reset(1): " << mybits.reset(1) << endl;
}

/*
mybits.set(): 1111
mybits.set(2,0): 1011
mybits.set(2): 1111
mybits.flip(): 0000
mybits.flip(1): 0010
mybits.reset(1): 0000
*/