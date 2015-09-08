[*back to contents*](https://github.com/gyuho/learn#contents)
<br>

# C++: concurrency

- [Reference](#reference)
- [`thread`](#thread)
- [`mutex`](#mutex)
- [`atomic`](#atomic)

[↑ top](#c-concurrency)
<br><br><br><br>
<hr>








#### Reference

- [Concurrency in C++11](https://www.classes.cs.uchicago.edu/archive/2013/spring/12300-1/labs/lab6/)

[↑ top](#c-concurrency)
<br><br><br><br>
<hr>









#### `thread`

```cpp
#include <iostream>
#include <vector>
#include <thread>
using namespace std;

int accum = 0;

void square(int x) {
	accum += x * x;
}

int main() {
	vector<thread> ths;
	for (int i = 1; i <= 20; i++) {
		ths.push_back(thread(&square, i));
	}

	// & to retrieve a reference and not a copy of the object
	// , since join changes the nature of the object.
	// for (auto& th : ths) {
	for (vector<thread>::iterator th = ths.begin(); th != ths.end(); th++)
		(*th).join(); // join() on each thread
		// blocks until the thread finishes
	cout << "accum = " << accum << endl;
}

/*
$ g++ -std=c++11 00_thread.cpp -pthread
accum = 2870


2870 is the correct answer
but:

$ for i in {1..1000}; do ./a.out; done | sort | uniq -c
      1 accum = 2645
      1 accum = 2674
      1 accum = 2749
      1 accum = 2834
      1 accum = 2866
    995 accum = 2870

bunch of race conditions!
We need mutex!
*/

```

[↑ top](#c-concurrency)
<br><br><br><br>
<hr>








#### `mutex`

```cpp
#include <iostream>
#include <vector>
#include <thread>
#include <mutex>
using namespace std;

int accum = 0;
mutex accum_mutex;

void square(int x) {
	int temp = x * x;
	accum_mutex.lock();
	accum += temp;
	accum_mutex.unlock();
}

int main() {
	vector<thread> ths;
	for (int i = 1; i <= 20; i++) {
		ths.push_back(thread(&square, i));
	}

	// & to retrieve a reference and not a copy of the object
	// , since join changes the nature of the object.
	// for (auto& th : ths) {
	for (vector<thread>::iterator th = ths.begin(); th != ths.end(); th++)
		(*th).join(); // join() on each thread
		// blocks until the thread finishes
	cout << "accum = " << accum << endl;
}

/*
$ g++ -std=c++11 01_mutex.cpp -pthread
accum = 2870


2870 is the correct answer
but:

$ for i in {1..1000}; do ./a.out; done | sort | uniq -c
   1000 accum = 2870

no race conditions!
*/

```

[↑ top](#c-concurrency)
<br><br><br><br>
<hr>





#### `atomic`

```cpp
#include <iostream>
#include <vector>
#include <thread>
#include <atomic>
using namespace std;

atomic<int> accum(0);

void square(int x) {
	accum += x * x;
}

int main() {
	vector<thread> ths;
	for (int i = 1; i <= 20; i++) {
		ths.push_back(thread(&square, i));
	}

	// & to retrieve a reference and not a copy of the object
	// , since join changes the nature of the object.
	// for (auto& th : ths) {
	for (vector<thread>::iterator th = ths.begin(); th != ths.end(); th++)
		(*th).join(); // join() on each thread
		// blocks until the thread finishes
	cout << "accum = " << accum << endl;
}

/*
$ g++ -std=c++11 02_atomic.cpp -pthread
accum = 2870


2870 is the correct answer
but:

$ for i in {1..1000}; do ./a.out; done | sort | uniq -c
   1000 accum = 2870

no race conditions!
*/

```

[↑ top](#c-concurrency)
<br><br><br><br>
<hr>
