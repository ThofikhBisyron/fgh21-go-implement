package main

import "fmt"

func calcharga(harga int, voucher string) int {
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
	return diskon
}
func calcjarak(jarak int) int {
	tarif := 5000

	if jarak > 2 {
		tarif = tarif + (jarak-2)*3000
	}
	return tarif
}
func calcpajak(pajak bool, harga int) int {
	biayapajak := 0
	if pajak {
		biayapajak = harga * 5 / 100
	}
	return biayapajak
}

func fazzfood(harga int, voucher string, jarak int, pajak bool) {
	diskon := calcharga(harga, voucher)
	tarif := calcjarak(jarak)
	biayapajak := calcpajak(pajak, harga)

	fmt.Println("harga : ", harga)

	fmt.Println("potongan : ", diskon)

	fmt.Println("biaya antar : ", tarif)

	fmt.Println("Pajak : ", biayapajak)

	var subtotal = harga - diskon + tarif + biayapajak
	fmt.Println("subtotal : ", subtotal)

}
func main() {
	fazzfood(75000, "FAZZFOOD50", 5, true)

}
