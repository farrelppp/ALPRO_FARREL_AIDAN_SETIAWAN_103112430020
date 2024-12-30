program sumganjil

kamus
    a, b, n, sum : integer
algoritma
    input(a, b)   
    a = a + ((a % 2) xor 1) 
    n = (b - a) / 2 + 1
    sum = n * (2 * a + (n - 1) * 2) / 2
    output(sum)
endprogram
