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