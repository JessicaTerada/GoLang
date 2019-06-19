package main

import "fmt"

type person struct{
	first string
	last string
}

type school struct{
	person
	old bool
	first string
}

func main(){
	sc1 := school{
		person: person{
			first: "Carlos",
			last: "Silva",
		},
		old: true,
		first: "yes",
	}

	p1:= person{
		first: "Jessica",
		last: "Terada",
	}
	p2:= person{
		first: "Natalia",
		last: "Terada",
	}

	fmt.Println(p1.first)
	fmt.Println(p2.first)
	fmt.Println(sc1.person.first)
}
