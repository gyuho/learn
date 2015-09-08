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
