package main

import(
	"fmt"
)

func countDigits(num int) int{
	len:=0
	for num!=0{
		num = num/10
		len++
	}
	return len
}

func main(){
	var num int
	fmt.Println("Enter number:")
	fmt.Scanln(&num)

	fmt.Println(countDigits(num))
}