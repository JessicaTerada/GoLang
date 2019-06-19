package main

import (
	"fmt"
	"runtime"
)

var y string
type noInt int
var a noInt
var x2 int
func  main() {
	x := 42
	n, _ := fmt.Println("Hello world", x, true)
	fmt.Println(n)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	x=4

	//x = a nao é possivel pq a é tipo noInt, a é static
	x = int(a)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	fmt.Println("Hi")
	foo()
	fmt.Println("after foo")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()

	u := 42
	y := 43

	u,y = y,u
	println(u, y)
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}

//control flor:
// 1- sequence
// 2- loop
// 3- conditional