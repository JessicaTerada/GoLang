package main

import "fmt"

func main(){
	y := []int{42,43,44,45,46,47,48,49,50,51}

	fmt.Println(y[:5])
	fmt.Println(y[5:])
	fmt.Println(y[2:7])
	fmt.Println(y[1:])
	fmt.Println(y)

	y = append(y, 52)
	y = append(y, 53, 54, 55)
	fmt.Println(y)

	z := []int{56, 57, 58, 59, 60}
	y = append(y, z...)
	fmt.Println(y)

	y = append(y[:3], y[6:]...)
	fmt.Println(y)
}
