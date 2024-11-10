#ifndef DLL_H
#define DLL_H

#include <iostream>
#include <string>

using namespace std;

typedef struct elmtList *address;
typedef struct elmtList {
    string judul;
    address prev;
    address next;
} elmtList;
typedef struct list {
    address first;
    address last;
} list;

// Deklarasi fungsi dan prosedur
bool isEmpty(list L);
void createList(list &L);
address createNewElmt(string Judul);
void insertFirst(list &L, address P);
void insertAfter(list &L, address Prec, address P);
void insertLast(list &L, address P);
void deleteFirst(list &L, address &P);
void deleteAfter(list &L, address Prec, address &P);
void deleteLast(list &L, address &P);
void concat(list L1, list L2, list &L3);
address findLagu(string Judul, list L);
void removeLagu(string Judul, list &L);

#endif