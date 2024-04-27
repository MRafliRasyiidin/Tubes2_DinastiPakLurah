# Tugas Besar 2 Strategi Algoritma 

## Daftar Isi 
- [Tugas Besar 2 Strategi Algoritma](#tugas-besar2-straregi-algoritma)
    - [Daftar isi](#daftar-isi)
    - [PakLurahMencari](#pak-lurah-mencari)
    - [Deskripsi Permasalahan](#deskripsi-permasalahan)
    - [Algoritma yang digunakan](#algoritma-yang-digunakan)
    - [Struktur Program](#struktur-program)
    - [Menjalankan Program](#menjalankan-program)

## Pak Lurah Mencari
Made by Muhamad Rafli Rasyiidin, M. Hanief Fatkhan Nashrullah, Indraswara Galih Jayanegara, for STIMA *course*

## Deskripsi Permasalahan 
WikiRace atau Wiki Game adalah permainan online yang menggunakan Wikipedia sebagai sumber informasi. Dalam permainan ini, pemain diminta untuk mencari jalan dari satu artikel Wikipedia ke artikel lain yang telah ditentukan, biasanya dengan aturan mencapai artikel akhir dalam waktu sesingkat mungkin atau dengan menggunakan jumlah tautan (klik) paling sedikit.

## Algoritma yang digunakan
### Backtracking
Backtracking digunakan untuk mencari path dari awal link sampai dengan link tujuan
### IDS
Algoritma ini digunakan untuk mencari link tujuan dengan cara mengunjungi setiap link pada depth n, lalu mencari seluruh link pada depth n secara keseluruhan 
lalu saat seluruh link pada depth n sudah selesai ditelusuri semua maka kita akan menuju ke depth n+1
### BFS
Algoritma ini seperti queue kita mengunjungi link pertama lalu semua link yang ada pada link tersebut akan dimasukkan kedalam sebuah queue lalu, setelah link awal selesai
dikunjungi kita akan mengunjungi link berikutnya yang ada pada antrian

## Struktur Program 

```
| 
| 
|___ backend
|   |-- algorithm.go 
|   |-- backend.exe
|   |-- bfs.go
|   |-- go.mod
|   |-- go.sum
|   |-- ids.go
|   |-- main.go
|   |-- package-lock.json
|   |-- package.json
|___ frontend
|   |-- public
|   |-- src
|   |-- .gitignore
|   |-- package-lock.json
|   |-- package.json
|   |-- tailwind.config.js
|___ .gitignore
|___ README.md
```


# Menjalankan Program 
## install depedencies 
install dependencies
pastikan berada pada folder src pada folder frontend, lalu jalankan perintah 

```
npm install
```
Jalankan perintah
```
npm install pm2@latest -g
```
Jalankan frontend pada folder src di folder frontend 
```
npm run start
```
Jalankan backend pada folder backend 
1. Jalankan perintah
```
go build .
```
2. Jalankan perintah berikut
```
pm2 start 'backend.exe'
```

silahkan dinikmati 


# Anggota Kelompok
|Nama           | NIM 
|---------------|----------------| 
| Muhamad Rafli Rasyiidin | 13522088 |
| M. Hanief Fatkhan Nashrullah | 13522100 |
| Indraswara Galih Jayanegara | 13522119|
