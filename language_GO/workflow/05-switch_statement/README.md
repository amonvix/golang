# Switch Statements and Fallthrough in Go

This chapter introduces **`switch` statements** in Go and explains how they provide structured, readable branching for decision-making scenarios.

Switch statements offer a cleaner alternative to long `if / else if` chains and make intent explicit.

---

## Concepts Covered

- `switch` without an expression
- `switch` with a value expression
- Case matching rules
- Default cases
- `fallthrough` behavior
- Controlled execution flow
- Readable multi-branch logic

---

## Purpose

The goal of this chapter is to understand:

- When to use `switch` instead of chained conditionals
- How Go evaluates switch cases
- How boolean-based and value-based switches differ
- How `fallthrough` alters normal execution
- How to write clear and intentional branching logic

---

## Switch Statement Overview

Go supports two primary forms of `switch`:

### Expressionless Switch

A `switch` without an expression evaluates boolean conditions in each case.  
This form behaves like a structured alternative to `if / else if` chains.

Each case executes only when its condition evaluates to `true`.

---

### Value-Based Switch

A `switch` with an expression compares a single value against multiple cases.

This form improves readability when checking a variable against known constants.

---

## Fallthrough Behavior

By default, Go executes **only one case** in a `switch`.

Using `fallthrough` explicitly forces execution to continue into the next case, regardless of its condition.

This behavior makes control flow intentional and prevents accidental cascading logic.

The example demonstrates both:

- Full fallthrough chains
- Partial fallthrough without reaching the default case

---

## Example Behavior

- A boolean-based switch evaluates conditions in order
- A value-based switch matches specific values
- Fallthrough chains force sequential execution
- Default cases handle unmatched scenarios
- Each switch executes predictably and explicitly

---

## Key Takeaways

- `switch` improves readability over chained conditionals
- Expressionless switches evaluate boolean logic directly
- Value-based switches compare against known constants
- `fallthrough` must be explicit
- Go favors clear intent over implicit behavior

---

## Notes

Switch statements reduce cognitive load when branching logic grows.

When decisions multiply,  
structure beats repetition.
