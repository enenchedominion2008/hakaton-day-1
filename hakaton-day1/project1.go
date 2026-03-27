package main

import "fmt"

func main() {
start:
	red := "\033[31m"
	reset := "\033[0m"

	ot := "\033[32m"
	
	ma := "\033[33m"



	var n, a int

	fmt.Print(ot +"input first : "+ reset)
	cal1, err := fmt.Scanf("%v\n", &n)
	
	

	if err != nil {
		fmt.Println(red +"== Input an number =="+ reset,)
		
	}

	fmt.Print(ot +"input second : "+ reset,)
	cal2, err := fmt.Scanf("%v\n", &a)
	
	if err != nil {
		fmt.Println(red +"== Input a valid number =="+ reset,)
		
	}
	fmt.Println(cal1,cal2)

	fmt.Println(ma +"choose your way"+reset)
	fmt.Println(ma +"(1) add"+ reset)
	fmt.Println(ma +"(2) subtract"+ reset)
	fmt.Println(ma +"(3) divide" + reset)
	fmt.Println( ma +"(4) multiply"+ reset)
	fmt.Println(ma +"(5) HELP" + reset)
	fmt.Println(ma +"(6) QUIT"+ reset)

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

			fmt.Println(ma + "==final division answer ==:" + reset, fmt.Sprint(n/a),"//REMAINDER// == ",n%a)
		}

	} else if choose == "4" {
		fmt.Println(ma + "==final multiplication answer ==:" + reset, fmt.Sprint(n*a))

	} else if choose == "5" {
		fmt.Println(ot +"this is a special edition of calculator " + reset)
		fmt.Println(ot +"after inputing the values you choose either to add to subtract to divid or to exit" + reset)
		fmt.Println(ot +"either to add to subtract to divid or to exit"+ reset)
		fmt.Println(ot +"but to exit you have to retype same value twice and after that it will come out"+ reset)
		fmt.Println(ot +"to add press (1) \n to subtract (2) \n to divide (3) \n to multiply (4) \n to enter help (5)"+ reset)
		
	}else if choose == "6"{
		fmt.Println(ma +"==THANK YOU for calcutating==" + reset)
		return
	} else {
		fmt.Println(red +"error"+ reset)
	}
	goto start

}
