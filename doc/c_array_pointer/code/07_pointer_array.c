#include <stdio.h>

int main()
{
	int n0=0, n1=1, n2=2;
	int * arr[3] = {&n0, &n1, &n2};
	printf("%d\n", *arr[0]); // 0
	printf("%d\n", *arr[1]); // 1
	printf("%d\n", *arr[2]); // 2

	char * sArr[3] = {"AA", "BB", "CC"};
	printf("%s\n", sArr[0]); // AA
	printf("%s\n", sArr[1]); // BB
	printf("%s\n", sArr[2]); // CC

	return 0;
}
