#include <iostream>
#include <string>
using namespace std;

// Node for Parent List (Penumpang)
struct Penumpang {
    int id;
    string nama;
    int umur;
    string email;
    string jenis_kelamin;
    Penumpang* next;

    Penumpang(int id, string nama, int umur, string email, string jenis_kelamin)
        : id(id), nama(nama), umur(umur), email(email), jenis_kelamin(jenis_kelamin), next(nullptr) {}
};

// Node for Child List (Penerbangan)
struct Penerbangan {
    string code;
    string asal;
    string tujuan;
    string tanggal;
    string waktu;
    int kapasitas;
    Penerbangan* next;

    Penerbangan(string code, string asal, string tujuan, string tanggal, string waktu, int kapasitas)
        : code(code), asal(asal), tujuan(tujuan), tanggal(tanggal), waktu(waktu), kapasitas(kapasitas), next(nullptr) {}
};

// Node for Relation List (Booking)
struct Booking {
    string code;
    int id_penumpang;
    string id_penerbangan;
    int harga;
    string tanggal;
    Booking* next;
    Booking* prev;

    Booking(string code, int id_penumpang, string id_penerbangan, int harga, string tanggal)
        : code(code), id_penumpang(id_penumpang), id_penerbangan(id_penerbangan), harga(harga), tanggal(tanggal), next(nullptr), prev(nullptr) {}
};

class MultiLinkedList {
private:
    Penumpang* headPenumpang;
    Penerbangan* headPenerbangan;
    Booking* headBooking;

public:
    MultiLinkedList() : headPenumpang(nullptr), headPenerbangan(nullptr), headBooking(nullptr) {}

    // a. Insert element parent
    void insertPenumpang(int id, string nama, int umur, string email, string jenis_kelamin) {
        Penumpang* newPenumpang = new Penumpang(id, nama, umur, email, jenis_kelamin);
        newPenumpang->next = headPenumpang;
        headPenumpang = newPenumpang;
    }

    // b. Insert element child
    void insertPenerbangan(string code, string asal, string tujuan, string tanggal, string waktu, int kapasitas) {
        Penerbangan* newPenerbangan = new Penerbangan(code, asal, tujuan, tanggal, waktu, kapasitas);
        newPenerbangan->next = headPenerbangan;
        headPenerbangan = newPenerbangan;
    }

    // c. Insert element relation
    void insertBooking(string code, int id_penumpang, string id_penerbangan, int harga, string tanggal) {
        Booking* newBooking = new Booking(code, id_penumpang, id_penerbangan, harga, tanggal);
        newBooking->next = headBooking;
        if (headBooking != nullptr) {
            headBooking->prev = newBooking;
        }
        headBooking = newBooking;
    }

    // d. Delete element parent
    void deletePenumpang(int id) {
        Penumpang* curr = headPenumpang;
        Penumpang* prev = nullptr;

        while (curr != nullptr && curr->id != id) {
            prev = curr;
            curr = curr->next;
        }

        if (curr == nullptr) {
            cout << "Penumpang not found!\n";
            return;
        }

        if (prev == nullptr) {
            headPenumpang = curr->next;
        } else {
            prev->next = curr->next;
        }

        // Delete related bookings
        Booking* bcurr = headBooking;
        while (bcurr != nullptr) {
            if (bcurr->id_penumpang == id) {
                deleteBooking(bcurr->code);
            }
            bcurr = bcurr->next;
        }

        delete curr;
        cout << "Penumpang deleted successfully!\n";
    }

    // e. Delete element child
    void deletePenerbangan(string code) {
        Penerbangan* curr = headPenerbangan;
        Penerbangan* prev = nullptr;

        while (curr != nullptr && curr->code != code) {
            prev = curr;
            curr = curr->next;
        }

        if (curr == nullptr) {
            cout << "Penerbangan not found!\n";
            return;
        }

        if (prev == nullptr) {
            headPenerbangan = curr->next;
        } else {
            prev->next = curr->next;
        }

        // Delete related bookings
        Booking* bcurr = headBooking;
        while (bcurr != nullptr) {
            if (bcurr->id_penerbangan == code) {
                deleteBooking(bcurr->code);
            }
            bcurr = bcurr->next;
        }

        delete curr;
        cout << "Penerbangan deleted successfully!\n";
    }

