#include <iostream>
#include <forward_list>
using namespace std;

int main()
{	
	forward_list<int> mylist = { 34, 77, 16, 2 };

	cout << "mylist contains:";
	for ( auto it = mylist.begin(); it != mylist.end(); ++it )
		cout << ' ' << *it;
	cout << '\n';
	// mylist contains: 34 77 16 2
	

	mylist.front() = 11;
	cout << "mylist contains:";
	for ( auto it = mylist.begin(); it != mylist.end(); ++it )
		cout << ' ' << *it;
	cout << '\n';
	// mylist contains: 11 77 16 2
}
