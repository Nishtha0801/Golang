package main

import(
	"fmt"
)

func printReverse(num int){
	for num!=0{
		rem:= num%10
		num = num/10
		fmt.Print(rem)
	}
}

func main(){
	var num int
	fmt.Println("Enter number:")
	fmt.Scanln(&num)

	printReverse(num)
}