package fazz

import "fmt"

type Survey struct {
	name         string
	age          int
	gender       string
	isSmoker     bool
	cigarVariant []string
}

func main() {
	survey := []Survey{
		Survey{
			name:     "RS",
			age:      25,
			gender:   "Laki-Laki",
			isSmoker: true,
			cigarVariant: []string{
				"Marlboro",
				"Gudang Garam",
				"Esse",
				"Manchester",
			},
		},
		Survey{
			name:     "Ra",
			age:      24,
			gender:   "Perempuan",
			isSmoker: false,
			cigarVariant: []string{
				"Marlboro",
				"Gudang Garam",
				"Esse",
				"Manchester",
			},
		},
	}
	for i := 0; i < len(survey); i++ {
		fmt.Println("Responden", i+1)
		fmt.Println("Name:", survey[i].name)
		fmt.Println("Age:", survey[i].age)
		fmt.Println("Gender:", survey[i].gender)
		fmt.Println("Is Survey Smoker:", survey[i].isSmoker)
		fmt.Println("Cigar Variant:", survey[i].cigarVariant)
	}

}
