program tangkiair

kamus
    kapasitastank, totalvolume, volumeember : integer

algoritma
    input(kapasitastank)
    totalvolume = 0

    while true do
        input(volumeember)
        totalvolume = totalvolume + volumeember

        if totalvolume >= kapasitastank then
            output(true)
            break
        else
            output(false)
        endif
    endwhile
endprogram
