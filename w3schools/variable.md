Go Variables - Simple Summary
In Go, variables are used to store data. There are different types of variables:

int → for whole numbers (e.g. 123, -50)
float32 → for decimal numbers (e.g. 19.99, -3.14)
string → for text (e.g. "Hello World") – must be in double quotes
bool → for true or false only

How to Create (Declare) Variables
There are 2 main ways to declare variables in Go:

Using var (Traditional way)Govar variablename type = valueExample: var name string = "John"
Using := (Short way - most common)Govariablename := valueGo automatically guesses the type from the value.

Important Notes:

With var, you can declare a variable without giving it a value immediately.
With :=, you must give a value right away.
Go is smart — if you don’t write the type, it will figure it out by itself.

Simple Example:
Gopackage main
import "fmt"

func main() {
    var student1 string = "John"     // full way
    var student2 = "Jane"            // type inferred
    x := 2                           // short way

    fmt.Println(student1)
    fmt.Println(student2)
    fmt.Println(x)
}
Output:
textJohn
Jane
2