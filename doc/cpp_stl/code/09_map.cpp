#include <iostream>
#include <map>
using namespace std;

void updateMap1(map<char,int> cmap)
{
	cmap['O'] = 100;
}

void updateMap2(map<char,int>& cmap)
{
	cmap['O'] = 100;
}

void updateMap3(map<char,int>* cmap)
{
	// (X) cmap['O'] = 1000;
	(*cmap)['O'] = 1000;
}

int main()
{
	map<char,int> cmap;
	cmap['A'] = 100;
	cmap['B'] = 200;
	cmap['C'] = 300;
	cout << "cmap['X']: " << cmap['X'] << endl; // 0
	updateMap1(cmap); cout << "cmap['O']: " << cmap['O'] << endl; // 0
	updateMap2(cmap); cout << "cmap['O']: " << cmap['O'] << endl; // 100
	updateMap3(&cmap); cout << "cmap['O']: " << cmap['O'] << endl; // 1000

	for (map<char,int>::iterator it=cmap.begin(); it!=cmap.end(); ++it)
		cout << it->first << " => " << it->second << '\n';

	cout << endl;

	map<string,int> smap;
	smap["Hello"] = -100;
	smap["World"] = 200;
	smap["C++"] = 300;
	smap["Hello"] = 100;

	map<string,int>::iterator iter = smap.find("NOT");
	cout << (iter == smap.end()) << endl;
	// 1
	// 'NOT' does not exist in the map

	cout << endl;

	for (map<string,int>::iterator it=smap.begin(); it!=smap.end(); ++it)
		cout << it->first << " => " << it->second << '\n';
	cout << '\n';
	
	iter = smap.find("World");
	if (iter != smap.end())
		smap.erase(iter);
		cout << "Deleted" << endl;

	for (map<string,int>::iterator it=smap.begin(); it!=smap.end(); ++it)
		cout << it->first << " => " << it->second << '\n';
}

/*
A => 100
B => 200
C => 300

1

C++ => 300
Hello => 100
World => 200

Deleted
C++ => 300
Hello => 100
*/
