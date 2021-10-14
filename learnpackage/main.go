package main

import (
	"fmt"
	"learnpackage/simpleinterest"
	"log"
)

var p, r, t = 5000.0, 10.0, 1.0

/*
* init function to check if p, r and t are greater than zero
 */
func init() {
	println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {

	/*for i:=1;i<=5;i++{
		fmt.Println("inside main loop")
	}*/
	fmt.Println("Simple interest calculation")
	si := simpleinterest.Calculate(p, r, t)
	fmt.Println("Simple interest is", si)
//switch statement and fallthrough
	switch si {

	case 500:
		fmt.Println("case executed")
		fallthrough
	case 501:
		fmt.Println("case executed 2")
	case 503:
		fmt.Println("case executed 3")
	}
// array with for loop range
	arr:=[...]int{2,5,6,7,2}

	for i,v:=range arr{

		fmt.Println("index:",i,"value:",v)
	}
}
