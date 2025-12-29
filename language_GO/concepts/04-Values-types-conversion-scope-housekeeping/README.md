# Values, Types, Conversion, Scope & Housekeeping

## Inspecting Inferred Types at Runtime

This lesson builds on Go’s type inference by showing how inferred types can be **inspected at runtime** using formatted output.

The focus here is not changing types, but **observing what Go decided**.

---

## Concepts Covered

- Type inference with `:=`
- Numeric literals and default types
- Floating-point numbers (`float64`)
- Complex numbers (`complex128`)
- Runtime type inspection using `%T`
- Basic program structure and housekeeping

---

## Purpose

The goal of this exercise is to understand:

- How Go assigns types during compilation
- How to verify inferred types during execution
- How different literals affect type inference
- Why complex numbers behave differently from real numbers

---

## Explanation

Go infers variable types at compile time based on the assigned value:

- Integer literals default to `int`
- Decimal literals default to `float64`
- Any expression containing `i` becomes a complex number (`complex128`)

Using `fmt.Printf` with the `%T` verb allows the program to print the inferred type of each variable, confirming Go’s internal decisions.

---

## Example Behavior

- `v := 42` → inferred as `int`
- `i := 88` → inferred as `int`
- `f := 3.142` → inferred as `float64`
- `g := 0.867 + 5i` → inferred as `complex128`

Each variable’s type is printed directly to the terminal at runtime.

---

## Key Takeaways

- Type inference happens before the program runs
- `%T` is a diagnostic tool, not a type converter
- Literal representation matters more than variable names
- Complex numbers are native and strongly typed in Go

---

## Notes

This example is intentionally simple and focused on observation.  
No casting. No conversions. Just clarity.

If you don’t know the type, ask Go — it will tell you.
