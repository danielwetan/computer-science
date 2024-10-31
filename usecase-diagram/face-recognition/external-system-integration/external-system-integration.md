@startuml
title Use Case Diagram for Integrasi dengan Sistem Lain

actor "Sistem Penggajian" as PayrollSystem
actor "Sistem Penilaian Akademik" as AcademicSystem

rectangle "Sistem Kehadiran" {

    usecase "Integrasi dengan Sistem Penggajian" as UC1
    usecase "Integrasi dengan Sistem Penilaian Akademik" as UC2
    usecase "Menghitung Gaji Berdasarkan Absensi" as UC3
    usecase "Menghitung Nilai Partisipasi Berdasarkan Kehadiran" as UC4

    PayrollSystem -- UC1 : Mengakses data kehadiran
    UC1 .> UC3 : <<include>>

    AcademicSystem -- UC2 : Mengakses data kehadiran
    UC2 .> UC4 : <<include>>
}

@enduml
