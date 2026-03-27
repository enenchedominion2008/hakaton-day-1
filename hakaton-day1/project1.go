package main

import "fmt"

func main() {
start:
	red := "\033[31m"
	reset := "\033[0m"

	ot := "\033[32m"
	
	ma := "\033[33m"



	var n, a int

	fmt.Print("input first : ")
	cal1, err := fmt.Scanf("%v\n", &n)
	
	

	if err != nil {
		fmt.Println("== Input an number ==",cal1)
		goto start
	}

	fmt.Print("input second : ")
	cal2, err := fmt.Scanf("%v\n", &a)
	
	if err != nil {
		fmt.Println("== Input a valid number ==",cal2)
		goto start
	}

	fmt.Println("choose your way")
	fmt.Println("(1) add")
	fmt.Println("(2) subtract")
	fmt.Println("(3) divide")
	fmt.Println("(4) multiply")
	fmt.Println("(5) help menu")
	fmt.Println("(6) QUIT")

	var choose string
	fmt.Scanln(&choose)
	
	
	
	 if choose == "1" {
		fmt.Println(ma + "==final answer addition == :" + reset, fmt.Sprint(n+a))
	} else if choose == "2" {
		fmt.Println(ma +"==final answer suctraction == :"+ reset, fmt.Sprint(n-a)) 
	} else if choose == "3" {
		if a == 0 {
			fmt.Println(red +"=== error cannot be divided 0 ==="+ reset)
		} else {

			fmt.Println(ma + "==final division answer ==:" + reset, fmt.Sprint(n/a))
		}

	} else if choose == "4" {
		fmt.Println(red + "==final multiplication answer ==:" + reset, fmt.Sprint(n*a))

	} else if choose == "5" {
		fmt.Println(ot +"this is a special edition of calculator " + reset)
		fmt.Println(ot +"after inputing the values you choose either to add to subtract to divid or to exit" + reset)
		fmt.Println(ot +"either to add to subtract to divid or to exit"+ reset)
		fmt.Println(ot +"but to exit you have to retype same value twice and after that it will come out"+ reset)
		
	}else if choose == "6"{
		fmt.Println(ma +"==THANK YOU for calcutating==" + reset)
		return
	} else {
		fmt.Println(red +"error"+ reset)
	}
	goto start

}
