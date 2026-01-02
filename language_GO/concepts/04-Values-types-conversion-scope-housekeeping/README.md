# Values, Types, Conversion, Scope & Housekeeping

## Type Inference and Variable Scope in Go

This lesson expands on type inference by introducing **variable scope**, showing how Go controls where variables can be accessed and how visibility works across functions.

---

## Concepts Covered

- Type inference with `:=`
- Explicit variable declaration with `var`
- Package-level (global) scope
- Function-level (local) scope
- Block scope
- Access rules between functions
- Runtime type inspection using `%T`

---

## Purpose

The goal of this exercise is to understand:

- The difference between global and local variables
- How scope affects variable accessibility
- How Go enforces strict scope rules at compile time
- How type inference behaves across different scopes

---

## Explanation

Go determines variable scope based on **where the variable is declared**:

- Variables declared at the package level are accessible throughout the package
- Variables declared inside a function are only accessible within that function
- Variables declared inside a block are limited to that block
- Variables declared in one function are **not accessible** from another function unless passed explicitly

Type inference (`:=`) still applies normally, but **scope rules always win**.

---

## Example Behavior

- `var x int = 43`  
  → Package-level variable, accessible inside `main`

- `v := 42`, `i := 88`, `f := 3.142`, `g := 0.867 + 5i`  
  → Function-scoped variables inferred as `int`, `int`, `float64`, and `complex128`

- `z := 45` (inside another function)  
  → Only accessible within that function

Attempting to access a variable outside its scope results in a **compile-time error**, not a runtime error.

---

## Key Takeaways

- Scope is enforced at **compile time**
- Global variables are accessible inside functions, but not the other way around
- Type inference does not bypass scope rules
- `%T` is useful to inspect inferred types, not visibility
- Functions are natural scope boundaries in Go

---

## Notes

This example intentionally demonstrates both **what works** and **what fails**.  
If Go lets it compile, it’s valid.  
If it doesn’t, it’s protecting you — aggressively and correctly.
