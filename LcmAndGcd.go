package main

import(
	"fmt"
)
func  gcd(a int , b int) int{
	if a==0 {
		return b
	}
	return gcd(b % a, a)
}

func lcm (a int , b int) int{
	return (a / gcd(a, b)) * b
}

func main(){
	var a int
	var b int
	fmt.Println("enter a")
	fmt.Scanln(&a)
	fmt.Println("enter b")
	fmt.Scanln(&b)

	var action int
	fmt.Println("Enter 1 for lcm & 2 for gcd: ")
	fmt.Scanln(&action)

	if action == 1 {
       fmt.Println(lcm(a,b))
    } else if action == 2 {
        fmt.Println(gcd(a,b))
    } else {
        fmt.Println("No such command found")
    }

}