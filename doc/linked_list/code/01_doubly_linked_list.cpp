#include <iostream>
using namespace std;
 
struct Node
{
	int value;
	Node *prev, *next;
	Node(int y)
	{
		value = y;
		next = prev = NULL;
	}
};
 
class LinkedList
{
	Node *head;
	Node *tail;

	public:
		LinkedList()
		{ 
			head = NULL;
			tail = NULL;
		}

		~LinkedList()
		{
			destroyList();
		}

		void pushFront(int x);
		void pushBack(int x);
		void printNodesForward();
		void printNodesReverse();
		void destroyList();
};

void LinkedList::pushFront(int x)
{
	Node *nd = new Node(x);
	if( head == NULL)
	{
		head = nd;
		tail = nd;
	}
	else
	{
		head->prev = nd;
		nd->next = head;
		head = nd;
	}
}

void LinkedList::pushBack(int x)
{
	Node *nd = new Node(x);
	if( tail == NULL)
	{
		head = nd;
		tail = nd;
	}
	else
	{
		tail->next = nd;
		nd->prev = tail;
		tail = nd;
	}
}
 
void LinkedList::printNodesForward()
{
	Node *temp = head;
	cout << "\nNodes in forward order:" << endl;
	while(temp != NULL)
	{
		cout << temp->value << "   " ;
		temp = temp->next;
	}
}

void LinkedList::printNodesReverse()
{
	Node *temp = tail;
	cout << "\nNodes in reverse order :" << endl;
	while(temp != NULL)
	{
		cout << temp->value << "   " ;
		temp = temp->prev;
	}
}

void LinkedList::destroyList()
{
	Node *T = tail;
	while(T != NULL)
	{
		Node *T2 = T;
		T = T->prev;
		delete T2;
	}
	head = NULL;
	tail = NULL;
}

int main()
{
		LinkedList *list = new LinkedList();
		//append nodes to front of the list
		for( int i = 1 ; i < 4 ; i++)
			list->pushFront(i);
 
		list->printNodesForward();
		list->printNodesReverse();
 
		//append nodes to back of the list
		for( int i = 1 ; i < 4 ; i++)
			list->pushBack(i+10);

		cout << endl << endl;
		list->printNodesForward();
		list->printNodesReverse();
 
		cout << endl << endl;
		delete list;
}
 
/*
Nodes in forward order:
3   2   1   
Nodes in reverse order :
1   2   3   


Nodes in forward order:
3   2   1   11   12   13   
Nodes in reverse order :
13   12   11   1   2   3 
*/
