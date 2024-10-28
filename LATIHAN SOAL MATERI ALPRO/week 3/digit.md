program ini untuk menerima input bilangan bulat 3 digit dan mengecek setiap digitnya berurutan menurun dari kiri ke kanan atau tidak.

1. membaca input bilangan bulat tiga digit dari pengguna.
2. memisahkan bilangan tsb jdi 3 digit :
    > digit pertama ratusan = angka / 100
    > digit kedua puluhan = angka / 10
    > digit ketiga satuan = angka % 10
3. ngecek apakah setiap digit berada dalam urutan menurun dgn logika digit 1 > digit 2 && digit 2 > digit 3
4. mencetak true jika urutan menurun terpenuhi / false.