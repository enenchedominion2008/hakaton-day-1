package main

import "fmt"

func main() {
start:
	var cal1, cal2 int

	fmt.Print("input : ")
	fmt.Scan(&cal1)

	fmt.Print("input : ")
	fmt.Scan(&cal2)

	fmt.Println("choose your way")
	fmt.Println("(1) add")
	fmt.Println("(2) subtract")
	fmt.Println("(3) divide")
	fmt.Println("(4) multiply")

	var choose int
	fmt.Scanln(&choose)

	if choose == 1 {
		fmt.Println("final answer addition :", fmt.Sprint(cal1+cal2))
	} else if choose == 2 {
		fmt.Println("final subtraction answer :", fmt.Sprint(cal1-cal2))
	} else if choose == 3 {

		fmt.Println("final division answer :", fmt.Sprint(cal1/cal2))

	} else if choose == 4 {
		fmt.Println("final multiplication answer :", fmt.Sprint(cal1*cal2))
		

	} else {
		if cal1 %2 == 0 {
			fmt.Println("error division 0")
		}
		fmt.Println("error out of range")
	}

	goto start

}
