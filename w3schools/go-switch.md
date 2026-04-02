The switch Statement (The "Pick One" Menu)

Think of a switch statement as a more organized way to write a long list of if...else if statements. It’s like a menu: the computer looks at your value and jumps straight to the option that matches.
How It Works

    The Expression: You give it one thing to check (like a variable called day).

    The Cases: You list out the possible values. If day matches a case, Go runs that code.

    No "Break" Needed: In other languages (like Java or C++), you have to tell the computer to stop after a match. In Go, it’s smarter—it finds the match, runs the code, and quits automatically.

Key Features

    The default Option: This is your "none of the above" backup. If none of your cases match, the default code runs. (e.g., If the day is 8, it prints "Not a weekday").

    Multi-Case: You can group values together with commas.

        Example: case 6, 7: fmt.Println("It's the weekend!")

    Matching Types: You can't compare apples to oranges. If your variable is a number (int), your cases must be numbers. You'll get an error if you try to check if 3 equals "Saturday".

Fixing the Exercise

To make the exercise code work, you need to tell the switch what variable it's looking at and use the case keyword:
Go

package main   
import ("fmt") 

func main() {
  var day = 2
  // You need to put 'day' here and 'case' before the numbers
  switch day {
  case 1:
    fmt.Print("Saturday")
  case 2:
    fmt.Print("Sunday")    
  }
}
