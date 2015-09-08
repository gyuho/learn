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
