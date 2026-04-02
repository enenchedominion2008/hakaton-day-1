== CREAT SLICE ====

Go Slices - Simple Summary
Slices are like arrays, but more flexible — their length can grow and shrink.
They are used to store multiple values of the same type in one variable.
Ways to Create a Slice:

Directly (most common way):Gomyslice := []int{1, 2, 3}
myslice2 := []int{}     // empty slice
From an Array:Goarr := [6]int{10, 11, 12, 13, 14, 15}
myslice := arr[2:4]     // slice from index 2 to 4
Using make() function:Gomyslice := make([]int, 5, 10)   // length 5, capacity 10
myslice2 := make([]int, 5)      // length 5, capacity = length

Important Functions:

len(slice) → returns the number of elements (length)
cap(slice) → returns the maximum capacity the slice can grow to

Key Difference from Arrays:

Arrays have fixed length
Slices have dynamic length (can change)

Example Output:

Empty slice: length = 0, capacity = 0
Slice with values: length and capacity are usually equal to number of elements

== MODIFYSLICES==

. Getting and Changing Stuff

Think of a slice like a list.

    Indexes start at 0: So slice[0] is the first item, slice[1] is the second, and so on.

    Editing: You can swap an item out just by pointing to its number (e.g., prices[2] = 50).

2. Adding More Stuff

    append() is your best friend: Use it to add new items to the end of a list.

    Adding lists to lists: If you want to dump one slice into another, use three dots after the second one: append(list1, list2...). It’s like "unpacking" the second list into the first.

3. Changing Size

    Unlike regular arrays, slices are "stretchy." You can make them shorter by "re-slicing" them or longer by appending more data.

4. Don't Waste Memory

    The "Trap": If you take a tiny slice of a massive list, Go keeps the whole massive list in the background just in case.

    The Fix: Use copy(). This lets you grab only the items you actually need and put them in a fresh, small container, letting the computer throw away the big, bulky original.
