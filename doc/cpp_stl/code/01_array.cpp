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
