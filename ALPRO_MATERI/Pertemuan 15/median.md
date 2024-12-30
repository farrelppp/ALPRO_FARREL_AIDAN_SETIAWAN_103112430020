program median

kamus
    y, median : integer
    numbers : array[1..9] of integer

algoritma
    input(y)

    for i = 1 to 9 do
        input(numbers[i])
    endfor

    urutkan(numbers)  

    median = numbers[5]  

    output("Median adalah: ", median)
endprogram
