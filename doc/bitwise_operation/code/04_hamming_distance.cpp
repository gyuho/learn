#include <iostream>
#include <cstring>
// #include <string.h>
#include <stdexcept>
using namespace std;

int hammingDistance(char txt1[], char txt2[]) {
	if (strlen(txt1) != strlen(txt2))
		throw invalid_argument("Undefined for sequences of unequal length");
	int count = 0;
	size_t size = strlen(txt1);
	for (int idx=0; idx < size; ++idx)
	{
		char b1 = txt1[idx];
		char b2 = txt2[idx];
		unsigned int xorBit = b1 ^ b2;

		for (int x=xorBit; x > 0; x >>= 1)
		{
			if (int(x & 1) == 1)
				count++;
		}
	}
	return count;
}

int main()
{
	char txt1[] = {'H', 'e', 'l', 'l', 'o', '\0'};
	char txt2[] = "Hello";
	// final null character ('\0') is appended automatically
	cout << (txt1 == txt2) << endl; // 0
	
	cout << "hammingDistance: " << hammingDistance(txt1, txt2) << endl;
	// hammingDistance: 0 

	char txt3[] = "A";
	char txt4[] = "a";
	cout << hammingDistance(txt3, txt4) << endl; // 1

	char txt5[] = "aaa";
	char txt6[] = "aba";
	cout << hammingDistance(txt5, txt6) << endl; // 2

	strncpy(txt6, "aBa", sizeof(txt6));
	cout << hammingDistance(txt5, txt6) << endl; // 3

	char txt7[] = "karolin";
	char txt8[] = "kathrin";
	cout << hammingDistance(txt7, txt8) << endl; // 9
}
