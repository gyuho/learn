[*back to problems*](https://github.com/gyuho/learn/tree/master/doc/problems)
<br>

# Problem

Find all instances of the word `CODE` in the board:

![word_search](img/word_search.png)

**Input:**

```go
// Go
var (
	board = [][]string{
		{"A", "X", "F", "H", "K", "C", "O", "F", "Q", "R"},
		{"C", "U", "Y", "T", "X", "B", "V", "H", "F", "D"},
		{"U", "J", "X", "B", "O", "D", "E", "N", "D", "S"},
		{"B", "E", "N", "C", "X", "M", "L", "O", "I", "L"},
		{"Q", "B", "D", "O", "Z", "P", "K", "O", "C", "K"},
		{"C", "T", "H", "D", "Y", "X", "E", "R", "T", "M"},
		{"A", "O", "B", "E", "U", "C", "O", "D", "E", "E"},
		{"H", "A", "D", "F", "F", "P", "H", "P", "O", "W"},
		{"P", "L", "G", "E", "V", "F", "G", "I", "C", "V"},
		{"A", "T", "E", "A", "S", "X", "G", "J", "D", "B"},
	}

	target = []string{"C", "O", "D", "E"}
)

```

```cpp
// C++
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

```

<br>
**Output:** 9 

<br><br><br>
- [Reference](#reference)
- [Algorithm #1](#algorithm-1)
- [Solution #1, in Go](#solution-1-in-go)
- [Solution #1, in C++](#solution-1-in-c)

[↑ top](#problem)
<br><br><br><br>
<hr>



<br><br><br><br><br><br><br><br><br><br>
<br><br><br><br><br><br><br><br><br><br>
<br><br><br><br><br><br><br><br><br><br>

*Please do not at my soluition yet. Please try it by yourself first.*

<br><br><br><br><br><br><br><br><br><br>
<br><br><br><br><br><br><br><br><br><br>
<br><br><br><br><br><br><br><br><br><br>
<hr>



#### Reference

- [Module 2: Multidimensional Arrays](http://www.seas.gwu.edu/~drum/cs1112/lectures/module2/suppl/index.html)
- [Module 9: Recursion, Part II](http://www.seas.gwu.edu/~drum/cs1112/lectures/module9/module9.html)

[↑ top](#problem)
<br><br><br><br>
<hr>






#### Algorithm #1

**Use recursion for multiple related decisions.** This is like a maze:
*each position in board leads to 8 other choices*. You need to decide which
direction to move: *left*, *right*, *up*, *down*, or *diagonals*. It's
multiple related decisions, so it's natural to try *recursion*. I will try
brute-forcing with recursion at every single row and column. And later
see if we can do better.

<br>
**FIRST**. Specify the **base case** of *recursion*. End recursion when:

1. Have found all the previous letters (e.g. `C`, `O`, `D` for `CODE`), and 
   just found the last letter (`E` for `CODE`).
2. No more position to move. Out of array range, at the end of board.

<br>
**SECOND**. Decide what to pass around to *recursive* functions:

1. Slice(array) of target letters (e.g. `C`, `O`, `D`, `E`).
2. Letter index in the target slice. 
3. Position on board to search for the target letter,
   to tell which direction to move.
4. Storage for previously found paths for backtracking.
5. Storage for paths that have all target letters.

<br>
**THIRD**. Create a two dimensional array with values from the board.
Start with the *first letter as a target letter*, at the position
*(0, 0)* in a two dimensional array. And if it *finds the target letter*,
call *recursive* function *onto other directions*. Otherwise, keep moving from
*left-top to right-bottom* until it reaches the end *(N, N)*. After iteration,
count the paths from the map.

[↑ top](#problem)
<br><br><br><br>
<hr>




#### Solution #1, in Go

```go
package main

import (
	"fmt"
	"strings"
)

func make2DSlice(row, column int) [][]string {
	mat := make([][]string, row)
	// for i := 0; i < row; i++ {
	for i := range mat {
		mat[i] = make([]string, column)
	}
	return mat
}

var (
	board = [][]string{
		{"A", "X", "F", "H", "K", "C", "O", "F", "Q", "R"},
		{"C", "U", "Y", "T", "X", "B", "V", "H", "F", "D"},
		{"U", "J", "X", "B", "O", "D", "E", "N", "D", "S"},
		{"B", "E", "N", "C", "X", "M", "L", "O", "I", "L"},
		{"Q", "B", "D", "O", "Z", "P", "K", "O", "C", "K"},
		{"C", "T", "H", "D", "Y", "X", "E", "R", "T", "M"},
		{"A", "O", "B", "E", "U", "C", "O", "D", "E", "E"},
		{"H", "A", "D", "F", "F", "P", "H", "P", "O", "W"},
		{"P", "L", "G", "E", "V", "F", "G", "I", "C", "V"},
		{"A", "T", "E", "A", "S", "X", "G", "J", "D", "B"},
	}

	target = []string{"C", "O", "D", "E"}
)

// each recursion needs its own storage, so we need to
// make a copy of it.
func copyMap(m map[string]string) map[string]string {
	n := make(map[string]string)
	for k, v := range m {
		n[k] = v
	}
	return n
}

func search(
	target []string,
	letterIdx int,
	row int,
	col int,
	subPath map[string]string,
	found map[string]bool,
) {
	// base case
	if row < 0 || col < 0 || row > len(board)-1 || col > len(board[0])-1 {
		// not valid move
		// because it exceeds array(slice) range
		return
	}

	targetLetter := target[letterIdx]
	currentLetter := board[row][col]
	if targetLetter != currentLetter {
		return
	}
	letterIdx++
	subPath[currentLetter] = fmt.Sprintf("%d,%d", row, col)

	// found the path
	lastTargetLetter := target[len(target)-1]
	if targetLetter == lastTargetLetter && len(target) == len(subPath) {
		ts := []string{}
		for _, v := range target {
			ts = append(ts, subPath[v])
		}
		found[strings.Join(ts, "->")] = true
		return
	}

	// find the next letter
	search(target, letterIdx, row, col-1, copyMap(subPath), found)   // left
	search(target, letterIdx, row, col+1, copyMap(subPath), found)   // right
	search(target, letterIdx, row-1, col, copyMap(subPath), found)   // up
	search(target, letterIdx, row+1, col, copyMap(subPath), found)   // down
	search(target, letterIdx, row-1, col-1, copyMap(subPath), found) // diagonal
	search(target, letterIdx, row+1, col+1, copyMap(subPath), found) // diagonal
	search(target, letterIdx, row-1, col+1, copyMap(subPath), found) // diagonal
	search(target, letterIdx, row+1, col-1, copyMap(subPath), found) // diagonal
	return
}

func main() {
	found := make(map[string]bool)
	for row, val := range board {
		for col := range val {
			subPath := make(map[string]string)
			search(target, 0, row, col, subPath, found)
		}
	}
	for path := range found {
		fmt.Println(path)
	}
}

/*
3,3->2,4->2,5->2,6
5,0->6,1->7,2->8,3
6,5->6,6->6,7->6,8
6,5->6,6->6,7->5,6
8,8->7,8->6,7->6,8
8,8->7,8->6,7->5,6
3,3->4,3->4,2->3,1
5,0->6,1->7,2->6,3
3,3->4,3->5,3->6,3
*/

```

[↑ top](#problem)
<br><br><br><br>
<hr>




#### Solution #1, in C++

```cpp
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

```

[↑ top](#problem)
<br><br><br><br>
<hr>
