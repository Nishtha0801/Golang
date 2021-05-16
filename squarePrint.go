package main

import(
	"fmt"
)
func squarePrint(n int){
	nst:=n
	
	for row:=1;row<=n;row++ {
		for cst:=1;cst<=nst;cst++ {
			fmt.Print("*\t")
		}
		nst = n
		fmt.Println()
	}
}

func main(){
	var n int
	fmt.Println("Enter n:")
	fmt.Scanln(&n)

	squarePrint(n)
}