#include "dll.h"

bool isEmpty(list L) {
    return L.first == NULL;
}

void createList(list &L) {
    L.first = NULL;
    L.last = NULL;
}

address createNewElmt(string Judul) {
    address P = new elmtList;
    if (P != NULL) {
        P->judul = Judul;
        P->prev = NULL;
        P->next = NULL;
    }
    return P;
}

void insertFirst(list &L, address P) {
    if (isEmpty(L)) {
        L.first = P;
        L.last = P;
    } else {
        P->next = L.first;
        L.first->prev = P;
        L.first = P;
    }
}

void insertAfter(list &L, address Prec, address P) {
    if (Prec != NULL) {
        P->next = Prec->next;
        if (Prec->next != NULL) {
            Prec->next->prev = P;
        } else {
            L.last = P;
        }
        Prec->next = P;
        P->prev = Prec;
    }
}

void insertLast(list &L, address P) {
    if (isEmpty(L)) {
        insertFirst(L, P);
    } else {
        P->prev = L.last;
        L.last->next = P;
        L.last = P;
    }
}

void deleteFirst(list &L, address &P) {
    if (!isEmpty(L)) {
        P = L.first;
        if (L.first == L.last) {
            L.first = NULL;
            L.last = NULL;
        } else {
            L.first = P->next;
            L.first->prev = NULL;
        }
        delete P;
    }
}

void deleteAfter(list &L, address Prec, address &P) {
    if (Prec != NULL && Prec->next != NULL) {
        P = Prec->next;
        if (P == L.last) {
            L.last = Prec;
        } else {
            P->next->prev = Prec;
        }
        Prec->next = P->next;
        delete P;
    }
}

void deleteLast(list &L, address &P) {
    if (!isEmpty(L)) {
        P = L.last;
        if (L.first == L.last) {
            L.first = NULL;
            L.last = NULL;
        } else {
            L.last = P->prev;
            L.last->next = NULL;
        }
        delete P;
    }
}

void concat(list L1, list L2, list &L3) {
    createList(L3);
    if (!isEmpty(L1)) {
        L3.first = L1.first;
        L3.last = L1.last;
    }
    if (!isEmpty(L2)) {
        if (isEmpty(L3)) {
            L3.first = L2.first;
            L3.last = L2.last;
        } else {
            L3.last->next = L2.first;
            L2.first->prev = L3.last;
            L3.last = L2.last;
        }
    }
}

address findLagu(string Judul, list L) {
    address P = L.first;
    while (P != NULL) {
        if (P->judul == Judul) {
            return P;
        }
        P = P->next;
    }
    return NULL;
}

void removeLagu(string Judul, list &L) {
    address P = findLagu(Judul, L);
    if (P != NULL) {
        if (P == L.first) {
            deleteFirst(L, P);
        } else if (P == L.last) {
            deleteLast(L, P);
        } else {
            deleteAfter(L, P->prev, P);
        }
    }
}