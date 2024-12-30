program modus

kamus
    x, countzero, countx, modus : integer
    numbers : array [1..9] of integer

algoritma
    input(x)

    for i = 1 to 9 do
        input(numbers[i])
    endfor

    countzero = 0
    countx = 0

    for each number in numbers do
        if number = 0 then
            countzero = countzero + 1
        elseif number = x then
            countx = countx + 1
        endif
    endfor

    if countzero > countx then
        modus = 0
    else
        modus = x
    endif

    output("Modus = ", modus)
endprogram
