Go Maps (The "Labelled Storage" System)

A Map is like a collection of lockers. Each locker has a Key (the name on the label) and a Value (what’s inside). To get the stuff, you just need the right key!
1. Creating a Map

There are two main ways to start:

    The Quick Way: m := map[string]int{"Alice": 25}.

    The "Right" Way (Empty Map): Use make(), like m := make(map[string]int).

        Amateur Warning: If you just declare var m map[string]int without make and try to add items, your program will crash (panic)!

2. Using the Map

    Add/Update: m["year"] = 1964. It’s the same syntax for both!

    Delete: Use the delete() function: delete(m, "year").

    Length: Use len(m) to see how many items are inside.

3. The "Unordered" Rule

Maps are unordered. Even if you put items in as "One, Two, Three," Go might print them as "Three, One, Two." If you need things in a specific order, you have to keep a separate list (slice) of the keys.
4. Checking if a Key Exists

If you look for a key that isn't there, Go gives you a "zero value" (like 0 or ""). To be sure if it's actually missing, use this "two-value" trick:
Go

val, ok := m["color"] 
// ok will be 'true' if it exists, 'false' if it doesn't.

5. Maps are References

This is important: if you copy a map (b := a), they both point to the same data. If you change something in b, it changes in a too!
Quick Amateur Guide
Feature	Rule
Keys	Must be things that can be compared with == (strings, numbers, etc.). No slices or other maps!
Values	Can be anything you want.
Duplicates	Not allowed for Keys. Every label must be unique.
Looping	Use range. It gives you the key and the value in each round.
