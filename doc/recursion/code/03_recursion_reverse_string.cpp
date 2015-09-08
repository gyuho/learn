#include <iostream>
#include <string>
#include <algorithm>
#include <string.h>
using namespace std;

void reverseInPlace(char* str);
char* reverseReturn(char* str);
string reverseRecursive(string str);

int main()
{
	string str = "Hello World!";
	reverse(str.begin(), str.end());
	cout << str << endl; // !dlroW olleH

	char st1[] = "Hello World!";
	reverseInPlace(st1);
	cout << st1 << endl; // !dlroW olleH

	char st2[] = "Hello World!";
	char* rs2 = reverseReturn(st2);
	cout << rs2 << endl; // !dlroW olleH
	delete [] rs2;

	string st3 = "Hello World!";
	cout << reverseRecursive(st3) << endl; // !dlroW olleH
}

void reverseInPlace(char* str) {
	// unsigned integer type
	// type able to represent the size of any object in bytes
	size_t size = strlen(str);
	if (size < 2) {
		return;
	}
	for ( size_t i = 0, j = size - 1; i < j; i++, j-- ) {
		char tempChar = str[i];
		str[i] = str[j];
		str[j] = tempChar;
	}
}

// DO NOT define with char s[]
char* reverseReturn(char* str)
{
	int length = strlen(str);

	// char bts[length];
	char* bts = (char*)malloc(length);
	// Dynamic allocation needs to be 
	// deallocated manually

	int i, j;
	for (i=0, j=length-1; i < j; ++i, --j)
	{
		bts[i] = str[j];
		bts[j] = str[i];
	}

	return bts;
}

string reverseRecursive(string str)
{
	if (str.length() == 1)
		return str;
	return reverseRecursive(str.substr(1, str.length())) + str.at(0);
}
