package main

import(
	"fmt"
	"bufio"
	"os"
	
)

func main(){
	var x int = 0
	var ans string
	for x<2{
		Scanner := bufio.NewScanner(os.Stdin)
		Scanner.Scan()
		String1:= Scanner.Text()
		
		ans+=String1
		x++;
	}
	
	fmt.Println(ans)
}