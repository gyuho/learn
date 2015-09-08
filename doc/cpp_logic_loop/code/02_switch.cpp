#include <iostream>
using namespace std;

int main()
{
	while (true)
	{
		cout << "Type an integer: ";
		int selected;
		cin >> selected;

		switch (selected)
		{
			case 0:
				{
					cout  << 0
						<< endl;
					break; // break out of switch
				}
			case 1:
				{
					cout  << 1
						<< endl;
					// continue; 
					// without this, it continues on 2
				}
			case 2:
				{
					cout  << 2
						<< endl;
					continue; // continue on while-loop
				}
			case 3:
				{
					cout  << 3
						<< endl;
					continue;
				}
			default:
				{
					cout  << "selected "
						<< selected
						<< endl;
				}
		}
		cout << "Breaking out of loop!" << endl;
		break;
	}
}

/*
Type an integer: 3
3
Type an integer: 2
2
Type an integer: 1
1
2
Type an integer: 0
0
Breaking out of loop!

Type an integer: 3
3
Type an integer: 2
2
Type an integer: 1
1
2
Type an integer: 5
selected 5
Breaking out of loop!
*/
