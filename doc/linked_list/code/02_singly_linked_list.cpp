#include <iostream>
using namespace std;

class LinkedList{
	struct Node {
		int value;
		Node *next;
	};

	public:
		// constructor
		LinkedList(){
			head = NULL;
		}

		void pushFront(int val){
			Node *nd = new Node();
			nd->value = val;
			nd->next = head;
			head = nd;
		}

		int popFront(){
			Node *nd = head;
			int rv = nd->value;

			head = head->next;
			delete nd;
			return rv;
		}

	private:
		Node *head;
};

int main() {
	LinkedList list;

	list.pushFront(1);
	list.pushFront(2);
	list.pushFront(3);

	cout << list.popFront() << endl; // 3
	cout << list.popFront() << endl; // 2
	cout << list.popFront() << endl; // 1
}
