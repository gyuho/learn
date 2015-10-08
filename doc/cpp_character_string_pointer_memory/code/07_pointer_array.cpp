#include <stdio.h>

int main() 
{
	int arr[3] = {0, 1, 2};
	printf("Name of array: %p\n", arr);
	printf("&arr[0]: %p\n", &arr[0]);
	printf("&arr[1]: %p\n", &arr[1]);
	printf("&arr[2]: %p\n", &arr[2]);
	/*
		Name of array: 0x7ffc539d5a60
		&arr[0]: 0x7ffc539d5a60
		&arr[1]: 0x7ffc539d5a64
		&arr[2]: 0x7ffc539d5a68
	*/

	// int num = 0;
	// int* pnum = &num;
	// arr = pnum;

	// arr = &arr[1];
}

/*
07_pointer_array.cpp:13:6: error: incompatible types in assignment of ‘int*’ to ‘int [3]’
  arr = pnum;
      ^
07_pointer_array.cpp:15:6: error: incompatible types in assignment of ‘int*’ to ‘int [3]’
  arr = &arr[1];
      ^
*/
