package main

import "fmt"

var y string
type noInt int
var a noInt
func  main() {
	x := 42
	n, _ := fmt.Println("Hello world", x, true)
	fmt.Println(n)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	x=4

	//x = a nao é possivel pq a é tipo noInt, a é static
	x = int(a)

}