#include "menu.h"

void inputAngka(int &NUM, int x) {
  NUM = x;
}

void tambah(int &NUM, int x) {
  NUM = NUM + x;
}

void printHasil(int NUM) {
  cout << "Nilai saat ini: " << NUM << endl;
}

int selectMenu() {
  cout << "==== MENU ====" << endl;
  cout << "1. Menambah data baru" << endl;
  cout << "2. Menampilkan semua data" << endl;
  cout << "3. Mengembalikan data terkecil" << endl;
  cout << "4. Menambahkan ke tengah list" << endl;
  cout << "0. Exit" << endl;
  cout << "Pilihlah menu: " << endl;

  int input = 0;
  cin >> input;

  return input;
}