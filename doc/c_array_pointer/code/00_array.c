#include <stdio.h>

int main(void)
{
	int arr0[2];
	arr0[0] = 10;
	arr0[1] = 20;
	int sum0=0, i;
	for (i=0; i<2; i++)
		sum0 += arr0[i];
	printf("sum0: %d\n", sum0);
	// sum0: 30

	int len=50;
	double arr1[len];
	int j;
	for (j=0; j<len; j++)
		arr1[j] = (double)j;
	double sum1=0;
	int k;
	for (k=0; k<len; k++)
		sum1 += arr1[k];
	printf("sum1: %f\n", sum1);
	// sum1: 1225.000000

	int arr2[3]={0, 1, 2};
	int i2;
	for (i2=0; i2<3; i2++)
		printf("%d ", arr2[i2]);
	printf("\n");
	// 0 1 2
	
	int arr3[]={0, 1, 2}; // automatically sized to 3
	int i3;
	for (i3=0; i3<3; i3++)
		printf("%d ", arr3[i3]);
	printf("\n");
	// 0 1 2

	int arr4[5]={1}; // automatically fill up with 0
	int i4;
	for (i4=0; i4<5; i4++)
		printf("%d ", arr4[i4]);
	printf("\n");
	// 1 0 0 0 0 

	printf("arr4 sizeof: %ld\n", sizeof(arr4));             // arr4 sizeof: 20
	printf("int sizeof: %ld\n", sizeof(int));               // int sizeof: 4
	printf("arr4 length: %ld\n", sizeof(arr4)/sizeof(int)); // arr4 length: 5

	printf("\ntype 3 integers: ");
	int arr[3];
	scanf("%d", &arr[0]);
	scanf("%d", &arr[1]);
	scanf("%d", &arr[2]);
	int ia;
	for (ia=0; ia<3; ia++)
		printf("%d ", arr[ia]);
	printf("\n");
	// type 3 integers: 100 200 300
	// 100 200 300

    return 0;
}

