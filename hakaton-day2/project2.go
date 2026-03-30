package main

import (
    "fmt"
    "strconv"
    "strings"
)

func main() {

    col := "\033[1;32m"
    reset := "\033[0m"
    idk := "\033[1;33m"
    dom := "\033[1;31m"

start2:
    var help string

    fmt.Println(col + "welcome to our converter app please type HELP if you want to understand our app (OR TYPE ENTER TO CONTINUE) " + reset)
    fmt.Scan(&help)
    if help == "help" {
        fmt.Println(col + "wellcome to the help menu here is how to use  the converter app " + reset)
        fmt.Println(idk + "when you enter you will see a menu that will ask you to select the type of conversion you wantvto make" + reset)
        fmt.Println(idk + "(1) if you select the first it means that you want to convert hexadecimal to decimal eg 1E" + reset)
        fmt.Println(idk + "(2) the second is to convert binary to decimal eg 1100011" + reset)
        fmt.Println(idk + "(3) the third option is to convert decimal to binary " + reset)
        fmt.Println(dom + "NOTE : YOU CANNOT EXIT THE APP UNTIL YOU FINISH CONVERSIONS" + reset)

    }
    if help == "enter" {
        fmt.Println("redireting into main page .............................................")
        goto start1
    }
    if help != "enter" && help != "help" {
        fmt.Println(dom + "PLEASE either type" + reset)
        fmt.Println(idk + "ENTER or HELP" + reset)
    }
    goto start2

start:
start1:
    var choose string

    fmt.Println(col + "please input your choice of conversion" + reset)
    fmt.Println("(1) HexToDec")
    fmt.Println("(2) BinToDec")
    fmt.Println("(3) DecTobin")

    fmt.Scan(&choose)

    if choose != "1" && choose != "2" && choose != "3" {
        fmt.Println("Not a valid input")
        goto start
    }

    var value string
    fmt.Println("input your value")
    n, err := fmt.Scan(&value)
    if err != nil {
        fmt.Println("please input a valid value", n)
    }

    if choose == "1" {
        fmt.Println(col + "your value from hexadecimal to decimal :" + reset)
        val, err := strconv.ParseInt(value, 16, 64)
        if err != nil {
            fmt.Println("Not a valid hex")
        } else {
            fmt.Println(val)
        }
    }

    if choose == "2" {
        val, err := strconv.ParseInt(value, 2, 64)
        if err != nil {
            fmt.Println("not a valid bin")
        } else {
            fmt.Println(val)
        }
    }
    if choose == "3" {

        val, err := strconv.ParseInt(value, 10, 64)
        if err != nil {
            fmt.Println("is not a dec")
        } else {
            fmt.Println(strconv.FormatInt(val, 2))
            fmt.Println(strings.ToUpper(strconv.FormatInt(val, 16)))
        }
    }
    if choose == "4" {
        fmt.Println("THANK YOU FOR SPENDING TIME HERE")
        return
    }

    var quit string
    fmt.Println("DO you want to stop now or continue (YES to quit) or (NO to continue)")
    fmt.Scan(&quit)
    if quit == "yes" {
        fmt.Println(" 👏🏻  👏🏻 thank you for converting with us 🥹🥹 ")
        return
    } else {
        if quit == "no" {
            goto start1
        }
    }
    goto start
}