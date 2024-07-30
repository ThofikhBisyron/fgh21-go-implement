package mainproses

import (
	"fazztrack/demo/calc"
	"fmt"
)

func Fazzfood(harga int, voucher string, jarak int, pajak bool) {
	diskon := calc.Calcharga(harga, voucher)
	tarif := calc.Calcjarak(jarak)
	biayapajak := calc.Calcpajak(pajak, harga)

	fmt.Println("harga : ", harga)

	fmt.Println("potongan : ", diskon)

	fmt.Println("biaya antar : ", tarif)

	fmt.Println("Pajak : ", biayapajak)

	var subtotal = harga - diskon + tarif + biayapajak
	fmt.Println("subtotal : ", subtotal)

}
