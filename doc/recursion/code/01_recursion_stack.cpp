#include <iostream>
#include <string>
#include <map>
#include <cstdio>
#include <string.h>
using namespace std;

void r(int num) {
	if (num < 0)
	{
		return;
	}
	cout << "r with " << num << endl;
	r(num - 1);
}

char keys[] = {
	'A',
	'B',
	'C',
	'D',
	'E',
	'F',
	'G',
	'H',
	'I',
	'\0',
};

void recursion(int index, map<char,string>* rmap) {
	if (keys[index] == '\0')
	{
		cout << endl;
		cout << "recursion is done" << endl;
		cout << endl;
		return;
	}

	printf("beginning recursion with index %d / key %c / map %s\n", index, keys[index], (*rmap)[keys[index]].c_str());

	recursion(index+1, rmap);
	(*rmap)[keys[index]] = "done";

	printf("after     recursion with index %d / key %c / map %s\n", index, keys[index], (*rmap)[keys[index]].c_str());
}

int main()
{
	r(10);
	cout << endl;
	/*
		r with 10
		r with 9
		r with 8
		r with 7
		r with 6
		r with 5
		r with 4
		r with 3
		r with 2
		r with 1
		r with 0
	*/

	size_t length = strlen(keys);
	cout << length << endl; // 9

	map<char,string> executed;
	int i = 0;
	while (keys[i] != '\0'){
		executed[keys[i]] = "not yet";
		i++;
	}

	recursion(0, &executed);
}

/*
beginning recursion with index 0 / key A / map not yet
beginning recursion with index 1 / key B / map not yet
beginning recursion with index 2 / key C / map not yet
beginning recursion with index 3 / key D / map not yet
beginning recursion with index 4 / key E / map not yet
beginning recursion with index 5 / key F / map not yet
beginning recursion with index 6 / key G / map not yet
beginning recursion with index 7 / key H / map not yet
beginning recursion with index 8 / key I / map not yet

recursion is done

after     recursion with index 8 / key I / map done
after     recursion with index 7 / key H / map done
after     recursion with index 6 / key G / map done
after     recursion with index 5 / key F / map done
after     recursion with index 4 / key E / map done
after     recursion with index 3 / key D / map done
after     recursion with index 2 / key C / map done
after     recursion with index 1 / key B / map done
after     recursion with index 0 / key A / map done
*/
