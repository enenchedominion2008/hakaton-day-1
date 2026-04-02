GO COMMENTS

comments in golang are used to ignore lines of code for exampple if you want to put a refrence and you dont want it to affect the code   eg

// This is a comment
package main
import ("fmt")

func main() {
  // This is a comment
  fmt.Println("Hello World!")
}

we also have 


Go Multi-line Comments

multi-line comments start with /* and ends with */.

Any text between /* and */ will be ignored by the compiler:
Example
package main
import ("fmt")

func main() {
  /* The code below will print Hello World
  to the screen, and it is amazing */
  fmt.Println("Hello World!")
}

