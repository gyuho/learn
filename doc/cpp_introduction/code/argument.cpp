#include <iostream>
using namespace std;

int main(int argc, char* argv[])
{
    // Check the number of parameters
    if (argc < 2) {
        cerr << "Usage: " << argv[0] << " NAME" << endl;
        return 1;
    }
    // Print the user's name:
    cout << argv[0] << "says hello, " << argv[1] << "!" << endl;
}

/*
$ ./a.out  Gyu-Ho

Then

./a.outsays hello, Gyu-Ho!
*/