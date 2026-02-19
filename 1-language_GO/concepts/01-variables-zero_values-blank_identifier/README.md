# variable declaration initialization and multiple assignment in go


This example demonstrates the use of **variable declaration initialization and multiple assignment in go
** in Go.

## Concepts Covered

- short variable declaration using :=
- multiple assignment with ignored values using underscore
- zero value initialization for variables
- variable redeclaration rules within the same scope
- explicit variable declaration with var keyword


## Purpose

To illustrate different methods of declaring and initializing variables, including multiple assignment with ignored values, and to highlight redeclaration constraints and zero value defaults in Go.


## Notes

- The underscore (_) is used to discard unwanted values during multiple assignment.
- Redeclaring variables with := in the same scope requires at least one new variable to avoid compilation errors.
- Variables declared without explicit initialization receive the zero value of their type.
- Explicit variable declaration with var allows separation of declaration and initialization.
