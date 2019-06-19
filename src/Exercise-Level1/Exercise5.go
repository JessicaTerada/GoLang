package main

import "fmt"

type T1 int
var x1 T1
var y int
func main(){
	y = int(x1)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
