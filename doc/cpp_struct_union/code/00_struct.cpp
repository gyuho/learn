#include <iostream>
#include <stdio.h>
#include <string.h>
using namespace std;

int main()
{
	struct data
	{
		string company; 
		char   name[6];
		int    number;
	};

	data dt;
	dt.company = "Google";
	string str = "Hello";
	strcpy(dt.name, str.c_str());
	dt.number = 100;

	cout << dt.company << endl; // Google
	cout << dt.name << endl;    // Hello
	cout << dt.number << endl;  // 100
}
