package main

import "fmt"

type Bio struct {
	fullname string
	skillSet []SkillSet
}

type SkillSet struct {
	name string
}

func main() {

	bio := []Bio{
		{
			fullname: "fazz",
			skillSet: []SkillSet{
				{
					name: "JavaScript",
				},
				{
					name: "CSS",
				},
				{
					name: "PHP",
				},
				{
					name: "HTML",
				},
			},
		},
		{
			fullname: "Rs",
			skillSet: []SkillSet{
				{
					name: "JavaScript",
				},
				{
					name: "CSS",
				},
				{
					name: "PHP",
				},
				{
					name: "HTML",
				},
			},
		},
		{
			fullname: "Yahud",
			skillSet: []SkillSet{
				{
					name: "JavaScript",
				},
				{
					name: "CSS",
				},
				{
					name: "PHP",
				},
				{
					name: "HTML",
				},
			},
		},
		{
			fullname: "lalala",
			skillSet: []SkillSet{
				{
					name: "JavaScript",
				},
				{
					name: "CSS",
				},
				{
					name: "PHP",
				},
				{
					name: "HTML",
				},
			},
		},
	}
	fmt.Println(bio[3].skillSet[2].name)
	fmt.Println(bio[0].skillSet[0].name)
}
