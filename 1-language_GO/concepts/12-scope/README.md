# variable scope and shadowing in go


This example demonstrates the use of **variable scope and shadowing in go
** in Go.

## Concepts Covered

- package-level variable scope
- block-level variable scope
- variable shadowing within nested scopes
- method definition on struct types
- function invocation from external packages
- use of short variable declaration within blocks


## Purpose

To illustrate how variable scope and shadowing operate in Go, including the distinction between package-level and block-level variables, and how shadowing affects variable visibility. It also demonstrates method calls on struct types and invoking functions from imported packages.


## Notes

- Package-level variables are accessible throughout the package unless shadowed by inner scope variables.
- Shadowing creates a new variable that temporarily hides the outer variable within its scope.
- Methods defined on struct types can be called on struct instances to encapsulate behavior.
- Variables declared in control structures like if statements are scoped only within those blocks.
- External package functions must be explicitly imported and invoked with their package qualifier.
