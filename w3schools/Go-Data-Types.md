Go Data Types - Simple Summary
Data Type tells Go what kind of value a variable can store and how much space it needs.
Go is statically typed — once you set a variable’s type, it can’t change.
Go has 3 Basic Data Types:

bool → true or false
Numeric → numbers (includes int, float, and complex)
string → text (e.g. "Hello")

Example:
Govar a bool = true        // Boolean
var b int = 5            // Integer
var c float32 = 3.14     // Float
var d string = "Hi!"     // String

Go Boolean Data Type - Simple Summary

The bool type is used to store true or false only.
Default value of a bool is false.
You can declare it with or without the bool keyword.

Simple Example:
Govar b1 bool = true     // with type
var b2 = true          // type inferred
var b3 bool            // default is false
b4 := true             // short way
Output:
texttrue
true
false
true

