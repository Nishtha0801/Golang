package main
import (
	"fmt"
)
func isEven(num int) bool{
	if num%2==0{
		return true
	} else{
		return false
	}
}
func main(){
	var num int
	fmt.Println("Enter the number:")
	fmt.Scanln(&num)
	fmt.Println(isEven(num))
}