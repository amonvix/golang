# Conditional Execution in Go

This chapter introduces **conditional execution** in Go, showing how programs make decisions at runtime using `if`, `else if`, and `else` statements.

Conditionals form the foundation of control flow, allowing a program to react to values, comparisons, and boolean expressions.

---

## Concepts Covered

- `if` statements
- `if / else` branching
- `if / else if / else` chains
- Comparison operators
- Boolean expressions
- Modulo operator (`%`)
- Decision-making based on runtime values

---

## Purpose

The goal of this chapter is to understand:

- How Go evaluates conditions
- How execution branches based on boolean expressions
- How to structure simple and compound conditionals
- How comparison results drive program behavior

This chapter focuses on clarity and correctness rather than complexity.

---

## Conditional Logic Overview

Go evaluates conditional expressions using boolean logic.

An `if` statement executes its block when the condition evaluates to `true`.  
When the condition evaluates to `false`, execution moves to the next applicable branch.

You can chain conditions using `else if` to express multiples mutually exclusive paths.

---

## Comparison and Decisions

The example demonstrates comparisons such as:

- Less than (`<`)
- Equality (`==`)
- Greater than (`>`)

These comparisons allow the program to classify values relative to a known reference point.

The modulo operator (`%`) enables parity checks, such as determining whether a number is even or odd.

---

## Example Behavior

- The program compares numeric values against a fixed reference
- Different messages print depending on the comparison result
- A parity check determines whether a number is even or odd
- Only one conditional branch executes per decision chain

---

## Key Takeaways

- Conditionals control execution flow
- Boolean expressions drive decision-making
- Go favors explicit branching over implicit behavior
- Clear conditions improve readability and maintainability
- Simple conditionals form the basis for complex logic

---

## Notes

Conditional logic is one of the first tools used to shape program behavior.

Before loops, concurrency, or abstraction, a program must learn how to decide.
