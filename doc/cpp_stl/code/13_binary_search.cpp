#include <iostream>     // cout
#include <algorithm>    // binary_search, sort, find
#include <vector>       // vector
using namespace std;

bool compare (int i,int j) { return (i<j); }

int main () {
	int numbers[] = {1,2,3,4,5,4,3,2,1};
	vector<int> v(numbers, numbers+9);

	// using default comparison:
	sort (v.begin(), v.end());
	cout << "looking for a 3: ";
	if (binary_search (v.begin(), v.end(), 3))
		cout << "found!\n";
	else
		cout << "not found.\n";

	cout << "looking for a 6: ";
	sort (v.begin(), v.end(), compare);
	if (binary_search (v.begin(), v.end(), 6, compare))
		cout << "found!\n";
	else
		cout << "not found.\n";

	cout << "looking for a 5: ";
	vector<int>::iterator it = find (v.begin(), v.end(), 5);
	if (it != v.end())
		cout << "found! " << *it << endl;
	else
		cout << "not found.\n";
}

/*
looking for a 3: found!
looking for a 6: not found.
looking for a 5: found! 5
*/
