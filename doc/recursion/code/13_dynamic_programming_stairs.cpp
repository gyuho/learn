#include <iostream>
#include <map>
using namespace std;

int countDynamic(int n, map<int,int>& mm)
{
	if (n < 0)
		return 0;
	else if (n == 0)
		return 1;
	else if (mm[n] > 0)
		return mm[n];
	mm[n] = countDynamic(n-1, mm) + \
			countDynamic(n-2, mm) + \
			countDynamic(n-3, mm);
	return mm[n];
}

int main()
{
	map<int,int> mm;
	cout << countDynamic(10, mm) << endl;

	for (map<int,int>::iterator it=mm.begin(); it!=mm.end(); ++it)
		cout << it->first << " => " << it->second << '\n';
	cout << endl;
	/*
	274
	1 => 1
	2 => 2
	3 => 4
	4 => 7
	5 => 13
	6 => 24
	7 => 44
	8 => 81
	9 => 149
	10 => 274
	*/
}
