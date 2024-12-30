program X

kamus
    n, i, j : integer

algoritma
    input(n)

    if n % 2 = 0 then
        output("bilangan harus ganjil")
        return
    endif

    for i = 1 to n do
        for j = 1 to n do
            if i = j or i + j = n + 1 then
                output(i, tanpa newline)
            else
                output(" ", tanpa newline)
            endif
        endfor
        output(newline)
    endfor
endprogram