    // f. Delete element relation
    void deleteBooking(string code) {
        Booking* curr = headBooking;

        while (curr != nullptr && curr->code != code) {
            curr = curr->next;
        }

        if (curr == nullptr) {
            cout << "Booking not found!\n";
            return;
        }

        if (curr->prev != nullptr) {
            curr->prev->next = curr->next;
        } else {
            headBooking = curr->next;
        }

        if (curr->next != nullptr) {
            curr->next->prev = curr->prev;
        }

        delete curr;
        cout << "Booking deleted successfully!\n";
    }

    // g. Find element Parent
    Penumpang* findPenumpang(int id) {
        Penumpang* curr = headPenumpang;
        while (curr != nullptr) {
            if (curr->id == id)
                return curr;
            curr = curr->next;
        }
        return nullptr;
    }

    // h. Find element child
    Penerbangan* findPenerbangan(string code) {
        Penerbangan* curr = headPenerbangan;
        while (curr != nullptr) {
            if (curr->code == code)
                return curr;
            curr = curr->next;
        }
        return nullptr;
    }

    // i. Find apakah parent dan child tertentu memiliki relasi
    bool isRelated(int id, string code) {
        Booking* curr = headBooking;
        while (curr != nullptr) {
            if (curr->id_penumpang == id && curr->id_penerbangan == code)
                return true;
            curr = curr->next;
        }
        return false;
    }

    // j. Show all data di List Parent
    void displayPenumpang() {
        Penumpang* curr = headPenumpang;
        while (curr != nullptr) {
            cout << "ID: " << curr->id << ", Name: " << curr->nama << ", Age: " << curr->umur
                 << ", Email: " << curr->email << ", Gender: " << curr->jenis_kelamin << endl;
            curr = curr->next;
        }
    }

    // k. Show all data di List Child
    void displayPenerbangan() {
        Penerbangan* curr = headPenerbangan;
        while (curr != nullptr) {
            cout << "Code: " << curr->code << ", Origin: " << curr->asal << ", Destination: " << curr->tujuan
                 << ", Date: " << curr->tanggal << ", Time: " << curr->waktu << ", Capacity: " << curr->kapasitas
                 << endl;
            curr = curr->next;
        }
    }

    // l. Show data child dari parent tertentu
    void displayChildFromParent(int id) {
        Booking* curr = headBooking;
        while (curr != nullptr) {
            if (curr->id_penumpang == id) {
                Penerbangan* penerbangan = findPenerbangan(curr->id_penerbangan);
                if (penerbangan != nullptr) {
                    cout << "Code: " << penerbangan->code << ", Origin: " << penerbangan->asal << ", Destination: " << penerbangan->tujuan
                         << ", Date: " << penerbangan->tanggal << ", Time: " << penerbangan->waktu << endl;
                }
            }
            curr = curr->next;
        }
    }

    // m. Show data parent dari child tertentu
    void displayParentFromChild(string code) {
        Booking* curr = headBooking;
        while (curr != nullptr) {
            if (curr->id_penerbangan == code) {
                Penumpang* penumpang = findPenumpang(curr->id_penumpang);
                if (penumpang != nullptr) {
                    cout << "ID: " << penumpang->id << ", Name: " << penumpang->nama << ", Age: " << penumpang->umur
                         << ", Email: " << penumpang->email << ", Gender: " << penumpang->jenis_kelamin << endl;
                }
            }
            curr = curr->next;
        }
    }

