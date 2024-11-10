#include "dll.h"

int main() {
    list L1, L2, L3;

    createList(L1);
    createList(L2);
    createList(L3);

    address P;
    P = createNewElmt("Lagu 1");
    insertFirst(L1, P);
    P = createNewElmt("Lagu 2");
    insertLast(L1, P);
    P = createNewElmt("Lagu 3");
    insertAfter(L1, L1.first, P);

    P = createNewElmt("Lagu 4");
    insertFirst(L2, P);
    P = createNewElmt("Lagu 5");
    insertLast(L2, P);

    concat(L1, L2, L3);

    address current = L3.first;
    while (current != NULL) {
        cout << current->judul << " - ";
        current = current->next;
    }
    cout << endl;

    string judul = "Lagu 2";
    P = findLagu(judul, L3);
    if (P != NULL) {
        removeLagu(judul, L3);
        cout << judul << " telah dihapus." << endl;
    } else {
        cout << judul << " tidak ditemukan." << endl;
    }

    current = L3.first;
    while (current != NULL) {
        cout << current->judul << " - ";
        current = current->next;
    }
    cout << endl;

    return 0;
}