The for Loop (Doing Things Over and Over)

In Go, the for loop is the only loop available. It’s the tool you use when you want to repeat a block of code multiple times.
1. The Standard Three-Part Loop

The most common way to write a loop is by giving it three instructions:

    Start: i := 0 (Where do we begin?)

    Check: i < 5 (Should we keep going? If true, yes!)

    Step: i++ (What do we do after each round? Usually add 1).
    Shutterstock

2. Controlling the Loop (break and continue)

Sometimes you need to change the plan while the loop is running:

    continue: "Skip this one!" It stops the current round and jumps straight to the next one.

    break: "I'm out!" It stops the loop entirely and exits.

3. Nested Loops (Loops inside Loops)

You can put a loop inside another loop.

    The Rule: The "inner" loop finishes all its rounds for every single round of the "outer" loop.

    Example: If you have 2 adjectives and 3 fruits, the inner loop runs 3 times for the first adjective, then 3 times for the second.

4. The range Keyword (The Shortcut)

If you have a list (like an array or slice), range is the easiest way to go through it. It gives you two things:

    The Index: The position number (0, 1, 2...).

    The Value: The actual item (like "apple").

Amateur Tip: If you don't need one of those, use an underscore (_) to tell Go to ignore it. Go gets grumpy if you create a variable and don't use it!
Fixing the Exercise

To make the exercise code work, you need to add the Step (i++) so the loop doesn't run forever, and put the keyword for at the start:
Go

package main   
import ("fmt") 

func main() {
  // Added 'for' and 'i++'
  for i := 0; i < 6; i++ {
    fmt.Println(i)
  }
}
