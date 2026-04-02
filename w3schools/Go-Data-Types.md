Go Data Types - Simple Summary
Data Type tells Go what kind of value a variable can store and how much space it needs.
Go is statically typed — once you set a variable’s type, it can’t change.
Go has 3 Basic Data Types:

  ========= BOOL ==========

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
============ INTERGER =============================
Go Integer Data Types - Simple Summary
Integer types are used to store whole numbers (no decimals).
There are two main categories:
1. Signed Integers
These can store both positive and negative numbers.

int     → Default type (size depends on your system)
int8    → Small range: -128 to 127
int16   → Medium range: -32768 to 32767
int32   → Larger range: -2147483648 to 2147483647
int64   → Very large range

2. Unsigned Integers
These can store only positive numbers (0 and above).

uint    == Default unsigned type
uint8   ==0 to 255
uint16 == 0 to 65535
uint32  == 0 to 4294967295
uint64  == Very large positive numbers

============FLOAT===========================

Go Float Data Types - Simple Summary
Float types are used to store numbers with decimal points (like 35.5, -2.34, or 123.456).
Go has two float types:

float32 → Smaller range
Size: 32 bits
Can store from -3.4e+38 to 3.4e+38
float64 → Larger range (more commonly used)
Size: 64 bits
Can store from -1.7e+308 to 1.7e+308

Important Tip:
If you don’t specify the type, Go uses float64 by default.
Which one to use?

Use float32 when you need smaller numbers and want to save memory.
Use float64 when you need bigger numbers or more precision.

Example of Error:
You cannot store 3.4e+39 in float32 because it’s too big.

============STRING================

Go String Data Type - Simple Summary

The string type is used to store text (a sequence of characters).
String values must be written inside double quotes " ".

Example:
Govar txt1 string = "Hello!"     // with initial value
var txt2 string                // empty (default is empty string)
txt3 := "World 1"              // short declaration
Output:
textType: string, value: Hello!
Type: string, value: 
Type: string, value: World 1
Key Point:
Empty string ("") is the default value when you declare a string without giving it a value.
