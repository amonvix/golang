# unexported package variable declaration


This example demonstrates the use of **unexported package variable declaration
** in Go.

## Concepts Covered

- package-level variable
- unexported identifier
- variable scope restriction
- constant-like variable usage


## Purpose

To declare a package-scoped variable representing the average orbital speed of a planet, restricting its visibility to within the package.


## Notes

- Using an unexported variable limits access to internal package use only.
- Naming conventions with lowercase initial letters enforce encapsulation.
- Consider using a constant if the value is immutable to improve clarity and performance.
- Documenting variables aids maintainability despite limited visibility.
