package main

import(
	"fmt"
)

func isPrime(num int) bool{
	for i:=2;i*i<=num;i++{
		if num%i==0{
			return false;
		}
	}
	return true
}

func main(){
	var num int
	fmt.Println("Enter number:")
	fmt.Scanln(&num)
	
	fmt.Println(isPrime(num))

}