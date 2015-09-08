#include <iostream>
using namespace std;

int main()
{
	// string literals are regular arrays
	//
	// \0 is a null character
	char bt1[] = {'H', 'e', 'l', 'l', 'o', '\0'};
	char bt2[] = "Hello";
	// final null character ('\0') is appended automatically
	cout << (bt1 == bt2) << endl; // 0

	cout << bt1 << endl << bt2 << endl;
	// Hello
	// Hello
	
	int i = 0;
	while (bt1[i] != '\0'){
		cout << bt1[i];
		i++;
	}
	cout << endl; // Hello
	i = 0;
	while (bt2[i] != '\0'){
		cout << bt2[i];
		i++;
	}
	cout << endl; // Hello

	// Is character array mutable? Yes.
	bt1[0] = 'A';
	cout << bt1 << endl;
	// Aello

	typedef unsigned char BYTE;
	BYTE text[] = "text";
	cout << text << endl; // text
}
