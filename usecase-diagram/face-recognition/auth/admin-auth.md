@startuml
title Use Case Diagram for Admin - Pendaftaran dan Otentikasi Pengguna

actor "Admin" as Admin

rectangle "Sistem Pendaftaran dan Otentikasi" {

    usecase "Mendaftarkan Pengguna Baru" as UC5
    usecase "Mengelola Akun Pengguna" as UC6
    usecase "Membuat Akun Admin" as UC7
    usecase "Autentikasi Multi Faktor" as UC8

    Admin -- UC5 : Mendaftarkan Pengguna Baru
    Admin -- UC6 : Mengelola Akun Pengguna
    Admin -- UC7 : Membuat Akun Admin
    Admin -- UC8 : Verifikasi MFA
    
    UC5 .> UC8 : <<include>>
    UC6 .> UC8 : <<include>>
    UC7 .> UC8 : <<include>>
}

@enduml
