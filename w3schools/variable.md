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

Go Multiple Variables - Quick Summary
You can declare multiple variables in Go in 3 easy ways:

Same Line (Same Type)Govar a, b, c, d int = 1, 3, 5, 7
Same Line (Different Types)
Use var or := without specifying type:Govar a, b = 6, "Hello"
c, d := 7, "World!"
In a Block (Cleanest way for many variables)Govar (
    a int
    b int = 1
    c string = "hello"
)

Key Point:
If you use var and specify the type, all variables on that line must be the same type.

Go Variable Naming Rules (Simple Summary)
Rules for naming variables in Go:

Must start with a letter or underscore (_)
Cannot start with a number
Can only contain letters, numbers, and underscores (a-z, A-Z, 0-9, _)
Names are case-sensitive (age, Age, and AGE are different)
Cannot contain spaces
Cannot be a Go keyword
No limit on how long the name can be

Multi-Word Variable Names:

Camel Case: myVariableName
Pascal Case: MyVariableName
Snake Case: my_variable_name