[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C++: logic, loop

- [logic](#logic)
- [`if`](#if)
- [`switch`](#switch)
- [`for`](#for)
- [`while`](#while)
- [fizzbuzz](#fizzbuzz)

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### logic

```cpp
#include <iostream>
using namespace std;

int main()
{
	bool b1 = true;
	bool b2 = false;
	cout << b1 && b2; // 1
	cout << endl;
	cout << b1 || b2; // 1
	cout << endl;

	int a = 10, b = 15, c = 20;
	cout << "a < b is "
		<< (a < b)
		<< endl;
	// a < b is 1
	
	cout << "b > c is "
		<< (b > c)
		<< endl;
	// b > c is 0
}

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### `if` 

```cpp
#include <iostream>
using namespace std;

int main()
{
	int a = 100;

	if ( a == 1 )
	{
		cout << "a is 1" << endl;
	}
	else if ( a == 2 )
	{
		cout << "a is 2" << endl;
	}
	else
	{
		cout << "a is "
			<< a
			<< endl;
	}
	// a is 100
}

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### `switch`

```cpp
#include <iostream>
using namespace std;

int main()
{
	while (true)
	{
		cout << "Type an integer: ";
		int selected;
		cin >> selected;

		switch (selected)
		{
			case 0:
				{
					cout  << 0
						<< endl;
					break; // break out of switch
				}
			case 1:
				{
					cout  << 1
						<< endl;
					// continue; 
					// without this, it continues on 2
				}
			case 2:
				{
					cout  << 2
						<< endl;
					continue; // continue on while-loop
				}
			case 3:
				{
					cout  << 3
						<< endl;
					continue;
				}
			default:
				{
					cout  << "selected "
						<< selected
						<< endl;
				}
		}
		cout << "Breaking out of loop!" << endl;
		break;
	}
}

/*
Type an integer: 3
3
Type an integer: 2
2
Type an integer: 1
1
2
Type an integer: 0
0
Breaking out of loop!

Type an integer: 3
3
Type an integer: 2
2
Type an integer: 1
1
2
Type an integer: 5
selected 5
Breaking out of loop!
*/

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### `for`

```cpp
#include <iostream>
using namespace std;

int main()
{
	for ( int cnt = 0; cnt < 5; ++cnt )
	{
		cout << cnt
			<< endl;
	}
	// 0
	// 1
	// 2
	// 3
	// 4
}

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### `while`

```cpp
#include <iostream>
using namespace std;

int main()
{
	int cnt = 0;
	while ( cnt < 5 )
	{
		cout << cnt << endl;
		cnt++;
	}
	// 0
	// 1
	// 2
	// 3
	// 4
}

```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>


#### fizzbuzz

> Write a program that prints the numbers from 1 to 100. But for **multiples of
> three print “Fizz”** instead of the number and for the **multiples of**
> **_five_** **print “Buzz”. For numbers which are multiples of both**
> **_three_** and **_five_** **print “FizzBuzz”**.
>
> [**Fizz Buzz Test**](http://c2.com/cgi/wiki?FizzBuzzTest)


```cpp
#include <iostream>
using namespace std;

int main()
{
	for ( int i = 1; i < 101; i++ )
	{
		if ( i%15 == 0 )
		{
			cout << "FizzBuzz" << endl;
		}
		else if ( i%3 == 0 )
		{
			cout << "Fizz" << endl;
		}
		else if ( i%5 == 0 )
		{
			cout << "Buzz" << endl;
		}
		else 
		{
			cout << i << endl;
		}
	}
}

// ...
// Buzz
// 41
// Fizz
// 43
// 44
// FizzBuzz
// 46
// ...
 
```

[↑ top](#c-logic-loop)
<br><br><br><br><hr>