    // n. Show setiap data parent beserta data child yang berelasi dengannya
    void displayParentsWithChildren() {
        Penumpang* pCurr = headPenumpang;
        while (pCurr != nullptr) {
            cout << "Penumpang: " << pCurr->nama << endl;
            displayChildFromParent(pCurr->id);
            pCurr = pCurr->next;
        }
    }

    // o. Show setiap data child beserta data parent yang berelasi dengannya
    void displayChildrenWithParents() {
        Penerbangan* cCurr = headPenerbangan;
        while (cCurr != nullptr) {
            cout << "Penerbangan: " << cCurr->code << endl;
            displayParentFromChild(cCurr->code);
            cCurr = cCurr->next;
        }
    }

    // p. Count jumlah child element parent tertentu
    int countChildrenFromParent(int id) {
        int count = 0;
        Booking* curr = headBooking;
        while (curr != nullptr) {
            if (curr->id_penumpang == id) {
                count++;
            }
            curr = curr->next;
        }
        return count;
    }

    // q. Count jumlah parent yang dimiliki oleh child tertentu
    int countParentsFromChild(string code) {
        int count = 0;
        Booking* curr = headBooking;
        while (curr != nullptr) {
            if (curr->id_penerbangan == code) {
                count++;
            }
            curr = curr->next;
        }
        return count;
    }

    // r. Count element child yang tidak memiliki parent
    int countOrphanChildren() {
        int count = 0;
        Penerbangan* curr = headPenerbangan;
        while (curr != nullptr) {
            if (countParentsFromChild(curr->code) == 0) {
                count++;
            }
            curr = curr->next;
        }
        return count;
    }

    // s. Count element parent yang tidak memiliki child
    int countOrphanParents() {
        int count = 0;
        Penumpang* curr = headPenumpang;
        while (curr != nullptr) {
            if (countChildrenFromParent(curr->id) == 0) {
                count++;
            }
            curr = curr->next;
        }
        return count;
    }

    // t. Edit relasi /mengganti child dari parent tertentu dan mengganti parent dari child tertentu
    void editRelation(string bookingCode, int newParentID, string newChildCode) {
        Booking* curr = headBooking;
        while (curr != nullptr) {
            if (curr->code == bookingCode) {
                curr->id_penumpang = newParentID;
                curr->id_penerbangan = newChildCode;
                cout << "Relation updated successfully!\n";
                return;
            }
            curr = curr->next;
        }
        cout << "Booking not found!\n";
    }
};

