#include <iostream>
#include <stack>

using namespace std;

// Queue is First In First Out
// Stack is Last In First Out

int main()
{
	stack<string> s;
	s.push("A"); // add elements
	s.push("B");
	s.push("C");
	
	while (s.size() > 0)
	{
		// returns the top element first
		cout << s.top() << endl;
		
		// remove the element on top of the stack
		s.pop();
	}
	cout << endl;

	cout << "size of stack has become " << s.size() << endl;
	
	//  we cannot iterate through stack and queue
	//	for ( auto elem : q )
	//	{
	//		cout << "stack elements: " << elem;
	//	}
}

/*
C
B
A

size of stack has become 0
*/
