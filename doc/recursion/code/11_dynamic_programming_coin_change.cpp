#include <iostream>
#include <vector>
#include <limits.h>
using namespace std;

int getChange(int amount, int coins[], int coinSize)
{
	vector<int> storage(1);
	storage[0] = 0;
	for (int a = 1; a <= amount; ++a)
	{
		storage.push_back(INT_MAX);
		for (int i=0; i<coinSize; ++i)
		{
			int coint = coins[i];
			if (a >= coint)
			{
				if (storage[a] > 1 + storage[a-coint]) 
				{
					// retrieve from storage
					storage[a] = 1 + storage[a-coint];
				} 
			}
		}
	}
	return storage[amount];
}

int main()
{
	int coins[] = {1, 5, 7, 9, 11};
	size_t coinSize = sizeof(coins) / sizeof(*coins);

	cout << getChange(6, coins, coinSize) << endl; // 2
	//we need 2 coins(1 and 5) to make 6 cents

	cout << getChange(16, coins, coinSize) << endl; // 2
	//we need 2 coins(7 and 9) to make 16 cents

	cout << getChange(25, coins, coinSize) << endl;  // 3
	cout << getChange(250, coins, coinSize) << endl; // 24
}
