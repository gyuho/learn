#include <iostream>
using namespace std;

// protected members can be
// accessed by inheriting classes

class Polygon {
	protected:
		int width, height;
	public:
		void set (int a, int b) { width=a; height=b;}
};

class Rectangle: public Polygon {
	public:
		int area () { return width * height; }
};

class Triangle: public Polygon {
	public:
		int area () { return width * height / 2; }
 };
  
int main () {
	Rectangle rect;
	Triangle trgl;
	rect.set (50, 100);
	trgl.set (50, 100);
	cout << rect.area() << endl; // 5000
	cout << trgl.area() << endl; // 2500
}