[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C++: function

- [function](#function)

[↑ top](#c-function)
<br><br><br><br><hr>


#### function

```cpp
#include <iostream>
using namespace std;
 
// function declaration
int max(int num1, int num2);
 
int main ()
{
   // local variable declaration:
   int a = 1;
   int b = 2;
   int ret = max(a, b);
   cout << "Max is : " << ret << endl;
}
 
int max(int num1, int num2) 
{
   if (num1 > num2)
   	   return num1;
   return num2; 
}

```

[↑ top](#c-function)
<br><br><br><br><hr>
