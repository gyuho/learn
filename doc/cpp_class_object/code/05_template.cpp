#include <iostream>
using namespace std;

template <class T>
class myClass {
	T a, b;

	public:
		// constructor
		myClass (T first, T second) {
			a=first; 
			b=second;
		}

		T getMax ();
};

template <class T>
T myClass<T>::getMax ()
{
	T val;
	val = a>b? a : b;
	return val;
}

int main () {
	myClass <int> myObject (100, 75);
	cout << myObject.getMax() << endl;
	// 100
}
