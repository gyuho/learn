// http://cslibrary.stanford.edu/110/BinaryTrees.html
#include <iostream>
using namespace std;

struct node { 
	int data; 
	struct node* left; 
	struct node* right; 
};

/* 
 Helper function that allocates a new node 
 with the given data and NULL left and right 
 pointers. 
*/ 
struct node* newNode(int data) { 
	// new is like 'malloc' that allocates memory
	struct node* node = new(struct node);
	node->data = data; 
	node->left = NULL; 
	node->right = NULL;
	return node; 
} 
 
/* 
 Give a binary search tree and a number, inserts a new node 
 with the given number in the correct place in the tree. 
 Returns the new root pointer which the caller should 
 then use (the standard trick to avoid using reference 
 parameters). 
*/ 
struct node* insert(struct node* node, int data) { 
	// 1. If the tree is empty, return a new, single node 
	if (node == NULL) { 
		return newNode(data) ; 
	}
	else
	{ 
		// 2. Otherwise, recur down the tree
		if (data <= node->data)
			node->left = insert(node->left, data); 
		else
			node->right = insert(node->right, data);

		// return the (unchanged) node pointer 
		return node;
	} 
} 

/* 
 Given a binary search tree, print out 
 its data elements in increasing 
 sorted order. 
*/ 
void printTree(struct node* node) { 
	if (node == NULL)
		return;
	printTree(node->left);
	printf("%d ", node->data);
	printTree(node->right);
} 


/* 
 Given a binary tree, return true if a node 
 with the target data is found in the tree. Recurs 
 down the tree, chooses the left or right 
 branch by comparing the target to each node. 
*/ 
static int search(struct node* node, int target) { 
	// 1. Base case == empty tree 
	// in that case, the target is not found so return false 
	if (node == NULL) 
	{
		return false; 
	}
	else
	{
		// 2. see if found here 
		if (target == node->data)
			return true;
		else
			// 3. otherwise recur down the correct subtree 
			if (target < node->data)
				return search(node->left, target);
			else
				return search(node->right, target);  
	}
} 

int main()
{
	node* root = newNode(2);
	insert(root, 3);
	insert(root, 1);
	insert(root, 4);
	printTree(root);
	// 1 2 3 4
	cout << endl;

	cout << "search: " << search(root, 4) << endl;
	// 1
}
