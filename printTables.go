package main

import(
	"fmt"
)

func printTable(num int){
	
	for i:=1;i<=10;i++{
		multiply:= num*i
		fmt.Println( num , "x" , i , "=" , multiply)
	}
}

func printAllTables(num int) {
	for i:=1;i<=num;i++{
		printTable(i);
		fmt.Println();
	}
}

func main(){
	var num int
	fmt.Println("Enter the number until which you wanna print table:")
	fmt.Scanln(&num)
	printAllTables(num)
}