package main

import "fmt"

func main(){
	x3 := 65

	for {
		if x3 == 90{
			break
		}

		fmt.Printf("%v\t%#U\n", x3, x3)
		x3++
	}
}