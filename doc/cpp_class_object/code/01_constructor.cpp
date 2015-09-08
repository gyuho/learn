#include <iostream>
using namespace std;

class Rectangle {
	int width, height;
	public:
		
		// void set(int, int);
		
		// constructor
		Rectangle (int, int);

		int area() {
			return width * height;
		};

		void print() {
			cout << "Width is " << width << endl;
			cout << "Height is " << height << endl;
			cout << "Area is " << width * height << endl;
		};
};

// void Rectangle::set(int x, int y) {
// 	width = x;
// 	height = y;
// }

Rectangle::Rectangle (int x, int y) {
	width = x;
	height = y;
};

int main() 
{
	Rectangle rect(5, 10);
	rect.print();
}

/*
Width is 5
Height is 10
Area is 50
*/
