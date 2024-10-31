@startuml
title Use Case Diagram for Integrasi dengan Sistem Lain

actor "Sistem Penggajian" as PayrollSystem

rectangle "Sistem Kehadiran" {

    usecase "Integrasi dengan Sistem Penggajian" as UC1
    usecase "Menghitung Gaji Berdasarkan Absensi" as UC3

    PayrollSystem -- UC1 : Mengakses data kehadiran
    UC1 .> UC3 : <<include>>
}

@enduml
