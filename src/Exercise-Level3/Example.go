package main

import "fmt"

var x2 int
func main(){
	s:= "Hello"
	bs:= []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T", bs)

	for i:=0; i<len(s);i++{
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("")

	for i, v:= range s{
		fmt.Println(i, v)
	}

	for {
		if x2 > 9{
			break
		}
		fmt.Println(x2)
		x2++
	}
	fmt.Println("Done.")

	x3 := 32

	for {
		if x3 == 50{
			break
		}

		fmt.Printf("%v\t%#U\n", x3, x3)
		x3++
	}

	if x4:=42; x4 == 42{
		fmt.Println("two statement in one\n")
	}


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