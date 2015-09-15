[*back to problems*](https://github.com/gyuho/learn/tree/master/doc/problems)
<br>

# Problem

**Word-search game.** Implement a program that finds all instances of the word
`code` in the board *below*:

![word_search](img/word_search.png)

<br>
<br>
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



#### Algorithm #1

**Use recursion for multiple related decisions.** This is like a maze:
*each position in board leads to 8 other choices*. You need to decide which
direction to move: *left*, *right*, *up*, *down*, or *diagonals*. It's
multiple related decisions, so it's natural to try *recursion*. I will try
brute-forcing with recursion at every single row and column. And later
see if we can do better.

**First**, what would be the **base case** of *recursion*? You want to end the
recursion when you already have found all the previous letters
(e.g. `c`, `o`, `d` for `code`), and you just find the last letter
(`e` for `code`).

**Second**, decide what you need to carry around or pass out to *recursive*
function calls.


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
