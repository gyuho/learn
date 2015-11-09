#include <stdio.h>

int main()
{
	int arr[3] = {0, 1, 2};

	printf("arr: %p\n", arr);         // 0x7ffcb4d51010
	printf("&arr[0]: %p\n", &arr[0]); // 0x7ffcb4d51010
	printf("&arr[1]: %p\n", &arr[1]); // 0x7ffcb4d51014
	printf("&arr[2]: %p\n", &arr[2]); // 0x7ffcb4d51018

	// (X)
	// arr = &arr[1];

	return 0;
}
