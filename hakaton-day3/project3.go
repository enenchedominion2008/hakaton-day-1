package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input string

	
	var choose string
	fmt.Println("input your method of conversion")
	fmt.Println("uppercase")
	fmt.Println("lowecase")
	fmt.Scan(&choose)

	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)


	// fmt.Println("put your words")
	fmt.Scan(&reader)

	if choose == "1" {
		fmt.Println(strings.ToUpper(input))
	}

}
