#include <iostream>
#include <string>
using namespace std;

int main()
{
	string str = "Hello World!!";
	str.pop_back();
	cout << str << endl;
	// Hello World!

	cout << str.back() << endl;  // !
	cout << str.empty() << endl; // 0
}
