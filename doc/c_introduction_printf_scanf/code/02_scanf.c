#include <stdio.h>

int main(void)
{
	int num0, num1, num2;
	printf("Type 3 integers: ");
	scanf("%o, %d, %x", &num0, &num1, &num2);

	printf("\nWe got:\n");
	printf("Octet in Decimal  : %d\n", num0); // 
	printf("Decimal in Decimal: %d\n", num1); // 
	printf("Hex in Decimal    : %d\n", num2); // 

	float fnum;
	printf("\nType 1 float: ");
	scanf("%f", &fnum);
	printf("Float: %f\n", fnum);

    return 0;
}

/*
Type 3 integers: 10, 10, 10

We got:
Octet in Decimal  : 8
Decimal in Decimal: 10
Hex in Decimal    : 16

Type 1 float: 1.23e-5
Float: 0.000012
*/

