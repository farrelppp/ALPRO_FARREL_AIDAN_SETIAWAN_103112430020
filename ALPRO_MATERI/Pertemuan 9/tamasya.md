program tamasya

kamus
N : integer (jumlah orang)
Kapasitas : integer (kapasitas satu mobil, yaitu 7 orang)
Mobil : integer (jumlah mobil penuh)
Sisa : integer (jumlah orang tersisa setelah mobil penuh)

algoritma
Input(N)
Kapasitas = 7
Mobil = N div Kapasitas
Sisa = N mod Kapasitas
Jika N ≤ Kapasitas maka:
Output("1 mobil dengan", Sisa, "bangku kosong")
Jika tidak, maka:
Output(Mobil, "mobil penuh dan 1 mobil berisi", Sisa, "orang")
EndProgram