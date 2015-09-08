#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

void erase(vector<string>& values, int pos);
void print(vector<string> values);

int main()
{
	// declare the vector of integers
	// vector<int> iv;
	typedef vector<int> int_vec_t;

	int_vec_t iv;
	// iv.push_back(1);
	iv.push_back(2);
	iv.push_back(3);
	iv.push_back(4);
	iv.push_back(5);
	// push front
	iv.insert(iv.begin(), 1);

	iv.erase(iv.begin()+2);
	iv.erase(iv.begin(), iv.begin()+2);

	for (int_vec_t::iterator it = iv.begin(); it != iv.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
	// 4 5 

	for (auto& it : iv)
		cout << ' ' << it;
	cout << endl;
	// 4 5 
	
	vector<string> strVector;
	strVector.push_back("A");
	strVector.push_back("B");
	strVector.push_back("C");

	cout << strVector.size() << endl; // 3

	while (!strVector.empty())
	{
		strVector.pop_back();
	}

	cout << strVector.size() << endl; // 0

	cout << endl;

	vector<string> members(5);
	members[0] = "A";
	members[1] = "B";
	members[2] = "C";
	members[3] = "D";
	members[4] = "E";
	print(members);
	
	int pos;
	cout << "Remove which element? ";
	cin >> pos;
	
	erase(members, pos);
	print(members);

	int houses[] = {3, 15, 13, 4, 7};
	vector<int> hv;
	size_t size = sizeof(houses) / sizeof(houses[0]);
	for (int i=0; i<size; ++i)
	{
		hv.push_back(houses[i]);
	}
	vector<int> example(hv.begin()+1, hv.end());
	for (vector<int>::iterator it = example.begin(); it != example.end(); ++it)
		cout << ' ' << *it;
	cout << endl;
	//  15 13 4 7

	cout << "The smallest element is " << *min_element(houses, houses+5) << '\n';
	cout << "The largest element is "  << *max_element(houses, houses+5) << '\n';
	// The smallest element is 3
	// The largest element is 15

	cout << "The smallest element is " << *min_element(hv.begin(), hv.end()) << '\n';
	cout << "The largest element is "  << *max_element(hv.begin(), hv.end()) << '\n';
	// The smallest element is 3
	// The largest element is 15
}

/**
 Removes an element from an unordered vector.
 @param values a vector
 @param pos the position of the element to erase
 */
void erase(vector<string>& values, int pos)
{
	int last_pos = values.size() - 1;
	values[pos] = values[last_pos];
	values.pop_back();
}

/**
 Prints all elements in a vector.
 @param values the vector to print
 */
void print(vector<string> values)
{
	for (int i = 0; i < values.size(); i++)
		cout << "[" << i << "] " << values[i] << "\n";
}

/*
[0] A
[1] B
[2] C
[3] D
[4] E
Remove which element? 3
[0] A
[1] B
[2] C
[3] E
*/