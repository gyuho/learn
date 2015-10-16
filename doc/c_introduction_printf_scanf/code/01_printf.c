#include <stdio.h> // to use printf

int main(void)
{
	// In C, special characters can be single-quoted.
	char newLine = '\n';
    printf("\"%s\": %d%c", "Hello", 10, newLine);
	// "Hello": 10

    printf("%10sWorld\n", "Hello");
	//     HelloWorld
    
    int num9=9, num17=17;
    printf("Octet: %o %#o\n", num9, num9); // Octet: 11 011 
    printf("Hex: %x %#x\n", num17, num17); // Hex: 11 0x11

    double fnum=0.123456789;
    printf("Float: %f %e\n", fnum, fnum); // Float: 0.123457 1.234568e-01
    printf("Float: %.3f\n", fnum);        // Float: 0.123

    return 0;
}

