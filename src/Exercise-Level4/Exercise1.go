package main

import "fmt"

func main(){
	y := [5]int{1,2,3,4,5}

	for i, v := range  y{
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", y)


	fmt.Println("----slice")
	z := []int{1,2,3,4,5,6,7,8,9,10}
	for i, v := range  z{
		fmt.Println(i, v)
	}

}
