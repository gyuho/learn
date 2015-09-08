#include <iostream>
#include <vector>
using namespace std;

int rob(vector<int> houses)
{
	switch (houses.size())
	{
		case 0:
			return 0;
		case 1:
			return houses[0];
		case 2:
		{
			int case1 = houses[0];
			int case2 = houses[1];
			if (case1 >= case2)
				return case1;
			return case2;
		}
		case 3:
		{
			int case1 = houses[0] + houses[2];
			int case2 = houses[1];
			if (case1 >= case2)
				return case1;
			return case2;
		}
	}
	vector<int> nv2(houses.begin()+2, houses.end());
	vector<int> nv3(houses.begin()+3, houses.end());
	int case1 = houses[0]+rob(nv2);
	int case2 = houses[1]+rob(nv3);
	if (case1 >= case2)
		return case1;
	return case2;
}

int main()
{
	int houses[] = {3, 15, 13, 4, 7};
	vector<int> hv;
	size_t size = sizeof(houses) / sizeof(houses[0]);
	for (int i=0; i<size; ++i)
	{
		hv.push_back(houses[i]);
	}
	vector<int> example(hv.begin()+1, hv.end());
	for (vector<int>::iterator it = example.begin(); it != example.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
	//  15 13 4 7

	cout << "rob(hv): " << rob(hv) << endl; // 23
}
