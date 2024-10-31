@startuml
title Use Case Diagram for Karyawan/Mahasiswa - Pendaftaran dan Otentikasi Pengguna

actor "Karyawan/Mahasiswa" as Employee

rectangle "Sistem Pendaftaran dan Otentikasi" {

    usecase "Mendaftar Akun" as UC1
    usecase "Login ke Sistem" as UC2
    usecase "Autentikasi Multi Faktor" as UC3
    usecase "Mengatur Ulang Kata Sandi" as UC4

    Employee -- UC1 : Mendaftar
    Employee -- UC2 : Login
    Employee -- UC3 : Verifikasi MFA (opsional)
    Employee -- UC4 : Pemulihan kata sandi
    
    UC2 .> UC3 : <<include>>
    UC4 .> UC1 : <<include>>
}

@enduml
