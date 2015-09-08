#include <iostream>     // cout
#include <algorithm>    // sort
#include <vector>       // vector
using namespace std;

bool myfunction (int i,int j) {
	return (i<j);
}

struct myClass {
	bool operator() (int i,int j) {
		return (i<j);
	}
} myObject;

int main () {
	int numbers[] = {32,71,12,45,26,80,53,33};
	vector<int> v (numbers, numbers+8);

	// using default comparison (operator <):
	sort (v.begin(), v.begin()+4);
	//(12 32 45 71)26 80 53 33

	// using function as comp
	sort (v.begin()+4, v.end(), myfunction);
	// 12 32 45 71(26 33 53 80)

	// using object as comp
	sort (v.begin(), v.end(), myObject);
	//(12 26 32 33 45 53 71 80)

	cout << "contains:";
	for (vector<int>::iterator it=v.begin(); it!=v.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
}

// contains: 12 26 32 33 45 53 71 80
