package main

import(
	"fmt"
)

func main(){
	var num int
	fmt.Println("Enter number:")
	fmt.Scanln(&num)

	if num%2==0{
		fmt.Println("even")
	}else{
		fmt.Println("odd")
	}
}