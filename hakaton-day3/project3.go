package main 

import (
	"fmt"
	"strings"
	"bufio"
	"os"
)

func main()  {

	
	
	input :=  bufio.NewReader(os.Stdin)
	
	fmt.Println("input your code")
	fmt.Scan(&input)
	

	var choose string
	fmt.Println("input your method (1)"  )
	fmt.Scan(&choose)
	if choose == "1" {
		fmt.Println(strings.ToUpper(input))
	}
	
}