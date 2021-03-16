#include "stack.h"
#include <iostream>

using namespace std;

int main() {
	Stack s(10);
	try
	{
		while (true) s.push('c');
	}
	catch (Stack::Overflow)
	{
		cout << "Stack Overflow!" << endl;
		return 1;
	}
	return 0;
}