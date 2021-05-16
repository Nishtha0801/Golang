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
func printAllPrimesInRange(low int,high int){
	for i:=low;i<=high;i++{
		if(isPrime(i)){
			fmt.Println(i)
		}
	}
}

func main(){
	var low int
	fmt.Println("Enter low:")
	fmt.Scanln(&low)
	var high int
	fmt.Println("Enter high:")
	fmt.Scanln(&high)
	
	printAllPrimesInRange(low,high)

}