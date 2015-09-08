#include<iostream>
#include<math.h>
#include<stack>
#include<stdlib.h>

using namespace std;

void show(char fromPeg, char toPeg, int disk)
{
	cout << "Move the disk " << disk << " from \'" << fromPeg << "\' to \'" << toPeg << "\'" << endl;
}

void move(stack<int> &src, stack<int> &dest, char s, char d)
{
	if(src.empty())
	{
		int destTop = dest.top();
		dest.pop();
		src.push(destTop);
		show(d, s, destTop);
	}

	else if(dest.empty())
	{
		int srcTop = src.top();
		src.pop();
		dest.push(srcTop);
		show(s, d, srcTop);
	}

	else if(src.top() < dest.top())
	{
		int srcTop = src.top();
		src.pop();
		dest.push(srcTop);
		show(s, d, srcTop);
	}
	else if(src.top() > dest.top())
	{
		int destTop = dest.top();
		dest.pop();
		src.push(destTop);
		show(d, s, destTop);
	}
}

void hanoi(int num, stack<int> &src, stack<int> &aux, stack<int> &dest)
{
	int i;
	int total_moves = pow(2, num) - 1;
	char s = 'S', d = 'D', a = 'A';

	if(num%2 == 0)
	{
		char temp = a;
		a = d;
		d = temp;
	}

	for(i=num;i>=1;i--)
		src.push(i);

	for(i=1; i <=total_moves; ++i)
	{
		if(i%3 == 1)
			move(src, dest, s, d);
		else if(i%3 == 2)
			move(src, aux, s, a);
		else if(i%3 == 0)
			move(aux, dest, a ,d);
	}
}

int main()
{
	unsigned num = 5;

	stack<int> src;
	stack<int> dest;
	stack<int> aux;

	hanoi(num, src, aux, dest);
	return 0;
}

/*
Move the disk 1 from 'S' to 'D'
Move the disk 2 from 'S' to 'A'
Move the disk 1 from 'D' to 'A'
Move the disk 3 from 'S' to 'D'
Move the disk 1 from 'A' to 'S'
Move the disk 2 from 'A' to 'D'
Move the disk 1 from 'S' to 'D'
Move the disk 4 from 'S' to 'A'
Move the disk 1 from 'D' to 'A'
Move the disk 2 from 'D' to 'S'
Move the disk 1 from 'A' to 'S'
Move the disk 3 from 'D' to 'A'
Move the disk 1 from 'S' to 'D'
Move the disk 2 from 'S' to 'A'
Move the disk 1 from 'D' to 'A'
Move the disk 5 from 'S' to 'D'
Move the disk 1 from 'A' to 'S'
Move the disk 2 from 'A' to 'D'
Move the disk 1 from 'S' to 'D'
Move the disk 3 from 'A' to 'S'
Move the disk 1 from 'D' to 'A'
Move the disk 2 from 'D' to 'S'
Move the disk 1 from 'A' to 'S'
Move the disk 4 from 'A' to 'D'
Move the disk 1 from 'S' to 'D'
Move the disk 2 from 'S' to 'A'
Move the disk 1 from 'D' to 'A'
Move the disk 3 from 'S' to 'D'
Move the disk 1 from 'A' to 'S'
Move the disk 2 from 'A' to 'D'
Move the disk 1 from 'S' to 'D'
*/
