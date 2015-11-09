#include <stdio.h>

int main(void)
{
	int num = 1;
	int * pnum;   // declare pointer-type pnum
	pnum = &num;  // store the address of num to pnum
	printf("%p\n", pnum); // 0x7ffdfbe9a224

	double f = 1.5;
	double * fp = &f;
	*fp = 100.10;
	printf("%f\n", f); // 100.100000

	int * pt = NULL;    // assign NULL pointer if not sure what to store
	printf("%p\n", pt); // (nil)

    return 0;
}

