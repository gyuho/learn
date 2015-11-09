#include <stdio.h>

int main()
{
	int arr[3] = {0, 1, 2};

	printf("*arr: %d\n", *arr);  // 0
	*arr += 100;
	printf("*arr: %d\n", *arr);  // 100
	arr[0] -= 50;
	printf("*arr: %d\n", *arr);  // 50

	int * ptr = &arr[0];
	// int * ptr = arr;
	
	printf("%d %d\n", *ptr, *arr);     // 50 50
	printf("%d %d\n", ptr[0], arr[0]); // 50 50
	printf("%d %d\n", ptr[1], arr[1]); // 1 1
	printf("%d %d\n", ptr[2], arr[2]); // 2 2

	return 0;
}
