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
	fmt.Println("capitalize")
	fmt.Scan(&choose)

	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)


	
	fmt.Scan(&reader)

	if choose == "uppercase" {
		fmt.Println(strings.ToUpper(input))
	}
	if choose == "lowercase" {
		fmt.Println(strings.ToLower(input))
	}
	if choose == "cap" {
		input = strings.ToLower(input)
		fmt.Println(strings.Title(input))
	}
}
