package main

import(
	"fmt"
	
)

func main(){
	var l int
	var b int
	fmt.Println("Enter length:")
	fmt.Scanln(&l)
	fmt.Println("Enter breadth:")
	fmt.Scanln(&b)

	fmt.Println(l*b)

}