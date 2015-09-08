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
	cout << "Bitmasking:  x & 0x0F" << endl;
	x = toLu("10001101");
	z = x & 0x0F;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Bitmasking:  x & 0xF" << endl;
	x = toLu("10001101");
	z = x & 0xF;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Bitmasking:  x & 0x1F" << endl;
	x = toLu("100011111");
	z = x & 0x1F;
	printf("%10lu (decimal %d)\n", toBin(z), z);

	cout << endl;
	cout << "Bitmasking:  x & y" << endl;
	x = toLu("10101010011111");
	y = toLu("10101010000000");
	z = x & y;
	printf("%10lu (decimal %d)\n", toBin(z), z);
}

/*
AND:  x & y
  10001101 (decimal 141)
  01010111 (decimal 87)
       101 (decimal 5)

Bitmasking:  x & 0x0F
  10001101 (decimal 141)
      1101 (decimal 13)

Bitmasking:  x & 0xF
  10001101 (decimal 141)
      1101 (decimal 13)

Bitmasking:  x & 0x1F
 100011111 (decimal 287)
     11111 (decimal 31)

Bitmasking:  x & y
10101010011111 (decimal 10911)
10101010000000 (decimal 10880)
10101010000000 (decimal 10880)
*/