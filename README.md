# Storage Locker

Program locker interaktif sederhana untuk menyimpan kartu identitas seorang tamu ketika mengunjungi sebuah kantor. Sebisa mungkin hanya menggunakan fungsi dan library standar.

Program ini dibuat dan dites menggunakan Go 1.10.8

## Instalasi

Pastikan workspace dan $GOPATH sudah diset.
Download source code dari repository.
```
go get github.com/alfarisi/storage-locker
cd src/github.com/alfarisi/storage-locker
```

Anda dapat langsung menjalankan program ini dengan perintah :
```
go run main.go
```

Atau, dapat juga build source code terlebih dulu, lalu jalankan file binary yang dihasilkan.
```
go build
./storage-locker
```

## Daftar Perintah

Berikut adalah daftar perintah yang dapat digunakan. Harap diperhatikan bahwa semua perintah dan argumen bersifat **case sensitive**.

`init [n]`
Untuk menentukan jumlah locker. Jumlah n locker ditentukan pertama kali ketika program dijalankan dengan menginput
angka tertentu. Jika jumlah loker belum ditentukan, program tidak dapat dijalankan. n adalah bilangan positif antara 1 sampai 255.

`status`
Untuk menampilkan status dari masing-masing nomor loker.

`input [tipe identitas] [nomor identitas]`
Untuk memasukkan dan mencatat kartu identitas. Tipe identitas hanya bisa diisi dengan **KTP**, **SIM** atau **Other**. Ketika memasukkan kartu identitas, otomatis akan dipilih nomor locker paling kecil yang tersedia. Misal yang tersedia 2 dan 4, maka akan terpilih nomor 2. Jika memasukkan lagi, akan terpilih nomor 4.

`leave [nomor loker]`
Untuk mengosongkan nomor locker tertentu.

`find [nomor identitas]`
Perintah ini akan menampilkan nomor locker berdasar nomor identitas.

`search [tipe identitas]`
Akan menampilkan daftar nomor identitas sesuai tipe identitas yang dicari.

`help`
Untuk menampilkan daftar perintah yang dapat digunakan.

`exit`
Untuk mengakhiri program