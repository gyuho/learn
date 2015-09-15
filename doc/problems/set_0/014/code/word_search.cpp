#include <iostream>
#include <string.h>
#include <vector>
#include <map>
#include <string>
using namespace std;

char board [10][10] = {
	{'A', 'X', 'F', 'H', 'K', 'C', 'O', 'F', 'Q', 'R'},
	{'C', 'U', 'Y', 'T', 'X', 'B', 'V', 'H', 'F', 'D'},
	{'U', 'J', 'X', 'B', 'O', 'D', 'E', 'N', 'D', 'S'},
	{'B', 'E', 'N', 'C', 'X', 'M', 'L', 'O', 'I', 'L'},
	{'Q', 'B', 'D', 'O', 'Z', 'P', 'K', 'O', 'C', 'K'},
	{'C', 'T', 'H', 'D', 'Y', 'X', 'E', 'R', 'T', 'M'},
	{'A', 'O', 'B', 'E', 'U', 'C', 'O', 'D', 'E', 'E'},
	{'H', 'A', 'D', 'F', 'F', 'P', 'H', 'P', 'O', 'W'},
	{'P', 'L', 'G', 'E', 'V', 'F', 'G', 'I', 'C', 'V'},
	{'A', 'T', 'E', 'A', 'S', 'X', 'G', 'J', 'D', 'B'},
};

char target[] = {'C', 'O', 'D', 'E'};

void join(const vector<string>& v, string d, string& s) {
	s.clear();
	for (vector<string>::const_iterator p = v.begin(); p != v.end(); ++p) {
		s += *p;
		if (p != v.end() - 1) {
			s += d;
		}
	}
}

map<char,string> copyMap(map<char,string>* m)
{
	map<char,string> n;
    for (map<char,string>::iterator it=(*m).begin(); it!=(*m).end(); ++it)
    	n[it->first] = it->second;
    return n;
}

void search(
		char target[],
		int letterIdx,
		int row,
		int col,
		map<char, string>* subPath,
		map<string, int>* found)
{
	// base case
	if ((row < 0) || (col < 0) || (row > 9) || (col > 9)) {
		// not valid move
		// because it exceeds array(slice) range
		return;
	}

	char targetLetter = target[letterIdx];
	char currentLetter = board[row][col];
	if (targetLetter != currentLetter) {
		return;
	}
	letterIdx++;
	(*subPath)[currentLetter] = to_string(row) + "," + to_string(col);

	// found the path
	char lastTargetLetter = target[strlen(target)-1];
	if ((targetLetter == lastTargetLetter) && (strlen(target) == (*subPath).size())) {
		vector<string> ts;
		for (int i = 0; i < strlen(target); ++i) {
			char v = target[i];
			ts.push_back((*subPath)[v]);
		}
		string s;
		join(ts, "->", s);
		(*found)[s] = 1;
	}

	// find the next letter
	map<char,string> left = copyMap(subPath);
	map<char,string> right = copyMap(subPath);
	map<char,string> up = copyMap(subPath);
	map<char,string> down = copyMap(subPath);
	map<char,string> diagonal0 = copyMap(subPath);
	map<char,string> diagonal1 = copyMap(subPath);
	map<char,string> diagonal2 = copyMap(subPath);
	map<char,string> diagonal3 = copyMap(subPath);
	search(target, letterIdx, row, col-1, &left, found);   // left
	search(target, letterIdx, row, col+1, &right, found);   // right
	search(target, letterIdx, row-1, col, &up, found);   // up
	search(target, letterIdx, row+1, col, &down, found);   // down
	search(target, letterIdx, row-1, col-1, &diagonal0, found); // diagonal
	search(target, letterIdx, row+1, col+1, &diagonal1, found); // diagonal
	search(target, letterIdx, row-1, col+1, &diagonal2, found); // diagonal
	search(target, letterIdx, row+1, col-1, &diagonal3, found); // diagonal
	return;
}

int main()
{
	map<string,int> found;
	for (int row=0; row<10; ++row) {
		for (int col=0; col<10; ++col) {
			map<char,string> subPath;
			search(target, 0, row, col, &subPath, &found);
		}
	}
	for (map<string,int>::iterator it=found.begin(); it!=found.end(); ++it)
    	cout << it->first << endl;
}

/*
3,3->2,4->2,5->2,6
3,3->4,3->4,2->3,1
3,3->4,3->5,3->6,3
5,0->6,1->7,2->6,3
5,0->6,1->7,2->8,3
6,5->6,6->6,7->5,6
6,5->6,6->6,7->6,8
8,8->7,8->6,7->5,6
8,8->7,8->6,7->6,8
*/

