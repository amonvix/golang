# Go Character Combat Coordination with External Package Calls


This example demonstrates the use of **Go Character Combat Coordination with External Package Calls
** in Go.

## Concepts Covered

- Package import and aliasing
- Function invocation from external packages
- Sequential execution of combat-related actions
- Formatted output using fmt.Printf
- Global variable usage for character identification


## Purpose

Coordinate a sequence of combat actions by invoking attack, defense, and evasion functions from an external package, while providing formatted status output reflecting each action performed by a specified character.


## Notes

- The external package functions are assumed to handle the underlying combat mechanics independently.
- Global variable usage for character name simplifies repeated references but may limit concurrency and flexibility.
- Output formatting uses standard library functions to maintain clear and consistent messaging.
- Error handling and concurrency considerations are not addressed in this implementation.
