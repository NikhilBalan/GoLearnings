package main

import (
	"fmt"
)

//check two mapes equal or not by comparing each element
func main() {

	map1:=map[string]int{
		"nikhil":20,
		"saju":21,
		"salbin":22,
	}

	map2:=map[string]int{
		"nikhi":20,
		"saju":21,
		"salbin":22,
	}
	//flag:=true
	if len(map1)==len(map2){

		for i,v1:=range map1{
			value,ok:=map2[i]
			if ok {
				if v1!=value {
					fmt.Println("maps are not equal")
					//flag=false
					return
				}

			}else{
				fmt.Println("maps are not equal")
				//flag=false
				return
			}
		}
		fmt.Println("maps are equal")

	}
	//equal check using reflect
	/*flag:=reflect.DeepEqual(map1,map2)
	if flag{
		fmt.Println("maps are eql by reflect")
		return
	}
	fmt.Println("maps are not equal by reflect")*/


}
