program gaji

kamus
nama : string  
gajiPokok, tunjangan, potongan, totalGaji : real  

algoritma  
input(nama)  
input(gajiPokok)  
input(tunjangan)  
input(potongan)  

totalGaji = gajiPokok + tunjangan - potongan  

output("Total Gaji", nama, ":", totalGaji)  
endprogram

