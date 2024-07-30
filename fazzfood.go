package main

import "fmt"

func fazzfood(harga int, voucher string, jarak int, pajak bool) {
	fmt.Println("harga : ", harga)
	diskon := 0
	if voucher == "FAZZFOOD50" {
		if harga > 50000 {
			diskon = harga * 50 / 100
			if diskon > 50000 {
				diskon = 50000
			}
		}
	}
	if voucher == "DITRAKTIR60" {
		if harga > 25000 {
			diskon = harga * 60 / 100
			if diskon > 25000 {
				diskon = 30000
			}
		}
	}

	fmt.Println("potongan : ", diskon)
	tarif := 5000

	if jarak > 2 {
		tarif = tarif + (jarak-2)*3000
	}
	fmt.Println("biaya antar : ", tarif)

	biayapajak := 0
	if pajak {
		biayapajak = harga * 5 / 100
	}
	fmt.Println("Pajak : ", biayapajak)

	var subtotal = harga - diskon + tarif + biayapajak
	fmt.Println("subtotal : ", subtotal)

}
func main() {
	fazzfood(75000, "FAZZFOOD50", 5, true)

}
