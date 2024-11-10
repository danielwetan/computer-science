#ifndef STACK_H
#define STACK_H

typedef char infotype;

struct Stack {
    infotype info[15];
    int top;
};

void createStack(Stack &S);
bool isEmpty(Stack S);
bool isFull(Stack S);
void push(Stack &S, infotype x);
infotype pop(Stack &S);
void printStack(Stack S);

#endif