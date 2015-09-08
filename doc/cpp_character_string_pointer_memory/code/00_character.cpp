#include <iostream>
using namespace std;

int main()
{
	// string literals are regular arrays
	//
	// \0 is a null character
	char bt0[] = {'H', 'e', 'l', 'l', 'o', '\0'};
	char bt1[] = "Hello";
	char* bt2 = "Hello"; // deprecated conversion from string constant to ‘char*’
	cout << (bt0 == bt1) << endl; // 0

	cout << bt0 << endl << bt1 << endl << bt2 << endl;
	// Hello
	// Hello
	// Hello

	int i = 0;
	while (bt0[i] != '\0'){
		cout << bt0[i];
		i++;
	}
	cout << endl; // Hello
	i = 0;
	while (bt1[i] != '\0'){
		cout << bt1[i];
		i++;
	}
	cout << endl; // Hello

	// Is character array mutable? Yes.
	bt0[0] = 'A';
	cout << bt0 << endl;
	// Aello

	typedef unsigned char BYTE;
	BYTE text[] = "text";
	cout << text << endl; // text
}
