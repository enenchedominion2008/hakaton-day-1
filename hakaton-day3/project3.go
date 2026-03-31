package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func toupp(s string) string {
	
	return strings.ToUpper(s)
}

func Tolow(s string) string {
	return strings.ToLower(s)
}
func cap(s string) string {
	s = strings.ToLower(s)
	return strings.Title(s)
}
func snake(s string) string {
	if s != " " {
		return strings.ReplaceAll(s," ","_")
	}
	return s
}


func main() {
	

	
	var choose string
	fmt.Println("input your method of conversion")
	fmt.Println("upper")
	fmt.Println("low")
	fmt.Println("cap")
	fmt.Scan(&choose)
	var input string
	reader := bufio.NewReader(os.Stdin)
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	

	
	fmt.Scan(&reader)

	if choose == "upper" {
		fmt.Println(toupp(input))
	}
	if choose == "low" {
		fmt.Println(Tolow(input))
	}
	if choose == "cap" {
		
		fmt.Println(cap(input))
	}
	if choose == "snake"  {
		fmt.Println(snake(input))
	}
}
