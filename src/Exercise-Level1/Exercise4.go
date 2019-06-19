package main

import "fmt"

type jessica int
var x jessica
func main(){
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}