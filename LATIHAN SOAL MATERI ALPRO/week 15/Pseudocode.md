Program AMPM
Deklarasi:
  jam12, jam24: integer
Algoritma:
  Input(jam24)
  Jika jam24 mod 12 = 0 Maka
    jam12 ← 12
  Else
    jam12 ← jam24 mod 12
  EndIf

  Jika jam24 < 12 Maka
    Output(jam12, "AM")
  Else
    Output(jam12, "PM")
  EndIf
Akhir Program