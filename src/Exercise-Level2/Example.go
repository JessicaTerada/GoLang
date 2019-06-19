package main

import "fmt"

const (
	_ = iota
	kb= 1 << (iota * 10)
)

const (
	_ = iota
)

const (
	EST = -(5+iota)
	CST
	MST
	PST
	reset, reset2 = iota, iota

)

func main(){
	fmt.Printf("%d\t%b\n", kb, kb)

	fmt.Println(CST)
	fmt.Println(reset)
	fmt.Println(reset2)
}
