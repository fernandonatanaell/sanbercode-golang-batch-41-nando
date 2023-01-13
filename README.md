# sanbercode-golang-batch-41-nando

1. Pemanggilan route sesuai dengan soal. File route terletak di routers/webRouter.go
2. Kemudian tiap route memanggil controller untuk actionnya ketika dipanggil route tersebut. Letak dari controller ada pada controllers/*.go
3. Untuk eksekusi query ada pada folder repository/*.go
4. Route yang method POST, PUT, dan DELETE wajib diperiksa melalui middleware yang terletak di middleware/checkLogin.go
5. Semua entities/models struct terletak di dalam folder structs/*.go
6. Konfigurasi database terletak di dalam folder database/

*** Bangun Datar ***
1. Aksesnya sama seperti soal tidak ada bedanya.
2. Controller dari nomor ini terletak di dalam controllers/bangunDatarController.go

*** Category ***

0. Bisa CRUD
1. Untuk category route juga sesuai semua dengan soal
2. Untuk method POST category, yang dikirimkan dari JSON cukup kolom nama saja, selebihnya sudah otomatis saya setting
3. Update delete sama seperti soal, cuma pada saat update, kolom updated_at pada field itu juga akan terupdate
4. Kemudian untuk update dan delete juga ada pengecekan, apakah ID yang dimasukkan itu sudah ada atau belum di database

*** Book ***

0. Bisa CRUD
1. Untuk route juga sama, sesuai dengan soal
2. Field yang bisa di ditambahkan atau diubah oleh user juga sesuai dengan yang ada di soal
3. Jika update, kolom updated_at juga terupdate
4. Update dan delete juga ada pengecekan, apakah ID yang dimasukkan itu sudah ada atau belum di database
5. Dapat validasi baik image_url maupun release_year
6. Dapat generate thickness dari total_page

*** Middleware ***
1. Middleware juga dapat berjalan dengan lancar
2. Middleware tidak dibuatkan table, melainkan masih manual di file checkLogin.go
