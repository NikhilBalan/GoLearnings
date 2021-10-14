package main

import "fmt"

//map in golang
func main() {
//using make
	employeeMap:=make(map[string]int)
	employeeMap["nikhil"]=28
	employeeMap["vijay"]=37
	employeeMap["sree"]=26
	employeeMap["anu"]=18
	employeeMap["balan"]=65

	fmt.Println("map:",employeeMap)

	//initializing during declaration

	empMap:=map[string]int{
		"nikhil":28,
		"saju":20,
	}

	empMap["thankamani"]=55

	fmt.Println("map2:",empMap)
	//fetch from map
	fmt.Println("value of nikhil in empMap:",empMap["nikhil"])

	//iterating map using range

	for key,value:=range employeeMap{
		fmt.Printf("employeeSalary[%s] = %d\n",key,value)
	}

	//delete from map
	delete(employeeMap,"nikhil")
	fmt.Println("after delete nikhil employeeSalary:",employeeMap)

	//map with struct

	type employee struct{
		salary int
		country string
	}

	emp1:=employee{
		salary: 20000,
		country: "india",
	}
	emp2:=employee{
		salary: 25000,
		country: "us",
	}

	emp3:=employee{
		salary: 146000,
		country: "India",
	}

	empSalMap:=map[string]employee{
		"chacko":emp1,
		"salbin":emp2,
		"nikhil":emp3,
	}

	for key,value:=range empSalMap{
		fmt.Println("name:",key,"value sal:",value.salary,"country:",value.country)
	}
	fmt.Println("size of map:",len(empSalMap))
}
