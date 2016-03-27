[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# Go: red black tree

- [Reference](#reference)
- [Tree](#tree)
- [Left-Leaning Red-Black Tree (`llrb`)](#left-leaning-red-black-tree-llrb)
- [`llrb`: RotateToLeft](#llrb-rotatetoleft)
- [`llrb`: RotateToRight](#llrb-rotatetoright)
- [`llrb`: FlipColor](#llrb-flipcolor)
- [`llrb`: MoveRedFromRightToLeft](#llrb-moveredfromrighttoleft)
- [`llrb`: MoveRedFromLeftToRight](#llrb-moveredfromlefttoright)
- [`llrb`: Insert](#llrb-insert)
- [`llrb`: Search](#llrb-search)
- [`llrb`: Traverse](#llrb-traverse)
- [`llrb`: Delete](#llrb-delete)

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### Reference
- [*Left-leaning Red-Black
  Trees*](https://www.cs.princeton.edu/~rs/talks/LLRB/LLRB.pdf), [*Lecture
  Notes*](http://www.cs.princeton.edu/~rs/talks/LLRB/08Penn.pdf), [**_Video_**](https://www.youtube.com/watch?v=lKmLBOJXZHI) by [*Robert
  Sedgewick*](https://www.cs.princeton.edu/~rs/)
- [*LLRB implementation*](https://github.com/petar/GoLLRB) by [*Petar
  Maymounkov*](http://www.maymounkov.org/)
- [**_LLRB implementation_**](https://github.com/gyuho/llrb)
  by *me*

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### Tree

[*Binary Search Tree*](https://en.wikipedia.org/wiki/Binary_search_tree) is
**_not_** always **_balanced_**, as below:

![bst](img/bst.png)

This is still a valid *binary search tree* but not a [*balanced binary
tree*](https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree). The
**worst case time complexity** of **search** is **_`O(n)`_**, not *`O(log n)`*.
Likewise **average time complexity** of **_insertion_** and **_deletion_**
is **_`O(log n)`_**, *but* the **worst** case is **_`O(n)`_**.


<br><br>
Then what if we **maintain the balance of a binary search tree**? Tree would be
always be **balanced** so **guarantee** **_searching in `O(log n)`_**. This is
where [**red black
tree**](https://en.wikipedia.org/wiki/Red%E2%80%93black_tree)—a
[*self-balancing binary search
tree*](https://en.wikipedia.org/wiki/Self-balancing_binary_search_tree)—comes
in. Like a binary search tree, it is a *good data structure for searching
algorithms*.

![llrb](img/llrb.png)


<br><br>
And what if we have **more than two(binary) children** per node? It would be
[**N-ary tree**](https://en.wikipedia.org/wiki/K-ary_tree). And allowing
**_multiple branches_** per node **decreases tree height**, which means **less
operations are required for searching**—*faster lookup*. This is where
[**b-tree**](https://en.wikipedia.org/wiki/B-tree)—*generalization of a binary
search tree*—comes in. **Database can minimize the number of disk accesses for
data retrieval**.

![btree](img/btree.png)

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### Left-Leaning Red-Black Tree (`llrb`)

**_Left-Leaning Red Black Tree_** (or `llrb`) properties are as follows:

- Every **_node_**(or **_edge_**) is either black or red.
- Every **path from root to null** Node has the **same number of black nodes**.
- Red nodes lean left.
- Two red nodes in a row are not allowed.

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: RotateToLeft

Here's how to [**_`RotateToLeft`_**](https://godoc.org/github.com/gyuho/goraph/llrb#RotateToLeft):

![llrb_rotate_to_left_00](img/llrb_rotate_to_left_00.png)
![llrb_rotate_to_left_01](img/llrb_rotate_to_left_01.png)
![llrb_rotate_to_left_02](img/llrb_rotate_to_left_02.png)
![llrb_rotate_to_left_result](img/llrb_rotate_to_left_result.png)

<br>
And [code](http://play.golang.org/p/A7GbcIoZOT):

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
	root.Black = true
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

	Key   Interface
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
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
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// insert inserts nd2 with nd1 as a root.
func (nd1 *Node) insert(nd2 *Node) *Node {
	if nd1 == nil {
		return nd2
	}
	if nd1.Key.Less(nd2.Key) {
		// nd1 is smaller than nd2
		// nd1 < nd2
		nd1.Right = nd1.Right.insert(nd2)
	} else {
		// nd1 is greater than nd2
		// nd1 >= nd2
		nd1.Left = nd1.Left.insert(nd2)
	}
	// Balance from nd1
	return Balance(nd1)
}

// Insert inserts a Node to a Tree without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
//
// Note that it recursively balances from its parent nodes
// to the root node at the top.
//
// And make sure paint the Root black(not-red).
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)

	// Root node must be always black.
	tr.Root.Black = true
}

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
}

func main() {
	node3 := NewNode(Float64(3))
	node3.Black = true

	node1 := NewNode(Float64(1))
	node1.Black = true

	node13 := NewNode(Float64(13))
	node13.Black = false

	node9 := NewNode(Float64(9))
	node9.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = true

	tr := New(node3)
	tr.Root.Right = node13
	tr.Root.Right.Left = node9
	tr.Root.Right.Right = node17
	tr.Root.Left = node1
	/*
	        3(B)
	      /      \
	   1(B)      13(R)
	            /   \
	         9(B)  17(B)
	*/
	fmt.Println("Before tr.Root = RotateToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	tr.Root = RotateToLeft(tr.Root)
	/*
			   	   13(B)
			   	  /     \
			   3(R)     17(B)
			  /   \
		   1(B)   9(B)
	*/

	fmt.Println("After tr.Root = RotateToLeft(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before tr.Root = RotateToLeft(tr.Root)
	// [1(true)]
	// [[1(true)] 3(true) [[9(true)] 13(false) [17(true)]]]
	// [[9(true)] 13(false) [17(true)]]
	// After tr.Root = RotateToLeft(tr.Root)
	// [[1(true)] 3(false) [9(true)]]
	// [[[1(true)] 3(false) [9(true)]] 13(true) [17(true)]]
	// [17(true)]
}
```

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: RotateToRight

Here's how to [**_`RotateToRight`_**](https://godoc.org/github.com/gyuho/goraph/llrb#RotateToRight):

![llrb_rotate_to_right_00](img/llrb_rotate_to_right_00.png)
![llrb_rotate_to_right_01](img/llrb_rotate_to_right_01.png)
![llrb_rotate_to_right_02](img/llrb_rotate_to_right_02.png)
![llrb_rotate_to_right_result](img/llrb_rotate_to_right_result.png)

<br>
And [code](http://play.golang.org/p/8BuyravQi1):

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
	root.Black = true
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

	Key   Interface
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
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
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// insert inserts nd2 with nd1 as a root.
func (nd1 *Node) insert(nd2 *Node) *Node {
	if nd1 == nil {
		return nd2
	}
	if nd1.Key.Less(nd2.Key) {
		// nd1 is smaller than nd2
		// nd1 < nd2
		nd1.Right = nd1.Right.insert(nd2)
	} else {
		// nd1 is greater than nd2
		// nd1 >= nd2
		nd1.Left = nd1.Left.insert(nd2)
	}
	// Balance from nd1
	return Balance(nd1)
}

// Insert inserts a Node to a Tree without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
//
// Note that it recursively balances from its parent nodes
// to the root node at the top.
//
// And make sure paint the Root black(not-red).
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)

	// Root node must be always black.
	tr.Root.Black = true
}

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
}

func main() {
	node20 := NewNode(Float64(20))
	node20.Black = true

	node39 := NewNode(Float64(39))
	node39.Black = true

	node25 := NewNode(Float64(25))
	node25.Black = false

	node16 := NewNode(Float64(16))
	node16.Black = false

	node15 := NewNode(Float64(15))
	node15.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = true

	tr := New(node20)
	tr.Root.Right = node39
	tr.Root.Right.Left = node25
	tr.Root.Left = node16
	tr.Root.Left.Left = node15
	tr.Root.Left.Right = node17
	/*
	             20(B)
	            /     \
	       16(R)     39(B)
	       /   \       /
	   15(B)  17(B)  25(R)
	*/
	fmt.Println("Before tr.Root = RotateToRight(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	tr.Root = RotateToRight(tr.Root)
	/*
	       16(B)
	      /     \
	   15(B)     20(R)
	            /    \
	        17(B)     39(B)
	                  /
	                25(R)
	*/

	fmt.Println("After tr.Root = RotateToRight(tr.Root)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before tr.Root = RotateToRight(tr.Root)
	// [[15(true)] 16(false) [17(true)]]
	// [[[15(true)] 16(false) [17(true)]] 20(true) [[25(false)] 39(true)]]
	// [[25(false)] 39(true)]
	// After tr.Root = RotateToRight(tr.Root)
	// [15(true)]
	// [[15(true)] 16(true) [[17(true)] 20(false) [[25(false)] 39(true)]]]
	// [[17(true)] 20(false) [[25(false)] 39(true)]]
}
```
[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: FlipColor

And Here's how to [**_`FlipColor`_**](https://godoc.org/github.com/gyuho/goraph/llrb#FlipColor):

![llrb_flip_color](img/llrb_flip_color.png)

<br>
And [code](http://play.golang.org/p/fkfZKuahNT):

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
	root.Black = true
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

	Key   Interface
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
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
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// insert inserts nd2 with nd1 as a root.
func (nd1 *Node) insert(nd2 *Node) *Node {
	if nd1 == nil {
		return nd2
	}
	if nd1.Key.Less(nd2.Key) {
		// nd1 is smaller than nd2
		// nd1 < nd2
		nd1.Right = nd1.Right.insert(nd2)
	} else {
		// nd1 is greater than nd2
		// nd1 >= nd2
		nd1.Left = nd1.Left.insert(nd2)
	}
	// Balance from nd1
	return Balance(nd1)
}

// Insert inserts a Node to a Tree without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
//
// Note that it recursively balances from its parent nodes
// to the root node at the top.
//
// And make sure paint the Root black(not-red).
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)

	// Root node must be always black.
	tr.Root.Black = true
}

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
}

func main() {
	node3 := NewNode(Float64(3))
	node3.Black = true

	node1 := NewNode(Float64(1))
	node1.Black = true

	node13 := NewNode(Float64(13))
	node13.Black = true

	node9 := NewNode(Float64(9))
	node9.Black = true

	node17 := NewNode(Float64(17))
	node17.Black = true

	tr := New(node3)
	tr.Root.Right = node13
	tr.Root.Right.Left = node9
	tr.Root.Right.Right = node17
	tr.Root.Left = node1
	/*
	        3(B)
	      /      \
	   1(B)      13(B)
	            /   \
	         9(B)  17(B)
	*/
	fmt.Println("Before FlipColor(tr.Root.Right)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)

	FlipColor(tr.Root.Right)
	/*
	        3(B)
	      /      \
	   1(B)      13(R)
	            /   \
	         9(R)  17(R)
	*/

	fmt.Println("After FlipColor(tr.Root.Right)")
	fmt.Println(tr.Root.Left)
	fmt.Println(tr.Root)
	fmt.Println(tr.Root.Right)
	// Output:
	// Before FlipColor(tr.Root.Right)
	// [1(true)]
	// [[1(true)] 3(true) [[9(true)] 13(true) [17(true)]]]
	// [[9(true)] 13(true) [17(true)]]
	// After FlipColor(tr.Root.Right)
	// [1(true)]
	// [[1(true)] 3(true) [[9(false)] 13(false) [17(false)]]]
	// [[9(false)] 13(false) [17(false)]]
}
```

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: MoveRedFromRightToLeft

Here's how to [**_`MoveRedFromRightToLeft`_**](https://godoc.org/github.com/gyuho/goraph/llrb#MoveRedFromRightToLeft):

![llrb_move_red_from_right_to_left_00](img/llrb_move_red_from_right_to_left_00.png)
![llrb_move_red_from_right_to_left_01](img/llrb_move_red_from_right_to_left_01.png)
![llrb_move_red_from_right_to_left_02](img/llrb_move_red_from_right_to_left_02.png)
![llrb_move_red_from_right_to_left_03](img/llrb_move_red_from_right_to_left_03.png)
![llrb_move_red_from_right_to_left_04](img/llrb_move_red_from_right_to_left_04.png)
![llrb_move_red_from_right_to_left_05](img/llrb_move_red_from_right_to_left_05.png)
![llrb_move_red_from_right_to_left_result](img/llrb_move_red_from_right_to_left_result.png)

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: MoveRedFromLeftToRight

Here's how to [**_`MoveRedFromLeftToRight`_**](https://godoc.org/github.com/gyuho/goraph/llrb#MoveRedFromLeftToRight):

![llrb_move_red_from_left_to_right_00](img/llrb_move_red_from_left_to_right_00.png)
![llrb_move_red_from_left_to_right_01](img/llrb_move_red_from_left_to_right_01.png)
![llrb_move_red_from_left_to_right_02](img/llrb_move_red_from_left_to_right_02.png)
![llrb_move_red_from_left_to_right_03](img/llrb_move_red_from_left_to_right_03.png)
![llrb_move_red_from_left_to_right_04](img/llrb_move_red_from_left_to_right_04.png)
![llrb_move_red_from_left_to_right_result](img/llrb_move_red_from_left_to_right_result.png)

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: Insert

Here's how `llrb`(*Left-Leaning Red Black Tree*) inserts.
Note that `insertion` always sets the root as black at the end:

![llrb_insert_00](img/llrb_insert_00.png)
![llrb_insert_01](img/llrb_insert_01.png)
![llrb_insert_02](img/llrb_insert_02.png)
![llrb_insert_03](img/llrb_insert_03.png)
![llrb_insert_04](img/llrb_insert_04.png)
![llrb_insert_05](img/llrb_insert_05.png)
![llrb_insert_06](img/llrb_insert_06.png)
![llrb_insert_07](img/llrb_insert_07.png)
![llrb_insert_08](img/llrb_insert_08.png)
![llrb_insert_09](img/llrb_insert_09.png)
![llrb_insert_10](img/llrb_insert_10.png)
![llrb_insert_11](img/llrb_insert_11.png)
![llrb_insert_12](img/llrb_insert_12.png)
![llrb_insert_13](img/llrb_insert_13.png)
![llrb_insert_14](img/llrb_insert_14.png)
![llrb_insert_15](img/llrb_insert_15.png)
![llrb_insert_16](img/llrb_insert_16.png)
![llrb_insert_17](img/llrb_insert_17.png)
![llrb_insert_18](img/llrb_insert_18.png)
![llrb_insert_19](img/llrb_insert_19.png)

<br>
And [code](http://play.golang.org/p/nu_bXrICJv):

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
	root.Black = true
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

	Key   Interface
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
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
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// insert inserts nd2 with nd1 as a root.
func (nd1 *Node) insert(nd2 *Node) *Node {
	if nd1 == nil {
		return nd2
	}
	if nd1.Key.Less(nd2.Key) {
		// nd1 is smaller than nd2
		// nd1 < nd2
		nd1.Right = nd1.Right.insert(nd2)
	} else {
		// nd1 is greater than nd2
		// nd1 >= nd2
		nd1.Left = nd1.Left.insert(nd2)
	}
	// Balance from nd1
	return Balance(nd1)
}

// Insert inserts a Node to a Tree without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
//
// Note that it recursively balances from its parent nodes
// to the root node at the top.
//
// And make sure paint the Root black(not-red).
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)

	// Root node must be always black.
	tr.Root.Black = true
}

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// MoveRedFromRightToLeft moves Red Node
// from Right sub-Tree to Left sub-Tree.
// Left and Right children must be present
func MoveRedFromRightToLeft(nd *Node) *Node {
	FlipColor(nd)
	if isRed(nd.Right.Left) {
		nd.Right = RotateToRight(nd.Right)
		nd = RotateToLeft(nd)
		FlipColor(nd)
	}
	return nd
}

// MoveRedFromLeftToRight moves Red Node
// from Left sub-Tree to Right sub-Tree.
// Left and Right children must be present
func MoveRedFromLeftToRight(nd *Node) *Node {
	FlipColor(nd)
	if isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
		FlipColor(nd)
	}
	return nd
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

// FixUp fixes the balances of the Node.
func FixUp(nd *Node) *Node {
	if isRed(nd.Right) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
}

func main() {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13}
	for _, num := range nums {
		tr.Insert(NewNode(Float64(num)))
	}
	fmt.Println(tr)
	// [[1(true)] 3(true) [[9(false)] 13(true)]]
	/*
	     3
	    / \
	   1   13
	       /
	      9
	*/
}
```

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: Search

`search` is exactly the same as **Binary Search Tree**, as [here](http://play.golang.org/p/4uYTR03HVy).

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
	root.Black = true
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

	Key   Interface
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
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
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// insert inserts nd2 with nd1 as a root.
func (nd1 *Node) insert(nd2 *Node) *Node {
	if nd1 == nil {
		return nd2
	}
	if nd1.Key.Less(nd2.Key) {
		// nd1 is smaller than nd2
		// nd1 < nd2
		nd1.Right = nd1.Right.insert(nd2)
	} else {
		// nd1 is greater than nd2
		// nd1 >= nd2
		nd1.Left = nd1.Left.insert(nd2)
	}
	// Balance from nd1
	return Balance(nd1)
}

// Insert inserts a Node to a Tree without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
//
// Note that it recursively balances from its parent nodes
// to the root node at the top.
//
// And make sure paint the Root black(not-red).
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)

	// Root node must be always black.
	tr.Root.Black = true
}

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// MoveRedFromRightToLeft moves Red Node
// from Right sub-Tree to Left sub-Tree.
// Left and Right children must be present
func MoveRedFromRightToLeft(nd *Node) *Node {
	FlipColor(nd)
	if isRed(nd.Right.Left) {
		nd.Right = RotateToRight(nd.Right)
		nd = RotateToLeft(nd)
		FlipColor(nd)
	}
	return nd
}

// MoveRedFromLeftToRight moves Red Node
// from Left sub-Tree to Right sub-Tree.
// Left and Right children must be present
func MoveRedFromLeftToRight(nd *Node) *Node {
	FlipColor(nd)
	if isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
		FlipColor(nd)
	}
	return nd
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

// FixUp fixes the balances of the Node.
func FixUp(nd *Node) *Node {
	if isRed(nd.Right) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
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
func (tr *Tree) Max() *Node {
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

func main() {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13}
	for _, num := range nums {
		tr.Insert(NewNode(Float64(num)))
	}
	fmt.Println(tr.Search(Float64(9)))
	// [9(false)]
	/*
	     3
	    / \
	   1   13
	       /
	      9
	*/
}
```

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: Traverse

`traverse` is exactly the same as **Binary Search Tree**, as [here](http://play.golang.org/p/CvjjTVTTJD).

```go
package main

import (
	"bytes"
	"fmt"
)

// Tree contains a Root node of a binary search tree.
type Tree struct {
	Root *Node
}

// New returns a new Tree with its root Node.
func New(root *Node) *Tree {
	tr := &Tree{}
	root.Black = true
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

	Key   Interface
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
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
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// insert inserts nd2 with nd1 as a root.
func (nd1 *Node) insert(nd2 *Node) *Node {
	if nd1 == nil {
		return nd2
	}
	if nd1.Key.Less(nd2.Key) {
		// nd1 is smaller than nd2
		// nd1 < nd2
		nd1.Right = nd1.Right.insert(nd2)
	} else {
		// nd1 is greater than nd2
		// nd1 >= nd2
		nd1.Left = nd1.Left.insert(nd2)
	}
	// Balance from nd1
	return Balance(nd1)
}

// Insert inserts a Node to a Tree without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
//
// Note that it recursively balances from its parent nodes
// to the root node at the top.
//
// And make sure paint the Root black(not-red).
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)

	// Root node must be always black.
	tr.Root.Black = true
}

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// MoveRedFromRightToLeft moves Red Node
// from Right sub-Tree to Left sub-Tree.
// Left and Right children must be present
func MoveRedFromRightToLeft(nd *Node) *Node {
	FlipColor(nd)
	if isRed(nd.Right.Left) {
		nd.Right = RotateToRight(nd.Right)
		nd = RotateToLeft(nd)
		FlipColor(nd)
	}
	return nd
}

// MoveRedFromLeftToRight moves Red Node
// from Left sub-Tree to Right sub-Tree.
// Left and Right children must be present
func MoveRedFromLeftToRight(nd *Node) *Node {
	FlipColor(nd)
	if isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
		FlipColor(nd)
	}
	return nd
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

// FixUp fixes the balances of the Node.
func FixUp(nd *Node) *Node {
	if isRed(nd.Right) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
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
func (tr *Tree) Max() *Node {
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

func main() {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13}
	for _, num := range nums {
		tr.Insert(NewNode(Float64(num)))
	}

	buf1 := new(bytes.Buffer)
	ch1 := make(chan string)
	go tr.PreOrder(ch1) // root, left, right
	for {
		v, ok := <-ch1
		if !ok {
			break
		}
		buf1.WriteString(v)
		buf1.WriteString(" ")
	}
	fmt.Println("PreOrder:", buf1.String()) // PreOrder: 3 1 13 9

	buf2 := new(bytes.Buffer)
	ch2 := make(chan string)
	go tr.InOrder(ch2) // left, root, right
	for {
		v, ok := <-ch2
		if !ok {
			break
		}
		buf2.WriteString(v)
		buf2.WriteString(" ")
	}
	fmt.Println("InOrder:", buf2.String()) // InOrder: 1 3 9 13

	buf3 := new(bytes.Buffer)
	ch3 := make(chan string)
	go tr.PostOrder(ch3) // left, right, root
	for {
		v, ok := <-ch3
		if !ok {
			break
		}
		buf3.WriteString(v)
		buf3.WriteString(" ")
	}
	fmt.Println("PostOrder:", buf3.String()) // PostOrder: 1 9 13 3

	buf4 := new(bytes.Buffer)
	nodes4 := tr.LevelOrder() // from top-level
	for _, v := range nodes4 {
		buf4.WriteString(fmt.Sprintf("%v ", v.Key))
	}
	fmt.Println("LevelOrder:", buf4.String()) // LevelOrder: 3 1 13 9
}

// PreOrder traverses from Root, Left-SubTree, and Right-SubTree. (DFS)
func (tr *Tree) PreOrder(ch chan string) {
	preOrder(tr.Root, ch)
	close(ch)
}

func preOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	ch <- fmt.Sprintf("%v", nd.Key) // root
	preOrder(nd.Left, ch)           // left
	preOrder(nd.Right, ch)          // right
}

// ComparePreOrder returns true if two Trees are same with PreOrder.
func ComparePreOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.PreOrder(ch1)
	go t2.PreOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// InOrder traverses from Left-SubTree, Root, and Right-SubTree. (DFS)
func (tr *Tree) InOrder(ch chan string) {
	inOrder(tr.Root, ch)
	close(ch)
}

func inOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	inOrder(nd.Left, ch)            // left
	ch <- fmt.Sprintf("%v", nd.Key) // root
	inOrder(nd.Right, ch)           // right
}

// CompareInOrder returns true if two Trees are same with InOrder.
func CompareInOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.InOrder(ch1)
	go t2.InOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// PostOrder traverses from Left-SubTree, Right-SubTree, and Root.
func (tr *Tree) PostOrder(ch chan string) {
	postOrder(tr.Root, ch)
	close(ch)
}

func postOrder(nd *Node, ch chan string) {
	// leaf node
	if nd == nil {
		return
	}
	postOrder(nd.Left, ch)          // left
	postOrder(nd.Right, ch)         // right
	ch <- fmt.Sprintf("%v", nd.Key) // root
}

// ComparePostOrder returns true if two Trees are same with PostOrder.
func ComparePostOrder(t1, t2 *Tree) bool {
	ch1, ch2 := make(chan string), make(chan string)
	go t1.PostOrder(ch1)
	go t2.PostOrder(ch2)
	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

// LevelOrder traverses the Tree with Breadth First Search.
// (http://en.wikipedia.org/wiki/Tree_traversal#Breadth-first_2)
//
//	levelorder(root)
//	  q = empty queue
//	  q.enqueue(root)
//	  while not q.empty do
//	    node := q.dequeue()
//	    visit(node)
//	    if node.left ≠ null then
//	      q.enqueue(node.left)
//	    if node.right ≠ null then
//	      q.enqueue(node.right)
//
func (tr *Tree) LevelOrder() []*Node {
	visited := []*Node{}
	queue := []*Node{tr.Root}
	for len(queue) != 0 {
		nd := queue[0]
		queue = queue[1:len(queue):len(queue)]
		visited = append(visited, nd)
		if nd.Left != nil {
			queue = append(queue, nd.Left)
		}
		if nd.Right != nil {
			queue = append(queue, nd.Right)
		}
	}
	return visited
}
```

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>


#### `llrb`: Delete

`llrb` deletion is complicated compared to binary search tree,
because it needs to maintain the `llrb` properties of:

- Every **_node_**(or **_edge_**) is either black or red.
- Every **path from root to null** Node has the **same number of black nodes**.
- Red nodes lean left.
- Two red nodes in a row are not allowed.

That is, **deleting** a *red* node won't violate the properties. But, deleting
a *black* node can change the number of black nodes from root to null in some
paths. And it also takes a lot of rotations [when the implementation has no parental
node pointer](http://www.read.seas.harvard.edu/~kohler/notes/llrb.html).
Note that just like a newly-inserted node is *red*, we always want to delete
a *red* node. So, if the node is not *red*, we need to make it **red**.

First, let's look at how the tree changes after each deletion:

![llrb_delete_result_00](img/llrb_delete_result_00.png)
![llrb_delete_result_01](img/llrb_delete_result_01.png)
![llrb_delete_result_02](img/llrb_delete_result_02.png)
![llrb_delete_result_03](img/llrb_delete_result_03.png)

<br>
```
Delete Algorithm:
1. Start 'delete' from tree Root.

2. Call 'delete' method recursively on each Node from binary search path.
	- e.g. If the key to delete is greater than Root's key
		, call 'delete' on Right Node.

# start

3. Recursive 'tree.delete(nd, key)'

	if key < nd.Key:

		if nd.Left is empty:
			# then nothing to delete, so return nil
			return nd, nil

		if (nd.Left is Black) and (nd.Left.Left is Black):
			# then move Red from Right to Left to update nd
			nd = MoveRedFromRightToLeft(nd)

		# recursively call 'delete' to update nd.Left
		nd.Left, deleted = tr.delete(nd.Left, key)

	else if key >= nd.Key:

		if nd.Left is Red:
			# RotateToRight(nd) to update nd
			nd = RotateToRight(nd)

		if (key == nd.Key) and nd.Right is empty:
			# then return nil, nd.Key to recursively update nd
			return nil, nd.Key

		if (nd.Right is not empty)
		and (nd.Right is Black)
		and (nd.Right.Left is Black):
			# then move Red from Left to Right to update nd
			nd = MoveRedFromLeftToRight(nd)

		if (key == nd.Key):
			# then DeleteMin of nd.Right to update nd.Right
			nd.Right, subDeleted = DeleteMin(nd.Right)

			# and then update nd.Key with DeleteMin(nd.Right)
			deleted, nd.Key = nd.Key, subDeleted

		else if key != nd.Key:
			# recursively call 'delete' to update nd.Right
			nd.Right, deleted = tr.delete(nd.Right, key)

	# recursively FixUp upwards to Root
	return FixUp(nd), deleted

# end

4. If the tree's Root is not nil, set Root Black.

5. Return the Interface(nil if the key does not exist.)
```

<br>
![llrb_delete_00](img/llrb_delete_00.png)
![llrb_delete_01](img/llrb_delete_01.png)
![llrb_delete_02](img/llrb_delete_02.png)
![llrb_delete_03](img/llrb_delete_03.png)
![llrb_delete_04](img/llrb_delete_04.png)
![llrb_delete_05](img/llrb_delete_05.png)
![llrb_delete_06](img/llrb_delete_06.png)
![llrb_delete_07](img/llrb_delete_07.png)
![llrb_delete_08](img/llrb_delete_08.png)
![llrb_delete_09](img/llrb_delete_09.png)
![llrb_delete_10](img/llrb_delete_10.png)
![llrb_delete_11](img/llrb_delete_11.png)
![llrb_delete_12](img/llrb_delete_12.png)
![llrb_delete_13](img/llrb_delete_13.png)
![llrb_delete_14](img/llrb_delete_14.png)
![llrb_delete_15](img/llrb_delete_15.png)
![llrb_delete_16](img/llrb_delete_16.png)
![llrb_delete_17](img/llrb_delete_17.png)
![llrb_delete_18](img/llrb_delete_18.png)
![llrb_delete_19](img/llrb_delete_19.png)
![llrb_delete_20](img/llrb_delete_20.png)
![llrb_delete_21](img/llrb_delete_21.png)
![llrb_delete_22](img/llrb_delete_22.png)
![llrb_delete_23](img/llrb_delete_23.png)
![llrb_delete_24](img/llrb_delete_24.png)
![llrb_delete_25](img/llrb_delete_25.png)
![llrb_delete_26](img/llrb_delete_26.png)
![llrb_delete_27](img/llrb_delete_27.png)

<br>
And here's how it actually happens:

<br>
And [code](http://play.golang.org/p/AT-nBWV3Ve):

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
	root.Black = true
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

	Key   Interface
	Black bool // True when the color of parent link is black.
	// In Left-Leaning Red-Black tree, new nodes are always red
	// because the zero boolean value is false.
	// Null links are black.

	// Right is a right child Node.
	Right *Node
}

// NewNode returns a new Node.
func NewNode(key Interface) *Node {
	nd := &Node{}
	nd.Key = key
	nd.Black = false
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
	s += fmt.Sprintf("%v(%v)", nd.Key, nd.Black)
	if nd.Right != nil {
		s += " " + nd.Right.String()
	}
	return "[" + s + "]"
}

func isRed(nd *Node) bool {
	if nd == nil {
		return false
	}
	return !nd.Black
}

// insert inserts nd2 with nd1 as a root.
func (nd1 *Node) insert(nd2 *Node) *Node {
	if nd1 == nil {
		return nd2
	}
	if nd1.Key.Less(nd2.Key) {
		// nd1 is smaller than nd2
		// nd1 < nd2
		nd1.Right = nd1.Right.insert(nd2)
	} else {
		// nd1 is greater than nd2
		// nd1 >= nd2
		nd1.Left = nd1.Left.insert(nd2)
	}
	// Balance from nd1
	return Balance(nd1)
}

// Insert inserts a Node to a Tree without replacement.
// It does standard BST insert and colors the new link red.
// If the new red link is a right link, rotate left.
// If two left red links in a row, rotate to right and flip color.
// (https://youtu.be/lKmLBOJXZHI?t=20m43s)
//
// Note that it recursively balances from its parent nodes
// to the root node at the top.
//
// And make sure paint the Root black(not-red).
func (tr *Tree) Insert(nd *Node) {
	if tr.Root == nd {
		return
	}
	tr.Root = tr.Root.insert(nd)

	// Root node must be always black.
	tr.Root.Black = true
}

// RotateToLeft runs when there is a right-leaning link.
// tr.Root = RotateToLeft(tr.Root) overwrite the Root
// with the new top Node.
func RotateToLeft(nd *Node) *Node {
	if nd.Right.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Right child
	x := nd.Right
	nd.Right = x.Left
	x.Left = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// RotateToRight runs when there are two left red links in a row.
// tr.Root = RotateToRight(tr.Root) overwrite the Root
// with the new top Node.
func RotateToRight(nd *Node) *Node {
	if nd.Left.Black {
		panic("Can't rotate a black link")
	}

	// exchange x and nd
	// nd is parent node, x is Left child
	x := nd.Left
	nd.Left = x.Right
	x.Right = nd

	x.Black = nd.Black
	nd.Black = false

	return x
}

// FlipColor flips the color.
// Left and Right children must be present
func FlipColor(nd *Node) {
	// nd is parent node
	nd.Black = !nd.Black
	nd.Left.Black = !nd.Left.Black
	nd.Right.Black = !nd.Right.Black
}

// MoveRedFromRightToLeft moves Red Node
// from Right sub-Tree to Left sub-Tree.
// Left and Right children must be present
func MoveRedFromRightToLeft(nd *Node) *Node {
	FlipColor(nd)
	if isRed(nd.Right.Left) {
		nd.Right = RotateToRight(nd.Right)
		nd = RotateToLeft(nd)
		FlipColor(nd)
	}
	return nd
}

// MoveRedFromLeftToRight moves Red Node
// from Left sub-Tree to Right sub-Tree.
// Left and Right children must be present
func MoveRedFromLeftToRight(nd *Node) *Node {
	FlipColor(nd)
	if isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
		FlipColor(nd)
	}
	return nd
}

// Balance balances the Node.
func Balance(nd *Node) *Node {
	// nd is parent node
	if isRed(nd.Right) && !isRed(nd.Left) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

// FixUp fixes the balances of the Node.
func FixUp(nd *Node) *Node {
	if isRed(nd.Right) {
		nd = RotateToLeft(nd)
	}
	if isRed(nd.Left) && isRed(nd.Left.Left) {
		nd = RotateToRight(nd)
	}
	if isRed(nd.Left) && isRed(nd.Right) {
		FlipColor(nd)
	}
	return nd
}

type Float64 float64

// Less returns true if float64(a) < float64(b).
func (a Float64) Less(b Interface) bool {
	return a < b.(Float64)
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
func (tr *Tree) Max() *Node {
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

func main() {
	root := NewNode(Float64(1))
	tr := New(root)
	nums := []float64{3, 9, 13, 17, 20, 25, 39, 16, 15, 2, 2.5}
	for _, num := range nums {
		tr.Insert(NewNode(Float64(num)))
	}

	fmt.Println("Deleting start!")
	fmt.Println("Deleted", tr.Delete(Float64(39)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Key)
	fmt.Println()

	fmt.Println("Deleted", tr.Delete(Float64(20)))
	fmt.Println(tr.Root.Left.Key)
	fmt.Println(tr.Root.Key)
	fmt.Println(tr.Root.Right.Key)
	fmt.Println()

	/*
	   Deleting start!
	   Deleted 39
	   3
	   13
	   20

	   Deleted 20
	   3
	   13
	   16
	*/
}

// DeleteMin deletes the minimum-Key Node of the sub-Tree.
func DeleteMin(nd *Node) (*Node, Interface) {
	if nd == nil {
		return nil, nil
	}
	if nd.Left == nil {
		return nil, nd.Key
	}
	if !isRed(nd.Left) && !isRed(nd.Left.Left) {
		nd = MoveRedFromRightToLeft(nd)
	}
	var deleted Interface
	nd.Left, deleted = DeleteMin(nd.Left)
	return FixUp(nd), deleted
}

// DeleteMin deletes the minimum-Key Node of the Tree.
// It returns the minimum Key or nil.
func (tr *Tree) DeleteMin() Interface {
	var deleted Interface
	tr.Root, deleted = DeleteMin(tr.Root)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

// Delete deletes the node with the Key and returns the Key Interface.
// It returns nil if the Key does not exist in the tree.
//
//
//	Delete Algorithm:
//	1. Start 'delete' from tree Root.
//
//	2. Call 'delete' method recursively on each Node from binary search path.
//		- e.g. If the key to delete is greater than Root's key
//			, call 'delete' on Right Node.
//
//
//	# start
//
//	3. Recursive 'tree.delete(nd, key)'
//
//		if key < nd.Key:
//
//			if nd.Left is empty:
//				# then nothing to delete, so return nil
//				return nd, nil
//
//			if (nd.Left is Black) and (nd.Left.Left is Black):
//				# then move Red from Right to Left to update nd
//				nd = MoveRedFromRightToLeft(nd)
//
//			# recursively call 'delete' to update nd.Left
//			nd.Left, deleted = tr.delete(nd.Left, key)
//
//		else if key >= nd.Key:
//
//			if nd.Left is Red:
//				# RotateToRight(nd) to update nd
//				nd = RotateToRight(nd)
//
//			if (key == nd.Key) and nd.Right is empty:
//				# then return nil, nd.Key to recursively update nd
//				return nil, nd.Key
//
//			if (nd.Right is not empty)
//			and (nd.Right is Black)
//			and (nd.Right.Left is Black):
//				# then move Red from Left to Right to update nd
//				nd = MoveRedFromLeftToRight(nd)
//
//			if (key == nd.Key):
//				# then DeleteMin of nd.Right to update nd.Right
//				nd.Right, subDeleted = DeleteMin(nd.Right)
//
//				# and then update nd.Key with DeleteMin(nd.Right)
//				deleted, nd.Key = nd.Key, subDeleted
//
//			else if key != nd.Key:
//				# recursively call 'delete' to update nd.Right
//				nd.Right, deleted = tr.delete(nd.Right, key)
//
//		# recursively FixUp upwards to Root
//		return FixUp(nd), deleted
//
//	# end
//
//
//	4. If the tree's Root is not nil, set Root Black.
//
//	5. Return the Interface(nil if the key does not exist.)
//
func (tr *Tree) Delete(key Interface) Interface {
	var deleted Interface
	tr.Root, deleted = tr.delete(tr.Root, key)
	if tr.Root != nil {
		tr.Root.Black = true
	}
	return deleted
}

func (tr *Tree) delete(nd *Node, key Interface) (*Node, Interface) {
	if nd == nil {
		return nil, nil
	}

	var deleted Interface

	// if key is Less than nd.Key
	if key.Less(nd.Key) {

		// if key is Less than nd.Key
		// and nd.Left is empty
		if nd.Left == nil {

			// then nothing to delete
			// so return the nil
			return nd, nil
		}

		// if key is Less than nd.Key
		// and nd.Left is Black
		// and nd.Left.Left is Black
		if !isRed(nd.Left) && !isRed(nd.Left.Left) {

			// then MoveRedFromRightToLeft(nd)
			nd = MoveRedFromRightToLeft(nd)
		}

		// and recursively call tr.delete(nd.Left, key)
		nd.Left, deleted = tr.delete(nd.Left, key)

	} else {
		// if key is not Less than nd.Key
		//(or key is greater than or equal to nd.Key)
		//(or key >= nd.Key)

		// and nd.Left is Red
		if isRed(nd.Left) {

			// then RotateToRight(nd)
			nd = RotateToRight(nd)
		}

		// and nd.Key is not Less than key
		// (or nd.Key >= key)
		// (or key == nd.Key)
		// and nd.Right is empty
		if !nd.Key.Less(key) && nd.Right == nil {
			// then return nil to delete the key
			return nil, nd.Key
		}

		// and nd.Right is not empty
		// and nd.Right is Black
		// and nd.Right.Left is Black
		if nd.Right != nil && !isRed(nd.Right) && !isRed(nd.Right.Left) {
			// then MoveRedFromLeftToRight(nd)
			nd = MoveRedFromLeftToRight(nd)
		}

		// and key == nd.Key
		if !nd.Key.Less(key) {

			var subDeleted Interface

			// then DeleteMin(nd.Right)
			nd.Right, subDeleted = DeleteMin(nd.Right)
			if subDeleted == nil {
				panic("Unexpected nil value")
			}

			// and update nd.Key with DeleteMin(nd.Right)
			deleted, nd.Key = nd.Key, subDeleted

		} else {
			// if updated nd.Key is Less than key (nd.Key < key) to update nd.Right
			nd.Right, deleted = tr.delete(nd.Right, key)
		}
	}

	// recursively FixUp upwards to Root
	return FixUp(nd), deleted
}
```

[↑ top](#go-red-black-tree)
<br><br><br><br><hr>
