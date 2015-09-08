#include <iostream>
#include <unordered_map>
using namespace std;

int main() {
	unordered_map<string,double> mmap = {
		{"A", 1.0},
		{"B", 2.0},
		{"B", 20.0},
		{"C", 3.0}
	};

	unordered_map<string,double>::const_iterator got = mmap.find("Q");
	if (got != mmap.end())
		cout << got->first << " => " << got->second << endl;
	else
		cout << "Q is not found" << endl;
	// Q is not found
	
	cout << endl;

	unordered_map<string,double>::iterator it = mmap.find("B");
	if (it != mmap.end())
		cout << it->first << " => " << it->second << endl;
	else
		cout << "B is not found" << endl;
	// B => 2
	
	cout << endl;

	mmap["B"] = 30.0;

	unordered_map<string,double>::iterator ib = mmap.find("B");
	if (ib != mmap.end())
		cout << ib->first << " => " << ib->second << endl;
	else
		cout << "B is not found" << endl;
	// B => 30

	unordered_map<string,double>::iterator iter = mmap.find("B");
	if (iter != mmap.end())
		mmap.erase(iter);
		cout << "Deleted" << endl;

	for (unordered_map<string,double>::iterator it=mmap.begin(); it!=mmap.end(); ++it)
		cout << it->first << " => " << it->second << '\n';
}

/*
C => 3
A => 1
*/
