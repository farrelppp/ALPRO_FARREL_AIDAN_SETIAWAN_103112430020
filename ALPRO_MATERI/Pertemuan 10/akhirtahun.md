program akhirtahun

kamus
    totalbelanja, finaltotal, diskon : real
    bersediakartu, memperolehkartu, memperolehdiskon, memperolehcashback : boolean

algoritma
    input(totalbelanja)
    input(bersediakartu)

    finaltotal = totalbelanja

    if bersediakartu then
        memperolehkartu = true
    else
        memperolehkartu = false
    endif

    if totalbelanja >= 100000 then
        memperolehdiskon = true
        diskon = totalbelanja / 10
        finaltotal = finaltotal - diskon

        if totalbelanja >= 200000 then
            if memperolehkartu then
                memperolehcashback = true
                finaltotal = finaltotal - 75000
            else
                memperolehcashback = false
            endif
        else
            memperolehcashback = false
        endif
    else
        memperolehdiskon = false
        memperolehcashback = false
    endif

    output(memperolehkartu)
    output(memperolehdiskon)
    output(memperolehcashback)
    output(finaltotal)
endprogram
