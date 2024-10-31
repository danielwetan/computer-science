@startuml
title Use Case Diagram for Manajemen Jadwal dan Kehadiran

actor "Admin" as Admin
actor "Pengguna" as User

rectangle "Sistem Manajemen Jadwal dan Kehadiran" {

    usecase "Menetapkan Jadwal Kehadiran" as UC1
    usecase "Mengelola Hari Kerja, Cuti, dan Libur" as UC2

    usecase "Memeriksa Jadwal Kehadiran" as UC3

    Admin -- UC1 : Menetapkan jadwal
    Admin -- UC2 : Mengelola hari kerja

    User -- UC3 : Memeriksa jadwal

    UC1 .> UC2 : <<include>>
}
@enduml
