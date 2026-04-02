Making Decisions (Go Conditions)

In programming, "conditions" are just "Yes/No" questions. If the answer is true, the computer does one thing; if it’s false, it does something else (or nothing at all).
The Tools You Use

To build these questions, you use the operators we talked about before:

    Math symbols: <, >, ==, etc.

    Logic symbols: && (AND), || (OR), ! (NOT).

The "If" Family (Decision Making)

Here is how you actually write those decisions in your code, explained like a total amateur:
1. The Simple if

The "Only if this happens" rule.

    What it is: "If this is true, do the thing. If not, just ignore it and keep going."

    Example: If it is raining, open an umbrella.

2. The if...else

The "One way or the other" rule.

    What it is: "If the first thing is true, do Action A. Otherwise (else), do Action B."

    Example: If you have $5, buy a taco. Else, buy a soda.

3. The else if

The "Check the next thing" rule.

    What it is: Use this when you have a list of options. It checks the first one; if that’s false, it checks the next one, and so on.

    Example: * If the light is Red, stop.

        Else if the light is Yellow, slow down.

        Else, go!

4. The "Nested If"

The "Decision inside a decision" rule.

    What it is: You put an if statement inside another if statement. You only even look at the second question if the first one was "true."

    Example: * If you are at the front of the line...
    * (Now check this): If you have a ticket, you can enter.
    Shutterstock

5. The switch

The "Shortcut" rule.

    What it is: If you are checking the same variable for a ton of different values (like checking if a day is Monday, Tuesday, Wednesday...), a switch is much cleaner than writing ten else if blocks.

Amateur Summary:

    if: The basic "Gatekeeper."

    else: The "Backup plan."

    else if: The "Alternative option."

    nested if: The "Double-check."

Does that make the flow of the code a bit easier to visualize?
