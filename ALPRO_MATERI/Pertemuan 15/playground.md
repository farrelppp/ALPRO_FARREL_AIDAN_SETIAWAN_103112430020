program playground

kamus
    membership : string
    duration, tarif : integer
    tarifnormal, tarifkelebihan : constant integer = 65000, 20000

algoritma
    input(membership)
    input(duration)

    tarif = 0

    if duration > 2 then
        tarif = 2 * tarifnormal
        tarif = tarif + (duration - 2) * tarifkelebihan
    else
        tarif = duration * tarifnormal
    endif

    switch membership do
        case "Gold":
            tarif = tarif / 2
        case "Silver":
            tarif = tarif * 75 / 100
    endswitch

    output("IDR ", tarif)
endprogram
