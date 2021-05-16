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

func rotateDigits(num int , r int) int{
	digit:= countDigits(num)
	r%=digit
	if r<0 {
		r+=digit
	}

	mul:=1
	div:=1

	for i:=1;i<=digit;i++{
		if i<=r{
			div*=10
		}else{
			mul*=10
		}
	}

	a:= num%div
	b:= num/div

	return (a*mul+b)

}

func main(){
	var num int
	fmt.Println("Enter number:")
	fmt.Scanln(&num)
	var r int
	fmt.Println("Enter r:")
	fmt.Scanln(&r)

	fmt.Println(rotateDigits(num,r))
}