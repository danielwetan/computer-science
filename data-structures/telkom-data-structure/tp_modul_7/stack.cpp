#include "stack.h"
#include <iostream>

using namespace std;

void createStack(Stack &S) {
    S.top = -1;
}

bool isEmpty(Stack S) {
    return S.top == -1;
}

bool isFull(Stack S) {
    return S.top == 14;
}

void push(Stack &S, infotype x) {
    if (isFull(S)) {
        cout << "Stack overflow" << endl;
    } else {
        S.top++;
        S.info[S.top] = x;
    }
}

infotype pop(Stack &S) {
    if (isEmpty(S)) {
        cout << "Stack underflow" << endl;
        return '\0';
    } else {
        infotype x = S.info[S.top];
        S.top--;
        return x;
    }
}

void printStack(Stack S) {
    if (isEmpty(S)) {
        cout << "Stack is empty" << endl;
    } else {
        for (int i = S.top; i >= 0; i--) {
            cout << S.info[i] << " ";
        }
        cout << endl;
    }
}