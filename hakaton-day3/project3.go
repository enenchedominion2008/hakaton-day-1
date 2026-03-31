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
func reverse(text string) string {
    
    word := []rune(text)
    for i, j := 0, len(word)-1; i < j; i, j = i+1, j-1 {
        word[i], word[j] = word[j], word[i]
}
return string(word)
} 


func main() {
	
	start :
	var help string 
	fmt.Println("press (CONTINUE to use app or (HELP)")
	fmt.Scan(&help)
	if help == "help" {
		fmt.Println("i will")
		goto start
	} 
	if help == "continue" {
		goto start2
	}
	if help != "enter" && help != "continue" {
		fmt.Println("error please type (CONTINUE OR ENTER)")
	}
	start2 :
	var choose string
	fmt.Println("input your method of conversion")
	fmt.Println("upper :")
	fmt.Println("low :")
	fmt.Println("cap :")
	fmt.Println("reverse :")
	fmt.Scan(&choose)
	if choose != "cap" && choose != "low" && choose != "upper" && choose != "reverse" {
		fmt.Println("plase input something with the instructions")
		goto start2
	}
	

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
	if choose == "reverse" {
		fmt.Println(reverse(input))
	} 
	if choose == "quit" {
		fmt.Println("  Shutting down String Transformer. Goodbye.")
		return
	}
	goto start
}
