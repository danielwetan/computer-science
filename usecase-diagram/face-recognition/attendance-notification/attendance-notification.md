@startuml
title Use Case Diagram for Notifikasi Otomatis

actor "Pengguna" as User

rectangle "Sistem Notifikasi Otomatis" {

    usecase "Mengirim Notifikasi Kehadiran" as UC1
    usecase "Notifikasi Kehadiran" as UC2
    usecase "Notifikasi Keterlambatan" as UC3
    usecase "Notifikasi Ketidakhadiran" as UC4

    User -- UC1 : Menerima notifikasi
    UC1 .> UC2 : <<include>>
    UC1 .> UC3 : <<include>>
    UC1 .> UC4 : <<include>>
}

@enduml
