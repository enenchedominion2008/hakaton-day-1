Go Structs (The "Digital Folder")

A Struct (short for structure) is like a custom folder where you can keep different types of information together. While an Array only lets you store one type of thing (like a list of just numbers), a Struct lets you mix and match (like a name, an age, and a job).
1. Creating a Struct

To make one, you use the keywords type and struct. You’re basically defining a "template" for a record.

    Example: If you’re making a profile for a Person, you tell Go that every "Person" should have a name (string), an age (int), and a job (string).

2. Using the Dot Operator (.)

Once you’ve created a variable from your struct, you use a period/dot to reach inside and grab or change specific info.

    pers1.name = "Hege" — This sets the name.

    fmt.Println(pers1.name) — This reads the name.

3. Passing Structs to Functions

You can treat a whole struct like a single item. Instead of sending a function 10 different variables (name, age, salary, etc.), you just toss the entire Person struct into the function at once.

It makes your code much cleaner because the function knows exactly what's inside that "folder."
Amateur Documentation Summary

    Struct: A "collection" of different data types.

    Members: The individual items inside the struct (like age or name).

    Dot Operator: The key you use to "unlock" the info inside the struct.
