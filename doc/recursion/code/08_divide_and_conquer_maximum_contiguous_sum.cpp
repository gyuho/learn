#include <iostream>
#include <vector>
using namespace std;

int kadane(vector<int> nv)
{
	if (nv.size() == 0)
		return 0;
	int temp = 0;
	int maxSum = 0;
	for (vector<int>::iterator it = nv.begin(); it != nv.end(); ++it)
	{
		temp += *it;
		if (temp < 0)
		{
			temp = 0;	
		}
		else if (maxSum < temp)
		{
			maxSum = temp;
		}	
	}
	return maxSum;
}

int main()
{
	vector<int> nv;
	int nums[] = {-2, -5, 6, -2, 3, -10, 5, -6};
	size_t size = sizeof(nums) / sizeof(nums[0]);
	for (int i=0; i<size; i++)
		nv.push_back(nums[i]);
	for (vector<int>::iterator it = nv.begin(); it != nv.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
	cout << "kadane: " << kadane(nv) << endl;
}

/*
 -2 -5 6 -2 3 -10 5 -6
kadane: 7
*/
