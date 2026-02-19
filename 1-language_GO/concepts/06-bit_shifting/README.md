# bit shifting powers of two output formatting


This example demonstrates the use of **bit shifting powers of two output formatting
** in Go.

## Concepts Covered

- bit shifting operator for power of two calculations
- formatted output with fmt.Printf
- decimal and binary number representation
- tabular alignment using tab characters


## Purpose

To generate and display powers of two from 2^0 to 2^7 in both decimal and binary formats using bit shifting, with output aligned in columns for readability.


## Notes

- Using bit shifting (1 << n) is an efficient way to compute powers of two.
- The %b verb in fmt.Printf formats integers as binary strings.
- Tab characters (\t) help align output but may vary in appearance depending on terminal settings.
- Hardcoding each print statement limits scalability; a loop could improve maintainability.
