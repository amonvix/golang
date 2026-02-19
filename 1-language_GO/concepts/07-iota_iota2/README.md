# go iota enumeration and bit shifting for byte size constants


This example demonstrates the use of **go iota enumeration and bit shifting for byte size constants
** in Go.

## Concepts Covered

- iota enumerator for sequential constant generation
- bit shifting to define powers of two
- custom typed constants with underlying int type
- non-standard scaling factors in constant definitions
- printing decimal and binary representations of constants


## Purpose

To illustrate the use of iota for generating sequential constants and applying bit shifting to define byte size units with both standard and custom scaling factors, while demonstrating their numeric and binary values.


## Notes

- Skipping the zero value with a blank identifier ensures enumeration starts from 1
- Bit shifting by multiples of 10 corresponds to powers of 1024, useful for byte sizes
- Custom multipliers applied to shifted values deviate from standard byte size definitions
- Defining constants with a custom type improves type safety and clarity
- Printing binary representations aids in understanding bit-level values of constants
