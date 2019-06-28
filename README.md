# Storage Locker

Program locker interaktif sederhana untuk menyimpan kartu identitas seorang tamu ketika mengunjungi sebuah kantor. Sebisa mungkin hanya menggunakan fungsi dan library standar.

Program ini dibuat dan dites menggunakan Go 1.10.8

## Instalasi

Pastikan workspace dan $GOPATH sudah diset.

Download source code dari repository.
```
$ go get github.com/alfarisi/storage-locker
$ cd src/github.com/alfarisi/storage-locker
```

Anda dapat langsung menjalankan program ini dengan perintah :
```
$ go run main.go
```

Atau, dapat juga build source code terlebih dulu, lalu jalankan file binary yang dihasilkan.
```
$ go build
$ ./storage-locker
```

## Daftar Perintah

Berikut adalah daftar perintah yang dapat digunakan. Harap diperhatikan bahwa semua perintah dan argumen bersifat **case sensitive**.

1. `init [n]` : untuk menentukan jumlah locker. Jumlah n locker ditentukan pertama kali ketika program dijalankan dengan menginput angka tertentu. Jika jumlah locker belum ditentukan, program tidak dapat dijalankan. n adalah bilangan positif antara 1 sampai 255.

2. `status` : untuk menampilkan status dari masing-masing nomor locker.

3. `input [tipe_identitas] [nomor_identitas]` : untuk memasukkan dan mencatat kartu identitas. Tipe identitas hanya bisa diisi dengan **KTP**, **SIM** atau **Other**. Ketika memasukkan kartu identitas, otomatis akan dipilih nomor locker paling kecil yang tersedia. Misal yang tersedia 2 dan 4, maka akan terpilih nomor 2. Jika memasukkan lagi, akan terpilih nomor 4.

4. `leave [nomor_locker]` : untuk mengosongkan nomor locker tertentu.

5. `find [nomor_identitas]` : perintah ini akan menampilkan nomor locker berdasar nomor identitas.

6. `search [tipe_identitas]` : akan menampilkan daftar nomor identitas sesuai tipe identitas yang dicari.

7. `help` : untuk menampilkan daftar perintah yang dapat digunakan.

8. `exit` : untuk mengakhiri program.