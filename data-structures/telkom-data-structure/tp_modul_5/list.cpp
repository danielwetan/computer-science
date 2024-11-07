#include <iostream>
#include "list.h"
using namespace std;

void createList(List &L) {
  first(L) = NULL;
}

address allocate(infotype x) {
  address p = new elmlist;
  info(p) = x;
  next(p) = NULL;

  return p;
}

void insertFirst(List &L, address P) {
  next(P) = first(L);
  first(L) = P;
}

void insertLast(List &L, address P) {
  if (first(L) == NULL) {
    insertFirst(L, P);
  } else {
    address last = first(L);
    while (next(last) != NULL) {
      last = next(last);
    }
    next(last) = P;
  }
}

void printInfo(List L) {
  address p = first(L);
  while (p != NULL) {
    cout << info(p) << ", ";
    p = next(p);
  }

  cout << endl;
}

address findMin(List L) {
  if (first(L) == NULL) {
    return NULL;
  }

  address minNode = first(L);
  address p = next(minNode);

  while (p != NULL) {
    if (info(p) < info(minNode)) {
      minNode = p;
    }
    p = next(p);
  }

  return minNode;
}

void insertMiddle(List &L, int value) {
  address newNode = allocate(value);

  if (first(L) == NULL) {
    first(L) = newNode;
    return;
  }

  address slow = first(L);
  address fast = first(L);

  while (fast != NULL && next(fast) != NULL) {
    fast = next(next(fast));
    if (fast != NULL) {
      slow = next(slow);
    }
  }

  next(newNode) = next(slow);
  next(slow) = newNode;
}