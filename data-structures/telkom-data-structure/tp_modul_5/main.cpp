#include "menu.h"
#include "list.h"

int main() {
  // int NUM = 0, pilihan = 0, x;
  int pilihan = 0, x;
  address min;
  List L;
  createList(L);

  pilihan = selectMenu();
  while (pilihan != 0) {
    switch(pilihan) {
      case 1:
        cout << "Masukan angka: ";
        cin >> x;
        insertLast(L, allocate(x));
        cout << "Berhasil" << endl;
        break;

      case 2:
        printInfo(L);
        break;

      case 3:
        min = findMin(L);
        cout << min->info; 
        cout << endl;
        break;

      case 4:
        cout << "Masukan angka: ";
        cin >> x;
        insertMiddle(L, x);
        cout << "Berhasil" << endl;
        break;
    }

    cout << endl;
    pilihan = selectMenu();
  }

  cout << "ANDA TELAH KELUAR DARI PROGRAM" << endl;
  return 0;
}