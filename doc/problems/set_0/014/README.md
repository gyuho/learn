[*back to problems*](https://github.com/gyuho/learn/tree/master/doc/problems)
<br>

# Problem

**Word-search game.** Implement a program that finds all instances of the word
`code` in the board *below*:

![word_search](img/word_search.png)

<br><br>
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

**FIRST**, what would be the **base case** of *recursion*? You want to end the
recursion when:
1. Have found all the previous letters (e.g. `c`, `o`, `d` for `code`), and 
   have just found the last letter (`e` for `code`).
2. No more position to move (at the end of board).

<br>
**SECOND**, decide what you need to carry around or pass out to *recursive*
functions:
1. Target letter. A target letter being empty means you do not want to
   proceed the search anymore, therefore ending recursion. A target
2. Position on board (row, column in two dimensional array) to search for the
   target letter. This tells which direction to move.
3. You want to store previously found instances (full letters),
   so not to overcount.

<br>
**THIRD**, create a two dimensional array with values from the board.
Start with the first letter as a target letter, at the position
*(0, 0)* in a two dimensional array. And if it finds the target letter,
call *recursive* function onto other directions. Otherwise, keep moving from
*left-top to right-bottom* until it reaches the end.

[↑ top](#problem)
<br><br><br><br>
<hr>




#### Solution #1, in Go

```go

```

[↑ top](#problem)
<br><br><br><br>
<hr>




#### Solution #1, in C++

```cpp

```

[↑ top](#problem)
<br><br><br><br>
<hr>
