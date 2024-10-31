@startuml
title Use Case Diagram for Deteksi Wajah dan Pencatatan Kehadiran

actor "Pengguna" as User
actor "Kamera" as Camera

rectangle "Sistem Deteksi Wajah dan Pencatatan Kehadiran" {

    usecase "Deteksi Wajah Pengguna" as UC1
    usecase "Pencocokan Wajah dengan Data" as UC2
    usecase "Pencatatan Waktu Kehadiran" as UC3
    usecase "Pencatatan Waktu Masuk" as UC4
    usecase "Pencatatan Waktu Keluar" as UC5
    
    User -- UC1 : Berada di depan kamera
    Camera -- UC1 : Deteksi wajah
    UC1 .> UC2 : <<include>>
    UC2 .> UC3 : <<include>>
    UC3 .> UC4 : <<include>>
    UC3 .> UC5 : <<include>>
}

@enduml
