package main

import "fmt"

//pointer in go
func main() {

	b:=255
	var a *int=&b
	var c *int
	fmt.Printf("type of a is %T\n",a)
	fmt.Println("address of b is ",a)
	fmt.Println("val of c is ",c)
	c=&b
	fmt.Println("val of c after initialization",c)

	//new initialization

	size:=new(int)
	fmt.Printf("val of size is %d ,type of size is %T, address of size is %v\n",*size,size,size)
*size=56
	fmt.Printf("val of size is %d" ,*size)

//modification using pointer

d:=100
e:=&d
fmt.Println("\nval of b is",*e)
*e++
	fmt.Println("val of b after pointer modification is",*e,"\n")

 //passing pointer to function

 change(e)
fmt.Println("after change by function val of b",*e)

//returning pointer from function

f:=getPointer()
fmt.Println("returned pointer value",*f)



}

func getPointer() *int {
	i:=79
	return &i
}

func change(e *int) {
	*e++
}
