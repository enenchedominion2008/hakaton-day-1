Functions (The "Save for Later" Code)A function is like a recipe. You write it once, and then you can "cook" it whenever you need it. It won’t do anything until you specifically call its name.1. How to Make OneThe Keyword: Start with func.The Name: Give it a clear name (e.g., myMessage).The Body: Put your code inside curly braces { }.2. How to Use It (Calling)To make the function actually run, just type its name with parentheses: myMessage(). You can do this as many times as you want!Parameters and Arguments (Passing Info)Think of Parameters as "placeholders" inside the function, and Arguments as the "actual data" you toss into it when you call it.Parameters: Defined in the function signature. You must say what type they are (like string or int).Arguments: The specific value (like "Alice" or 25) you send to the function.Go// 'name' is the parameter
func sayHello(name string) {
  fmt.Println("Hello", name)
}

func main() {
  sayHello("Alice") // "Alice" is the argument
}
Return Values (Getting Info Back)Sometimes you don't just want a function to print something; you want it to give you a result back so you can use it later.The Setup: You have to tell Go what type of data the function will return.The Keyword: Use return to send the value back to whoever called it.Gofunc add(x int, y int) int {
  return x + y // Sends the sum back
}
Recursion (The Loop-de-Loop)Recursion is a bit of a brain-bender! It’s when a function calls itself.The Idea: It keeps calling itself until it reaches a "Base Case" (a stop sign). If you don't have a stop sign, the function will keep going forever until the computer crashes!Example: Calculating a factorial ($5!$) or walking through a folder that has more folders inside it.ShutterstockFixing the ExerciseTo complete your documentation exercise, you need to define the function and then call it inside main:Gopackage main   
import ("fmt")

// 1. Create the function
func myFunction() {
  fmt.Println("I just got executed!")
}

func main() {
  // 2. Call the function
  myFunction()
}
