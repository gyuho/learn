#include <iostream>
#include <queue>
#include <vector>
#include <functional>
using namespace std;

template<class T> using minpq = priority_queue<T, vector<T>, greater<T>>;

struct compare
{
	bool operator()(const int& l, const int& r)  
	{
		return l > r;  
	}
};

int main ()
{
	priority_queue<int> pq;
	pq.push(1000);
	pq.push(10);
	pq.push(20);
	pq.push(20);
	pq.push(17);
	pq.push(55);
	pq.push(15);
	pq.push(100);
	cout << "pq.top() is now " << pq.top() << endl; // pq.top() is now 20
	cout << "Popping out elements: ";
	while (!pq.empty())
	{
		cout << pq.top() << ' ';
		pq.pop();
	}
	cout << endl;
	cout << endl;

	minpq<int> minpq0;
	minpq0.push(1000);
	minpq0.push(10);
	minpq0.push(20);
	minpq0.push(20);
	minpq0.push(17);
	minpq0.push(55);
	minpq0.push(15);
	minpq0.push(100);
	cout << "minpq0.top() is now " << minpq0.top() << endl; // minpq0.top() is now 20
	cout << "Popping out elements: ";
	while (!minpq0.empty())
	{
		cout << minpq0.top() << ' ';
		minpq0.pop();
	}
	cout << endl;
	cout << endl;

	priority_queue<int, vector<int>, compare > minpq1;
	minpq1.push(1000);
	minpq1.push(10);
	minpq1.push(20);
	minpq1.push(20);
	minpq1.push(17);
	minpq1.push(55);
	minpq1.push(15);
	minpq1.push(100);
	cout << "minpq1.top() is now " << minpq1.top() << endl; // minpq1.top() is now 20
	cout << "Popping out elements: ";
	while (!minpq1.empty())
	{
		cout << minpq1.top() << ' ';
		minpq1.pop();
	}
	cout << endl;
	cout << endl;

	priority_queue<int, vector<int>, greater<int>> minpq2;
	minpq2.push(1000);
	minpq2.push(10);
	minpq2.push(20);
	minpq2.push(20);
	minpq2.push(17);
	minpq2.push(55);
	minpq2.push(15);
	minpq2.push(100);
	cout << "minpq2.top() is now " << minpq2.top() << endl; // minpq2.top() is now 20
	cout << "Popping out elements: ";
	while (!minpq2.empty())
	{
		cout << minpq2.top() << ' ';
		minpq2.pop();
	}
	cout << endl;
	cout << endl;

	priority_queue<int, vector<int>, less<int>> maxpq;
	maxpq.push(1000);
	maxpq.push(10);
	maxpq.push(20);
	maxpq.push(20);
	maxpq.push(17);
	maxpq.push(55);
	maxpq.push(15);
	maxpq.push(100);
	cout << "maxpq.top() is now " << maxpq.top() << endl; // maxpq.top() is now 20
	cout << "Popping out elements: ";
	while (!maxpq.empty())
	{
		cout << maxpq.top() << ' ';
		maxpq.pop();
	}
	cout << endl;
}

/*
pq.top() is now 1000
Popping out elements: 1000 100 55 20 20 17 15 10 

minpq0.top() is now 10
Popping out elements: 10 15 17 20 20 55 100 1000 

minpq1.top() is now 10
Popping out elements: 10 15 17 20 20 55 100 1000 

minpq2.top() is now 10
Popping out elements: 10 15 17 20 20 55 100 1000 

maxpq.top() is now 1000
Popping out elements: 1000 100 55 20 20 17 15 10 
*/
