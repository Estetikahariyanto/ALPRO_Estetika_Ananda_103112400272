program ini untuk menerima input bilangan bulat antara 0 hingga 999 dan misah jadi digit ratusan, puluhan, satuan
cara kerja program :
1. meminta input bilangan bulat dari pengguna.
2. memeriksa bilangan tersebut apakah berada diantara 0 hingga 999.
3. jika syarat nya benar, program akan membagi bilangan tersebut ke dalam 3 digit.
> **digit ratusan (d1)** = mengambil bagian ratusan dari bilangan dengan `x / 100` .
> **digit puluhan (d2)** = mengambil bagianpuluhan dari bilangan dengan `(x / 10) % 10` .
> **digit satuan (d3)** = mengambil bagian satuan dari bilangan dengan `x % 10` .
kalau bilangan diluar rentang 0 hingga 999, maka akan tampil pesan peringatan.