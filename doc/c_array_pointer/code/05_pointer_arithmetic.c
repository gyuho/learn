#include <stdio.h>

int main()
{
	int * pt0 = 0x0010;  // increase by 4
	printf("%p %p\n", pt0+1, pt0+2); // 0x14 0x18

	double * pt1 = 0x0010;  // increase by 8
	printf("%p %p\n", pt1+1, pt1+2); // 0x18 0x20

	int arr[3] = {0, 1, 2};
	int * ptr = arr;
	printf("%d %d %d\n", *ptr, *(ptr+1), *(ptr+2));
	// 0 1 2
	
	printf("%d\n", arr[2] == *(arr+2)); // 1

	return 0;
}
