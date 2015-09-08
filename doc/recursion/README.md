[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

## Recursion

- [Reference](#reference)
- [recursion, factorial](#recursion-factorial)
- [recursion, convert to binary number](#recursion-convert-to-binary-number)
- [recursion, closure, fibonacci](#recursion-closure-fibonacci)
- [recursion, reverse string](#recursion-tower-hanoi)
- [recursion, tower hanoi](#recursion-tower-hanoi)
- [recursion, binary search tree](#recursion-binary-search-tree)
- [divide and conquer, merge sort](#divide-and-conquer-merge-sort)
- [divide and conquer, quick sort](#divide-and-conquer-quick-sort)
- [divide and conquer, maximum contiguous sum](#divide-and-conquer-maximum-contiguous-sum)
- [dynamic programming, coin change](#dynamic-programming-coin-change)
- [dynamic programming, rob houses](#dynamic-programming-rob-houses)
- [dynamic programming, stairs](#dynamic-programming-stairs)
- [dynamic programming, longest common subsequence](#dynamic-programming-longest-common-subsequence)

[↑ top](#recursion)
<br><br><br><br>
<hr>



#### Reference

- [Recursion](https://en.wikipedia.org/wiki/Recursion)
- [Tail call](https://en.wikipedia.org/wiki/Tail_call)
- [Recursion And Tail Calls In Go](http://www.goinggo.net/2013/09/recursion-and-tail-calls-in-go_26.html)
- [Closure (computer
  science)](https://simple.wikipedia.org/wiki/Closure_(computer_science))
- [Closure (computer
  programming)](https://en.wikipedia.org/wiki/Closure_(computer_programming))
- [Dynamic programming](https://en.wikipedia.org/wiki/Dynamic_programming)
- [Memoization](https://en.wikipedia.org/wiki/Memoization)
- [Greedy algorithm](https://en.wikipedia.org/wiki/Greedy_algorithm)
- [Divide and conquer algorithms](https://en.wikipedia.org/wiki/Divide_and_conquer_algorithms)
- [Divide and conquer, Princeton CS](http://www.cs.princeton.edu/~wayne/cs423/lectures/divide-and-conquer-4up.pdf)

[↑ top](#recursion)
<br><br><br><br>
<hr>





#### recursion, factorial

[Recursion](https://en.wikipedia.org/wiki/Recursion) is simple but can 
be confusing. It is kind of like *iteration* but with **self-referencing**.
Once a *recursive* function gets called outside first, it keeps **calling
itself**. Most important is to specify when to end the recursion. Otherwise it
recurs forever and causes stack overflow (run out of memory).
So use carefully!

<br>
Easiest example of recursion would be *factorial*:

```go
package main

import "fmt"

func factorialWithIteration(num int) int {
	result := 1
	if num == 0 {
		return result
	}
	for i := 2; i <= num; i++ {
		result *= i
	}
	return result
}

func factorial(num int) int {
	if num <= 1 {
		fmt.Println("returning: 1")
		return 1
	}
	fmt.Println("returning:", num, num-1)
	return num * factorial(num-1)
}

func main() {
	fmt.Println(factorialWithIteration(5)) // 120
	fmt.Println(factorial(5))              // 120
}

/*
returning: 5 4
returning: 4 3
returning: 3 2
returning: 2 1
returning: 1
*/

```

```cpp
#include <iostream>
using namespace std;

long factorial(int num) {
	if (num == 0)
		return 1;
	return num * factorial(num - 1);
}

int main()
{
	cout << factorial(5) << endl; // 120
}

```

<br>
If you print out all function calls:

```
factorial(5)

	returning: 5 4
	returning: 4 3
	returning: 3 2
	returning: 2 1
	returning: 1
	...
```

it's cumulative as:

```
return 5 * factorial(5-1)
	return 5 * (4 * factorial(4-1))
		return 5 * (4 * (3 * factorial(3-1)))
			return 5 * (4 * (3 * (2 * (factorial(2-1)))))
				return 5 * 4 * 3 * 2 * 1
```

<br>
[Call stack](https://en.wikipedia.org/wiki/Call_stack) is a
**stack or LIFO (last in, first out)** data structure to keep track of
active [subroutines](https://en.wikipedia.org/wiki/Subroutine) of a
computer program. And a **call stack** is composed of **stack frames**.
A **stack frame** is literally a frame of data that gets pushed onto
the call stack. And a stack frame usually represent function calls and
its arguments. So each recursive function call allocates another stack frame,
in addition to cost of executing multiplication. 
*Recursion code* looks simpler but brings memory overhead. 

<br>
And there is no *memoization* or *caching* of **previous computations**.
Can we do better?

1. We can also **store** the previous factorial results in hash table,
in order not to repeat the same computation. We can do this with [**_dynamic
programming_**](https://en.wikipedia.org/wiki/Dynamic_programming).

2. We can reduce memory consumption with [*tail
   recursion*](https://en.wikipedia.org/wiki/Tail_call), which is often used in
   *functional programming languages*.

[↑ top](#recursion)
<br><br><br><br>
<hr>





#### recursion, convert to binary number

```go
package main

import "fmt"

func toBinaryNumber(num uint64) uint64 {
	fmt.Println("calling on", num)
	if num == 0 {
		return 0
	}
	return (num % 2) + 10*toBinaryNumber(num/2)
}

func main() {
	fmt.Println(toBinaryNumber(15))
	/*
	   calling on 15
	   calling on 7
	   calling on 3
	   calling on 1
	   calling on 0
	   1111
	*/
}

```

```cpp
#include <iostream>
using namespace std;

long unsigned int toBinaryNumber(long unsigned int num) {
	if (num == 0)
		return 0;
	return (num % 2) + 10*toBinaryNumber(num/2);
}

int main()
{
	cout << toBinaryNumber(15) << endl; // 1111
}

```



[↑ top](#recursion)
<br><br><br><br>
<hr>






#### recursion, closure, fibonacci

```go
package main

import "fmt"

// fib returns nth fibonacci number.
func fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d, ", fib(uint(i)))
	}
	fmt.Println()
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377,
}

```

```cpp
#include <iostream>
using namespace std;

long unsigned int fib(long unsigned int num) {
	if (num == 0)
		return 0;
	else if (num == 1)
		return 1;
	else
		return fib(num-1) + fib(num-2);
}

int main()
{
	for (int i=0; i<15; ++i)
		cout << fib(i) << ", ";
	cout << endl;
	// 0, 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 
}

```


<br>
[Closure](https://en.wikipedia.org/wiki/Closure_(computer_programming))
is a function with its own environment and at least on bound variable.
Closures are used when *functions* are first-class objects in the language,
and can be passed to, or returned from higher-order functions.
For example, in Go, you would:

```go
package main

import "fmt"

func fib() func() int {
	v1, v2 := 0, 1
	return func() int {
		v1, v2 = v2, v1+v2
		return v1
	}
}

func main() {
	f := fib()
	for i := 0; i < 15; i++ {
		fmt.Printf("%d, ", f())
	}
	fmt.Println()
	// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610,
}

```

[↑ top](#recursion)
<br><br><br><br>
<hr>



#### recursion, reverse string

```go
package main

import "fmt"

func reverse(str string) string {
	if len(str) == 0 {
		return str
	}
	return reverse(str[1:]) + string(str[0])
}

func reverseInRune(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	fmt.Println(reverse("Hello"))       // olleH
	fmt.Println(reverseInRune("Hello")) // olleH
}

/*
reverse("Hello")
(reverse("ello")) + "H"
((reverse("llo")) + "e") + "H"
(((reverse("lo")) + "l") + "e") + "H"
((((reverse("o")) + "l") + "l") + "e") + "H"
(((("o") + "l") + "l") + "e") + "H"
"olleH"
*/

```

```cpp
#include <iostream>
#include <string>
#include <algorithm>
#include <string.h>
using namespace std;

void reverseInPlace(char* str);
char* reverseReturn(char* str);
string reverseRecursive(string str);

int main()
{
	string str = "Hello World!";
	reverse(str.begin(), str.end());
	cout << str << endl; // !dlroW olleH

	char st1[] = "Hello World!";
	reverseInPlace(st1);
	cout << st1 << endl; // !dlroW olleH

	char st2[] = "Hello World!";
	char* rs2 = reverseReturn(st2);
	cout << rs2 << endl; // !dlroW olleH
	delete [] rs2;

	string st3 = "Hello World!";
	cout << reverseRecursive(st3) << endl; // !dlroW olleH
}

void reverseInPlace(char* str) {
	// unsigned integer type
	// type able to represent the size of any object in bytes
	size_t size = strlen(str);
	if (size < 2) {
		return;
	}
	for ( size_t i = 0, j = size - 1; i < j; i++, j-- ) {
		char tempChar = str[i];
		str[i] = str[j];
		str[j] = tempChar;
	}
}

// DO NOT define with char s[]
char* reverseReturn(char* str)
{
	int length = strlen(str);

	// char bts[length];
	char* bts = (char*)malloc(length);
	// Dynamic allocation needs to be 
	// deallocated manually

	int i, j;
	for (i=0, j=length-1; i < j; ++i, --j)
	{
		bts[i] = str[j];
		bts[j] = str[i];
	}

	return bts;
}

string reverseRecursive(string str)
{
	if (str.length() == 1)
		return str;
	return reverseRecursive(str.substr(1, str.length())) + str.at(0);
}

```

[↑ top](#recursion)
<br><br><br><br>
<hr>





#### recursion, tower hanoi

```go
package main

import "fmt"

// move num disks from src to dest
func hanoi(num int, src, aux, dest string) {
	if num > 0 {
		if num == 1 {
			fmt.Printf("%v -> %v\n", src, dest)
		} else {
			// order matters!
			hanoi(num-1, src, dest, aux)
			hanoi(1, src, aux, dest)
			hanoi(num-1, aux, src, dest)
		}
	}
}

/*
A -> C
A -> B
C -> B
A -> C
B -> A
B -> C
A -> C
*/

func main() {
	hanoi(3, "A", "B", "C")
}

```

```cpp
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

```

[↑ top](#recursion)
<br><br><br><br>
<hr>





#### recursion, binary search tree

```go
package main

import "fmt"

// Tree contains a Root node of a binary search tree.
type Tree struct {
	Root *Node
}

// New returns a new Tree with its root Node.
func New(root *Node) *Tree {
	tr := &Tree{}
	tr.Root = root
	return tr
}

// Interface represents a single object in the tree.
type Interface interface {
	// Less returns true when the receiver item(key)
	// is less than the given(than) argument.
	Less(than Interface) bool
}

// Node is a Node and a Tree itself.
type Node struct {
	// Left is a left child Node.
	Left *Node

	Key Interface

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	return nd
}

func (tr *Tree) String() string {
	return tr.Root.String()
}

func (nd *Node) String() string {
	if nd == nil {
		return "[]"
	}
	s := ""
	if nd.Left != nil {
		s += nd.Left.String() + " "
	}
	s += fmt.Sprintf("%v", nd.Key)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

// Insert inserts a Node to a Tree without replacement.
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)
}

func (nd *Node) insert(node *Node) *Node {
	if nd == nil {
		return node
	}
	if nd.Key.Less(node.Key) {
		nd.Right = nd.Right.insert(node)
	} else {
		nd.Left = nd.Left.insert(node)
	}
	return nd
}

// Min returns the minimum key Node in the tree.
func (tr Tree) Min() *Node {
	nd := tr.Root
	if nd == nil {
		return nil
	}
	for nd.Left != nil {
		nd = nd.Left
	}
	return nd
}

// Max returns the maximum key Node in the tree.
func (tr Tree) Max() *Node {
	nd := tr.Root
	if nd == nil {
		return nil
	}
	for nd.Right != nil {
		nd = nd.Right
	}
	return nd
}

// Search does binary-search on a given key and returns the first Node with the key.
func (tr Tree) Search(key Interface) *Node {
	nd := tr.Root
	// just updating the pointer value (address)
	for nd != nil {
		if nd.Key == nil {
			break
		}
		switch {
		case nd.Key.Less(key):
			nd = nd.Right
		case key.Less(nd.Key):
			nd = nd.Left
		default:
			return nd
		}
	}
	return nil
}

// SearchChan does binary-search on a given key and return the first Node with the key.
func (tr Tree) SearchChan(key Interface, ch chan *Node) {
	searchChan(tr.Root, key, ch)
	close(ch)
}

func searchChan(nd *Node, key Interface, ch chan *Node) {
	// leaf node
	if nd == nil {
		return
	}
	// when equal
	if !nd.Key.Less(key) && !key.Less(nd.Key) {
		ch <- nd
		return
	}
	searchChan(nd.Left, key, ch)  // left
	searchChan(nd.Right, key, ch) // right
}

// SearchParent does binary-search on a given key and returns the parent Node.
func (tr Tree) SearchParent(key Interface) *Node {
	nd := tr.Root
	parent := new(Node)
	parent = nil
	// just updating the pointer value (address)
	for nd != nil {
		if nd.Key == nil {
			break
		}
		switch {
		case nd.Key.Less(key):
			parent = nd // copy the pointer(address)
			nd = nd.Right
		case key.Less(nd.Key):
			parent = nd // copy the pointer(address)
			nd = nd.Left
		default:
			return parent
		}
	}
	return nil
}

type nodeStruct struct {
	ID    string
	Value int
}

func (n nodeStruct) Less(b Interface) bool {
	return n.Value < b.(nodeStruct).Value
}

func main() {
	root := NewNode(nodeStruct{"A", 5})
	tr := New(root)
	tr.Insert(NewNode(nodeStruct{"B", 3}))
	tr.Insert(NewNode(nodeStruct{"C", 17}))
	fmt.Printf("%s\n", tr)
	// [[{B 3}] {A 5} [{C 17}]]

	fmt.Println(tr.Search(nodeStruct{"A", 5}))   // [[{B 3}] {A 5} [{C 17}]]
	fmt.Println(tr.Search(nodeStruct{Value: 3})) // [{B 3}]
	fmt.Println(tr.Search(nodeStruct{"C", 17}))  // [{C 17}]

	ch := make(chan *Node)
	go tr.SearchChan(nodeStruct{"A", 5}, ch)
	fmt.Println(<-ch) // [[{B 3}] {A 5} [{C 17}]]

	fmt.Println(tr.Max()) // [{C 17}]
	fmt.Println(tr.Min()) // [{B 3}]
}

```

```cpp
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

```

[↑ top](#recursion)
<br><br><br><br>
<hr>





#### divide and conquer, merge sort

```go
package main

import "fmt"

// O(n)
func merge(s1, s2 []int) []int {
	final := make([]int, len(s1)+len(s2))
	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] <= s2[j] {
			final[i+j] = s1[i]
			i++
			continue
		}
		final[i+j] = s2[j]
		j++
	}
	for i < len(s1) {
		final[i+j] = s1[i]
		i++
	}
	for j < len(s2) {
		final[i+j] = s2[j]
		j++
	}
	return final
}

// O(n * log n)
// Recursively splits the array into subarrays, until only one element.
// From each subarray, merge them into a sorted array.
func mergeSort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}
	idx := len(slice) / 2
	left := mergeSort(slice[:idx])
	right := mergeSort(slice[idx:])
	return merge(left, right)
}

// O(n * log n)
// Recursively splits the array into subarrays, until only one element.
// From each subarray, merge them into a sorted array.
func mergeSortConcurrency(slice []int, ch chan []int) {
	if len(slice) < 2 {
		ch <- slice
		return
	}

	idx := len(slice) / 2
	ch1, ch2 := make(chan []int), make(chan []int)

	go mergeSortConcurrency(slice[:idx], ch1)
	go mergeSortConcurrency(slice[idx:], ch2)

	left := <-ch1
	right := <-ch2

	// close after waiting to receive all
	close(ch1)
	close(ch2)

	ch <- merge(left, right)
}

func main() {
	sliceA := []int{9, -13, 4, -2, 3, 1, -10, 21, 12}
	fmt.Println(mergeSort(sliceA))

	sliceB := []int{9, -13, 4, -2, 3, 1, -10, 21, 12}
	ch := make(chan []int)
	go mergeSortConcurrency(sliceB, ch)
	fmt.Println(<-ch)
}

```

```cpp
#include <iostream>     // cout
#include <algorithm>    // merge, sort
#include <vector>       // vector
using namespace std;

int main () {
	int first[] = {5,10,15,20,25};
	int second[] = {50,40,30,20,10};
	vector<int> v(10);

	sort(first,first+5);
	sort(second,second+5);
	merge(first, first+5, second, second+5, v.begin());

	for (vector<int>::iterator it=v.begin(); it!=v.end(); ++it)
		cout << ' ' << *it;
	cout << '\n';
	//  5 10 10 15 20 20 25 30 40 50
}

```

[↑ top](#recursion)
<br><br><br><br>
<hr>



#### divide and conquer, quick sort

```go
package main

import "fmt"

/*
Go supports multiple assignment
Swap(&slice[i+1], &slice[lidx])

func Swap(a *int, b *int) {
	temp := *a
	*a = *b
	*b = temp
}

Pseudocode from CLRS

quickSort(A, p, r)
	if p < r
		q = partition(A, p, r)
		quickSort(A, p, q-1)
		quickSort(A, q+1, r)

Partition(A, p, r)
	x = A[r]
	i = p - 1
	for j = p to r - 1
		if A[j] =< x
			i = i + 1
			exchange A[i] with A[j]
	exchange A[i+1] with A[r]
	return i+1
*/

// O(n * log n), in place with O(log(n)) stack space
// First choose a pivot element.
// And partition the array around the pivot.
// Around pivot, bigger ones moved to the right.
// Smaller ones moved to the left.
// Repeat (Recursion)
func quickSort(slice []int, fidx, lidx int) {
	if fidx < lidx {
		mid := partition(slice, fidx, lidx)
		quickSort(slice, fidx, mid-1)
		quickSort(slice, mid+1, lidx)
	}
}

// O(n)
// partition literally partition an integer array around a pivot.
// Elements bigger than pivot move to the right.
// Smaller ones move to left.
// It returns the index of the pivot element.
// After partition, the pivot element is places
// where it should be in the final sorted array.
// fidx = index of first element, usually 0
// lidx = index of last element, usually slice.size - 1
func partition(slice []int, fidx int, lidx int) int {
	x := slice[lidx]
	i := fidx - 1

	for j := fidx; j < lidx; j++ {
		if slice[j] <= x {
			i++
			slice[i], slice[j] = slice[j], slice[i] // to swap
		}
	}
	slice[i+1], slice[lidx] = slice[lidx], slice[i+1]
	return i + 1
}

func main() {
	slice := []int{9, -13, 4, -2, 3, 1, -10, 21, 12}
	quickSort(slice, 0, len(slice)-1)
	fmt.Println(slice)
	// [-13 -10 -2 1 3 4 9 12 21]
}

```

```cpp
// http://geeksquiz.com/quick-sort/
#include <stdio.h>
 
// A utility function to swap two elements
void swap(int* a, int* b)
{
	int t = *a;
	*a = *b;
	*b = t;
}
 
/* This function takes last element as pivot, places the pivot element at its
   correct position in sorted array, and places all smaller (smaller than pivot)
   to left of pivot and all greater elements to right of pivot */
int partition (int arr[], int l, int h)
{
	int x = arr[h];    // pivot
	int i = (l - 1);  // Index of smaller element
 
	for (int j = l; j <= h- 1; j++)
	{
		// If current element is smaller than or equal to pivot 
		if (arr[j] <= x)
		{
			i++;    // increment index of smaller element
			swap(&arr[i], &arr[j]);  // Swap current element with index
		}
	}
	swap(&arr[i + 1], &arr[h]);  
	return (i + 1);
}
 
/* arr[] --> Array to be sorted, l  --> Starting index, h  --> Ending index */
void quickSort(int arr[], int l, int h)
{
	if (l < h)
	{
		int p = partition(arr, l, h); /* Partitioning index */
		quickSort(arr, l, p - 1);
		quickSort(arr, p + 1, h);
	}
}
 
void printArray(int arr[], int size)
{
	int i;
	for (i=0; i < size; i++)
		printf("%d ", arr[i]);
	printf("\n");
}
 
int main()
{
	int arr[] = {9, -13, 4, -2, 3, 1, -10, 21, 12};
	int n = sizeof(arr) / sizeof(arr[0]);
	quickSort(arr, 0, n-1);
	printArray(arr, n);
	// -13 -10 -2 1 3 4 9 12 21
}

```

[↑ top](#recursion)
<br><br><br><br>
<hr>



#### divide and conquer, maximum contiguous sum

```go
/*
Maximum Contiguous Subarray(substring)
-100, 1, 2 => 1, 2

Kadane Algorithm Dynamic Programming: O ( n )

Divide and Conquer method: O ( n lg n )
maximum of the following
	getMCS(A, begin, mid)
	getMCS(A, mid+1, end)
	getMCS(crossing)
*/
package main

import (
	"fmt"
	"math"
)

func main() {
	s := []int{-2, -5, 6, -2, 3, -10, 5, -6}
	fmt.Println(getMCS(s, 0, len(s)-1)) // 7
	fmt.Println(kadane(s))              // [6 -2 3] 7
}

func kadane(slice []int) ([]int, int) {
	if len(slice) == 0 {
		return []int{}, 0
	}
	temp, maxSum := 0, 0
	lastIdx := 0
	for i, v := range slice {
		temp += v
		if temp < 0 {
			temp = 0 // reset
			continue
		}
		if maxSum < temp {
			maxSum = temp
			lastIdx = i
		}
	}
	check := 0
	firstIdx := 0
	for j := lastIdx; j > 0; j-- {
		check += slice[j]
		if maxSum == check {
			firstIdx = j
		}
	}
	return slice[firstIdx : lastIdx+1], maxSum
}

func max(more ...int) int {
	max := more[0]
	for _, elem := range more {
		if max < elem {
			max = elem
		}
	}
	return max
}

func getMCS(slice []int, first, last int) int {
	if first == last {
		return slice[first]
	}
	mid := (first + last) / 2
	return max(
		getMCS(slice, first, mid),
		getMCS(slice, mid+1, last),
		getMCSAcross(slice, first, mid, last),
	)
}

func getMCSAcross(slice []int, first, mid, last int) int {
	sum1 := 0
	leftSum := math.MinInt32
	for i := mid; first <= i; i-- {
		sum1 += slice[i]
		if leftSum < sum1 {
			leftSum = sum1
		}
	}
	sum2 := 0
	rightSum := math.MinInt32
	for i := mid + 1; i <= last; i++ {
		sum2 += slice[i]
		if rightSum < sum2 {
			rightSum = sum2
		}
	}
	return leftSum + rightSum
}

```

```cpp
#include <iostream>
#include <vector>
using namespace std;

int kadane(vector<int> nv)
{
	if (nv.size() == 0)
		return 0;
	int temp = 0;
	int maxSum = 0;
	for (vector<int>::iterator it = nv.begin(); it != nv.end(); ++it)
	{
		temp += *it;
		if (temp < 0)
		{
			temp = 0;	
		}
		else if (maxSum < temp)
		{
			maxSum = temp;
		}	
	}
	return maxSum;
}

int main()
{
	vector<int> nv;
	int nums[] = {-2, -5, 6, -2, 3, -10, 5, -6};
	size_t size = sizeof(nums) / sizeof(nums[0]);
	for (int i=0; i<size; i++)
		nv.push_back(nums[i]);
	for (vector<int>::iterator it = nv.begin(); it != nv.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
	cout << "kadane: " << kadane(nv) << endl;
}

/*
 -2 -5 6 -2 3 -10 5 -6
kadane: 7
*/

```

<br>
Note that the [`kadane`](https://en.wikipedia.org/wiki/Maximum_subarray_problem)
algorithm uses a simple version of dynamic programming
using previous computation.

[↑ top](#recursion)
<br><br><br><br>
<hr>


#### dynamic programming, coin change

> Dynamic programming (usually referred to as DP ) is a very powerful technique
> to solve a particular class of problems. It demands very elegant formulation
> of the approach and simple thinking and the coding part is very easy. The
> idea is very simple, If you have solved a problem with the given input, then
> save the result for future reference, so as to avoid solving the same problem
> again.. shortly 'Remember your Past'.  If the given problem can be broken
> up in to smaller sub-problems and these smaller subproblems are in turn
> divided in to still-smaller ones, and in this process, if you observe some
> over-lappping subproblems, then its a big hint for DP. Also, the optimal
> solutions to the subproblems contribute to the optimal solution of the given
> problem (referred to as the Optimal Substructure Property).
>
> [*Tutorial for Dynamic
> Programming*](https://www.codechef.com/wiki/tutorial-dynamic-programming) *by
> Codechef*

So dynamic prgramming breaks down a complex problem into simpler sub-problems.
And it saves each computation, bottom-up or top-to-bottom, and consturct the
solution from the previously found ones.

```go
package main

import (
	"fmt"
	"math"
)

func getChange(amount int, coins []int) int {
	storage := []int{0}
	for a := 1; a <= amount; a++ {
		storage = append(storage, math.MaxInt32)
		for _, coin := range coins {
			if a >= coin {
				if storage[a] > 1+storage[a-coin] {
					// retrieve from storage
					storage[a] = 1 + storage[a-coin]
				}
			}
		}
	}
	return storage[amount]
}

// Find this minimum number of coins needed to make change fo x amount
func main() {
	coins := []int{1, 5, 7, 9, 11}

	fmt.Println(getChange(6, coins)) // 2
	//we need 2 coins(1 and 5) to make 6 cents

	fmt.Println(getChange(16, coins)) // 2
	//we need 2 coins(7 and 9) to make 16 cents

	fmt.Println(getChange(25, coins))  // 3
	fmt.Println(getChange(250, coins)) // 24
}

```

```cpp
#include <iostream>
#include <vector>
#include <limits.h>
using namespace std;

int getChange(int amount, int coins[], int coinSize)
{
	vector<int> storage(1);
	storage[0] = 0;
	for (int a = 1; a <= amount; ++a)
	{
		storage.push_back(INT_MAX);
		for (int i=0; i<coinSize; ++i)
		{
			int coint = coins[i];
			if (a >= coint)
			{
				if (storage[a] > 1 + storage[a-coint]) 
				{
					// retrieve from storage
					storage[a] = 1 + storage[a-coint];
				} 
			}
		}
	}
	return storage[amount];
}

int main()
{
	int coins[] = {1, 5, 7, 9, 11};
	size_t coinSize = sizeof(coins) / sizeof(*coins);

	cout << getChange(6, coins, coinSize) << endl; // 2
	//we need 2 coins(1 and 5) to make 6 cents

	cout << getChange(16, coins, coinSize) << endl; // 2
	//we need 2 coins(7 and 9) to make 16 cents

	cout << getChange(25, coins, coinSize) << endl;  // 3
	cout << getChange(250, coins, coinSize) << endl; // 24
}

```

[↑ top](#recursion)
<br><br><br><br>
<hr>



#### dynamic programming, rob houses

```go
// http://codercareer.blogspot.com/2013/02/no-44-maximal-stolen-values.html
// Problem: Maximize the money to rob without robbing adjacent houses.
//
// Example 1. Rob 22(=15+7), not 11(=3+1+7)
// 					 Can't rob 3 + 15
//
// ___[ 3 ]___[ 15 ]___[ 1 ]___[ 4 ]___[ 7 ]___
//
//
// Example 2. Rob 11(=3+1+7), not 9(=5+4)
//
// ___[ 3 ]___[ 5 ]___[ 1 ]___[ 4 ]___[ 7 ]___
//
package main

import "fmt"

// rob returns the maximum sum of money that we can rob from the input slice.
//
// Dynamic Programming:
// if slice contains more than 3 elements
// return the maximum in two cases:
//	1. Choose the first element
//	2. Choose the second
//
func rob(houses []int) int {
	switch len(houses) {
	case 0:
		return 0
	case 1:
		return houses[0]
	case 2:
		return max(houses...)
	case 3:
		case1 := houses[0] + houses[2]
		case2 := houses[1]
		if case1 >= case2 {
			return case1
		}
		return case2
	}
	return max(
		houses[0]+rob(houses[2:]),
		houses[1]+rob(houses[3:]),
	)
}

func main() {
	testCases := []struct {
		Houses   []int
		MaxMoney int
	}{
		{[]int{3}, 3},                                       // 3 = 3
		{[]int{3, 15}, 15},                                  // 15 = 15
		{[]int{3, 15, 13}, 16},                              // 16 = 3 + 13
		{[]int{3, 15, 13, 5}, 20},                           // 20 = 15 + 5
		{[]int{3, 15, 13, 4, 7}, 23},                        // 23 = 3 + 13 + 7
		{[]int{3, 15, 1, 4, 7}, 22},                         // 22 = 15 + 7
		{[]int{3, 5, 1, 4, 7}, 12},                          // 12 = 5 + 7
		{[]int{7, 2, 4, 3, 1, 4}, 15},                       // 15 = 7 + 4 + 4
		{[]int{7, 2, 4, 3, 1, 2}, 13},                       // 13 = 7 + 4 + 2
		{[]int{1, 7, 12, 3, 1, 2}, 15},                      // 15 = 1 + 12 + 2 = 15
		{[]int{1, 7, 12, 3, 1, 2, 8, 3}, 22},                // 22 = 1 + 12 + 1 + 8
		{[]int{1, 7, 12, 3, 1, 11, 8, 1}, 25},               // 25 = 1 + 12 + 11 + 1
		{[]int{6, 7, 2, 1, 3, 9, 6, 12, 1, 2, 6, 7, 1}, 38}, // 38 = ?
	}
	for idx, testCase := range testCases {
		maxmoney1 := rob(testCase.Houses)
		maxmoney2 := testCase.MaxMoney
		fmt.Printf("Max: %3d / Original Slice: %v\n", testCase.MaxMoney, testCase.Houses)
		if maxmoney1 != maxmoney2 {
			fmt.Printf("WRONG %2d: %v != %v / %+v\n", idx, maxmoney1, maxmoney2, testCase)
		}
	}
}

func max(more ...int) int {
	max := more[0]
	for _, elem := range more {
		if max < elem {
			max = elem
		}
	}
	return max
}

```

```cpp
#include <iostream>
#include <vector>
using namespace std;

int rob(vector<int> houses)
{
	switch (houses.size())
	{
		case 0:
			return 0;
		case 1:
			return houses[0];
		case 2:
		{
			int case1 = houses[0];
			int case2 = houses[1];
			if (case1 >= case2)
				return case1;
			return case2;
		}
		case 3:
		{
			int case1 = houses[0] + houses[2];
			int case2 = houses[1];
			if (case1 >= case2)
				return case1;
			return case2;
		}
	}
	vector<int> nv2(houses.begin()+2, houses.end());
	vector<int> nv3(houses.begin()+3, houses.end());
	int case1 = houses[0]+rob(nv2);
	int case2 = houses[1]+rob(nv3);
	if (case1 >= case2)
		return case1;
	return case2;
}

int main()
{
	int houses[] = {3, 15, 13, 4, 7};
	vector<int> hv;
	size_t size = sizeof(houses) / sizeof(houses[0]);
	for (int i=0; i<size; ++i)
	{
		hv.push_back(houses[i]);
	}
	vector<int> example(hv.begin()+1, hv.end());
	for (vector<int>::iterator it = example.begin(); it != example.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
	//  15 13 4 7

	cout << "rob(hv): " << rob(hv) << endl; // 23
}

```

[↑ top](#recursion)
<br><br><br><br>
<hr>



#### dynamic programming, stairs

```go
// A child is running up a staircase with n steps,
// and can hop either 1 step, 2 steps, or 3 steps at a time.
// Implement a method to count how many possible ways
// the child can run up the stairs.
package main

import "fmt"

func countRecursive(n int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	}
	return countRecursive(n-1) + countRecursive(n-2) + countRecursive(n-3)
}

func countDynamic(n int, mm map[int]int) int {
	if n < 0 {
		return 0
	} else if n == 0 {
		return 1
	} else if mm[n] > 0 {
		return mm[n]
	}
	mm[n] = countDynamic(n-1, mm) +
		countDynamic(n-2, mm) +
		countDynamic(n-3, mm)
	return mm[n]
}

func main() {
	fmt.Println(countRecursive(10)) // 274

	nmap := make(map[int]int)
	fmt.Println(countDynamic(10, nmap)) // 274
	fmt.Println(nmap)
	// map[2:2 4:7 6:24 8:81 10:274 1:1 3:4 5:13 7:44 9:149]
}

```

```cpp
#include <iostream>
#include <map>
using namespace std;

int countDynamic(int n, map<int,int>& mm)
{
	if (n < 0)
		return 0;
	else if (n == 0)
		return 1;
	else if (mm[n] > 0)
		return mm[n];
	mm[n] = countDynamic(n-1, mm) + \
			countDynamic(n-2, mm) + \
			countDynamic(n-3, mm);
	return mm[n];
}

int main()
{
	map<int,int> mm;
	cout << countDynamic(10, mm) << endl;

	for (map<int,int>::iterator it=mm.begin(); it!=mm.end(); ++it)
		cout << it->first << " => " << it->second << '\n';
	cout << endl;
	/*
	274
	1 => 1
	2 => 2
	3 => 4
	4 => 7
	5 => 13
	6 => 24
	7 => 44
	8 => 81
	9 => 149
	10 => 274
	*/
}

```

[↑ top](#recursion)
<br><br><br><br>
<hr>



#### dynamic programming, longest common subsequence

```go
/*
Longest Common Subsequence

Subsequence needs not be contiguous
X = BDCABA
Y = ABCBDAB => LCS is B C B

Dynamic Programming method : O ( m * n )
*/
package main

import "fmt"

func main() {
	fmt.Println(LCS("BDCABA", "ABCBDAB"))
	// 4 BDAB

	fmt.Println(LCS("AXXBCDXXX", "XFXHXKX"))
	// 4 XXXX

	fmt.Println(LCS("AGGTABTABTABTAB", "GXTXAYBTABTABTAB"))
	// 13 GTABTABTABTAB

	fmt.Println(LCS("AGGTABGHSRCBYJSVDWFVDVSBCBVDWFDWVV", "GXTXAYBRGDVCBDVCCXVXCWQRVCBDJXCVQSQQ"))
	// 14 GTABGCBVWVCBDV
}

func LCS(str1, str2 string) (int, string) {
	size1 := len(str1)
	size2 := len(str2)
	mat := create2D(size1+1, size2+1)
	i, j := 0, 0
	for i = 0; i <= size1; i++ {
		for j = 0; j <= size2; j++ {
			if i == 0 || j == 0 {
				mat[i][j] = 0
			} else if str1[i-1] == str2[j-1] {
				mat[i][j] = mat[i-1][j-1] + 1
			} else {
				mat[i][j] = max(mat[i-1][j], mat[i][j-1])
			}
		}
	}
	return mat[size1][size2], backTrack(mat, str1, str2, size1-1, size2-1)
}

func backTrack(mat [][]int, str1, str2 string, i, j int) string {
	if i == -1 || j == -1 {
		return ""
	} else if str1[i] == str2[j] {
		return backTrack(mat, str1, str2, i-1, j-1) + string(str1[i])
	}
	if mat[i+1][j] > mat[i][j+1] {
		return backTrack(mat, str1, str2, i, j-1)
	}
	return backTrack(mat, str1, str2, i-1, j)
}

func max(more ...int) int {
	max := more[0]
	for _, elem := range more {
		if max < elem {
			max = elem
		}
	}
	return max
}

func create2D(size1, size2 int) [][]int {
	mat := make([][]int, size1)
	for i := range mat {
		mat[i] = make([]int, size2)
	}
	return mat
}

```

```cpp
// http://www.geeksforgeeks.org/printing-longest-common-subsequence/
#include <iostream>
#include <cstring>
#include <cstdlib>
using namespace std;
 
/* Returns length of LCS for X[0..m-1], Y[0..n-1] */
void LCS( char *X, char *Y, int m, int n )
{
	int L[m+1][n+1];

	/* Following steps build L[m+1][n+1] in bottom up fashion. Note
			that L[i][j] contains length of LCS of X[0..i-1] and Y[0..j-1] */
	for (int i=0; i<=m; i++)
	{
		for (int j=0; j<=n; j++)
		{
			if (i == 0 || j == 0)
				L[i][j] = 0;
			else if (X[i-1] == Y[j-1])
				L[i][j] = L[i-1][j-1] + 1;
			else
				L[i][j] = max(L[i-1][j], L[i][j-1]);
		}
	}

	// Following code is used to print LCS
	int index = L[m][n];

	// Create a character array to store the lcs string
	char lcs[index+1];
	lcs[index] = '\0'; // Set the terminating character

	// Start from the right-most-bottom-most corner and
	// one by one store characters in lcs[]
	int i = m, j = n;
	while (i > 0 && j > 0)
	{
			// If current character in X[] and Y are same, then
			// current character is part of LCS
			if (X[i-1] == Y[j-1])
			{
					lcs[index-1] = X[i-1]; // Put current character in result
					i--; j--; index--;     // reduce values of i, j and index
			}
 
			// If not same, then find the larger of two and
			// go in the direction of larger value
			else if (L[i-1][j] > L[i][j-1])
				i--;
			else
				j--;
	}

	// Print the lcs
	cout << "LCS of " << X << " and " << Y << " is " << lcs << endl;
}

int main()
{
	char X[] = "AGGTAB";
	char Y[] = "GXTXAYB";
	int m = strlen(X);
	int n = strlen(Y);
	LCS(X, Y, m, n);
	// LCS of AGGTAB and GXTXAYB is GTAB
}

```

[↑ top](#recursion)
<br><br><br><br>
<hr>
