Jawablah pertanyaan-pertanyaan berikut: 

A.  Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah 
eksekusi program tersebut sesuai spesifikasi soal? 

Jawab :
jika nambernilai 80.1, program akan memeriksa nilai namsesuai dengan kondisi if-else. Berdasarkan kondisi dalam kode 80.1 tidak memenuhi nam >= 88, jadi tidak akan mendapatkan nilai A.
80.1 memenuhi nam >= 72.5, sehingga program ditetapkan nmk = "AB".
Jadi, keluaran dari program adalah AB.


B. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur 
program seharusnya! 

Jawab :
Kode menggunakan kondisi if-else bertingkat tanpa batasan yang tepat untuk nilai nam. Misalnya, kondisi if nam >= 72.5akan berlaku untuk semua nilai namyang lebih besar dari 72.5, bahkan jika nilai tersebut lebih kecil dari batas Byang sebenarnya (88 > nam> 72.5).
Seharusnya, setiap kondisi mencakup jarak nilai yang lebih tepat sesuai tabel. Misalnya, Buntuk 65 <= nam < 72.5dan seterusnya.
Program Alur Seharusnya : Alur yang benar adalah menggunakan rentang nilai untuk setiap kondisi penilaian, seperti:

88 <= nam <= 100untuk nilaiA
72.5 <= nam < 88untuk nilaiAB
65 <= nam < 72.5untuk nilaiB
dan seterusnya sesuai tabel.


C. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. 
Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’. 


package main

import (
	"fmt"
)

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Tentukan NMK berdasarkan rentang nilai yang sesuai
	if nam >= 88 && nam <= 100 {
		nmk = "A"
	} else if nam >= 72.5 && nam < 88 {
		nmk = "AB"
	} else if nam >= 65 && nam < 72.5 {
		nmk = "B"
	} else if nam >= 57.5 && nam < 65 {
		nmk = "BC"
	} else if nam >= 50 && nam < 57.5 {
		nmk = "C"
	} else if nam >= 40 && nam < 50 {
		nmk = "D"
	} else {
		nmk = "E"
	}

	fmt.Println("Nilai mata kuliah:", nmk)
}