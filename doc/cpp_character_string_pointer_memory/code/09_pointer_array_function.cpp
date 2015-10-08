#include <iostream>
using namespace std;

// (X) C++ does not allow to pass an entire
// array as an argument to a function.
// However, You can pass a pointer to an array.
int getSize1(int nums[])
{
	size_t size = sizeof(nums) / sizeof(nums[0]);
	return size;
}

// (X) C++ does not allow to pass an entire
// array as an argument to a function.
// However, You can pass a pointer to an array.
int getSize2(int* nums)
{
	size_t size = sizeof(nums) / sizeof(nums[0]);
	return size;
}

void printArray(int arr[], int size)
{
	int i;
	for (i=0; i < size; i++)
		printf("%d ", arr[i]);
	printf("\n");
}

int main()
{
	int nums[] = {0, 1, 3};
	cout << "*nums: " << *nums << endl;
	cout << "*(nums+2): " << *(nums+2) << endl;
	size_t size1 = sizeof(nums) / sizeof(nums[0]);
	size_t size2 = sizeof(nums) / sizeof(*nums);
	cout << "size1: " << size1 << endl;
	cout << "size2: " << size2 << endl;
	cout << "getSize1: " << getSize1(nums) << endl;
	cout << "getSize2: " << getSize2(nums) << endl;
	printArray(nums, 3); // 0 1 3
}

/*
*nums: 0
*(nums+2): 3
size1: 3
size2: 3
getSize1: 2
getSize2: 2
*/
