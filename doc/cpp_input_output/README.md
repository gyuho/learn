[*back to contents*](https://github.com/gyuho/learn#contents)<br>

# C++: input, output

- [input, output](#input-output)
- [exist](#exist)
- [write, read](#write-read)

[↑ top](#c-input-output)
<br><br><br><br><hr>


#### input, output

```cpp
#include <iostream>
using namespace std;

int main ()
{
	int i;
	cout << "Please enter an integer value: ";
	cin >> i;
	cout << "The value you entered is " << i;
	cout << " and its double is " << i*2 << ".\n";
}

/*
Please enter an integer value: 111
The value you entered is 111 and its double is 222.
*/

```

[↑ top](#c-input-output)
<br><br><br><br><hr>


#### exist

```cpp
#include <iostream>
using namespace std;
#include <sys/stat.h>

inline bool isExist (const string& name) {
	struct stat buffer;
	return (stat (name.c_str(), &buffer) == 0); 
}

int main()
{
	cout << isExist("./testdata/sample.txt") << endl; // 1
}

```

[↑ top](#c-input-output)
<br><br><br><br><hr>


#### write, read

```cpp
#include <iostream>
#include <fstream>
// #include <stdio.h>
#include <cstdio>
using namespace std;

int main () {

	string fpath = "example.txt";

	ofstream fw;
	fw.open (fpath);
	fw << "Hello World!\n";
	fw << "C++\n";
	fw.close();

	string line;
	ifstream fi (fpath);
	if (fi.is_open())
	{
		while ( getline (fi, line) )
		{
			cout << line << endl;
		}
		fi.close();
	}
	else cout << "Unable to open file" << endl;

	if( remove( fpath.c_str() ) != 0 )
		perror( "Error deleting file" );
	else
		puts( "File successfully deleted" );
}

/*
Hello World!
C++
File successfully deleted
*/

```

[↑ top](#c-input-output)
<br><br><br><br><hr>

