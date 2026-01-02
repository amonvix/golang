# Bitwise Operations and Untyped Constants in Go

This lesson demonstrates how Go handles **bitwise operations** and **untyped numeric constants**, highlighting how extremely large values can exist safely at compile time and adapt to context when used.

---

## Concepts Covered

- Bitwise shift operators (`<<`, `>>`)
- Untyped numeric constants
- Compile-time constant evaluation
- Context-dependent typing
- Function argument type requirements
- Integer vs floating-point usage
- Go’s constant precision model

---

## Purpose

The goal of this exercise is to understand:

- How bitwise shifts work in Go
- Why Go constants can represent extremely large values
- How untyped constants adapt to required types
- How Go enforces type safety without losing flexibility

---

## Explanation

In Go, constants are evaluated at **compile time** and are **untyped by default**.

This allows expressions like:

- Shifting `1` left by `100` bits to create a huge constant
- Shifting that value back down to a manageable size
- Passing the same constant to functions that expect different types

The constant only receives a concrete type **when context requires it**, such as when passed as a function argument.

---

## Example Behavior

- `Big := 1 << 100`  
  → Creates an untyped constant far larger than any built-in numeric type

- `Small := Big >> 99`  
  → Reduces the value to a size that fits safely into standard types

- `needInt(Small)`  
  → `Small` is interpreted as `int`

- `needFloat(Small)` and `needFloat(Big)`  
  → Both values are interpreted as `float64`

No explicit type conversion is required.

---

## Key Takeaways

- Constants in Go are not bound by machine word size
- Untyped constants adapt to their usage context
- Bitwise shifts are evaluated at compile time
- Type safety is enforced without sacrificing expressiveness
- Go avoids implicit runtime conversions

---

## Notes

This example highlights one of Go’s most powerful design choices:  
**strict typing with flexible constants**.

If it compiles, it’s safe.  
If it doesn’t, Go stopped you for a good reason.
