#include <iostream>
#include <unordered_set>
using namespace std;

int main() {
	unordered_set<int> myset;
	for (int i=1; i<5; ++i)
		myset.insert(i*10);
	for (int i=1; i<5; ++i)
		myset.insert(i*10);

	cout << "myset size: " << myset.size() << endl;
	cout << endl;

	for (unordered_set<int>::iterator it=myset.begin(); it!=myset.end(); ++it)
		cout << *it << ' ';
	cout << endl;
	cout << endl;

	for (int i=1; i<5; ++i)
		cout << i << " count: " << myset.count(i) << endl;
	cout << endl;
	for (int i=1; i<5; ++i)
		cout << i*10 << " count: " << myset.count(i*10) << endl;
	cout << endl;
	unordered_set<int>::iterator it = myset.find(30);
	if (it != myset.end())
		cout << "30 is in the set";
	cout << endl;
}

/*
myset size: 4

40 30 20 10

1 count: 0
2 count: 0
3 count: 0
4 count: 0

10 count: 1
20 count: 1
30 count: 1
40 count: 1

30 is in the set
*/
