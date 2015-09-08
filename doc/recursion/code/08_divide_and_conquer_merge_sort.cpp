#include <iostream>     // cout
#include <algorithm>    // merge, sort
#include <vector>       // vector
using namespace std;

int main () {
	int first[] = {5,10,15,20,25};
	int second[] = {50,40,30,20,10};
	vector<int> v(10);

	sort(first,first+5);
	sort(second,second+5);
	merge(first, first+5, second, second+5, v.begin());

	for (vector<int>::iterator it=v.begin(); it!=v.end(); ++it)
		cout << ' ' << *it;
	cout << '\n';
	//  5 10 10 15 20 20 25 30 40 50
}
