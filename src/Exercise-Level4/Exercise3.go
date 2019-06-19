package main

import "fmt"

func main() {
	y := make([]string, 10, 10)
	y = []string{"Sao Paulo", "Parana", "Santa Catarina", "Minas Gerais", "Rio de Janeiro", "Bahia",
		"Amazonas", "Recife", "Goias", "Mato Grosso"}

	for i, v := range  y{
		fmt.Println(i, v)
	}

	x1 := []string{"a", "b", "c"}
	x2 := []string{"d", "e", "f"}
	z := [][]string{x1, x2}

	for i , v := range z{
		fmt.Println("Record: ", i)
		for j, w := range v{
			fmt.Println("Index position %v value %v", j, w)
		}
	}

	fmt.Println("-----------------map")
	m := map[string][]string{
		"Matematica" : {"Jessica", "Maria"},
		"Ingles" : {"Joao", "Carla"},
		"Geografia" : {"Lola", "Luis"},
	}

	m["Fisica"] = []string{"Roberto", "Arthur", "Samantha"}
	delete(m, "Geografia")

	for k, v := range m{
		fmt.Println("Record for", k)
		for i, v2 := range v{
			fmt.Println("\t", i, v2)
		}
	}


}
