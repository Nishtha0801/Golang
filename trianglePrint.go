package main

import(
	"fmt"
)
func trianglePrint(n int){
	nst:=1
	
	for row:=1;row<=n;row++ {
		for cst:=1;cst<=nst;cst++ {
			fmt.Print("*\t")
		}
		nst++
		fmt.Println()
	}
}

func main(){
	var n int
	fmt.Println("Enter n:")
	fmt.Scanln(&n)

	trianglePrint(n)
}