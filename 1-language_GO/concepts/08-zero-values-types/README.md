# variable declaration and initialization in Go


This example demonstrates the use of **variable declaration and initialization in Go
** in Go.

## Concepts Covered

- package main and import usage
- zero-value initialization for variables
- explicit variable declaration with type
- short variable declaration syntax
- multiple variable declaration and initialization
- blank identifier usage to discard values
- formatted printing with fmt.Printf and type reflection


## Purpose

To illustrate different ways to declare, initialize, and print variables of various types, including the use of zero values, explicit types, short declarations, multiple assignments, and discarding unwanted values.


## Notes

- Variables declared without explicit initialization receive their zero value automatically.
- Short variable declaration (:=) infers the variable type from the assigned value.
- The blank identifier (_) can be used to ignore unwanted values in multiple assignments.
- Use fmt.Printf with %T to verify the concrete type of a variable at runtime.