int main() {
    MultiLinkedList ml;

    // a. Insert element parent
    cout << "a. Insert element parent" << endl;
    ml.insertPenumpang(1, "Alice", 30, "alice@example.com", "Female");
    ml.insertPenumpang(2, "Bob", 25, "bob@example.com", "Male");
    ml.displayPenumpang();
    cout << "=======" << endl << endl;

    // b. Insert element child
    cout << "b. Insert element child" << endl;
    ml.insertPenerbangan("FL123", "NY", "LA", "2024-12-20", "10:00", 200);
    ml.insertPenerbangan("FL456", "SF", "TX", "2024-12-22", "15:00", 150);
    ml.displayPenerbangan();
    cout << "=======" << endl << endl;

    // c. Insert element relation
    cout << "c. Insert element relation" << endl;
    ml.insertBooking("BK001", 1, "FL123", 500, "2024-12-15");
    ml.insertBooking("BK002", 2, "FL456", 300, "2024-12-16");
    ml.displayParentsWithChildren();
    cout << "=======" << endl << endl;

    // d. Delete element parent
    cout << "d. Delete element parent" << endl;
    ml.deletePenumpang(2);
    ml.displayPenumpang();
    ml.displayParentsWithChildren();
    cout << "=======" << endl << endl;

    // e. Delete element child
    cout << "e. Delete element child" << endl;
    ml.deletePenerbangan("FL123");
    ml.displayPenerbangan();
    ml.displayChildrenWithParents();
    cout << "=======" << endl << endl;

    // f. Delete element relation
    cout << "f. Delete element relation" << endl;
    ml.deleteBooking("BK002");
    ml.displayParentsWithChildren();
    ml.displayChildrenWithParents();
    cout << "=======" << endl << endl;

    // g. Find element Parent
    cout << "g. Find element Parent" << endl;
    Penumpang* foundParent = ml.findPenumpang(1);
    if (foundParent) {
        cout << "Found: ID: " << foundParent->id << ", Name: " << foundParent->nama << endl;
    } else {
        cout << "Parent not found!" << endl;
    }
    cout << "=======" << endl << endl;

    // h. Find element child
    cout << "h. Find element child" << endl;
    Penerbangan* foundChild = ml.findPenerbangan("FL456");
    if (foundChild) {
        cout << "Found: Code: " << foundChild->code << ", Origin: " << foundChild->asal << ", Destination: " << foundChild->tujuan << endl;
    } else {
        cout << "Child not found!" << endl;
    }
    cout << "=======" << endl << endl;

    // i. Find apakah parent dan child tertentu memiliki relasi
    cout << "i. Find apakah parent dan child tertentu memiliki relasi" << endl;
    if (ml.isRelated(1, "FL456")) {
        cout << "Parent and child are related!" << endl;
    } else {
        cout << "No relation found!" << endl;
    }
    cout << "=======" << endl << endl;

    // j. Show all data di List Parent
    cout << "j. Show all data di List Parent" << endl;
    ml.displayPenumpang();
    cout << "=======" << endl << endl;

    // k. Show all data di List Child
    cout << "k. Show all data di List Child" << endl;
    ml.displayPenerbangan();
    cout << "=======" << endl << endl;

    // l. Show data child dari parent tertentu
    cout << "l. Show data child dari parent tertentu" << endl;
    ml.displayChildFromParent(1);
    cout << "=======" << endl << endl;

    // m. Show data parent dari child tertentu
    cout << "m. Show data parent dari child tertentu" << endl;
    ml.displayParentFromChild("FL456");
    cout << "=======" << endl << endl;

    // n. Show setiap data parent beserta data child yang berelasi dengannya
    cout << "n. Show setiap data parent beserta data child yang berelasi dengannya" << endl;
    ml.displayParentsWithChildren();
    cout << "=======" << endl << endl;

    // o. Show setiap data child beserta data parent yang berelasi dengannya
    cout << "o. Show setiap data child beserta data parent yang berelasi dengannya" << endl;
    ml.displayChildrenWithParents();
    cout << "=======" << endl << endl;

    // p. Count jumlah child element parent tertentu
    cout << "p. Count jumlah child element parent tertentu" << endl;
    cout << "Total children for parent ID 1: " << ml.countChildrenFromParent(1) << endl;
    cout << "=======" << endl << endl;

    // q. Count jumlah parent yang dimiliki oleh child tertentu
    cout << "q. Count jumlah parent yang dimiliki oleh child tertentu" << endl;
    cout << "Total parents for flight FL456: " << ml.countParentsFromChild("FL456") << endl;
    cout << "=======" << endl << endl;

    // r. Count element child yang tidak memiliki parent
    cout << "r. Count element child yang tidak memiliki parent" << endl;
    cout << "Total orphan children: " << ml.countOrphanChildren() << endl;
    cout << "=======" << endl << endl;

    // s. Count element parent yang tidak memiliki child
    cout << "s. Count element parent yang tidak memiliki child" << endl;
    cout << "Total orphan parents: " << ml.countOrphanParents() << endl;
    cout << "=======" << endl << endl;

    // t. Edit relasi / mengganti child dari parent tertentu dan mengganti parent dari child tertentu
    cout << "t. Edit relasi" << endl;
    ml.insertBooking("BK003", 1, "FL456", 600, "2024-12-17");
    ml.editRelation("BK003", 2, "FL123");
    ml.displayParentsWithChildren();
    cout << "=======" << endl << endl;

    return 0;
}
