package calc

func Calcharga(harga int, voucher string) int {
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
func Calcjarak(jarak int) int {
	tarif := 5000

	if jarak > 2 {
		tarif = tarif + (jarak-2)*3000
	}
	return tarif
}
func Calcpajak(pajak bool, harga int) int {
	biayapajak := 0
	if pajak {
		biayapajak = harga * 5 / 100
	}
	return biayapajak
}
