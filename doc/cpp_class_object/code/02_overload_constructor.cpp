#include <iostream>
using namespace std;

class Rectangle {
	int width, height;
	public:
		
		// constructor
		Rectangle (int, int);

		// overload constructor
		Rectangle();
		
		void print() {
			cout << "Width is " << width << endl;
			cout << "Height is " << height << endl;
			cout << "Area is " << width * height << endl;
		};
};


Rectangle::Rectangle (int x, int y) {
	width = x;
	height = y;
};


// constructor overloading
// just like function overloading
Rectangle::Rectangle() {
	width = 100;
	height = 200;
}


int main() 
{
	Rectangle rect(5, 10);
	rect.print();

	Rectangle rb;
	rb.print();
}

/*
Width is 5
Height is 10
Area is 50

Width is 100
Height is 200
Area is 20000
*/
