package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	s:="Señor"

	for _,f:=range s{
		fmt.Printf("%c ",f)

	}
	fmt.Println()


	//counting string length

	word1 := "Señor"
	fmt.Printf("String: %s\n", word1)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word1))
	fmt.Printf("Number of bytes: %d\n", len(word1))


	//string equaltiy
	str1:="go"
	str2:="go"
	if str1==str2{
		fmt.Println("str1 and str2 are equal")
	}

	//string concatination

	str3:=str1+" "+str2
	fmt.Println(str3)
	//string modofication using SprintF
	str4:=fmt.Sprintf("%s%s",str1,str2)
	fmt.Println(str4)

str5:=modifyString([]rune(str4))

fmt.Println(str5)


}
//mutating the string
func modifyString(str4 []rune) string {

	str4[0]='a'
	str:=string (str4)
	for i:=range str4{
		fmt.Printf("%c ",str4[i])
	}
	return str

}
