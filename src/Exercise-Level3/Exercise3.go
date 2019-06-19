package main

import "fmt"

func main(){
	//the default may apear anywhere.
	// the first expression that triggers execution folow left-to-right and top to bottom
	//in gothe switch can be a expression, to write if else if else chain as a switch
	x5 := 1
	switch{
	case x5 == 2, x5==3:
		fmt.Println("no")
	case (2 == 4):
		fmt.Println("also no")
	default:
		fmt.Println("default")
	case (x5 == 1):
		fmt.Println("yes")
	case true:
		fmt.Println("YES")
	}
}
