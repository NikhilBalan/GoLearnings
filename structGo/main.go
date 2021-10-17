package main

import "fmt"



//struct go

type Address struct {
	city  string
	state string
}

type Address1 struct {
	city  string
	state string
}

type Person struct {
	name    string
	age     int
	 Address
}
func main() {

	type Employee struct {

		name string
		age int
		sal float64
	}

	emp1:=Employee{
		name:"nikhil",
		age:28,
	}

	emp2:=Employee{"Nikhil",28,78.9}
	fmt.Println("emp1:",emp1)
	fmt.Println("emp2:",emp2)
	// accessing each element
	fmt.Println("emp2 name:",emp2.name)
	fmt.Println("emp2 age:",emp2.age)

	//anonymous struct

	emp3:= struct {
		name string
		id int
		age int
	}{
		name:"chandhu",
		id:2890,
		age:56,

	}

	fmt.Println("emp3:",emp3)
	fmt.Println("emp3 id:",emp3.id)
//nested struct
	person1:=Person{
		name:    "nikki",
		age:     35,
		Address: Address{
			city:  "london",
			state: "uk",
		},
	}

	fmt.Println("person1: ",person1.city)

	//equality check

	adr1:=Address{
		city:  "KOCHI",
		state: "KOCHI",
	}

	adr2:=Address{
		city:  "KOCHI",
		state: "KOCHI",
	}

	if adr1==adr2{
		fmt.Println("equal")
	} else{
		fmt.Println("not equal")
	}

}
