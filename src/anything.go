package main

import "fmt"

func main() {
	fmt.Println("Hi")
	foo()
	fmt.Println("after foo")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	bar()
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
