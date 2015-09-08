#include <iostream>
#include <queue>

using namespace std;

// Queue is First In First Out
// Stack is Last In First Out

int main()
{
	queue<string> q;
	q.push("A"); // add elements
	q.push("B");
	q.push("C");
	
	while (q.size() > 0)
	{
		// returns the oldest, first element
		cout << q.front() << endl;
		
		// remove the element on top of the stack
		q.pop();
	}
	cout << endl;
	cout << "size of queue has become " << q.size() << endl;

	//  we cannot iterate through stack and queue
	//	for ( auto elem : q )
	//	{
	//		cout << "stack elements: " << elem;
	//	}
}

/*
A
B
C

size of queue has become 0
*/
