#include <iostream>
using namespace std;

class Rectangle {
	int width, height;
	public:
		void set(int, int);
		int area() {
			return width * height;
		};
		void print() {
			cout << "Width is " << width << endl;
			cout << "Height is " << height << endl;
			cout << "Area is " << width * height << endl;
		};
};

void Rectangle::set(int x, int y) {
	width = x;
	height = y;
}

int main() 
{
	Rectangle rect;
	rect.print();
	rect.set(5, 10);
	cout << "Area: " << rect.area() << endl;
	rect.print();
}

/*
Width is 4196944
Height is 0
Area is 0

Area: 50

Width is 5
Height is 10
Area is 50
*/
