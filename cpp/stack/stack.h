#pragma once
class Stack {
	char* v;
	int top;
	int max_size;

public:
	class Overflow {};
	class Underflow {};
	class BadSize {};

	Stack(int s);
	~Stack();

	void push(char);
	char pop();
};