@startuml
title Use Case Diagram for Laporan Kehadiran

actor "Admin" as Admin

rectangle "Sistem Laporan Kehadiran" {

    usecase "Menghasilkan Laporan Kehadiran" as UC1
    usecase "Memilih Periode Laporan" as UC2
    usecase "Filter Laporan" as UC3
    usecase "Unduh Laporan (Excel/PDF)" as UC4

    Admin -- UC1 : Menghasilkan Laporan
    UC1 .> UC2 : <<include>>
    UC1 .> UC3 : <<include>>
    UC1 .> UC4 : <<include>>
}

@enduml
