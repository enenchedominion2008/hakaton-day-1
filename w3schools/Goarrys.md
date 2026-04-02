Go Arrays - Simple Summary
An Array is used to store multiple values of the same type in one variable.
How to Declare an Array:
Go// With length
var arr = [3]int{1, 2, 3}
arr2 := [5]int{10, 20, 30, 40, 50}

// With inferred length (...)
var arr3 = [...]int{1, 2, 3}
arr4 := [...]string{"Apple", "Banana"}
Key Points:

Arrays have a fixed size (you can’t change the length later).
Index starts at 0 (first element is arr[0]).
You can change values: prices[2] = 50
Uninitialized elements get default values (0 for int, "" for string).
Use len(arr) to get the number of elements.

Simple Examples:

Access element: fmt.Println(arr[0])
Change element: arr[2] = 99
Initialize specific positions: [5]int{1:10, 3:50}

Exercise Answer:
Govar cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
fmt.Print(cars)
