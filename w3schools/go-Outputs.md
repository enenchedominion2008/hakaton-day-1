Go Output Functions - Quick Summary
Go has 3 main functions to print text:

fmt.Print()
Prints values without a newline at the end.
Adds a space between arguments only if they are not strings.
Example: fmt.Print("Hello", "World") → HelloWorld

fmt.Println()
Prints values with a space between them.
Automatically adds a new line (\n) at the end.
Best for normal printing.
Example: fmt.Println("Hello", "World") → Hello World

fmt.Printf()
Used for formatted printing.
Uses verbs like %v (value) and %T (type).
Example:Gofmt.Printf("Value: %v, Type: %T\n", 15, 15)


Quick Tip:

Use Println() for simple output.
Use Printf() when you want to control the format.


