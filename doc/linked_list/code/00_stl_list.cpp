#include <iostream>
#include <list>
using namespace std;

int main()
{
	double myDs[] = {1.2, -5.3, 10.2, 1, -100.23};
	list<double> myList (myDs, myDs+5);

	cout << "myList contains: " << endl;
	for (list<double>::iterator it=myList.begin(); it!=myList.end(); ++it)
		cout << " " << *it;
	cout << endl;

	myList.front() = -100.55;

	cout << "myList contains: " << endl;
	for (list<double>::iterator it=myList.begin(); it!=myList.end(); ++it)
		cout << " " << *it;
	cout << endl;

	for (int i=0; i<=1000; ++i)
		myList.push_back(i);

	int sum(0);
	while (!myList.empty())
	{
		sum += myList.front();
		myList.pop_front();
	}

	cout << "total: " << sum << endl; 
}

/*
myList contains: 
 1.2 -5.3 10.2 1 -100.23
myList contains: 
 -100.55 -5.3 10.2 1 -100.23
total: 500307
*/
