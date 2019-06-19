package main

import "fmt"

func main(){
	fmt.Println("----------array")
	//Unlike in C/C++ (where arrays act like pointers) and Java (where arrays are object references),
	// arrays in Go are values.
	//(1) assigning one array to another copies all of the elements, and
	//(2) if you pass an array to a function, it will receive a copy of the array (not a pointer or reference to it).
	var x[5] int
	x[3] = 42
	fmt.Println(x)
	fmt.Println(len(x))

	fmt.Println("----------slice")
	//y:= type{values} //composite literal, which is an expression that creates a new instance
	// each time its evaluated
	//The zero value of a slice is nil.
	//slice allows you to group together values of the same type
	//Slices, on the other hand, are much more flexible, powerful, and convenient than arrays.
	// Unlike arrays, slices can be resized using the built-in append function.
	// Further, slices are reference types, meaning that they are cheap to assign
	// and can be passed to other functions without having to create a new copy of its underlying array.
	// Lastly, the functions in Go’s standard library all use slices rather than arrays in their public APIs.
	y := []int{4,57,42}
	fmt.Println(y)
	fmt.Println(y[1:])

	fmt.Println("----------slice add")
	y = append(y, 73)
	z:= []int{236,31}
	//... means all the values in z
	y = append(y, z...)

	fmt.Println("----------slice delete array")
	y = append(y[:2], y[4:]...)

	fmt.Println("----------another way to show the slice")
	for i, v := range  y{
		fmt.Println(i, v)
	}

	fmt.Println("----------slice make")
	//slice has the length and a capacity, the capacity is the number of elements un the underlying array
	//the make will create a zero value array, in this case 2 elements 0
	z1 := make([]int, 2, 3)
	fmt.Println(z1)
	fmt.Println(len(z1))
	fmt.Println(cap(z1))
	//z1 = append(z1, 1000) won't work, the capacity is two

	fmt.Println("----------dimensional slice")
	a1 := []string{"a", "b", "c"}
	a2 := []string{"1", "2", "3"}

	a3 := [][]string{a1,a2}
	fmt.Println(a3)

	fmt.Println("----------map")
	m := map[string]int{
		"Amarelo": 1,
		"Vermelho": 2,
		"Rosa" : 2,
	}
	fmt.Println(m)
	fmt.Println(m["Vermelho"])

	//if the key doesnt exist, the value will return 0
	v, ok := m["Verde"]
	fmt.Println(ok)
	fmt.Println(v)

	//to distinguish a missing entry from a zero value, you can descriminate with a form of multiple assigment
	//for convention this is called as ok
	if v, ok := m["Amarelo"]; ok{
		fmt.Println("Key: ", v)
	}

	//the range iterates through all entries of an array, slice, string or map, or values received on a channel
	//with one exception: if at most one iteration variable is present and len(x) is constant, the range expression is not evaluated.
	fmt.Println("----------add map")
	m["Azul"] = 3
	for k, v:= range m{
		fmt.Println(k, v)
	}

	x1 := []int{4,2,3,9}
	for k, v:= range x1{
		fmt.Println(k, v)
	}

	fmt.Println("----------delete")
	delete(m, "Rosa")
	//if a key doesnt exist no errors are thrown
	delete(m, "Roxo")
	fmt.Println(m)
	if v, ok := m["Roxo"]; ok{
		fmt.Println("Key: ", v)
		delete(m, "Roxo")
	}
}


//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	f := new(File)
//	f.fd = fd
//	f.name = name
//	f.dirinfo = nil
//	f.nepipe = 0
//	return f
//}

//composite literal
//func NewFile(fd int, name string) *File {
//	if fd < 0 {
//		return nil
//	}
//	f := File{fd, name, nil, 0}
//	return &f
//}