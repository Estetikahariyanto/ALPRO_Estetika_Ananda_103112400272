mulai
    deklarasikan variabel N sebagai integer

    tulis "masukkan bilangan bulat positif N: "
    baca N

    jika N <= 0 atau input tidak valid :
        tulis "input tidak valid. harap masukkan bilanganbulat positif."
        selesai

        untuk setiap bilangan d dari 1 sampai N :
            jika N mod d = 0, maka:
                isfactor ← true
            jika tidak:
                isfactor ← false
            tulis d dan isfactor

    selesai..