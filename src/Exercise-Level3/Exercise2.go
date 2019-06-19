package main

import "fmt"

func main(){
	bd := 1993
	count := -1
	for bd <= 2019{
		bd++
		count++
	}
	fmt.Println(count)
}
