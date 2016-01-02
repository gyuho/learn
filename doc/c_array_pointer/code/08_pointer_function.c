#include <stdio.h>

// argument is passed by value!
// make sure to pass pointer if you want to update!
void updateArray(int * arrPt, int len, int delta)
{
	int i;
	for (i=0; i<len; i++)
		arrPt[i] += delta;
}

void printArray(int * arrPt, int len)
{
	int i;
	for (i=0; i<len; i++)
		printf("%d ", arrPt[i]);
	printf("\n");
}

int main()
{
	int arr[3] = {1,2,3};
	updateArray(arr, sizeof(arr) / sizeof(int), 10);
	printArray(arr, sizeof(arr) / sizeof(int));
	// 11 12 13

	return 0;
}
