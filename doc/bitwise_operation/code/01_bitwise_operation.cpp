#include <iostream>
#include <stdio.h>
#include <string>
#include <bitset>
using namespace std;

long unsigned int toBin(long unsigned int num) {
	if (num == 0)
		return 0;
	return (num % 2) + 10*toBin(num/2);
}

long unsigned int toLu(string bstr) {
	// auto size = bstr.length();
	// to_ulong converts to unsigned long integer

	// Templates are evaluated when compiled
	const int size = 100;
	long unsigned int num = bitset<size>(bstr).to_ulong();
	printf("%10s (decimal %lu)\n", bstr.c_str(), num);
	return num;
}

int main()
{
	cout << toBin(5) << endl;
	// 101

	cout << bitset<100>("101").to_ulong() << endl;
	// 5

	cout << toLu("101") << endl;
	//       101 (decimal 5)

	cout << endl;
	cout << "AND:  x & y" << endl;
	unsigned int x = toLu("10001101");
	unsigned int y = toLu("01010111");
	unsigned int z = x & y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "OR:  x | y" << endl;
	x = toLu("10001101");
	y = toLu("01010111");
	z = x | y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "XOR:  x ^ y" << endl;
	x = toLu("0101");
	y = toLu("0011");
	z = x ^ y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "NOT(bit complement):  ^x  or  ~x" << endl;
	x = toLu("0101");
	// ~ flips every bit.
	/*
	NOT 0000 0000 0000 0000 0000 0000 0000 0101
	-------------------------------------------
	    1111 1111 1111 1111 1111 1111 1111 1010
	*/	
	z = ~x;
	printf("%10lu (decimal %d)\n", toBin(z), z);
	z = x ^ 0xF;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "AND NOT:  x &^ y  or  x &~ y" << endl;
	x = toLu("0101");
	y = toLu("0011");
	z = x &~ y;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Left Shift:  x << 1"<< endl;
	x = toLu("1010");
	y = x << 1;
	printf("%10lu (decimal %d)\n", toBin(y), y);

	cout << endl;
	cout << "Right Shift:  x >> 1" << endl;
	x = toLu("1010");
	y = x >> 1;
	printf("%10lu (decimal %d)\n", toBin(y), y);
}

/*
AND:  x & y
  10001101 (decimal 141)
  01010111 (decimal 87)
       101 (decimal 5)

OR:  x | y
  10001101 (decimal 141)
  01010111 (decimal 87)
  11011111 (decimal 223)

XOR:  x ^ y
      0101 (decimal 5)
      0011 (decimal 3)
       110 (decimal 6)

NOT(bit complement):  ^x  or  ~x
      0101 (decimal 5)
13368089053625086306 (decimal -6)
      1010 (decimal 10)

AND NOT:  x &^ y  or  x &~ y
      0101 (decimal 5)
      0011 (decimal 3)
       100 (decimal 4)

Left Shift:  x << 1
      1010 (decimal 10)
     10100 (decimal 20)

Right Shift:  x >> 1
      1010 (decimal 10)
       101 (decimal 5)
*/
