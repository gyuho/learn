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
