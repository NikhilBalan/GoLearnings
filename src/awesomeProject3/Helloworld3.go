package main

import (
	"fmt"
	"math"
	"unsafe"
)

func main(){
	var age int=30
	fmt.Println("my age is",age)
	age=28
	fmt.Println("my new age is",age)
	age =55
	fmt.Println("my current age is",age)

	//type inference
	var newAge=35
	fmt.Println("my final age is",newAge)

//multiple variable declaration

var firstAge,scondName =20,"nikki"
fmt.Println("first age:",firstAge,"second age:",scondName)

var(
	firstName="nikhil"
	lastName="balan"
	newId=25
)
fmt.Println("first name:",firstName,"last name:",lastName,newId)

//different variables without init value
var (
	name string
	oldAge int
)
fmt.Println("name",name,"age123",oldAge)

//shorthand variable declaration
count:=20
fmt.Println("count is:",count)

//shorthand variable declaration can only used to delcare atleast one new variable

a,b:=10.6,20.8
fmt.Println("a:",a,"b:",b)
b,c:=25.0,35.6
	fmt.Println("c:",c,"b:",b)

b,c=1.5,2.7
	fmt.Println("c:",c,"b:",b)

d:=math.Min(a,b)
fmt.Println("minimum:",d)

y,z:=15,35
fmt.Printf("type of y and z is %T,size of y and z is %d",y,unsafe.Sizeof(z))
fmt.Println()
//string variables

fname:="nikhil"
lname:="balan"
tName:=fname+" "+lname
fmt.Println("actual name is:",tName)

j:=10
k:=21.2
fmt.Println(j+int(k))
fmt.Println(float64(j))

//constant
const ca=50
fmt.Println(ca)
//function call
tot:=calculateBill(100,5)
fmt.Println("total price:",tot)

//function with multiple return values
area,perimeter:=rectangleCalc(10,5)
fmt.Println("area:",area,"perimeter:",perimeter)

	area1,perimeter1:=rectProps(10.6,5.2)
	fmt.Println("area:",area1,"perimeter:",perimeter1)
}

func calculateBill(price ,no int) int {
//	var totalPrice=price*no
	return price*no
}

func rectangleCalc(length,width int) (area int,perimtr int){
	 area=length*width
	 perimtr=(length+width)*2
	return area,perimtr
}

func rectProps(length, width float64)(area3, perimeter3 float64) {
	area3 = length * width
	perimeter3 = (length + width) * 2
	return //no explicit return value
}


