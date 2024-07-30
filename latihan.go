package main

import "fmt"

type Bio struct {
	fullName string
	age      int
	hobbies  []string
}

func main() {
	// var hobbies []string = []string{
	// 	"programing",
	// 	"musikan",
	// 	"melawan",
	// 	"membantai",
	// }
	// hobbies[0] = "test"
	// fmt.Println(hobbies)
	var bio Bio = Bio{
		fullName: "FazzTrack",
		age:      25,
		hobbies: []string{
			"Programing",
			"Hikking",
		},
	}
	fmt.Println(bio)
}
