#include <iostream>     // cout
#include <algorithm>    // min
using namespace std;

int main () {
	cout << "min(1,2)==" << min(1,2) << endl;
	cout << "min(2,1)==" << min(2,1) << endl;
	cout << "min('a','z')==" << min('a','z') << endl;
	cout << "min(3.14,2.72)==" << min(3.14,2.72) << endl;

	cout << "max(1,2)==" << max(1,2) << endl;
	cout << "max(2,1)==" << max(2,1) << endl;
	cout << "max('a','z')==" << max('a','z') << endl;
	cout << "max(3.14,2.72)==" << max(3.14,2.72) << endl;
}

/*
min(1,2)==1
min(2,1)==1
min('a','z')==a
min(3.14,2.72)==2.72
max(1,2)==2
max(2,1)==2
max('a','z')==z
max(3.14,2.72)==3.14
*/