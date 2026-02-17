# Zero Values, Short Declarations, and the Blank Identifier in Go

This lesson introduces **zero values** in Go and explores how variables behave when declared without explicit initialization, alongside short declarations and the blank identifier.

---

## Concepts Covered

- Zero values in Go
- Package-level variable initialization
- Short variable declaration (`:=`)
- Multiples variable assignment
- Explicit type declaration
- Runtime type inspection with `%T`
- Blank identifier (`_`)

---

## Purpose

The goal of this exercise is to understand:

- What zero values are and why they exist
- How Go initializes variables by default
- The difference between declared and inferred values
- How to discard unwanted values in a safe way
- How Go enforces clarity in variable usage

---

## Explanation

In Go, the compiler automatically assigns a zero value to every variable that you declare without initialization.

For example:

- Numeric types default to `0`
- Floating-point types default to `0.0`
- Strings default to an empty string
- Booleans default to `false`

This design eliminates uninitialized memory and makes code safer by default.

---

### Short Declarations

Using `:=`, Go infers both the **type** and **value** at the same time.  
This is the most common way to declare variables inside functions.

---

### Explicit Types

You can still declare variables with explicit types when precision or intent matters, such as using `float32` instead of the default `float64`.

---

### Blank Identifier

The blank identifier (`_`) allows you to intentionally ignore values returned by assignments or functions.

This avoids unused variable errors while keeping code explicit and readable.

---

## Example Behavior

- A package-level `int` variable defaults to `0`
- Variables declared with `:=` receive inferred types
- A single line can initialize multiples variables
- Explicit types override default inference
- Discarded values do not exist in scope

---

## Key Takeaways

- Zero values are a core safety feature in Go
- Variables are always initialized
- Type inference favors clarity and correctness
- Explicit typing is still available when needed
- The blank identifier is intentional, not a hack

---

## Notes

Zero values remove an entire class of bugs common in other languages.  
Go chooses predictability over surprise — every time.
