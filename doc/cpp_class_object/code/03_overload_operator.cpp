#include <iostream>
using namespace std;

class myData {
	public:
		int x, y;
		myData () {}
		myData (int a, int b) : x(a), y(b) {}
};

myData operator+ (const myData& lhs, const myData& rhs) {
	myData temp;
	temp.x = lhs.x + rhs.x;
	temp.y = lhs.y + rhs.y;
	return temp;
}

int main () {
	myData d1 (3,1);
	myData d2 (1,2);
	myData result;
	result = d1 + d2;
	cout << result.x << ',' << result.y << endl;
	// 4,3
}