#include "stack.h"
#include <iostream>
using namespace std;

int main() {
    Stack myStack;
    createStack(myStack);

    int nim = 1304211006;
    int sisa = nim % 4;

    switch (sisa) {
        case 0:
            push(myStack, 'I'); push(myStack, 'F'); push(myStack, 'L');
            push(myStack, 'A'); push(myStack, 'B'); push(myStack, 'J');
            push(myStack, 'A'); push(myStack, 'Y'); push(myStack, 'A');
            break;
        case 1:
            push(myStack, 'H'); push(myStack, 'A'); push(myStack, 'L');
            push(myStack, 'O'); push(myStack, ' '); push(myStack, 'B');
            push(myStack, 'A'); push(myStack, 'N'); push(myStack, 'D');
            push(myStack, 'U'); push(myStack, 'N'); push(myStack, 'G');
            break;
        case 2:
            push(myStack, 'P'); push(myStack, 'E'); push(myStack, 'R');
            push(myStack, 'C'); push(myStack, 'A'); push(myStack, 'Y');
            push(myStack, 'A'); push(myStack, ' '); push(myStack, 'D');
            push(myStack, 'I'); push(myStack, 'R'); push(myStack, 'I');
            break;
        case 3:
            push(myStack, 'S'); push(myStack, 'T'); push(myStack, 'R');
            push(myStack, 'U'); push(myStack, 'K'); push(myStack, 'T');
            push(myStack, 'U'); push(myStack, 'R'); push(myStack, ' ');
            push(myStack, 'D'); push(myStack, 'A'); push(myStack, 'T');
            push(myStack, 'A');
            break;
        default:
            cout << "NIM tidak valid" << endl;
            return 1;
    }

    cout << "Isi stack awal: ";
    printStack(myStack);
    cout << "Isi stack setelah pop: ";
    while (!isEmpty(myStack)) {
        pop(myStack);
    }
    printStack(myStack);
    cout << endl;
    return 0;
}