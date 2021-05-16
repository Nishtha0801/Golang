package main

import(
	"fmt"
)

func main(){
	var marks int
	fmt.Println("Enter marks:")
	fmt.Scanln(&marks)
	
	if marks>90{
		fmt.Println("Excellent")
	} else if marks>80{
		fmt.Println("good")
	} else if marks>70{
		fmt.Println("fair")
	} else if marks>60{
		fmt.Println("meets expections")
	} else{
		fmt.Println("below par")
	}
}