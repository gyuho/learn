#include <iostream>
using namespace std;
 
// function declaration
int max(int num1, int num2);
 
int main ()
{
   // local variable declaration:
   int a = 1;
   int b = 2;
   int ret = max(a, b);
   cout << "Max is : " << ret << endl;
}
 
int max(int num1, int num2) 
{
   if (num1 > num2)
   	   return num1;
   return num2; 
}
