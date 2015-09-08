[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# C++: class, object

- [Object Oriented Programming](#object-oriented-programming)
- [`class`](#class)
- [`class`: `private`, `protected`, `public`](#class-private-protected-public)
- [constructor](#constructor)
- [overload constructor](#overload-constructor)
- [overload operator](#overload-operator)
- [inheritance](#inheritance)
- [template](#template)

[↑ top](#c-class-object)
<br><br><br><br>
<hr>









### Object Oriented Programming

*Steve Jobs* explains [here](http://www.edibleapple.com/2011/10/29/steve-jobs-explains-object-oriented-programming/):

> **_Objects_** are like people. They’re living, breathing 
> things that have knowledge inside them about 
> **how to do things** and **have memory** inside them
> so they can remember things.

An object can contains:

- how the object does things - `method`
- memory to store things

[↑ top](#c-class-object)
<br><br><br><br>
<hr>








#### `class`

```cpp
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

```

[↑ top](#c-class-object)
<br><br><br><br>
<hr>











#### `class`: `private`, `protected`, `public`

```cpp
class MyClass {
    private:
        int privateMember;
    protected:
        int protectedMember;
    public:
        int publicMember;
};
```

- `private`: no class but `MyClass` can access this.
- `protected`: most accessible level is at the members inherited from `MyClass`.
- `public`: everything can access.

[↑ top](#c-class-object)
<br><br><br><br>
<hr>












#### constructor

```cpp
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

```

[↑ top](#c-class-object)
<br><br><br><br>
<hr>








#### overload constructor

```cpp
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

```

[↑ top](#c-class-object)
<br><br><br><br>
<hr>









#### overload operator

```cpp
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

```

[↑ top](#c-class-object)
<br><br><br><br>
<hr>












#### inheritance

```cpp
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

```

[↑ top](#c-class-object)
<br><br><br><br>
<hr>









#### template

```cpp
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

```

[↑ top](#c-class-object)
<br><br><br><br>
<hr>