# Go Package Function Invocation and Output


This example demonstrates the use of **Go Package Function Invocation and Output
** in Go.

## Concepts Covered

- Importing and utilizing external Go packages
- Function calls from imported packages
- Standard output printing with fmt.Println
- Handling multiple function return values
- Package-level side-effect function invocation


## Purpose

Invoke and display results from multiple functions provided by an external package, demonstrating integration and output of package functionalities within a Go application. Additionally, execute a package function with side effects or internal processing without capturing its return value.


## Notes

- Ensure the external package is properly installed and accessible in the module's environment.
- Handle potential panics or errors if the imported functions have side effects or internal state changes.
- Output order reflects the sequence of function calls, which may be relevant if functions depend on internal package state.
- Avoid ignoring returned values if they are critical for further processing or error handling.
