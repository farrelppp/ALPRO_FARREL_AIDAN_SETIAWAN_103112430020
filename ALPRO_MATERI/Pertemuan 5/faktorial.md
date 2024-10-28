program faktorial

kamus
n, faktorial : integer
i : integer

algoritma
input(n)
jika n < 0 maka
    output("Bilangan harus positif!")
    kembali
akhir jika

faktorial = 1
untuk i dari 2 hingga n lakukan
    faktorial = faktorial * i
akhir untuk

output(faktorial)
endprogram