# Hands-on: Sequential and Conditional Control Flow

Create a program that uses **both SEQUENTIAL and CONDITIONAL** control flow.

## Requirements

- Create a random integer between **0 and 250**
- Store the value in a variable named `x`
  - Function reference:
    - `func Intn(n int) int`
    - `rand.Intn(n)`
- Print the **name** and **value** of the variable
- Use an `if` statement to evaluate the value of `x`:

### Conditional Rules

- If the value is between **0 and 100**
  - Print: `between 0 and 100`
- If the value is between **101 and 200**
  - Print: `between 101 and 200`
- If the value is between **201 and 250**
  - Print: `between 201 and 250`

## Inclusive vs Exclusive

Regarding `rand.Intn()`:

- Does it include **0** in its output?
- Does it include **250** in its output?

Demonstrate this behavior in code using the numbers **0–3**.

## Hint

- Use the logical AND operator: `&&`
