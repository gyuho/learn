[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# bash, vim

I primarily work with Ubuntu, bash, vim.

- [Reference](#reference)
- [ubuntu](#ubuntu)
- [terminal](#terminal)
- [tmux](#tmux)
- [`vim`](#vim)

[↑ top](#bash-vim)
<br><br><br><br>
<hr>





#### Reference

- [*Learn Enough Command Line to Be Dangerous by Michael Hartl*](http://www.learnenough.com/command-line)
- [*The art of command line*](https://github.com/jlevy/the-art-of-command-line)

[↑ top](#bash-vim)
<br><br><br><br>
<hr>










#### ubuntu

- <kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>t</kbd> : `launch terminal`
- <kbd>print</kbd> : `Take a screenshot of an area`
- <kbd>alt</kbd> + <kbd>space</kbd> : `switch to previous (input) source`
- <kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>m</kbd> : `maximize window`
- <kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>r</kbd> : `restore window`
- <kbd>alt</kbd> + <kbd>q</kbd> : `close window`
- <kbd>ctrl</kbd> + <kbd>alt</kbd> + <kbd>v</kbd> : `maximize window vertically`

[↑ top](#bash-vim)
<br><br><br><br>
<hr>







#### terminal

- <kbd>ctrl</kbd> + <kbd>t</kbd> : `new tab`
- <kbd>ctrl</kbd> + <kbd>w</kbd> : `close tab`
- <kbd>alt</kbd> + <kbd>q</kbd> : `close window`
- <kbd>ctrl</kbd> + <kbd>c</kbd> : `copy`
- <kbd>ctrl</kbd> + <kbd>v</kbd> : `paste`
- <kbd>ctrl</kbd> + <kbd>shift</kbd> + <kbd>+</kbd> : `zoom in`
- <kbd>ctrl</kbd> + <kbd>-</kbd> : `zoom out`

[↑ top](#bash-vim)
<br><br><br><br>
<hr>







#### tmux

- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>t</kbd> : `show time`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>%</kbd> : `split vertically`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>"</kbd> : `split horizontally`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>arrow</kbd> : `switch`
- <kbd>ctrl</kbd> + <kbd>b</kbd> + <kbd>arrow</kbd> : `resize`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, `:resize-pane -R 10` : `resize`
- <kbd>ctrl</kbd> + <kbd>b</kbd>, <kbd>x</kbd> : `kill pane`

[↑ top](#bash-vim)
<br><br><br><br>
<hr>






#### `vim`

```
           l
          /
         k
        /
       j
      /
     h
```

- <kbd>i</kbd>  : `insert mode`
- <kbd>o</kbd>  : `insert a new line and start insert mode`
- <kbd>:</kbd> + <kbd>1</kbd> : `beginning of a file`
- <kbd>G</kbd> : `end of a file`
- <kbd>0</kbd> : `beginning of a line`
- <kbd>$</kbd> : `end of a line`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>s</kbd> : `split horizontally`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>v</kbd> : `split vertically`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>arrow</kbd> : `move between panels`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>ctrl</kbd> + <kbd>w</kbd> : `move between panels`
- <kbd>ctrl</kbd> + <kbd>w</kbd>, <kbd>q</kbd> : `close a panel`
- <kbd>:vertical resize +10</kbd> : `horizontal increase`
- <kbd>F2</kbd> : `toggle NerdTree`
- <kbd>F3</kbd> : `cycle NerdTree`
- <kbd>F3</kbd>, <kbd>m</kbd>, <kbd>a</kbd> : `create a new file from NerdTree`
- <kbd>shift</kbd> + <kbd>r</kbd> : `refresh NerdTree`
- <kbd>g</kbd> + <kbd>c</kbd> : `toggle comments`
- <kbd>d</kbd> + <kbd>d</kbd> : `delete the current line`
- <kbd>d</kbd> + <kbd>w</kbd> : `delete the next word`
- <kbd>d</kbd> + <kbd>→</kbd> : `delete the character after cursor`
- <kbd>d</kbd> + <kbd>←</kbd> : `delete the character before cursor`
- <kbd>b</kbd> : `beginning of present or previous word`
- <kbd>w</kbd> : `beginning of next word`
- <kbd>e</kbd> : `end of present word`
- <kbd>ea</kbd> : `append at the end of word`
- <kbd>A</kbd> : `append at the end of line`
- <kbd>b</kbd> + <kbd>v</kbd> + <kbd>e</kbd> : `select the current word`
- <kbd>d</kbd> + <kbd>a</kbd> + <kbd>w</kbd> : `delete the current word`
- <kbd>c</kbd> + <kbd>a</kbd> + <kbd>w</kbd> : `delete the current word, and insert`
- <kbd>y</kbd> : `copy the selected characters`
- <kbd>yy</kbd> : `copy the entire line`
- <kbd>“</kbd> + <kbd>+</kbd> + <kbd>y</kbd> : `copy to clipboard`
- <kbd>u</kbd> : `undo`
- <kbd>ctrl</kbd> + <kbd>r</kbd> : `undo the undo`
- <kbd>P</kbd> : `paste before the cursor`
- <kbd>p</kbd> : `paste after`
- <kbd>/</kbd>, <kbd>n</kbd> : `find, and next`
- <kbd>:%s/old/new/gc</kbd> : `find, and replace`
- <kbd>.</kbd> : `prepend in comparison with the line above`
- <kbd>ctrl</kbd> + <kbd>n</kbd>, <kbd>i</kbd> : `expand selection, and edit`
- <kbd>ctrl</kbd> + <kbd>z</kbd> : `go back to terminal`
- <kbd>fg</kbd> : `come back to vim`
- <kbd>ctrl</kbd> + <kbd>p</kbd> : `in insert mode, autocomplete`
- <kbd>ctrl</kbd> + <kbd>p</kbd> : `in visual mode, fuzzy file search`
- <kbd>ctrl</kbd> + <kbd>p</kbd>, <kbd>p</kbd> : `in visual mode, repeat fuzzy file search`
- <kbd>space</kbd> + <kbd>w</kbd> : `:w`
- <kbd>space</kbd> + <kbd>q</kbd> : `:q`

[↑ top](#bash-vim)
<br><br><br><br>
<hr>
