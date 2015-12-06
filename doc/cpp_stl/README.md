[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C++: STL

- [Reference](#reference)
- [`array`](#array)
- [`vector`](#vector)
- [`list`](#list)
- [`forward_list`](#forward_list)
- [`stack`](#stack)
- [`queue`](#queue)
- [`heap`](#heap)
- [`priority_queue`](#priority_queue)
- [`map`](#map)
- [`unordered_map`](#unordered_map)
- [`set`](#set)
- [`unordered_set`](#unordered_set)
- [`binary_search`](#binary_search)
- [`min`, `max`](#min-max)
- [`sort`](#sort)

[↑ top](#c-stl)
<br><br><br><br><hr>


#### Reference

- [Arrays C++](https://msdn.microsoft.com/en-us/library/7wkxxx2e.aspx)
- [cppreference.com](http://en.cppreference.com/w/)
- [cplusplus.com/reference](http://www.cplusplus.com/reference/)

[↑ top](#c-stl)


#### `array`

```cpp
#include <iostream>
#include <stdio.h>

int main()
{
	int sz = 3;
	int* arr = new int(sz);

	for (int i=0; i<sz; i++)
		arr[i] = 100;

	std::cout << arr << std::endl;
	// 0xd6f010
	
	for (int i=0; i<sz; i++)
		printf ("%d\n", arr[i]);
	// 100
	// 100
	// 100
	
	double twoArray[][4] = { 
	   { 32.19, 47.29, 31.99, 19.11 },
	   { 11.29, 22.49, 33.47, 17.29 },
	   { 41.97, 22.09,  9.76, 22.55 }  
	};
	std::cout << twoArray << std::endl;
	
	printf ("Done\n");

	// dynamic array must be deleted
	delete [] arr;


	// with array
	size_t size = 10;

	// static(or local) arrays are created on the stack
	// and get destroyed automatically after function exit.
	// They have a fixed size.
	int staticArray[10];

	// dynamic arrays are stored on the heap.
	// They can have anysize but you need to allocate,
	// free them manually.
	int* dynamicArray = new int[size];

	for (int i=0; i<10; i++)
	{
		staticArray[i] = i;
		dynamicArray[i] = i;
	}
	
	for (int i=0; i<10; i++) 
	{
		std::cout << staticArray[i] << std::endl;
		std::cout << dynamicArray[i] << std::endl;
	}

	delete[] dynamicArray;
}

```

```cpp
#include <iostream>
#include <array>
using namespace std;

// we cannot resize the array container
// while we can resize vector container with push_back, pop_back

int main()
{
	int arr1[10] = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
	for ( auto elem : arr1 )
	{
		cout << elem << ", ";
	}
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	cout << endl;
	// If you need to modify the original
	// we need auto&
	for ( auto& elem : arr1 )
	{
		cout << elem << ", ";
	}
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,
	
	cout << endl;
	array<int, 10> arr2;
	for (int i=0 ; i < 10 ; i++)
	{
		arr2.at(i) = i+1;
	}
	for ( auto elem : arr2 )
	{
		cout << elem << ", ";
	}
	cout << endl;
	// 1, 2, 3, 4, 5, 6, 7, 8, 9, 10,

	cout << endl;
	cout << "auto elem : arr2" << endl;
	for ( auto elem : arr2 )
	{
		elem = 100;
	}
	for ( auto elem : arr2 )
	{
		cout << elem << ", ";
	}
	cout << endl;
	
	cout << endl;
	cout << "auto& elem : arr2" << endl;
	for ( auto& elem : arr2 )
	{
		elem = 100;
	}
	for ( auto elem : arr2 )
	{
		cout << elem << ", ";
	}
	cout << endl;
	
	cout << endl;
	cout << "auto it = arr2.begin()" << endl;
	// no need to put &
	for ( auto it = arr2.begin(); it != arr2.end(); ++it )
		*it = 1000;
	for ( auto it = arr2.begin(); it != arr2.end(); ++it )
	{
		cout << *it << ", ";
	}
	cout << endl;
	for ( array<int,10>::iterator it = arr2.begin(); it != arr2.end(); ++it )
	{
		cout << *it << ", ";
	}
	cout << endl;
		
	cout << endl;
	cout << "The size of array container is " << arr2.size() << endl;
}

/*
auto elem : arr2
1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 

auto& elem : arr2
100, 100, 100, 100, 100, 100, 100, 100, 100, 100, 

auto it = arr2.begin()
1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 
1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 1000, 

The size of array container is 10
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `vector`

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void erase(vector<string>& values, int pos);
void print(vector<string> values);

int main()
{
	// declare the vector of integers
	// vector<int> iv;
	typedef vector<int> int_vec_t;

	int_vec_t iv;
	// iv.push_back(1);
	iv.push_back(2);
	iv.push_back(3);
	iv.push_back(4);
	iv.push_back(5);
	// push front
	iv.insert(iv.begin(), 1);

	iv.erase(iv.begin()+2);
	iv.erase(iv.begin(), iv.begin()+2);

	for (int_vec_t::iterator it = iv.begin(); it != iv.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
	// 4 5 

	for (auto& it : iv)
		cout << ' ' << it;
	cout << endl;
	// 4 5 
	
	vector<string> strVector;
	strVector.push_back("A");
	strVector.push_back("B");
	strVector.push_back("C");

	cout << strVector.size() << endl; // 3

	while (!strVector.empty())
	{
		strVector.pop_back();
	}

	cout << strVector.size() << endl; // 0

	cout << endl;

	vector<string> members(5);
	members[0] = "A";
	members[1] = "B";
	members[2] = "C";
	members[3] = "D";
	members[4] = "E";
	print(members);
	
	int pos;
	cout << "Remove which element? ";
	cin >> pos;
	
	erase(members, pos);
	print(members);

	int houses[] = {3, 15, 13, 4, 7};
	vector<int> hv;
	size_t size = sizeof(houses) / sizeof(houses[0]);
	for (int i=0; i<size; ++i)
	{
		hv.push_back(houses[i]);
	}
	vector<int> example(hv.begin()+1, hv.end());
	for (vector<int>::iterator it = example.begin(); it != example.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
	//  15 13 4 7

	cout << "The smallest element is " << *min_element(houses, houses+5) << '\n';
	cout << "The largest element is "  << *max_element(houses, houses+5) << '\n';
	// The smallest element is 3
	// The largest element is 15

	cout << "The smallest element is " << *min_element(hv.begin(), hv.end()) << '\n';
	cout << "The largest element is "  << *max_element(hv.begin(), hv.end()) << '\n';
	// The smallest element is 3
	// The largest element is 15
}

/**
 Removes an element from an unordered vector.
 @param values a vector
 @param pos the position of the element to erase
 */
void erase(vector<string>& values, int pos)
{
	int last_pos = values.size() - 1;
	values[pos] = values[last_pos];
	values.pop_back();
}

/**
 Prints all elements in a vector.
 @param values the vector to print
 */
void print(vector<string> values)
{
	for (int i = 0; i < values.size(); i++)
		cout << "[" << i << "] " << values[i] << "\n";
}

/*
[0] A
[1] B
[2] C
[3] D
[4] E
Remove which element? 3
[0] A
[1] B
[2] C
[3] E
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `list`

`list` implements doubly linked list:

```cpp
#include <iostream>
#include <list>
using namespace std;

int main()
{
	double myDs[] = {1.2, -5.3, 10.2, 1, -100.23};
	list<double> myList (myDs, myDs+5);

	cout << "myList contains: " << endl;
	for (list<double>::iterator it=myList.begin(); it!=myList.end(); ++it)
		cout << " " << *it;
	cout << endl;

	myList.front() = -100.55;

	cout << "myList contains: " << endl;
	for (list<double>::iterator it=myList.begin(); it!=myList.end(); ++it)
		cout << " " << *it;
	cout << endl;

	for (int i=0; i<=1000; ++i)
		myList.push_back(i);

	int sum(0);
	while (!myList.empty())
	{
		sum += myList.front();
		myList.pop_front();
	}

	cout << "total: " << sum << endl; 
}

/*
myList contains: 
 1.2 -5.3 10.2 1 -100.23
myList contains: 
 -100.55 -5.3 10.2 1 -100.23
total: 500307
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `forward_list`

`forward_list` implements singly linked list:

```cpp
#include <iostream>
#include <forward_list>
using namespace std;

int main()
{	
	forward_list<int> mylist = { 34, 77, 16, 2 };

	cout << "mylist contains:";
	for ( auto it = mylist.begin(); it != mylist.end(); ++it )
		cout << ' ' << *it;
	cout << '\n';
	// mylist contains: 34 77 16 2
	

	mylist.front() = 11;
	cout << "mylist contains:";
	for ( auto it = mylist.begin(); it != mylist.end(); ++it )
		cout << ' ' << *it;
	cout << '\n';
	// mylist contains: 11 77 16 2
}

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `stack`

```cpp
#include <iostream>
#include <stack>

using namespace std;

// Queue is First In First Out
// Stack is Last In First Out

int main()
{
	stack<string> s;
	s.push("A"); // add elements
	s.push("B");
	s.push("C");
	
	while (s.size() > 0)
	{
		// returns the top element first
		cout << s.top() << endl;
		
		// remove the element on top of the stack
		s.pop();
	}
	cout << endl;

	cout << "size of stack has become " << s.size() << endl;
	
	//  we cannot iterate through stack and queue
	//	for ( auto elem : q )
	//	{
	//		cout << "stack elements: " << elem;
	//	}
}

/*
C
B
A

size of stack has become 0
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `queue`

```cpp
#include <iostream>
#include <queue>

using namespace std;

// Queue is First In First Out
// Stack is Last In First Out

int main()
{
	queue<string> q;
	q.push("A"); // add elements
	q.push("B");
	q.push("C");
	
	while (q.size() > 0)
	{
		// returns the oldest, first element
		cout << q.front() << endl;
		
		// remove the element on top of the stack
		q.pop();
	}
	cout << endl;
	cout << "size of queue has become " << q.size() << endl;

	//  we cannot iterate through stack and queue
	//	for ( auto elem : q )
	//	{
	//		cout << "stack elements: " << elem;
	//	}
}

/*
A
B
C

size of queue has become 0
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `heap`

```cpp
#include <iostream>     // cout
#include <algorithm>    // make_heap, pop_heap, push_heap, sort_heap
#include <vector>       // vector
using namespace std;

int main () {
	int myints[] = {10,20,30,5,15};
	vector<int> v(myints,myints+5);

	cout << "befire max heap    : " << v.front() << endl;
	make_heap (v.begin(),v.end());
	cout << "initial max heap   : " << v.front() << endl;

	pop_heap (v.begin(),v.end());
	v.pop_back();
	cout << "max heap after pop : " << v.front() << endl;

	v.push_back(99);
	push_heap (v.begin(),v.end());
	cout << "max heap after push: " << v.front() << endl;

	// heap-sort
	sort_heap (v.begin(),v.end());

	cout << "final sorted range :";
	for (unsigned i=0; i<v.size(); i++)
		cout << ' ' << v[i];
	cout << endl;
}

/*
befire max heap    : 10
initial max heap   : 30
max heap after pop : 20
max heap after push: 99
final sorted range : 5 10 15 20 99
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `priority_queue`

```cpp
#include <iostream>
#include <queue>
#include <vector>
#include <functional>
using namespace std;

template<class T> using minpq = priority_queue<T, vector<T>, greater<T>>;

struct compare
{
	bool operator()(const int& l, const int& r)  
	{
		return l > r;  
	}
};

int main ()
{
	priority_queue<int> pq;
	pq.push(1000);
	pq.push(10);
	pq.push(20);
	pq.push(20);
	pq.push(17);
	pq.push(55);
	pq.push(15);
	pq.push(100);
	cout << "pq.top() is now " << pq.top() << endl; // pq.top() is now 20
	cout << "Popping out elements: ";
	while (!pq.empty())
	{
		cout << pq.top() << ' ';
		pq.pop();
	}
	cout << endl;
	cout << endl;

	minpq<int> minpq0;
	minpq0.push(1000);
	minpq0.push(10);
	minpq0.push(20);
	minpq0.push(20);
	minpq0.push(17);
	minpq0.push(55);
	minpq0.push(15);
	minpq0.push(100);
	cout << "minpq0.top() is now " << minpq0.top() << endl; // minpq0.top() is now 20
	cout << "Popping out elements: ";
	while (!minpq0.empty())
	{
		cout << minpq0.top() << ' ';
		minpq0.pop();
	}
	cout << endl;
	cout << endl;

	priority_queue<int, vector<int>, compare > minpq1;
	minpq1.push(1000);
	minpq1.push(10);
	minpq1.push(20);
	minpq1.push(20);
	minpq1.push(17);
	minpq1.push(55);
	minpq1.push(15);
	minpq1.push(100);
	cout << "minpq1.top() is now " << minpq1.top() << endl; // minpq1.top() is now 20
	cout << "Popping out elements: ";
	while (!minpq1.empty())
	{
		cout << minpq1.top() << ' ';
		minpq1.pop();
	}
	cout << endl;
	cout << endl;

	priority_queue<int, vector<int>, greater<int>> minpq2;
	minpq2.push(1000);
	minpq2.push(10);
	minpq2.push(20);
	minpq2.push(20);
	minpq2.push(17);
	minpq2.push(55);
	minpq2.push(15);
	minpq2.push(100);
	cout << "minpq2.top() is now " << minpq2.top() << endl; // minpq2.top() is now 20
	cout << "Popping out elements: ";
	while (!minpq2.empty())
	{
		cout << minpq2.top() << ' ';
		minpq2.pop();
	}
	cout << endl;
	cout << endl;

	priority_queue<int, vector<int>, less<int>> maxpq;
	maxpq.push(1000);
	maxpq.push(10);
	maxpq.push(20);
	maxpq.push(20);
	maxpq.push(17);
	maxpq.push(55);
	maxpq.push(15);
	maxpq.push(100);
	cout << "maxpq.top() is now " << maxpq.top() << endl; // maxpq.top() is now 20
	cout << "Popping out elements: ";
	while (!maxpq.empty())
	{
		cout << maxpq.top() << ' ';
		maxpq.pop();
	}
	cout << endl;
}

/*
pq.top() is now 1000
Popping out elements: 1000 100 55 20 20 17 15 10 

minpq0.top() is now 10
Popping out elements: 10 15 17 20 20 55 100 1000 

minpq1.top() is now 10
Popping out elements: 10 15 17 20 20 55 100 1000 

minpq2.top() is now 10
Popping out elements: 10 15 17 20 20 55 100 1000 

maxpq.top() is now 1000
Popping out elements: 1000 100 55 20 20 17 15 10 
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `map`

```cpp
#include <iostream>
#include <map>
using namespace std;

void updateMap1(map<char,int> cmap)
{
	cmap['O'] = 100;
}

void updateMap2(map<char,int>& cmap)
{
	cmap['O'] = 100;
}

void updateMap3(map<char,int>* cmap)
{
	// (X) cmap['O'] = 1000;
	(*cmap)['O'] = 1000;
}

int main()
{
	map<char,int> cmap;
	cmap['A'] = 100;
	cmap['B'] = 200;
	cmap['C'] = 300;
	cout << "cmap['X']: " << cmap['X'] << endl; // 0
	updateMap1(cmap); cout << "cmap['O']: " << cmap['O'] << endl; // 0
	updateMap2(cmap); cout << "cmap['O']: " << cmap['O'] << endl; // 100
	updateMap3(&cmap); cout << "cmap['O']: " << cmap['O'] << endl; // 1000

	for (map<char,int>::iterator it=cmap.begin(); it!=cmap.end(); ++it)
		cout << it->first << " => " << it->second << '\n';

	cout << endl;

	map<string,int> smap;
	smap["Hello"] = -100;
	smap["World"] = 200;
	smap["C++"] = 300;
	smap["Hello"] = 100;

	map<string,int>::iterator iter = smap.find("NOT");
	cout << (iter == smap.end()) << endl;
	// 1
	// 'NOT' does not exist in the map

	cout << endl;

	for (map<string,int>::iterator it=smap.begin(); it!=smap.end(); ++it)
		cout << it->first << " => " << it->second << '\n';
	cout << '\n';
	
	iter = smap.find("World");
	if (iter != smap.end())
		smap.erase(iter);
		cout << "Deleted" << endl;

	for (map<string,int>::iterator it=smap.begin(); it!=smap.end(); ++it)
		cout << it->first << " => " << it->second << '\n';
}

/*
A => 100
B => 200
C => 300

1

C++ => 300
Hello => 100
World => 200

Deleted
C++ => 300
Hello => 100
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `unordered_map`

```cpp
#include <iostream>
#include <unordered_map>
using namespace std;

int main() {
	unordered_map<string,double> mmap = {
		{"A", 1.0},
		{"B", 2.0},
		{"B", 20.0},
		{"C", 3.0}
	};

	unordered_map<string,double>::const_iterator got = mmap.find("Q");
	if (got != mmap.end())
		cout << got->first << " => " << got->second << endl;
	else
		cout << "Q is not found" << endl;
	// Q is not found
	
	cout << endl;

	unordered_map<string,double>::iterator it = mmap.find("B");
	if (it != mmap.end())
		cout << it->first << " => " << it->second << endl;
	else
		cout << "B is not found" << endl;
	// B => 2
	
	cout << endl;

	mmap["B"] = 30.0;

	unordered_map<string,double>::iterator ib = mmap.find("B");
	if (ib != mmap.end())
		cout << ib->first << " => " << ib->second << endl;
	else
		cout << "B is not found" << endl;
	// B => 30

	unordered_map<string,double>::iterator iter = mmap.find("B");
	if (iter != mmap.end())
		mmap.erase(iter);
		cout << "Deleted" << endl;

	for (unordered_map<string,double>::iterator it=mmap.begin(); it!=mmap.end(); ++it)
		cout << it->first << " => " << it->second << '\n';
}

/*
C => 3
A => 1
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `set`

```cpp
#include <iostream>
#include <set>
using namespace std;

int main() {
	set<int> myset;
	for (int i=1; i<5; ++i)
		myset.insert(i*10);
	for (int i=1; i<5; ++i)
		myset.insert(i*10);

	cout << "myset size: " << myset.size() << endl;
	cout << endl;

	for (set<int>::iterator it=myset.begin(); it!=myset.end(); ++it)
		cout << *it << ' ';
	cout << endl;
	cout << endl;

	for (int i=1; i<5; ++i)
		cout << i << " count: " << myset.count(i) << endl;
	cout << endl;
	for (int i=1; i<5; ++i)
		cout << i*10 << " count: " << myset.count(i*10) << endl;
	cout << endl;
	set<int>::iterator it = myset.find(30);
	if (it != myset.end())
		cout << "30 is in the set";
	cout << endl;
}

/*
myset size: 4

10 20 30 40 

1 count: 0
2 count: 0
3 count: 0
4 count: 0

10 count: 1
20 count: 1
30 count: 1
40 count: 1

30 is in the set
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `unordered_set`

```cpp
#include <iostream>
#include <unordered_set>
using namespace std;

int main() {
	unordered_set<int> myset;
	for (int i=1; i<5; ++i)
		myset.insert(i*10);
	for (int i=1; i<5; ++i)
		myset.insert(i*10);

	cout << "myset size: " << myset.size() << endl;
	cout << endl;

	for (unordered_set<int>::iterator it=myset.begin(); it!=myset.end(); ++it)
		cout << *it << ' ';
	cout << endl;
	cout << endl;

	for (int i=1; i<5; ++i)
		cout << i << " count: " << myset.count(i) << endl;
	cout << endl;
	for (int i=1; i<5; ++i)
		cout << i*10 << " count: " << myset.count(i*10) << endl;
	cout << endl;
	unordered_set<int>::iterator it = myset.find(30);
	if (it != myset.end())
		cout << "30 is in the set";
	cout << endl;
}

/*
myset size: 4

40 30 20 10

1 count: 0
2 count: 0
3 count: 0
4 count: 0

10 count: 1
20 count: 1
30 count: 1
40 count: 1

30 is in the set
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `binary_search`

```cpp
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

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `min`, `max`

```cpp
#include <iostream>     // cout
#include <algorithm>    // min
using namespace std;

int main () {
	cout << "min(1,2)==" << min(1,2) << endl;
	cout << "min(2,1)==" << min(2,1) << endl;
	cout << "min('a','z')==" << min('a','z') << endl;
	cout << "min(3.14,2.72)==" << min(3.14,2.72) << endl;

	cout << "max(1,2)==" << max(1,2) << endl;
	cout << "max(2,1)==" << max(2,1) << endl;
	cout << "max('a','z')==" << max('a','z') << endl;
	cout << "max(3.14,2.72)==" << max(3.14,2.72) << endl;
}

/*
min(1,2)==1
min(2,1)==1
min('a','z')==a
min(3.14,2.72)==2.72
max(1,2)==2
max(2,1)==2
max('a','z')==z
max(3.14,2.72)==3.14
*/

```

[↑ top](#c-stl)
<br><br><br><br><hr>


#### `sort`

```cpp
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

```

[↑ top](#c-stl)
<br><br><br><br><hr>
