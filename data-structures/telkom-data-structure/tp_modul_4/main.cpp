#include <iostream>
#include "list.h"
using namespace std;

int main() {
  List L;

  // 1. Panggilah create list
  createList(L);

  // 2. Buat sintak menanyakan angka pertama yang ingin diinputkan user ke list
  int x1;
  cout << "Masukan angka pertama yang ingin dimasukan ke list: ";
  cin >> x1;

  // 3. Panggil fungsi allocate agar data tersebut dijadikan elemen
  address P1 = allocate(x1);

  // 4. Panggil procedure insert first yang telah dibuat
  insertFirst(L, P1);

  // 5. Panggil procedure show info untuk mengecek apakah angka tersebut berhasil menjadi elemen di list
  cout << "Elemen list saat ini: ";
  printInfo(L);
  cout << endl;

  // 6. Buat kembali sintak no 2 s/d no 5 untuk data angka kedua dari user
  int x2;
  cout << "Masukan angka kedua yang ingin dimasukan ke list: ";
  cin >> x2;
  address P2 = allocate(x2);
  insertFirst(L, P2);
  cout << "Elemen list saat ini: ";
  printInfo(L);
  cout << endl;

  // 7. Buat kemabil sintak no 2 s/d no 5 untuk data ketiga dari user
  int x3;
  cout << "Masukan angka ketiga yang ingin dimasukan ke list: ";
  cin >> x3;
  address P3 = allocate(x3);
  insertFirst(L, P3);
  cout << "Elemen list saat ini: ";
  printInfo(L);
  cout << endl;

  return 0;
}